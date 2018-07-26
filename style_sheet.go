package xlsx

import (
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/hash"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/number_format"
	"reflect"
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

		ss.file.MarkAsUpdated()
	}

	return ss
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

	for id, f := range *ss.ml.NumberFormats {
		ss.numberIndex[hash.NumberFormat(f).Hash()] = id
	}
}

func (ss *StyleSheet) buildXfIndexes() {
	if ss.ml.CellXfs == nil {
		ss.ml.CellXfs = &[]*ml.StyleRef{}
	}

	var (
		font *ml.Font
		border *ml.Border
		fill *ml.Fill
		number *ml.NumberFormat
	)

	//build xf indexes
	for id, xf := range *ss.ml.CellXfs {
		font = (*ss.ml.Fonts)[xf.FontId]
		border = (*ss.ml.Borders)[xf.BorderId]
		fill = (*ss.ml.Fills)[xf.FillId]

		if numberFormat.IsBuiltIn(xf.NumFmtId) {
			//create built-in pseudo type
			number = &ml.NumberFormat{}
			*number = numberFormat.New(xf.NumFmtId, "")
		} else {
			//lookup for existing type
			for _, num := range *ss.ml.NumberFormats {
				if num.ID == xf.NumFmtId {
					number = num
					break
				}
			}
		}

		key := hash.Style(font, fill, xf.Alignment, number, xf.Protection, border)
		ss.xfIndex[key] = id
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

	xfID, ok := ss.xfIndex[f.Key()]
	if !ok {
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

		//we don't need this one, because we don't have task 'render final style', so for our case there is no 'override' of direct style
		styleXfID := 0

		//now let's try to get xf index of direct formatting - cellXf
		cellXf := &ml.StyleRef{
			XfId:              styleXfID,
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

		xfID = -1

		if ss.ml.CellXfs == nil {
			//file can be new or missing this part of information, so let's fix it
			ss.ml.CellXfs = &[]*ml.StyleRef{}
		} else {
			//try to find at already existing cellXf
			for id, xf := range *ss.ml.CellXfs {
				if reflect.DeepEqual(xf, cellXf) {
					xfID = id
					break
				}
			}
		}

		//add a new cellXf
		if xfID == -1 {
			xfID = len(*ss.ml.CellXfs)
			*ss.ml.CellXfs = append(*ss.ml.CellXfs, cellXf)
			ss.file.MarkAsUpdated()
		}

		//link this format with xfID
		ss.xfIndex[f.Key()] = xfID
	}

	return format.StyleRefID(xfID)
}

func (ss *StyleSheet) addFontIfRequired(font *ml.Font) int {
	if font == nil {
		return 0
	}

	fontID := -1

	if ss.ml.Fonts == nil {
		//file can be new or missing this part of information, so let's fix it
		ss.ml.Fonts = &[]*ml.Font{}
	} else {
		//get id of font
		for id, f := range *ss.ml.Fonts {
			if reflect.DeepEqual(f, font) {
				fontID = id
				break
			}
		}
	}

	//add a new one, if there is no such font
	if fontID == -1 {
		fontID = len(*ss.ml.Fonts)
		*ss.ml.Fonts = append(*ss.ml.Fonts, font)
		ss.file.MarkAsUpdated()
	}

	return fontID
}

func (ss *StyleSheet) addFillIfRequired(fill *ml.Fill) int {
	if fill == nil {
		return 0
	}

	fillID := -1

	if ss.ml.Fills == nil {
		//file can be new or missing this part of information, so let's fix it
		ss.ml.Fills = &[]*ml.Fill{}
	} else {
		//get id of fill
		for id, f := range *ss.ml.Fills {
			if reflect.DeepEqual(f, fill) {
				fillID = id
				break
			}
		}
	}

	//add a new one, if there is no such fill
	if fillID == -1 {
		fillID = len(*ss.ml.Fills)
		*ss.ml.Fills = append(*ss.ml.Fills, fill)
		ss.file.MarkAsUpdated()
	}

	return fillID
}

func (ss *StyleSheet) addBorderIfRequired(border *ml.Border) int {
	if border == nil {
		return 0
	}

	borderID := -1

	if ss.ml.Borders == nil {
		//file can be new or missing this part of information, so let's fix it
		ss.ml.Borders = &[]*ml.Border{}
	} else {
		//get id of border
		for id, b := range *ss.ml.Borders {
			if reflect.DeepEqual(b, border) {
				borderID = id
				break
			}
		}
	}

	//add a new one, if there is no such border
	if borderID == -1 {
		borderID = len(*ss.ml.Borders)
		*ss.ml.Borders = append(*ss.ml.Borders, border)
		ss.file.MarkAsUpdated()
	}

	return borderID
}

func (ss *StyleSheet) addNumFormatIfRequired(numFormat *ml.NumberFormat) int {
	if numFormat == nil {
		return 0
	}

	numID := -1

	if ss.ml.NumberFormats == nil {
		//file can be new or missing this part of information, so let's fix it
		ss.ml.NumberFormats = &[]*ml.NumberFormat{}
	} else {
		//get id of number format
		for id, b := range *ss.ml.NumberFormats {
			if reflect.DeepEqual(b, numFormat) {
				numID = id
				break
			}
		}
	}

	//add a new one, if there is no such number format
	if numID == -1 {
		numID = len(*ss.ml.NumberFormats)
		*ss.ml.NumberFormats = append(*ss.ml.NumberFormats, numFormat)
		ss.file.MarkAsUpdated()
	}

	return numID
}
