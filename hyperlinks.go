package xlsx

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/helpers"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
)

type hyperlinkManager struct {
	sheet          *sheetInfo
	defaultStyleID format.DirectStyleID
}

func newHyperlinkManager(sheet *sheetInfo) *hyperlinkManager {
	//attach hyperlinks object if required
	if sheet.ml.Hyperlinks == nil {
		var links []*ml.Hyperlink
		sheet.ml.Hyperlinks = &links
	}

	return &hyperlinkManager{sheet: sheet, defaultStyleID: -1}
}

func (m *hyperlinkManager) Add(ref types.Ref, link interface{}) (format.DirectStyleID, error) {
	//check if hyperlink has style and if not, then add default
	if m.defaultStyleID == -1 {
		//we need to add default named style for hyperlink
		defaultStyleID := m.sheet.workbook.doc.AddFormatting(format.New(
			format.NamedStyle(format.NamedStyleHyperlink),
			format.Font.Default,
			format.Font.Underline(format.UnderlineTypeSingle),
			format.Font.Color("#0563C1"),
		))

		m.defaultStyleID = defaultStyleID
	}

	//resolve HyperlinkInfo if required
	var object *types.HyperlinkInfo
	if url, ok := link.(string); ok {
		object = types.NewHyperlink(types.Hyperlink.ToUrl(url))
	} else if pointer, ok := link.(*types.HyperlinkInfo); ok {
		object = pointer
	} else if value, ok := link.(types.HyperlinkInfo); ok {
		object = &value
	} else {
		panic("unsupported type of hyperlink, only string or types.HyperlinkInfo is allowed")
	}

	//prepare hyperlink info
	hyperlink, styleID, err := helpers.FromHyperlink(object)
	if err != nil {
		return 0, err
	}

	//if link has external target, then add relation for it
	if len(hyperlink.RID) > 0 {
		m.sheet.attachRelationshipsIfRequired()

		//lookup for already existing targets to get RID
		rid := m.sheet.relationships.GetIdByTarget(string(hyperlink.RID))

		//looks like target is new, let's create it and use
		if rid = m.sheet.relationships.GetIdByTarget(string(hyperlink.RID)); len(rid) == 0 {
			_, rid = m.sheet.relationships.AddLink(internal.RelationTypeHyperlink, string(hyperlink.RID))
		}

		hyperlink.RID = rid
	}

	//add source Ref info
	hyperlink.Bounds = ref.ToBounds()
	*m.sheet.ml.Hyperlinks = append(*m.sheet.ml.Hyperlinks, hyperlink)

	//if there are custom styles, then use it otherwise use default hyperlink styles
	if styleID == 0 {
		styleID = m.defaultStyleID
	}

	return styleID, nil
}

//if there is a hyperlink for provided ref, then return it.
func (m *hyperlinkManager) Get(ref types.CellRef) *types.HyperlinkInfo {
	panic(errorNotSupported)
}

func (m *hyperlinkManager) Remove(ref types.CellRef) {
	panic(errorNotSupported)
}
