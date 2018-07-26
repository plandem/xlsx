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
	xfIndex  map[string]int
	dxfIndex map[string]int

	//hash -> index for types
	borderIndex map[string]int
	fillIndex   map[string]int
	fontIndex   map[string]int
	numberIndex map[string]int

	doc  *Spreadsheet
	file *ooxml.PackageFile
}

func newStyleSheet(f interface{}, doc *Spreadsheet) *StyleSheet {
	ss := &StyleSheet{
		doc:         doc,
		xfIndex:     make(map[string]int),
		dxfIndex:    make(map[string]int),
		borderIndex: make(map[string]int),
		fillIndex:   make(map[string]int),
		fontIndex:   make(map[string]int),
		numberIndex: make(map[string]int),
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

func (ss *StyleSheet) addDefaults() {
	//TODO: research about default items for a new XLSX
	ss.ml.Fills = &[]*ml.Fill{{
		Pattern: &ml.PatternFill{
			Type: format.PatternTypeNone,
		},
	}}

	ss.ml.Borders = &[]*ml.Border{{
		Left:   &ml.BorderSegment{},
		Right:  &ml.BorderSegment{},
		Top:    &ml.BorderSegment{},
		Bottom: &ml.BorderSegment{},
	}}

	ss.ml.Fonts = &[]*ml.Font{{
		Family: 2,
		Size:   12.0,
		Name:   "Calibri",
		Scheme: "minor",
	}}

	ss.ml.CellXfs = &[]*ml.StyleRef{{
		XfId:              0,
		FontId:            0,
		FillId:            0,
		BorderId:          0,
		NumFmtId:          0,
	}}
}

func (ss *StyleSheet) buildFontIndexes() {
	if ss.ml.Fonts == nil {
		ss.ml.Fonts = &[]*ml.Font{}
	}

	for id, f := range *ss.ml.Fonts {
		ss.fontIndex[hash.Font(f).Hash()] = id
	}
}

func (ss *StyleSheet) buildFillIndexes() {
	if ss.ml.Fills == nil {
		ss.ml.Fills = &[]*ml.Fill{}
	}

	for id, f := range *ss.ml.Fills {
		ss.fillIndex[hash.Fill(f).Hash()] = id
	}
}

func (ss *StyleSheet) buildBorderIndexes() {
	if ss.ml.Borders == nil {
		ss.ml.Borders = &[]*ml.Border{}
	}

	for id, f := range *ss.ml.Borders {
		ss.borderIndex[hash.Border(f).Hash()] = id
	}
}

func (ss *StyleSheet) buildNumberIndexes() {
	if ss.ml.NumberFormats == nil {
		ss.ml.NumberFormats = &[]*ml.NumberFormat{}
	}

	//N.B.: NumberFormat uses ID, not indexes
	for _, f := range *ss.ml.NumberFormats {
		ss.numberIndex[hash.NumberFormat(f).Hash()] = f.ID
	}
}

func (ss *StyleSheet) buildXfIndexes() {
	if ss.ml.CellXfs == nil {
		ss.ml.CellXfs = &[]*ml.StyleRef{}
	}

	//build xf indexes
	for id, xf := range *ss.ml.CellXfs {
		ss.xfIndex[hash.StyleRef(xf).Hash()] = id
	}
}

//buildIndexes process already existing styles and build indexed for it
func (ss *StyleSheet) buildIndexes() {
	ss.buildBorderIndexes()
	ss.buildFillIndexes()
	ss.buildFontIndexes()
	ss.buildNumberIndexes()
	ss.buildXfIndexes()
	//ss.buildDxfIndexes()
}

/*
func (ss *StyleSheet) addDXF(f *format.StyleFormat) int {
	ss.file.LoadIfRequired(ss.buildIndexes)

	//--- inline styles?
	//<xsd:element name="font" type="CT_Font" minOccurs="0" maxOccurs="1"/>
	//<xsd:element name="numFmt" type="CT_NumFmt" minOccurs="0" maxOccurs="1"/>
	//<xsd:element name="fill" type="CT_Fill" minOccurs="0" maxOccurs="1"/>
	//<xsd:element name="alignment" type="CT_CellAlignment" minOccurs="0" maxOccurs="1"/>
	//<xsd:element name="border" type="CT_Border" minOccurs="0" maxOccurs="1"/>
	//<xsd:element name="protection" type="CT_CellProtection" minOccurs="0" maxOccurs="1"/>
	//<xsd:element name="extLst" type="CT_ExtensionList" minOccurs="0" maxOccurs="1"/>

	dxf_id, ok := ss.getDXFByKey(f.Key())
	if !ok {
		//copy info from styles
		log.Println(f.Key())
	}

	return dxf_id
}
*/

func (ss *StyleSheet) addXF(f *format.StyleFormat) format.StyleRefID {
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

	cellXf := &ml.StyleRef{
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
	key := hash.StyleRef(cellXf).Hash()
	if id, ok := ss.xfIndex[key]; ok {
		return format.StyleRefID(id)
	}

	//add a new one and return related id
	nextID := len(*ss.ml.CellXfs)
	*ss.ml.CellXfs = append(*ss.ml.CellXfs, cellXf)
	ss.xfIndex[key] = nextID
	ss.file.MarkAsUpdated()
	return format.StyleRefID(nextID)
}

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

func (ss *StyleSheet) addNumFormatIfRequired(number *ml.NumberFormat) int {
	//if there is no information, then use default
	if number == nil {
		return 0
	}

	//if is built-in format then return id
	if numberFormat.IsBuiltIn(number.ID) {
		return number.ID
	}

	//return id of already existing information. hash ignores ID for non built-in types
	key := hash.NumberFormat(number).Hash()
	if id, ok := ss.numberIndex[key]; ok {
		return id
	}

	//N.B.: NumberFormat uses ID, not indexes
	nextID := numberFormat.LastReservedID + len(*ss.ml.NumberFormats) + 1
	number.ID = nextID

	*ss.ml.NumberFormats = append(*ss.ml.NumberFormats, number)
	ss.numberIndex[key] = nextID
	ss.file.MarkAsUpdated()
	return nextID
}
