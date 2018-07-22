package ml

import (
	"github.com/plandem/ooxml/ml"
)

//StyleSheet is a direct mapping of XSD CT_Stylesheet
type StyleSheet struct {
	XMLName       ml.Name          `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main styleSheet"`
	NumberFormats *[]*NumberFormat `xml:"numFmts>numFmt,omitempty"`
	Fonts         *[]*Font         `xml:"fonts>font,omitempty"`
	Fills         *[]*Fill         `xml:"fills>fill,omitempty"`
	Borders       *[]*Border       `xml:"borders>border,omitempty"`
	CellStyleXfs  *[]*StyleRef     `xml:"cellStyleXfs>xf,omitempty"`
	CellXfs       *[]*StyleRef     `xml:"cellXfs>xf,omitempty"`
	CellStyles    *[]*CellStyle    `xml:"cellStyles>cellStyle,omitempty"`
	Dxfs          *ml.Reserved     `xml:"dxfs,omitempty"`
	TableStyles   *ml.Reserved     `xml:"tableStyles,omitempty"`
	Colors        *ml.Reserved     `xml:"colors,omitempty"`
	ExtLst        *ml.Reserved     `xml:"extLst,omitempty"`
}

//NumberFormat is a direct mapping of XSD CT_NumFmt
type NumberFormat struct {
	ID   int    `xml:"numFmtId,attr"`
	Code string `xml:"formatCode,attr"`
}

//Font is a direct mapping of XSD CT_Font
type Font struct {
	Name      ml.Property       `xml:"name,omitempty"`
	Charset   ml.PropertyInt    `xml:"charset,omitempty"`
	Family    FontFamilyType    `xml:"family,omitempty"`
	Bold      ml.PropertyBool   `xml:"b,omitempty"`
	Italic    ml.PropertyBool   `xml:"i,omitempty"`
	Strike    ml.PropertyBool   `xml:"strike,omitempty"`
	Shadow    ml.PropertyBool   `xml:"shadow,omitempty"`
	Condense  ml.PropertyBool   `xml:"condense,omitempty"`
	Extend    ml.PropertyBool   `xml:"extend,omitempty"`
	Color     *Color            `xml:"color,omitempty"`
	Size      ml.PropertyDouble `xml:"sz,omitempty"`
	Underline UnderlineType     `xml:"u,omitempty"`
	VAlign    FontVAlignType    `xml:"vertAlign,omitempty"`
	Scheme    FontSchemeType    `xml:"scheme,omitempty"`
}

//Color is a direct mapping of XSD CT_Color
type Color struct {
	Auto    bool             `xml:"auto,attr,omitempty"`
	Indexed ml.OptionalIndex `xml:"indexed,attr,omitempty"`
	RGB     string           `xml:"rgb,attr,omitempty"`
	Theme   ml.OptionalIndex `xml:"theme,attr,omitempty"`
	Tint    *float64         `xml:"tint,attr,omitempty"` //optional double
}

//Fill is a direct mapping of XSD CT_Fill
type Fill struct {
	Pattern  *PatternFill `xml:"patternFill,omitempty"`
	Gradient *ml.Reserved `xml:"gradientFill,omitempty"`
}

//PatternFill is a direct mapping of XSD CT_PatternFill
type PatternFill struct {
	Color      *Color      `xml:"fgColor,omitempty"`
	Background *Color      `xml:"bgColor,omitempty"`
	Type       PatternType `xml:"patternType,attr,omitempty"`
}

//Border is a direct mapping of XSD CT_Border
type Border struct {
	Left         *BorderSegment `xml:"left,omitempty"`  //WTF at ECMA-376 Edition 2/3/4/5, there is 'start', but no 'left'?!?!
	Right        *BorderSegment `xml:"right,omitempty"` //WTF at ECMA-376 Edition 2/3/4/5, there is 'end', but no 'right'?!?!
	Top          *BorderSegment `xml:"top,omitempty"`
	Bottom       *BorderSegment `xml:"bottom,omitempty"`
	Diagonal     *BorderSegment `xml:"diagonal,omitempty"`
	Vertical     *BorderSegment `xml:"vertical,omitempty"`
	Horizontal   *BorderSegment `xml:"horizontal,omitempty"`
	DiagonalUp   bool           `xml:"diagonalUp,attr,omitempty"`
	DiagonalDown bool           `xml:"diagonalDown,attr,omitempty"`
	Outline      bool           `xml:"outline,attr,omitempty"`
}

//BorderSegment is a direct mapping of XSD CT_BorderPr
type BorderSegment struct {
	Color *Color          `xml:"color,omitempty"`
	Type  BorderStyleType `xml:"style,attr,omitempty"`
}

//CellStyle is a direct mapping of XSD CT_CellStyle
type CellStyle struct {
	Name          string           `xml:"name,attr,omitempty"`
	XfId          int              `xml:"xfId,attr"`
	BuiltinId     ml.OptionalIndex `xml:"builtinId,attr,omitempty"`
	ILevel        uint             `xml:"iLevel,attr,omitempty"`
	Hidden        bool             `xml:"hidden,attr,omitempty"`
	CustomBuiltin bool             `xml:"customBuiltin,attr,omitempty"`
	ExtLst        *ml.Reserved     `xml:"extLst,omitempty"`
}

//StyleRef is a direct mapping of XSD CT_Xf
type StyleRef struct {
	NumFmtId          int             `xml:"numFmtId,attr"`
	FontId            int             `xml:"fontId,attr"`
	FillId            int             `xml:"fillId,attr"`
	BorderId          int             `xml:"borderId,attr"`
	XfId              int             `xml:"xfId,attr"`
	QuotePrefix       bool            `xml:"quotePrefix,attr,omitempty"`
	PivotButton       bool            `xml:"pivotButton,attr,omitempty"`
	ApplyNumberFormat bool            `xml:"applyNumberFormat,attr,omitempty"`
	ApplyFont         bool            `xml:"applyFont,attr,omitempty"`
	ApplyFill         bool            `xml:"applyFill,attr,omitempty"`
	ApplyBorder       bool            `xml:"applyBorder,attr,omitempty"`
	ApplyAlignment    bool            `xml:"applyAlignment,attr,omitempty"`
	ApplyProtection   bool            `xml:"applyProtection,attr,omitempty"`
	Alignment         *CellAlignment  `xml:"alignment,omitempty"`
	Protection        *CellProtection `xml:"protection,omitempty"`
	ExtLst            *ml.Reserved    `xml:"extLst,omitempty"`
}

//CellProtection is a direct mapping of XSD CT_CellProtection
type CellProtection struct {
	Locked bool `xml:"locked,attr,omitempty"`
	Hidden bool `xml:"hidden,attr,omitempty"`
}

//CellAlignment is a direct mapping of XSD CT_CellAlignment
type CellAlignment struct {
	Horizontal      HAlignType `xml:"horizontal,attr,omitempty"`
	Vertical        VAlignType `xml:"vertical,attr,omitempty"`
	TextRotation    int        `xml:"textRotation,attr,omitempty"`
	WrapText        bool       `xml:"wrapText,attr,omitempty"`
	Indent          int        `xml:"indent,attr,omitempty"`
	RelativeIndent  int        `xml:"relativeIndent,attr,omitempty"`
	JustifyLastLine bool       `xml:"justifyLastLine,attr,omitempty"`
	ShrinkToFit     bool       `xml:"shrinkToFit,attr,omitempty"`
	ReadingOrder    int        `xml:"readingOrder,attr,omitempty"`
}

//StyleRefID is helper type do forbid usage of integers directly and getting valid ID for StyleFormat via style-sheet
type StyleRefID int
