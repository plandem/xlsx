package xlsx

import (
	"errors"
	"fmt"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
	_ "unsafe"
)

//go:linkname fromHyperlinkInfo github.com/plandem/xlsx/types.fromHyperlinkInfo
func fromHyperlinkInfo(info *types.HyperlinkInfo) (hyperlink *ml.Hyperlink, styleID format.DirectStyleID, err error)

//go:linkname toHyperlinkInfo github.com/plandem/xlsx/types.toHyperlinkInfo
func toHyperlinkInfo(hyperlink *ml.Hyperlink, targetInfo string, styleID format.DirectStyleID) *types.HyperlinkInfo

type hyperlinks struct {
	sheet          *sheetInfo
	defaultStyleID format.DirectStyleID
}

//newHyperlinks creates an object that implements hyperlinks functionality
func newHyperlinks(sheet *sheetInfo) *hyperlinks {
	return &hyperlinks{sheet: sheet, defaultStyleID: -1}
}

func (h *hyperlinks) initIfRequired() {
	//attach hyperlinks object if required
	if h.sheet.ml.Hyperlinks == nil {
		var links []*ml.Hyperlink
		h.sheet.ml.Hyperlinks = &links
	}
}

//Add adds a new hyperlink info for provided bounds, where link can be string or HyperlinkInfo
func (h *hyperlinks) Add(bounds types.Bounds, link interface{}) (format.DirectStyleID, error) {
	h.initIfRequired()

	//check if hyperlink has style and if not, then add default
	if h.defaultStyleID == -1 {
		//we need to add default named style for hyperlink
		defaultStyleID := h.sheet.workbook.doc.AddFormatting(format.New(
			format.NamedStyle(format.NamedStyleHyperlink),
			format.Font.Default,
			format.Font.Underline(format.UnderlineTypeSingle),
			format.Font.Color("#0563C1"),
		))

		h.defaultStyleID = defaultStyleID
	}

	//resolve HyperlinkInfo if required
	var object *types.HyperlinkInfo
	if target, ok := link.(string); ok {
		object = types.NewHyperlink(types.Hyperlink.ToTarget(target))
	} else if pointer, ok := link.(*types.HyperlinkInfo); ok {
		object = pointer
	} else if value, ok := link.(types.HyperlinkInfo); ok {
		object = &value
	} else {
		return format.DefaultDirectStyle, errors.New("unsupported type of hyperlink, only string or types.HyperlinkInfo is allowed")
	}

	//let's check existing hyperlinks for overlapping bounds
	hyperlinkIndex := -1
	for linkIndex, link := range *h.sheet.ml.Hyperlinks {
		if link.Bounds.Equals(bounds) {
			hyperlinkIndex = linkIndex
		} else if link.Bounds.Overlaps(bounds) {
			return format.DefaultDirectStyle, errors.New(fmt.Sprintf("intersection of different hyperlinks is not allowed, %s intersects with %s", link.Bounds, bounds))
		}
	}

	//prepare hyperlink info
	hyperlink, styleID, err := fromHyperlinkInfo(object)
	if err != nil {
		return format.DefaultDirectStyle, err
	}

	//exceeded Excel limit for total hyperlinks
	if len(*h.sheet.ml.Hyperlinks) >= internal.ExcelHyperlinkLimit {
		return format.DefaultDirectStyle, errors.New(fmt.Sprintf("exceeds Excel limit (%d) for total number of hyperlinks per worksheet", internal.ExcelHyperlinkLimit))
	}

	//if link has external target, then add relation for it
	if len(hyperlink.RID) > 0 {
		h.sheet.attachRelationshipsIfRequired()

		//lookup for already existing targets to get RID
		rid := h.sheet.relationships.GetIdByTarget(string(hyperlink.RID))

		//looks like target is new, let's create it and use
		if rid = h.sheet.relationships.GetIdByTarget(string(hyperlink.RID)); len(rid) == 0 {
			_, rid = h.sheet.relationships.AddLink(internal.RelationTypeHyperlink, string(hyperlink.RID))
		}

		hyperlink.RID = rid
	}

	//add source Ref info
	hyperlink.Bounds = bounds
	if hyperlinkIndex == -1 {
		//add a new hyperlink
		*h.sheet.ml.Hyperlinks = append(*h.sheet.ml.Hyperlinks, hyperlink)
	} else {
		//update existing hyperlink
		(*h.sheet.ml.Hyperlinks)[hyperlinkIndex] = hyperlink
	}

	//if there are custom styles, then use it otherwise use default hyperlink styles
	if styleID == format.DefaultDirectStyle {
		styleID = h.defaultStyleID
	}

	return styleID, nil
}

//Get returns a resolved hyperlink info for provided ref or nil if there is no any hyperlink
func (h *hyperlinks) Get(ref types.CellRef) *types.HyperlinkInfo {
	h.initIfRequired()

	links := *h.sheet.ml.Hyperlinks
	if len(links) > 0 {
		cIdx, rIdx := ref.ToIndexes()
		for _, link := range links {
			if link.Bounds.Contains(cIdx, rIdx) {
				cell := h.sheet.sheet.CellByRef(ref)
				styleID := cell.ml.Style
				return toHyperlinkInfo(link, h.sheet.relationships.GetTargetById(string(link.RID)), styleID)
			}
		}
	}

	return nil
}

//Remove removes hyperlink info for bounds
func (h *hyperlinks) Remove(bounds types.Bounds) {
	h.initIfRequired()

	links := *h.sheet.ml.Hyperlinks
	if len(links) > 0 {
		newLinks := make([]*ml.Hyperlink, 0, len(links))
		for _, link := range links {
			if !link.Bounds.Overlaps(bounds) {
				//copy only non overlapping bounds
				newLinks = append(newLinks, link)
			}
		}

		h.sheet.ml.Hyperlinks = &newLinks
	}
}

func (h *hyperlinks) pack() *[]*ml.Hyperlink {
	//hyperlinks must have at least one object
	if h.sheet.ml.Hyperlinks != nil && len(*h.sheet.ml.Hyperlinks) == 0 {
		h.sheet.ml.Hyperlinks = nil
	}

	return h.sheet.ml.Hyperlinks
}
