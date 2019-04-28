package types

import (
	"errors"
	"fmt"
	sharedML "github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/validator"
	"log"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

type HyperlinkInfo struct {
	hyperlink *ml.Hyperlink
	styleID   format.DirectStyleID
	linkType  hyperlinkType
}

type hyperlinkOption func(o *HyperlinkInfo)
type hyperlinkType byte

const (
	hyperlinkTypeUnknown hyperlinkType = iota
	hyperlinkTypeWorkbook
	hyperlinkTypeWeb
	hyperlinkTypeEmail
	hyperlinkTypeFile
)

//Hyperlink is a 'namespace' for all possible settings for hyperlink
var Hyperlink hyperlinkOption

//NewHyperlink creates and returns a new HyperlinkInfo object that holds settings for hyperlink and related styles
func NewHyperlink(options ...hyperlinkOption) *HyperlinkInfo {
	i := &HyperlinkInfo{
		hyperlink: &ml.Hyperlink{},
	}
	i.Set(options...)
	return i
}

//Set sets new options for hyperlink
func (i *HyperlinkInfo) Set(options ...hyperlinkOption) {
	for _, o := range options {
		o(i)
	}
}

func (i *HyperlinkInfo) Validate() error {
	//TODO: add validators here
	switch i.linkType {
	case hyperlinkTypeUnknown:
		return errors.New("unknown type of hyperlink")
	case hyperlinkTypeWorkbook:
		if len(i.hyperlink.Location) == 0 {
			return errors.New("unknown location at target workbook")
		}
	case hyperlinkTypeWeb:
		if len(i.hyperlink.RID) > internal.UrlLimit {
			return errors.New(fmt.Sprintf("url exceeded maximum allowed length (%d chars)", internal.UrlLimit))
		}
	case hyperlinkTypeEmail:
		if len(i.hyperlink.RID) > internal.UrlLimit {
			return errors.New(fmt.Sprintf("email exceeded maximum allowed length (%d chars)", internal.UrlLimit))
		}
	case hyperlinkTypeFile:
		if len(i.hyperlink.RID) > internal.UrlLimit {
			return errors.New(fmt.Sprintf("link to file exceeded maximum allowed length (%d chars)", internal.UrlLimit))
		}
	}

	return nil
}

func (i *HyperlinkInfo) Formatting() format.DirectStyleID {
	return i.styleID
}

func (i *HyperlinkInfo) Target() string {
	switch i.linkType {
	case hyperlinkTypeFile:
		fallthrough
	case hyperlinkTypeEmail:
		return string(i.hyperlink.RID)
	case hyperlinkTypeWeb:
		if len(i.hyperlink.Location) > 0 {
			return fmt.Sprintf("%s#%s", i.hyperlink.RID, i.hyperlink.Location)
		}

		return string(i.hyperlink.RID)

	case hyperlinkTypeWorkbook:
		path := string(i.hyperlink.RID)
		location := i.hyperlink.Location

		if len(location) > 0 && location[0] == '#' {
			location = location[1:]
		}

		if len(location) > 0 {
			return fmt.Sprintf("%s#%s", path, location)
		}

		return path
	}

	return ""
}

func (o *hyperlinkOption) Formatting(styleID format.DirectStyleID) hyperlinkOption {
	return func(i *HyperlinkInfo) {
		i.styleID = styleID
	}
}

func (o *hyperlinkOption) Tooltip(tip string) hyperlinkOption {
	return func(i *HyperlinkInfo) {
		i.hyperlink.Tooltip = tip
	}
}

func (o *hyperlinkOption) Display(display string) hyperlinkOption {
	return func(i *HyperlinkInfo) {
		i.hyperlink.Display = display
	}
}

func (o *hyperlinkOption) ToMail(address, subject string) hyperlinkOption {
	return func(i *HyperlinkInfo) {
		if len(subject) > 0 {
			i.hyperlink.RID = sharedML.RID(fmt.Sprintf("mailto:%s?subject=%s", address, subject))
		} else {
			i.hyperlink.RID = sharedML.RID(fmt.Sprintf("mailto:%s", address))
		}

		i.linkType = hyperlinkTypeEmail
	}
}

func (o *hyperlinkOption) ToUrl(address string) hyperlinkOption {
	return func(i *HyperlinkInfo) {
		if u, err := url.Parse(address); err == nil {
			if u.Fragment != "" {
				i.hyperlink.Location = u.Fragment
				u.Fragment = ""
			}

			i.hyperlink.RID = sharedML.RID(u.String())
			i.linkType = hyperlinkTypeWeb
		}
	}
}

func (o *hyperlinkOption) ToFile(fileName string) hyperlinkOption {
	return func(i *HyperlinkInfo) {
		//change the directory separator from Unix to DOS
		fileName = strings.Replace(fileName, "/", "\\", -1)

		//add the file:/// URI to the url for Windows style "C:/" link and network shares
		if matched, err := regexp.MatchString(`^((\w+:.*)|(\\))`, fileName); matched && err == nil {
			fileName = "file:///" + fileName
		}

		//convert a '.\dir\filename' link to 'dir\filename'
		re := regexp.MustCompile(`^\.\\`)
		fileName = re.ReplaceAllString(fileName, "")

		i.hyperlink.RID = sharedML.RID(fileName)

		//workbook can be internal(same) or external
		if ext := filepath.Ext(fileName); ext == ".xlsx" || ext == ".xls" {
			i.linkType = hyperlinkTypeWorkbook
		} else {
			i.linkType = hyperlinkTypeFile
		}
	}
}

func (o *hyperlinkOption) ToRef(ref Ref, sheetName string) hyperlinkOption {
	return func(i *HyperlinkInfo) {
		i.linkType = hyperlinkTypeWorkbook

		if len(ref) > 0 {
			if len(sheetName) > 0 {
				//sheet + ref
				i.hyperlink.Location = fmt.Sprintf("#%s!%s", escapeLocation(sheetName), ref)
			} else {
				//ref only, cell be cell or bookmark
				i.hyperlink.Location = fmt.Sprintf("#%s", escapeLocation(string(ref)))
			}
		}
	}
}

/*

../Budgets/Annual/Budget2010.xlsx
../Budgets/Annual/Budget2010.xlsx#'Sheet3'
../Budgets/Annual/Budget2010.xlsx#'Sheet3'!G43
../Budgets/Annual/Budget2010.xlsx#DeptTotals


		-other sheet
		=HYPERLINK("#Sheet2!A1", "Sheet2")
		=HYPERLINK("#'Price list'!A1", "Price list")

		-same sheet
		=HYPERLINK("#A1", "Go to cell A1")

		-other local workbook
		=HYPERLINK("D:\Source data\Book3.xlsx", "Book3")
		=HYPERLINK("[D:\Source data\Book3.xlsx]Sheet2!A1", "Book3")

		-other network workbook
		=HYPERLINK("\\SERVER1\Svetlana\Price list.xlsx", "Price list")
		=HYPERLINK("[\\SERVER1\Svetlana\Price list.xlsx]Sheet4!A1", "Price list")

		- url
		=HYPERLINK("https://www.ablebits.com","Go to Ablebits.com")

		-email
		=HYPERLINK("mailto:support@ablebits.com","Drop us an email")

		- file
		=HYPERLINK("D:\Word files\Price list.docx","Price list")
*/

//ToTarget is very close to HYPERLINK function of Excel
// https://support.office.com/en-us/article/hyperlink-function-333c7ce6-c5ae-4164-9c47-7de9b76f577f
//
// a) to resource:
//	"location" or "[location]"
// b) to target at resource
//	"[location]target or "location#target"
// Here are some examples of supported values
//	=HYPERLINK("http://example.microsoft.com/report/budget report.xlsx", "Click for report")
//	=HYPERLINK("[http://example.microsoft.com/report/budget report.xlsx]Annual!F10", D1)
//	=HYPERLINK("[http://example.microsoft.com/report/budget report.xlsx]'First Quarter'!DeptTotal", "Click to see First Quarter Department Total")
//	=HYPERLINK("[http://example.microsoft.com/Annual Report.docx]QrtlyProfits", "Quarterly Profit Report")
//	=HYPERLINK("\\FINANCE\Statements\1stqtr.xlsx", D5)
//	=HYPERLINK("D:\FINANCE\1stqtr.xlsx", H10)
//	=HYPERLINK("[C:\My Documents\Mybook.xlsx]Totals")
//	=HYPERLINK("[Book1.xlsx]Sheet1!A10","Go to Sheet1 > A10")
//	=HYPERLINK("[Book1.xlsx]January!A10","Go to January > A10")
func (o *hyperlinkOption) ToTarget(target string) hyperlinkOption {
	return func(i *HyperlinkInfo) {
		if validator.IsURL(target) {
			log.Printf("url => %s", target)
			i.Set(Hyperlink.ToUrl(target))
		} else if validator.IsMailTo(target) {
			email, subject := "", ""
			log.Printf("mailto => %s", target)
			i.Set(Hyperlink.ToMail(email, subject))
		} else if validator.IsEmail(target) {
			log.Printf("email => %s", target)
			i.Set(Hyperlink.ToMail(target, ""))
		} else if validator.IsFilePath(target) {
			log.Printf("file => %s", target)
			i.Set(Hyperlink.ToFile(target))
		}
	}
}
//
//func (o *hyperlinkOption) ToLocation(location string) hyperlinkOption {
//	return func(i *HyperlinkInfo) {
//
//	}
//}

//private method used by hyperlinks manager to unpack HyperlinkInfo
func fromHyperlinkInfo(info *HyperlinkInfo) (hyperlink *ml.Hyperlink, styleID format.DirectStyleID, err error) {
	if err = info.Validate(); err != nil {
		return
	}

	styleID = info.styleID
	hyperlink = info.hyperlink
	return
}

//private method used by hyperlinks manager to pack HyperlinkInfo
func toHyperlinkInfo(hyperlink *ml.Hyperlink, targetInfo string, styleID format.DirectStyleID) *HyperlinkInfo {
	return NewHyperlink(
		Hyperlink.Formatting(styleID),
		Hyperlink.ToTarget(targetInfo+hyperlink.Location),
	)
}

func escapeLocation(location string) string {
	// TODO: escape location (research what kind of escaping Excel is expecting)
	return `'` + strings.Replace(location, `'`, `\'`, -1) + `'`
}