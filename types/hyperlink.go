package types

import (
	"errors"
	"fmt"
	sharedML "github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"net/url"
)

type HyperlinkInfo struct {
	hyperlink *ml.Hyperlink
	styleID   format.DirectStyleID
	linkType  hyperlinkType
}

type hyperlinkOption func(o *HyperlinkInfo)
type hyperlinkType byte

const (
	hyperlinkTypeUnknown = iota
	hyperlinkTypeWorkbook
	hyperlinkTypeWeb
	hyperlinkTypeEmail
	hyperlinkTypeFile
)

//Hyperlink is a 'namespace' for all possible settings for hyperlink
var Hyperlink hyperlinkOption

//NewHyperlink
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
	}

	return nil
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
		i.hyperlink.RID = sharedML.RID(fmt.Sprintf("mailto:%s?subject=%s", address, subject))
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

/*

../Budgets/Annual/Budget2010.xlsx
../Budgets/Annual/Budget2010.xlsx#'Sheet3'
../Budgets/Annual/Budget2010.xlsx#'Sheet3'!G43
../Budgets/Annual/Budget2010.xlsx#DeptTotals


- file
=HYPERLINK("D:\Word files\Price list.docx","Price list")

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
*/
func (o *hyperlinkOption) ToFile(fileName string) hyperlinkOption {
	return func(i *HyperlinkInfo) {
		//UNC path
		i.hyperlink.RID = sharedML.RID(fileName)

		//workbook can be internal(same) or external
		if i.linkType != hyperlinkTypeWorkbook {
			i.linkType = hyperlinkTypeFile
		}
	}
}

func (o *hyperlinkOption) ToSheet(sheetName string) hyperlinkOption {
	return func(i *HyperlinkInfo) {
		i.linkType = hyperlinkTypeWorkbook
		//i.hyperlink.Location = ""
	}
}

func (o *hyperlinkOption) ToRef(ref Ref) hyperlinkOption {
	return func(i *HyperlinkInfo) {
		i.linkType = hyperlinkTypeWorkbook
		//i.hyperlink.Location = ""
	}
}
