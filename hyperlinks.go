package xlsx

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
)

//TODO: implement remove/refresh functionality
type hyperlinkManager struct {
	sheet *sheetInfo
	defaultStyleID format.StyleID
}

//newHyperlinkManager creates an object that implements hyperlink cells functionality
func newHyperlinkManager(sheet *sheetInfo) *hyperlinkManager {
	//we need to add default named style for hyperlink
	defaultStyleID := sheet.workbook.doc.AddFormatting(format.New(
		format.NamedStyle(format.NamedStyleHyperlink),
		format.Font.Default,
		format.Font.Underline(format.UnderlineTypeSingle),
		format.Font.Color("#0563C1"),
	))

	//attach hyperlinks object if required
	if sheet.ml.Hyperlinks == nil {
		var links []*ml.Hyperlink
		sheet.ml.Hyperlinks = &links
	}

	return &hyperlinkManager{
		sheet: sheet,
		defaultStyleID: defaultStyleID,
	}
}

//if there is a hyperlink for provided ref, then return it.
func (m *hyperlinkManager) Get(ref types.CellRef) int {
	panic(errorNotSupported)

	if m.sheet.ml.Hyperlinks != nil {
		hyperlinks := *m.sheet.ml.Hyperlinks
		if len(hyperlinks) > 0 {
			for _, mc := range hyperlinks {
				if mc.Bounds.ContainsRef(ref) {
					//TODO: create hyperlink, populate with related data and return
					break
				}
			}
		}
	}

	return 0
}

//Bounds   types.Bounds `xml:"ref,attr"`
//Location string       `xml:"location,attr,omitempty"`
//Tooltip  string       `xml:"tooltip,attr,omitempty"`
//Display  string       `xml:"display,attr,omitempty"`
//RID      ml.RID       `xml:"id,attr,omitempty"`
func (m *hyperlinkManager) Add(ref types.Ref, link interface{}) (format.StyleID, error) {
	//var (
	//location string
	//tooltip  string
	//display  string
	//style *format.StyleFormat
	//)

	//if location, ok := link.(string); ok {
	//	simple version provided - location only
	//	_ = location
	//} /*else if hlink, ok := link.(types.Hyperlink); ok {
	// advanced version provided
	//_ = hlink
	//} else {
	//	panic("a")
	//}

	//external hyperlinks
	url := link.(string)
	m.sheet.attachRelationshipsIfRequired()
	_, rid := m.sheet.relationships.AddLink(internal.RelationTypeHyperlink, url)
	_ = rid

	*m.sheet.ml.Hyperlinks = append(*m.sheet.ml.Hyperlinks, &ml.Hyperlink{
		Bounds:   ref.ToBounds(),
		RID:      rid,
	})

	return m.defaultStyleID, nil
}

func (m *hyperlinkManager) Remove(ref types.CellRef) {

}
