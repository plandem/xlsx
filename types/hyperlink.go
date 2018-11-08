package types

import (
	"errors"
	"fmt"
	sharedML "github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
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
	/*
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

	switch i.linkType {
	case hyperlinkTypeUnknown:
		return errors.New("unknown type of hyperlink")
	case hyperlinkTypeWorkbook:
		if len(i.hyperlink.Location) == 0 {
			return errors.New("unknown location at target workbook")
		}
	case hyperlinkTypeWeb:
		if len(i.hyperlink.RID) > internal.ExcelUrlLimit {
			return errors.New(fmt.Sprintf("url exceeded maximum allowed length (%d chars)", internal.ExcelUrlLimit))
		}
	case hyperlinkTypeEmail:
		if len(i.hyperlink.RID) > internal.ExcelUrlLimit {
			return errors.New(fmt.Sprintf("email exceeded maximum allowed length (%d chars)", internal.ExcelUrlLimit))
		}
	case hyperlinkTypeFile:
		if len(i.hyperlink.RID) > internal.ExcelUrlLimit {
			return errors.New(fmt.Sprintf("link to file exceeded maximum allowed length (%d chars)", internal.ExcelUrlLimit))
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
				// TODO: escape sheetName (research what kind of escaping Excel is expecting)
				sheetName = strings.Replace(sheetName, `'`, `\'`, -1)
				i.hyperlink.Location = fmt.Sprintf("#'%s'!%s", sheetName, ref)
			} else {
				//ref only
				i.hyperlink.Location = fmt.Sprintf("#%s", ref)
			}
		}
	}
}

/*

../Budgets/Annual/Budget2010.xlsx
../Budgets/Annual/Budget2010.xlsx#'Sheet3'
../Budgets/Annual/Budget2010.xlsx#'Sheet3'!G43
../Budgets/Annual/Budget2010.xlsx#DeptTotals
*/
func (o *hyperlinkOption) ToTarget(target string) hyperlinkOption {
	return func(i *HyperlinkInfo) {
		if u, err := url.Parse(target); err == nil {
			//if u.Fragment != "" {
			//	i.hyperlink.Location = u.Fragment
			//	u.Fragment = ""
			//}
			//
			//i.hyperlink.RID = sharedML.RID(u.String())
			log.Printf("%+v, %+v", u, err)
			i.linkType = hyperlinkTypeWeb
		}
	}
}

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
func toHyperlinkInfo(hyperlink *ml.Hyperlink, targetInfo string, styleID format.DirectStyleID) (info *HyperlinkInfo) {
	info = NewHyperlink(
		Hyperlink.Formatting(styleID),
	)

	//TODO: create a HyperlinkInfo with resolved internal types, i.e. linkType
	//info.hyperlink = hyperlinkInfo
	return
}
