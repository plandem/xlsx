package xlsx

import (
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/hash"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/number_format"
)

//StyleSheet is a higher level object that wraps ml.StyleSheet with functionality
type StyleSheet struct {
	ml ml.StyleSheet

	//hash -> index for styles
	styleIndex      map[string]format.StyleID
	diffStyleIndex  map[string]format.DiffStyleID
	namedStyleIndex map[string]format.NamedStyleID

	//hash -> index for types
	borderIndex map[string]int
	fillIndex   map[string]int
	fontIndex   map[string]int
	numberIndex map[string]int

	//hash for typed number formats
	typedStyles map[numberFormat.Type]format.StyleID

	doc  *Spreadsheet
	file *ooxml.PackageFile
}

func newStyleSheet(f interface{}, doc *Spreadsheet) *StyleSheet {
	ss := &StyleSheet{
		doc:             doc,
		styleIndex:      make(map[string]format.StyleID),
		diffStyleIndex:  make(map[string]format.DiffStyleID),
		namedStyleIndex: make(map[string]format.NamedStyleID),
		borderIndex:     make(map[string]int),
		fillIndex:       make(map[string]int),
		fontIndex:       make(map[string]int),
		numberIndex:     make(map[string]int),
		typedStyles:     make(map[numberFormat.Type]format.StyleID),
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
func (ss *StyleSheet) addDefaults() {
	//TODO: research more about default items for a new XLSX
	//..

	//add default types
	ss.ml.Fills = &[]*ml.Fill{
		{
			Pattern: &ml.PatternFill{
				Type: format.PatternTypeNone,
			},
		},
		{
			Pattern: &ml.PatternFill{
				Type: format.PatternTypeGray125,
			},
		},
	}

	ss.ml.Borders = &[]*ml.Border{{
		Left:   &ml.BorderSegment{},
		Right:  &ml.BorderSegment{},
		Top:    &ml.BorderSegment{},
		Bottom: &ml.BorderSegment{},
	}}

	ss.ml.Fonts = &[]*ml.Font{{
		Family: format.FontFamilySwiss,
		Scheme: format.FontSchemeMinor,
		Name:   "Calibri",
		Size:   11.0,
		//Color: ml.Color{Theme: 1}
	}}

	//add default ref for CellStyleXfs
	ss.ml.CellStyleXfs = &[]*ml.Style{{
		FontId:   0,
		FillId:   0,
		BorderId: 0,
		NumFmtId: 0,
	}}

	//add default ref for CellXfs
	ss.ml.CellXfs = &[]*ml.Style{{
		XfId:     0,
		FontId:   0,
		FillId:   0,
		BorderId: 0,
		NumFmtId: 0,
	}}

	//add default ref for CellStyles
	index := 0
	ss.ml.CellStyles = &[]*ml.NamedStyle{{
		Name:      "Normal",
		XfId:      0,
		BuiltinId: &index,
	}}
}

//adds a number formats for each type of number format if required. These styles will be used by cell's typed SetXXX methods
func (ss *StyleSheet) addTypedStylesIfRequired() {
	if len(ss.typedStyles) == 0 {
		for _, t := range []numberFormat.Type{
			numberFormat.General,
			numberFormat.Integer,
			numberFormat.Float,
			numberFormat.Date,
			numberFormat.Time,
			numberFormat.DateTime,
			numberFormat.DeltaTime,
		} {
			id, _ := numberFormat.Default(t)
			ss.typedStyles[t] = ss.addStyle(format.New(format.NumberFormatID(id)))
		}

		ss.file.MarkAsUpdated()
	}
}

//resolveNumberFormat returns resolved NumberFormat code for styleID
func (ss *StyleSheet) resolveNumberFormat(id format.StyleID) string {
	style := (*ss.ml.CellXfs)[id]

	//return code for built-in number format
	if number := numberFormat.Normalize(ml.NumberFormat{ID: style.NumFmtId}); len(number.Code) > 0 {
		return number.Code
	}

	//try to lookup through custom formats and find same ID
	for _, f := range *ss.ml.NumberFormats {
		if style.NumFmtId == f.ID {
			return f.Code
		}
	}

	//N.B.: wtf is going on?! non built-in and not existing id?
	_, code := numberFormat.Default(numberFormat.General)
	return code
}

//build indexes for fonts
func (ss *StyleSheet) buildFontIndexes() {
	if ss.ml.Fonts == nil {
		ss.ml.Fonts = &[]*ml.Font{}
	}

	for id, f := range *ss.ml.Fonts {
		ss.fontIndex[hash.Font(f).Hash()] = id
	}
}

//build indexes for fill
func (ss *StyleSheet) buildFillIndexes() {
	if ss.ml.Fills == nil {
		ss.ml.Fills = &[]*ml.Fill{}
	}

	for id, f := range *ss.ml.Fills {
		ss.fillIndex[hash.Fill(f).Hash()] = id
	}
}

//build indexes for border
func (ss *StyleSheet) buildBorderIndexes() {
	if ss.ml.Borders == nil {
		ss.ml.Borders = &[]*ml.Border{}
	}

	for id, f := range *ss.ml.Borders {
		ss.borderIndex[hash.Border(f).Hash()] = id
	}
}

//build indexes for number formats
func (ss *StyleSheet) buildNumberIndexes() {
	if ss.ml.NumberFormats == nil {
		ss.ml.NumberFormats = &[]*ml.NumberFormat{}
	}

	//N.B.: NumberFormat uses ID, not indexes
	for _, f := range *ss.ml.NumberFormats {
		ss.numberIndex[hash.NumberFormat(f).Hash()] = f.ID
	}
}

//build indexes for styles
func (ss *StyleSheet) buildStyleIndexes() {
	if ss.ml.CellXfs == nil {
		ss.ml.CellXfs = &[]*ml.Style{}
	}

	for id, xf := range *ss.ml.CellXfs {
		ss.styleIndex[hash.Style(xf).Hash()] = format.StyleID(id)
	}
}

//build indexes for differential styles
func (ss *StyleSheet) buildDiffStyleIndexes() {
	if ss.ml.Dxfs == nil {
		ss.ml.Dxfs = &[]*ml.DiffStyle{}
	}

	for id, dxf := range *ss.ml.Dxfs {
		ss.diffStyleIndex[hash.DiffStyle(dxf).Hash()] = format.DiffStyleID(id)
	}
}

//build indexes for all indexes
func (ss *StyleSheet) buildIndexes() {
	ss.buildBorderIndexes()
	ss.buildFillIndexes()
	ss.buildFontIndexes()
	ss.buildNumberIndexes()
	ss.buildStyleIndexes()
	ss.buildDiffStyleIndexes()
}

//adds a differential style
func (ss *StyleSheet) addDiffStyle(f *format.StyleFormat) format.DiffStyleID {
	ss.file.LoadIfRequired(ss.buildIndexes)

	//get settings for style
	font, fill, alignment, numFormat, protection, border := f.Settings()

	dXf := &ml.DiffStyle{
		Font:         font,
		Fill:         fill,
		Border:       border,
		NumberFormat: numFormat,
		Alignment:    alignment,
		Protection:   protection,
	}

	//return id of already existing information
	key := hash.DiffStyle(dXf).Hash()
	if id, ok := ss.diffStyleIndex[key]; ok {
		return format.DiffStyleID(id)
	}

	//add a new one and return related id
	nextID := format.DiffStyleID(len(*ss.ml.Dxfs))
	*ss.ml.Dxfs = append(*ss.ml.Dxfs, dXf)
	ss.diffStyleIndex[key] = nextID
	ss.file.MarkAsUpdated()
	return nextID
}

//adds a style
func (ss *StyleSheet) addStyle(f *format.StyleFormat) format.StyleID {
	ss.file.LoadIfRequired(ss.buildIndexes)

	//get settings and add information if required
	font, fill, alignment, numFormat, protection, border := f.Settings()
	fontID := ss.addFontIfRequired(font)
	fillID := ss.addFillIfRequired(fill)
	borderID := ss.addBorderIfRequired(border)
	numID := ss.addNumFormatIfRequired(numFormat)

	/*
		Note to remember excel internals:
		---
		cell.s = cellXfs.index
		cellXfs.xfId = cellStyleXf.index
		cellStyle.xfId = cellStyleXf.index
	*/

	cellXf := &ml.Style{
		XfId:              0, //we don't need this one, because we don't have task 'render final style', so for our case there is no 'override' of direct style
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

	//return id of already existing information
	key := hash.Style(cellXf).Hash()
	if id, ok := ss.styleIndex[key]; ok {
		return id
	}

	//add a new one and return related id
	nextID := format.StyleID(len(*ss.ml.CellXfs))
	*ss.ml.CellXfs = append(*ss.ml.CellXfs, cellXf)
	ss.styleIndex[key] = nextID
	ss.file.MarkAsUpdated()
	return nextID
}

//adds a new font if required
func (ss *StyleSheet) addFontIfRequired(font *ml.Font) int {
	//if there is no information, then use default
	if font == nil {
		return 0
	}

	//return id of already existing information
	key := hash.Font(font).Hash()
	if id, ok := ss.fontIndex[key]; ok {
		return id
	}

	//add a new one and return related id
	nextID := len(*ss.ml.Fonts)
	*ss.ml.Fonts = append(*ss.ml.Fonts, font)
	ss.fontIndex[key] = nextID
	ss.file.MarkAsUpdated()
	return nextID
}

//adds a new fill if required
func (ss *StyleSheet) addFillIfRequired(fill *ml.Fill) int {
	//if there is no information, then use default
	if fill == nil {
		return 0
	}

	//return id of already existing information
	key := hash.Fill(fill).Hash()
	if id, ok := ss.fillIndex[key]; ok {
		return id
	}

	//add a new one and return related id
	nextID := len(*ss.ml.Fills)
	*ss.ml.Fills = append(*ss.ml.Fills, fill)
	ss.fillIndex[key] = nextID
	ss.file.MarkAsUpdated()
	return nextID
}

//adds a new border if required
func (ss *StyleSheet) addBorderIfRequired(border *ml.Border) int {
	//if there is no information, then use default
	if border == nil {
		return 0
	}

	//return id of already existing information
	key := hash.Border(border).Hash()
	if id, ok := ss.borderIndex[key]; ok {
		return id
	}

	//add a new one and return related id
	nextID := len(*ss.ml.Borders)
	*ss.ml.Borders = append(*ss.ml.Borders, border)
	ss.borderIndex[key] = nextID
	ss.file.MarkAsUpdated()
	return nextID
}

//adds a new number format if required
func (ss *StyleSheet) addNumFormatIfRequired(number *ml.NumberFormat) int {
	//if there is no information, then use default
	if number == nil {
		return 0
	}

	//if is built-in format then return id
	if numberFormat.IsBuiltIn(number.ID) {
		return number.ID
	}

	//Return id of already existing information.
	//N.B.: Supposed that for custom format we have -1 as code, so hash should be same for new/existing custom format
	key := hash.NumberFormat(number).Hash()
	if id, ok := ss.numberIndex[key]; ok {
		return id
	}

	//try to lookup through custom formats and find same code
	for _, f := range *ss.ml.NumberFormats {
		if number.Code == f.Code {
			return f.ID
		}
	}

	//N.B.: NumberFormat uses ID, not indexes
	nextID := numberFormat.LastReservedID + len(*ss.ml.NumberFormats) + 1
	number.ID = nextID

	*ss.ml.NumberFormats = append(*ss.ml.NumberFormats, number)
	ss.numberIndex[key] = nextID
	ss.file.MarkAsUpdated()
	return nextID
}
