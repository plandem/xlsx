// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"github.com/plandem/xlsx/format/conditional"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
	"strings"

	// to link unexported
	_ "unsafe"
)

//go:linkname fromConditionalFormat github.com/plandem/xlsx/format/conditional.from
func fromConditionalFormat(f *conditional.Info) (*ml.ConditionalFormatting, []*styles.Info)

type conditionals struct {
	sheet *sheetInfo
}

//newConditionals creates an object that implements conditional formatting functionality
func newConditionals(sheet *sheetInfo) *conditionals {
	return &conditionals{sheet: sheet}
}

func (c *conditionals) initIfRequired() {
	//attach conditionals if required
	if c.sheet.ml.ConditionalFormatting == nil {
		var conditionals []*ml.ConditionalFormatting
		c.sheet.ml.ConditionalFormatting = &conditionals
	}
}

//Add adds a conditional formatting with attaching additional refs if required
func (c *conditionals) Add(ci *conditional.Info, refs []types.Ref) error {
	c.initIfRequired()

	//attach additional refs, if required
	if len(refs) > 0 {
		ci.Set(conditional.Refs(refs...))
	}

	if err := ci.Validate(); err != nil {
		return err
	}

	info, formats := fromConditionalFormat(ci)
	if info != nil {
		var startCol, startRow int

		//some rules require starting CellRef in formula
		for i, b := range info.Bounds {
			if i == 0 {
				startCol, startRow = b.FromCol, b.FromRow
			} else if b.FromCol < startCol {
				startCol = b.FromCol
			} else if b.FromRow < startRow {
				startRow = b.FromRow
			}
		}

		startCellRef := types.CellRefFromIndexes(startCol, startRow)

		for i, styleInfo := range formats {
			if styleInfo != nil {
				//add a new diff styles
				styleID := c.sheet.workbook.doc.styleSheet.addDiffStyle(styleInfo)
				info.Rules[i].Style = &styleID
			}

			//finalize rules
			for _, ruleInfo := range info.Rules {
				for i, formula := range ruleInfo.Formula {
					ruleInfo.Formula[i] = ml.Formula(strings.ReplaceAll(string(formula), ":cell:", string(startCellRef)))
				}
			}

			//add a new conditional
			*c.sheet.ml.ConditionalFormatting = append(*c.sheet.ml.ConditionalFormatting, info)
		}
	}

	return nil
}

//Remove deletes a conditional formatting from refs
func (c *conditionals) Remove(refs []types.Ref) {
	panic(errorNotSupported)
}

//Resolve checks if requested cIdx and rIdx related to any conditionals formatting and returns it
func (c *conditionals) Resolve(cIdx, rIdx int) *conditional.Info {
	//TODO: Populate format.Info with required information
	panic(errorNotSupported)
}

func (c *conditionals) pack() {
	//conditionals must have at least one object
	if c.sheet.ml.ConditionalFormatting != nil && len(*c.sheet.ml.ConditionalFormatting) == 0 {
		c.sheet.ml.ConditionalFormatting = nil
	}
}
