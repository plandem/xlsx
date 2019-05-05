package xlsx

import (
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/hash"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/number_format"
	_ "unsafe"
)

//go:linkname fromStyleFormat github.com/plandem/xlsx/format.fromStyleFormat
func fromStyleFormat(f *format.StyleFormat) (font *ml.Font, fill *ml.Fill, alignment *ml.CellAlignment, numFormat *ml.NumberFormat, protection *ml.CellProtection, border *ml.Border, namedInfo *ml.NamedStyleInfo)

//StyleSheet is a higher level object that wraps ml.StyleSheet with functionality
type StyleSheet struct {
	ml ml.StyleSheet

	//hash -> index for styles
	directStyleIndex map[hash.Code]format.DirectStyleID
	diffStyleIndex   map[hash.Code]format.DiffStyleID
	namedStyleIndex  map[hash.Code]format.NamedStyleID

	//hash -> index for types
	borderIndex map[hash.Code]int
	fillIndex   map[hash.Code]int
	fontIndex   map[hash.Code]int
	numberIndex map[hash.Code]int

	//hash for typed number formats
	typedStyles map[numberFormat.Type]format.DirectStyleID

	doc  *Spreadsheet
	file *ooxml.PackageFile
}

func newStyleSheet(f interface{}, doc *Spreadsheet) *StyleSheet {
	ss := &StyleSheet{
		doc:              doc,
		directStyleIndex: make(map[hash.Code]format.DirectStyleID),
		diffStyleIndex:   make(map[hash.Code]format.DiffStyleID),
		namedStyleIndex:  make(map[hash.Code]format.NamedStyleID),
		borderIndex:      make(map[hash.Code]int),
		fillIndex:        make(map[hash.Code]int),
		fontIndex:        make(map[hash.Code]int),
		numberIndex:      make(map[hash.Code]int),
		typedStyles:      make(map[numberFormat.Type]format.DirectStyleID),
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
	ss.ml.Fills.Items = append(ss.ml.Fills.Items,
		&ml.Fill{
			Pattern: &ml.PatternFill{
				Type: format.PatternTypeNone,
			},
		},
		&ml.Fill{
			Pattern: &ml.PatternFill{
				Type: format.PatternTypeGray125,
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
		Family: format.FontFamilySwiss,
		Scheme: format.FontSchemeMinor,
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
		def := format.NewStyles(
			format.NamedStyle(format.NamedStyleNormal),
			format.Font.Default,
		)
	*/
}

//build indexes for all indexes
func (ss *StyleSheet) buildIndexes() {
	//build indexes for fonts
	for id, f := range ss.ml.Fonts.Items {
		ss.fontIndex[hash.Font(f).Hash()] = id
	}

	//build indexes for fill
	for id, f := range ss.ml.Fills.Items {
		ss.fillIndex[hash.Fill(f).Hash()] = id
	}

	//build indexes for border
	for id, f := range ss.ml.Borders.Items {
		ss.borderIndex[hash.Border(f).Hash()] = id
	}

	//build indexes for number formats
	for _, f := range ss.ml.NumberFormats.Items {
		//N.B.: NumberFormat uses ID, not indexes
		ss.numberIndex[hash.NumberFormat(f).Hash()] = f.ID
	}

	//build indexes for named styles
	for id, xf := range ss.ml.CellStyleXfs.Items {
		ss.namedStyleIndex[hash.NamedStyle(xf).Hash()] = format.NamedStyleID(id)
	}

	//build indexes for direct styles
	for id, xf := range ss.ml.CellXfs.Items {
		ss.directStyleIndex[hash.DirectStyle(xf).Hash()] = format.DirectStyleID(id)
	}

	//build indexes for differential styles
	for id, dxf := range ss.ml.Dxfs.Items {
		ss.diffStyleIndex[hash.DiffStyle(dxf).Hash()] = format.DiffStyleID(id)
	}
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
			ss.typedStyles[t] = ss.addStyle(format.NewStyles(format.NumberFormatID(id)))
		}

		ss.file.MarkAsUpdated()
	}
}

//resolveNumberFormat returns resolved NumberFormat code for styleID
func (ss *StyleSheet) resolveNumberFormat(id ml.DirectStyleID) string {
	style := ss.ml.CellXfs.Items[id]

	//return code for built-in number format
	if number := numberFormat.Normalize(ml.NumberFormat{ID: style.NumFmtId}); len(number.Code) > 0 {
		return number.Code
	}

	//try to lookup through custom formats and find same ID
	for _, f := range ss.ml.NumberFormats.Items {
		if style.NumFmtId == f.ID {
			return f.Code
		}
	}

	//N.B.: wtf is going on?! non built-in and not existing id?
	_, code := numberFormat.Default(numberFormat.General)
	return code
}

//resolveDirectStyle returns resolved StyleFormat for DirectStyleID
func (ss *StyleSheet) resolveDirectStyle(id ml.DirectStyleID) *format.StyleFormat {
	if id == 0 {
		return nil
	}

	cellStyle := ss.ml.CellXfs.Items[id]
	style := &format.StyleFormat{}
	_ = cellStyle

	//TODO: Populate format.StyleFormat with required information
	panic(errorNotSupported)

	return style
}

//adds a differential style
func (ss *StyleSheet) addDiffStyle(f *format.StyleFormat) format.DiffStyleID {
	ss.file.LoadIfRequired(ss.buildIndexes)

	//get settings for style
	font, fill, alignment, numFormat, protection, border, _ := fromStyleFormat(f)

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
	nextID := format.DiffStyleID(len(ss.ml.Dxfs.Items))
	ss.ml.Dxfs.Items = append(ss.ml.Dxfs.Items, dXf)
	ss.diffStyleIndex[key] = nextID
	ss.file.MarkAsUpdated()
	return nextID
}

//add a named style if required
func (ss *StyleSheet) addNamedStyleIfRequired(namedInfo *ml.NamedStyleInfo, style ml.Style) ml.NamedStyleID {
	if namedInfo == nil {
		return 0
	}

	namedStyle := ml.NamedStyle(style)
	key := hash.NamedStyle(&namedStyle).Hash()

	//TODO: check if it's possible to have 2 same built-styles

	//if there is already same styles, then use it
	if id, ok := ss.namedStyleIndex[key]; ok {
		namedInfo.XfId = ml.NamedStyleID(id)
	} else {
		//add a new style
		nextID := format.NamedStyleID(len(ss.ml.CellStyleXfs.Items))
		ss.ml.CellStyleXfs.Items = append(ss.ml.CellStyleXfs.Items, &namedStyle)
		ss.namedStyleIndex[key] = nextID

		//add style info
		namedInfo.XfId = ml.NamedStyleID(nextID)
		ss.ml.CellStyles.Items = append(ss.ml.CellStyles.Items, namedInfo)
	}

	//add named info
	ss.file.MarkAsUpdated()
	return namedInfo.XfId
}

//adds a style. Style can be Direct or Named. Depends on settings.
func (ss *StyleSheet) addStyle(f *format.StyleFormat) format.DirectStyleID {
	ss.file.LoadIfRequired(ss.buildIndexes)

	//get settings and add information if required
	font, fill, alignment, numFormat, protection, border, namedInfo := fromStyleFormat(f)
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

	XfId := ml.NamedStyleID(0)
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

	//add named style if required and get related XfId
	XfId = ss.addNamedStyleIfRequired(namedInfo, style)

	cellXf := &ml.DirectStyle{
		XfId:  XfId,
		Style: style,
	}

	//return id of already existing information
	key := hash.DirectStyle(cellXf).Hash()
	if id, ok := ss.directStyleIndex[key]; ok {
		return id
	}

	//add a new one and return related id
	nextID := format.DirectStyleID(len(ss.ml.CellXfs.Items))
	ss.ml.CellXfs.Items = append(ss.ml.CellXfs.Items, cellXf)
	ss.directStyleIndex[key] = nextID
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
	nextID := len(ss.ml.Fonts.Items)
	ss.ml.Fonts.Items = append(ss.ml.Fonts.Items, font)
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
	nextID := len(ss.ml.Fills.Items)
	ss.ml.Fills.Items = append(ss.ml.Fills.Items, fill)
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
	nextID := len(ss.ml.Borders.Items)
	ss.ml.Borders.Items = append(ss.ml.Borders.Items, border)
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
	for _, f := range ss.ml.NumberFormats.Items {
		if number.Code == f.Code {
			return f.ID
		}
	}

	//N.B.: NumberFormat uses ID, not indexes
	nextID := numberFormat.LastReservedID + len(ss.ml.NumberFormats.Items) + 1
	number.ID = nextID

	ss.ml.NumberFormats.Items = append(ss.ml.NumberFormats.Items, number)
	ss.numberIndex[key] = nextID
	ss.file.MarkAsUpdated()
	return nextID
}
