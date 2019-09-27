// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"errors"
	"fmt"
	sharedML "github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
	"github.com/plandem/xlsx/types/hyperlink"

	// to link unexported
	_ "unsafe"
)

//
////go:linkname fromHyperlinkInfo github.com/plandem/xlsx/types/hyperlink.from
//func fromHyperlinkInfo(info *hyperlink.Info) (hyperlink *ml.Hyperlink, format interface{}, err error)
//
////go:linkname toHyperlinkInfo github.com/plandem/xlsx/types/hyperlink.to
//func toHyperlinkInfo(hyperlink *ml.Hyperlink, targetInfo string, styleID styles.DirectStyleID) *hyperlink.Info

type hyperlinks struct {
	sheet          *sheetInfo
	defaultStyleID styles.DirectStyleID
}

//newHyperlinks creates an object that implements hyperlinks functionality
func newHyperlinks(sheet *sheetInfo) *hyperlinks {
	return &hyperlinks{sheet: sheet, defaultStyleID: -1}
}

//Add adds a new hyperlink info for provided bounds, where link can be string or Info
func (h *hyperlinks) Add(bounds types.Bounds, link interface{}) (interface{}, error) {
	//check if hyperlink has style and if not, then add default
	if h.defaultStyleID == -1 {
		//we need to add default named style for hyperlink
		defaultStyleID := h.sheet.workbook.doc.AddStyles(styles.New(
			styles.NamedStyle(styles.NamedStyleHyperlink),
			styles.Font.Default,
			styles.Font.Underline(styles.UnderlineTypeSingle),
			styles.Font.Color("#0563C1"),
		))

		h.defaultStyleID = defaultStyleID
	}

	//resolve Info if required
	var object *hyperlink.Info
	if target, ok := link.(string); ok {
		object = hyperlink.New(hyperlink.ToTarget(target))
	} else if pointer, ok := link.(*hyperlink.Info); ok {
		object = pointer
	} else if value, ok := link.(hyperlink.Info); ok {
		object = &value
	} else {
		return styles.DefaultDirectStyle, errors.New("unsupported type of hyperlink, only string or hyperlink.Info is allowed")
	}

	//let's check existing hyperlinks for overlapping bounds
	hyperlinkIndex := -1
	for linkIndex, link := range h.sheet.ml.Hyperlinks.Items {
		if link.Bounds.Equals(bounds) {
			hyperlinkIndex = linkIndex
		} else if link.Bounds.Overlaps(bounds) {
			return styles.DefaultDirectStyle, fmt.Errorf("intersection of different hyperlinks is not allowed, %s intersects with %s", link.Bounds, bounds)
		}
	}

	//prepare hyperlink info
	hLink, format, err := hyperlink.From(object)
	if err != nil {
		return styles.DefaultDirectStyle, err
	}

	//exceeded Excel limit for total hyperlinks
	if len(h.sheet.ml.Hyperlinks.Items) >= internal.ExcelHyperlinkLimit {
		return styles.DefaultDirectStyle, fmt.Errorf("exceeds Excel limit (%d) for total number of hyperlinks per worksheet", internal.ExcelHyperlinkLimit)
	}

	//if link has external target, then add relation for it
	if len(hLink.RID) > 0 {
		h.sheet.attachRelationshipsIfRequired()

		var rid sharedML.RID

		//lookup for already existing targets to get RID
		if rid = h.sheet.relationships.GetIdByTarget(string(hLink.RID)); len(rid) == 0 {
			//looks like target is new, let's create it and use
			_, rid = h.sheet.relationships.AddLink(internal.RelationTypeHyperlink, string(hLink.RID))
		}

		hLink.RID = rid
	}

	//add source Ref info
	hLink.Bounds = bounds
	if hyperlinkIndex == -1 {
		//add a new hyperlink
		h.sheet.ml.Hyperlinks.Items = append(h.sheet.ml.Hyperlinks.Items, hLink)
	} else {
		//update existing hyperlink
		h.sheet.ml.Hyperlinks.Items[hyperlinkIndex] = hLink
	}

	//if there are custom styles, then use it otherwise use default hyperlink styles
	if format == nil {
		return h.defaultStyleID, nil
	}

	return format, nil
}

//Get returns a resolved hyperlink info for provided ref or nil if there is no any hyperlink
func (h *hyperlinks) Get(ref types.CellRef) *hyperlink.Info {
	links := h.sheet.ml.Hyperlinks.Items
	if len(links) > 0 {
		cIdx, rIdx := ref.ToIndexes()
		for _, link := range links {
			if link.Bounds.Contains(cIdx, rIdx) {
				cell := h.sheet.sheet.CellByRef(ref)
				styleID := cell.ml.Style
				return hyperlink.To(link, h.sheet.relationships.GetTargetById(string(link.RID)), styleID)
			}
		}
	}

	return nil
}

//Remove removes hyperlink info for bounds
func (h *hyperlinks) Remove(bounds types.Bounds) {
	if len(h.sheet.ml.Hyperlinks.Items) > 0 {
		newLinks := make([]*ml.Hyperlink, 0, len(h.sheet.ml.Hyperlinks.Items))

		for _, link := range h.sheet.ml.Hyperlinks.Items {
			if !link.Bounds.Overlaps(bounds) {
				//copy only non overlapping bounds
				newLinks = append(newLinks, link)
			}
		}

		h.sheet.ml.Hyperlinks.Items = newLinks
	}
}
