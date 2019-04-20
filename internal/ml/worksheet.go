package ml

import (
	"github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//Worksheet is a direct mapping of XSD CT_Worksheet
type Worksheet struct {
	XMLName               ml.Name         `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main worksheet"`
	SheetPr               *ml.Reserved    `xml:"sheetPr,omitempty"`
	Dimension             *SheetDimension `xml:"dimension,omitempty"`
	SheetViews            *SheetViews     `xml:"sheetViews,omitempty"`
	SheetFormatPr         *ml.Reserved    `xml:"sheetFormatPr,omitempty"`
	Cols                  *[]*Col         `xml:"cols>col,omitempty"` //we HAVE TO remove 'cols' if there is no any 'col'
	SheetData             []*Row          `xml:"sheetData>row"`
	SheetCalcPr           *ml.Reserved    `xml:"sheetCalcPr,omitempty"`
	SheetProtection       *ml.Reserved    `xml:"sheetProtection,omitempty"`
	ProtectedRanges       *ml.Reserved    `xml:"protectedRanges,omitempty"`
	Scenarios             *ml.Reserved    `xml:"scenarios,omitempty"`
	AutoFilter            *ml.Reserved    `xml:"autoFilter,omitempty"`
	SortState             *ml.Reserved    `xml:"sortState,omitempty"`
	DataConsolidate       *ml.Reserved    `xml:"dataConsolidate,omitempty"`
	CustomSheetViews      *ml.Reserved    `xml:"customSheetViews,omitempty"`
	MergeCells            *[]*MergeCell   `xml:"mergeCells>mergeCell,omitempty"`
	PhoneticPr            *ml.Reserved    `xml:"phoneticPr,omitempty"`
	ConditionalFormatting *ml.Reserved    `xml:"conditionalFormatting,omitempty"`
	DataValidations       *ml.Reserved    `xml:"dataValidations,omitempty"`
	Hyperlinks            *ml.Reserved    `xml:"hyperlinks,omitempty"`
	PrintOptions          *ml.Reserved    `xml:"printOptions,omitempty"`
	PageMargins           *ml.Reserved    `xml:"pageMargins,omitempty"`
	PageSetup             *ml.Reserved    `xml:"pageSetup,omitempty"`
	HeaderFooter          *ml.Reserved    `xml:"headerFooter,omitempty"`
	RowBreaks             *ml.Reserved    `xml:"rowBreaks,omitempty"`
	ColBreaks             *ml.Reserved    `xml:"colBreaks,omitempty"`
	CustomProperties      *ml.Reserved    `xml:"customProperties,omitempty"`
	CellWatches           *ml.Reserved    `xml:"cellWatches,omitempty"`
	IgnoredErrors         *ml.Reserved    `xml:"ignoredErrors,omitempty"`
	SmartTags             *ml.Reserved    `xml:"smartTags,omitempty"`
	Drawing               *ml.Reserved    `xml:"drawing,omitempty"`
	DrawingHF             *ml.Reserved    `xml:"drawingHF,omitempty"`
	Picture               *ml.Reserved    `xml:"picture,omitempty"`
	OleObjects            *ml.Reserved    `xml:"oleObjects,omitempty"`
	Controls              *ml.Reserved    `xml:"controls,omitempty"`
	WebPublishItems       *ml.Reserved    `xml:"webPublishItems,omitempty"`
	TableParts            *ml.Reserved    `xml:"tableParts,omitempty"`
	ExtLst                *ml.Reserved    `xml:"extLst,omitempty"`
}

//SheetDimension is a direct mapping of XSD CT_SheetDimension
type SheetDimension struct {
	Bounds primitives.Bounds `xml:"ref,attr"`
}

//Col is a direct mapping of XSD CT_Col
type Col struct {
	Min          int     `xml:"min,attr"`
	Max          int     `xml:"max,attr"`
	Width        float32 `xml:"width,attr,omitempty"`
	Style        StyleID `xml:"style,attr,omitempty"`
	Hidden       bool    `xml:"hidden,attr,omitempty"`
	BestFit      bool    `xml:"bestFit,attr,omitempty"`
	CustomWidth  bool    `xml:"customWidth,attr,omitempty"`
	Phonetic     bool    `xml:"phonetic,attr,omitempty"`
	OutlineLevel uint8   `xml:"outlineLevel,attr,omitempty"`
	Collapsed    bool    `xml:"collapsed,attr,omitempty"`
}

//Row is a direct mapping of XSD CT_Row
type Row struct {
	Cells        []*Cell      `xml:"c"`
	ExtLst       *ml.Reserved `xml:"extLst,omitempty"`
	Ref          int          `xml:"r,attr,omitempty"` //1-based index
	Spans        string       `xml:"spans,attr,omitempty"`
	Style        StyleID      `xml:"s,attr,omitempty"`
	CustomFormat bool         `xml:"customFormat,attr,omitempty"`
	Height       float32      `xml:"ht,attr,omitempty"`
	Hidden       bool         `xml:"hidden,attr,omitempty"`
	CustomHeight bool         `xml:"customHeight,attr,omitempty"`
	OutlineLevel uint8        `xml:"outlineLevel,attr,omitempty"`
	Collapsed    bool         `xml:"collapsed,attr,omitempty"`
	ThickTop     bool         `xml:"thickTop,attr,omitempty"`
	ThickBot     bool         `xml:"thickBot,attr,omitempty"`
	Phonetic     bool         `xml:"ph,attr,omitempty"`
}

//Cell is a direct mapping of XSD CT_Cell
type Cell struct {
	Formula   *CellFormula        `xml:"f,omitempty"`
	Value     string              `xml:"v,omitempty"`
	InlineStr *StringItem         `xml:"is,omitempty"`
	ExtLst    *ml.Reserved        `xml:"extLst,omitempty"`
	Ref       primitives.CellRef  `xml:"r,attr"`
	Style     StyleID             `xml:"s,attr,omitempty"`
	Type      primitives.CellType `xml:"t,attr,omitempty"`
	Cm        ml.OptionalIndex    `xml:"cm,attr,omitempty"`
	Vm        ml.OptionalIndex    `xml:"vm,attr,omitempty"`
	Ph        bool                `xml:"ph,attr,omitempty"`
}

//CellFormula is a direct mapping of XSD CT_CellFormula
type CellFormula struct {
	Content string                     `xml:",chardata"`
	T       primitives.CellFormulaType `xml:"t,attr,omitempty"` //default 'normal'
	Aca     bool                       `xml:"aca,attr,omitempty"`
	Bounds  primitives.Bounds          `xml:"ref,attr,omitempty"`
	Dt2D    bool                       `xml:"dt2D,attr,omitempty"`
	Dtr     bool                       `xml:"dtr,attr,omitempty"`
	Del1    bool                       `xml:"del1,attr,omitempty"`
	Del2    bool                       `xml:"del2,attr,omitempty"`
	R1      primitives.CellRef         `xml:"r1,attr,omitempty"`
	R2      primitives.CellRef         `xml:"r2,attr,omitempty"`
	Ca      bool                       `xml:"ca,attr,omitempty"`
	Si      ml.OptionalIndex           `xml:"si,attr,omitempty"`
	Bx      bool                       `xml:"bx,attr,omitempty"`
}

//MergeCell is a direct mapping of XSD CT_MergeCell
type MergeCell struct {
	Bounds primitives.Bounds `xml:"ref,attr"`
}

//SheetViews is a direct mapping of XSD CT_SheetViews
type SheetViews struct {
	SheetView []*SheetView `xml:"sheetView,omitempty"`
	ExtLst    *ml.Reserved `xml:"extLst,omitempty"`
}

//SheetView is a direct mapping of XSD CT_SheetView
type SheetView struct {
	Pane                     *ml.Reserved       `xml:"pane,omitempty"`
	Selection                *ml.Reserved       `xml:"selection,omitempty"`
	PivotSelection           *ml.Reserved       `xml:"pivotSelection,omitempty"`
	ExtLst                   *ml.Reserved       `xml:"extLst,omitempty"`
	WindowProtection         bool               `xml:"windowProtection,attr,omitempty"`
	ShowFormulas             bool               `xml:"showFormulas,attr,omitempty"`
	ShowGridLines            bool               `xml:"showGridLines,attr,omitempty"`
	ShowRowColHeaders        bool               `xml:"showRowColHeaders,attr,omitempty"`
	ShowZeros                bool               `xml:"showZeros,attr,omitempty"`
	RightToLeft              bool               `xml:"rightToLeft,attr,omitempty"`
	TabSelected              bool               `xml:"tabSelected,attr,omitempty"`
	ShowRuler                bool               `xml:"showRuler,attr,omitempty"`
	ShowOutlineSymbols       bool               `xml:"showOutlineSymbols,attr,omitempty"`
	DefaultGridColor         bool               `xml:"defaultGridColor,attr,omitempty"`
	ShowWhiteSpace           bool               `xml:"showWhiteSpace,attr,omitempty"`
	View                     string             `xml:"view,attr,omitempty"` //ST_SheetViewType
	TopLeftCell              primitives.CellRef `xml:"topLeftCell,attr,omitempty"`
	ColorId                  uint               `xml:"colorId,attr,omitempty"`
	ZoomScale                uint               `xml:"zoomScale,attr,omitempty"`
	ZoomScaleNormal          uint               `xml:"zoomScaleNormal,attr,omitempty"`
	ZoomScaleSheetLayoutView uint               `xml:"zoomScaleSheetLayoutView,attr,omitempty"`
	ZoomScalePageLayoutView  uint               `xml:"zoomScalePageLayoutView,attr,omitempty"`
	WorkbookViewId           uint               `xml:"workbookViewId,attr"`
}
