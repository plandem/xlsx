// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/index"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/number_format"

	// to link unexported
	_ "unsafe"
)

//
////go:linkname fromStyleFormat github.com/plandem/xlsx/format/styles.from
//func fromStyleFormat(f *styles.Info) (font *ml.Font, fill *ml.Fill, alignment *ml.CellAlignment, numFormat *ml.NumberFormat, protection *ml.CellProtection, border *ml.Border, namedInfo *ml.NamedStyleInfo)

//styleSheet is a higher level object that wraps ml.StyleSheet with functionality
type styleSheet struct {
	ml ml.StyleSheet

	//hash -> index for styles
	directStyleIndex index.Index
	diffStyleIndex   index.Index
	namedStyleIndex  index.Index

	//hash -> index for types
	borderIndex index.Index
	fillIndex   index.Index
	fontIndex   index.Index
	numberIndex index.Index

	//hash for typed number formats
	typedStyles map[number.Type]styles.DirectStyleID

	doc  *Spreadsheet
	file *ooxml.PackageFile
}

func newStyleSheet(f interface{}, doc *Spreadsheet) *styleSheet {
	ss := &styleSheet{
		doc:         doc,
		typedStyles: make(map[number.Type]styles.DirectStyleID),
	}

	ss.file = ooxml.NewPackageFile(doc.pkg, f, &ss.ml, nil)

	if ss.file.IsNew() {
		ss.doc.pkg.ContentTypes().RegisterContent(ss.file.FileName(), internal.ContentTypeStyles)
		ss.doc.relationships.AddFile(internal.RelationTypeStyles, ss.file.FileName())
		ss.file.MarkAsUpdated()
		ss.addDefaults()
		ss.buildIndexes()
	}

	return ss
}

//adds a default items for new created xlsx
func (ss *styleSheet) addDefaults() {
	//TODO: research more about default items for a new XLSX
	//..

	//add default types
	ss.ml.Fills.Items = append(ss.ml.Fills.Items,
		&ml.Fill{
			Pattern: &ml.PatternFill{
				Type: styles.PatternTypeNone,
			},
		},
		&ml.Fill{
			Pattern: &ml.PatternFill{
				Type: styles.PatternTypeGray125,
			},
		},
	)

	ss.ml.Borders.Items = append(ss.ml.Borders.Items, &ml.Border{
		Left:   &ml.BorderSegment{},
		Right:  &ml.BorderSegment{},
		Top:    &ml.BorderSegment{},
		Bottom: &ml.BorderSegment{},
	})

	ss.ml.Fonts.Items = append(ss.ml.Fonts.Items, &ml.Font{
		Family: styles.FontFamilySwiss,
		Scheme: styles.FontSchemeMinor,
		Name:   "Calibri",
		Size:   11.0,
		//Color: ml.Color{Theme: 1}
	})

	//add default ref for CellStyleXfs
	ss.ml.CellStyleXfs.Items = append(ss.ml.CellStyleXfs.Items, &ml.NamedStyle{
		FontId:   0,
		FillId:   0,
		BorderId: 0,
		NumFmtId: 0,
	})

	//add default ref for CellXfs
	ss.ml.CellXfs.Items = append(ss.ml.CellXfs.Items, &ml.DirectStyle{
		XfId: ml.NamedStyleID(0),
		Style: ml.Style{
			FontId:   0,
			FillId:   0,
			BorderId: 0,
			NumFmtId: 0,
		},
	})

	//add default ref for CellStyles
	index := 0
	ss.ml.CellStyles.Items = append(ss.ml.CellStyles.Items, &ml.NamedStyleInfo{
		Name:      "Normal",
		XfId:      ml.NamedStyleID(0),
		BuiltinId: &index,
	})

	/*
		TODO: replace hardcoded defaults with format
		def := format.New(
			format.NamedStyle(format.NamedStyleNormal),
			format.Font.Default,
		)
	*/
}

//build indexes for all indexes
func (ss *styleSheet) buildIndexes() {
	//build indexes for fonts
	for id, f := range ss.ml.Fonts.Items {
		_ = ss.fontIndex.Add(f, id)
	}

	//build indexes for fill
	for id, f := range ss.ml.Fills.Items {
		_ = ss.fillIndex.Add(f, id)
	}

	//build indexes for border
	for id, f := range ss.ml.Borders.Items {
		_ = ss.borderIndex.Add(f, id)
	}

	//build indexes for number formats
	for _, f := range ss.ml.NumberFormats.Items {
		//N.B.: NumberFormat uses ID, not indexes
		_ = ss.numberIndex.Add(f, f.ID)
	}

	//build indexes for named styles
	for id, xf := range ss.ml.CellStyleXfs.Items {
		_ = ss.namedStyleIndex.Add(xf, id)
	}

	//build indexes for direct styles
	for id, xf := range ss.ml.CellXfs.Items {
		_ = ss.directStyleIndex.Add(xf, id)
	}

	//build indexes for differential styles
	for id, dxf := range ss.ml.Dxfs.Items {
		_ = ss.diffStyleIndex.Add(dxf, id)
	}
}

//adds a number formats for each type of number format if required. These styles will be used by cell's typed SetXXX methods
func (ss *styleSheet) addTypedStylesIfRequired() {
	if len(ss.typedStyles) == 0 {
		for _, t := range []number.Type{
			number.General,
			number.Integer,
			number.Float,
			number.Date,
			number.Time,
			number.DateTime,
			number.DeltaTime,
		} {
			id, _ := number.Default(t)
			ss.typedStyles[t] = ss.addStyle(styles.New(styles.NumberFormatID(id)))
		}

		ss.file.MarkAsUpdated()
	}
}

//resolveNumberFormat returns resolved NumberFormat code for styleID
func (ss *styleSheet) resolveNumberFormat(id ml.DirectStyleID) string {
	style := ss.ml.CellXfs.Items[id]

	//return code for built-in number format
	if n := number.Normalize(ml.NumberFormat{ID: style.NumFmtId}); len(n.Code) > 0 {
		return n.Code
	}

	//try to lookup through custom formats and find same ID
	for _, f := range ss.ml.NumberFormats.Items {
		if style.NumFmtId == f.ID {
			return f.Code
		}
	}

	//N.B.: wtf is going on?! non built-in and not existing id?
	_, code := number.Default(number.General)
	return code
}

//resolveDirectStyle returns resolved Info for DirectStyleID
func (ss *styleSheet) resolveDirectStyle(id ml.DirectStyleID) *styles.Info {
	if id == 0 {
		return nil
	}

	panic(errorNotSupported)

	//cellStyle := ss.ml.CellXfs.Items[id]
	//style := &styles.Info{}
	//_ = cellStyle

	//TODO: Populate format.Info with required information
	//return style
}

//adds a differential style
func (ss *styleSheet) addDiffStyle(f *styles.Info) styles.DiffStyleID {
	ss.file.LoadIfRequired(ss.buildIndexes)

	//get settings for style
	font, fill, alignment, numFormat, protection, border, _ := styles.From(f)

	dXf := &ml.DiffStyle{
		Font:         font,
		Fill:         fill,
		Border:       border,
		NumberFormat: numFormat,
		Alignment:    alignment,
		Protection:   protection,
	}

	//return id of already existing information
	if id, ok := ss.diffStyleIndex.Get(dXf); ok {
		return styles.DiffStyleID(id)
	}

	//add a new one and return related id
	nextID := styles.DiffStyleID(len(ss.ml.Dxfs.Items))
	ss.ml.Dxfs.Items = append(ss.ml.Dxfs.Items, dXf)
	_ = ss.diffStyleIndex.Add(dXf, int(nextID))
	ss.file.MarkAsUpdated()
	return nextID
}

//add a named style if required
func (ss *styleSheet) addNamedStyleIfRequired(namedInfo *ml.NamedStyleInfo, style ml.Style) ml.NamedStyleID {
	if namedInfo == nil {
		return 0
	}

	namedStyle := ml.NamedStyle(style)

	//TODO: check if it's possible to have 2 same built-styles

	//if there is already same styles, then use it
	if id, ok := ss.namedStyleIndex.Get(&namedStyle); ok {
		namedInfo.XfId = ml.NamedStyleID(id)
	} else {
		//add a new style
		nextID := styles.NamedStyleID(len(ss.ml.CellStyleXfs.Items))
		ss.ml.CellStyleXfs.Items = append(ss.ml.CellStyleXfs.Items, &namedStyle)
		_ = ss.namedStyleIndex.Add(&namedStyle, int(nextID))

		//add style info
		namedInfo.XfId = ml.NamedStyleID(nextID)
		ss.ml.CellStyles.Items = append(ss.ml.CellStyles.Items, namedInfo)
	}

	//add named info
	ss.file.MarkAsUpdated()
	return namedInfo.XfId
}

//adds a style. Style can be Direct or Named. Depends on settings.
func (ss *styleSheet) addStyle(f *styles.Info) styles.DirectStyleID {
	ss.file.LoadIfRequired(ss.buildIndexes)

	//get settings and add information if required
	font, fill, alignment, numFormat, protection, border, namedInfo := styles.From(f)
	fontID := ss.addFontIfRequired(font)
	fillID := ss.addFillIfRequired(fill)
	borderID := ss.addBorderIfRequired(border)
	numID := ss.addNumFormatIfRequired(numFormat)

	/*
		Note to remember excel internals:
		---
		cell.s = cellXfs.index  => DirectStyleID
		cellXfs.xfId = cellStyleXf.index => NamedStyleID
		cellStyle.xfId = cellStyleXf.index => NamedStyleID
	*/

	style := ml.Style{
		FontId:            fontID,
		FillId:            fillID,
		BorderId:          borderID,
		NumFmtId:          numID,
		Alignment:         alignment,
		Protection:        protection,
		ApplyFont:         fontID > 0,
		ApplyBorder:       borderID > 0,
		ApplyFill:         fillID > 0,
		ApplyNumberFormat: numID > 0,
		ApplyAlignment:    alignment != nil,
		ApplyProtection:   protection != nil,
	}

	//add named style if required and get related xfid
	xfid := ss.addNamedStyleIfRequired(namedInfo, style)

	cellXf := &ml.DirectStyle{
		XfId:  xfid,
		Style: style,
	}

	//return id of already existing information
	if id, ok := ss.directStyleIndex.Get(cellXf); ok {
		return styles.DirectStyleID(id)
	}

	//add a new one and return related id
	nextID := styles.DirectStyleID(len(ss.ml.CellXfs.Items))
	ss.ml.CellXfs.Items = append(ss.ml.CellXfs.Items, cellXf)
	_ = ss.directStyleIndex.Add(cellXf, int(nextID))
	ss.file.MarkAsUpdated()
	return nextID
}

//adds a new font if required
func (ss *styleSheet) addFontIfRequired(font *ml.Font) int {
	//if there is no information, then use default
	if font == nil {
		return 0
	}

	//return id of already existing information
	if id, ok := ss.fontIndex.Get(font); ok {
		return id
	}

	//add a new one and return related id
	nextID := len(ss.ml.Fonts.Items)
	ss.ml.Fonts.Items = append(ss.ml.Fonts.Items, font)
	_ = ss.fontIndex.Add(font, nextID)
	ss.file.MarkAsUpdated()
	return nextID
}

//adds a new fill if required
func (ss *styleSheet) addFillIfRequired(fill *ml.Fill) int {
	//if there is no information, then use default
	if fill == nil {
		return 0
	}

	//return id of already existing information
	if id, ok := ss.fillIndex.Get(fill); ok {
		return id
	}

	//add a new one and return related id
	nextID := len(ss.ml.Fills.Items)
	ss.ml.Fills.Items = append(ss.ml.Fills.Items, fill)
	_ = ss.fillIndex.Add(fill, nextID)
	ss.file.MarkAsUpdated()
	return nextID
}

//adds a new border if required
func (ss *styleSheet) addBorderIfRequired(border *ml.Border) int {
	//if there is no information, then use default
	if border == nil {
		return 0
	}

	//return id of already existing information
	if id, ok := ss.borderIndex.Get(border); ok {
		return id
	}

	//add a new one and return related id
	nextID := len(ss.ml.Borders.Items)
	ss.ml.Borders.Items = append(ss.ml.Borders.Items, border)
	_ = ss.borderIndex.Add(border, nextID)
	ss.file.MarkAsUpdated()
	return nextID
}

//adds a new number format if required
func (ss *styleSheet) addNumFormatIfRequired(n *ml.NumberFormat) int {
	//if there is no information, then use default
	if n == nil {
		return 0
	}

	//if is built-in format then return id
	if number.IsBuiltIn(n.ID) {
		return n.ID
	}

	//Return id of already existing information.
	//N.B.: Supposed that for custom format we have -1 as code, so hash should be same for new/existing custom format
	if id, ok := ss.numberIndex.Get(n); ok {
		return id
	}

	//try to lookup through custom formats and find same code
	for _, f := range ss.ml.NumberFormats.Items {
		if n.Code == f.Code {
			return f.ID
		}
	}

	//N.B.: NumberFormat uses ID, not indexes
	nextID := number.LastReservedID + len(ss.ml.NumberFormats.Items) + 1
	n.ID = nextID

	ss.ml.NumberFormats.Items = append(ss.ml.NumberFormats.Items, n)
	_ = ss.numberIndex.Add(n, nextID)
	ss.file.MarkAsUpdated()
	return nextID
}
