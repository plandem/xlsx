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
}

//newHyperlinkManager creates an object that implements hyperlink cells functionality
func newHyperlinkManager(sheet *sheetInfo) *hyperlinkManager {
	return &hyperlinkManager{
		sheet: sheet,
	}
}

//Bounds   types.Bounds `xml:"ref,attr"`
//Location string       `xml:"location,attr,omitempty"`
//Tooltip  string       `xml:"tooltip,attr,omitempty"`
//Display  string       `xml:"display,attr,omitempty"`
//RID      ml.RID       `xml:"id,attr,omitempty"`
func (hm *hyperlinkManager) add(ref types.Ref, link interface{}) (format.StyleID, error) {
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
	hm.sheet.attachRelationshipsIfRequired()
	_, rid := hm.sheet.relationships.AddLink(internal.RelationTypeHyperlink, url)

	if hm.sheet.ml.Hyperlinks == nil {
		var links []*ml.Hyperlink
		hm.sheet.ml.Hyperlinks = &links
		_ = rid
	}

	//*hm.sheet.ml.Hyperlinks = append(*hm.sheet.ml.Hyperlinks, &ml.Hyperlink{
	//	Bounds:   ref.ToBounds(),
	//	RID:      rid,
	//})

	return 0, nil
}
