package xlsx

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
)

func addNewStyles(xl *Spreadsheet, t *testing.T) format.DirectStyleID {
	require.NotNil(t, xl)

	require.Equal(t, 1, len(xl.styleSheet.directStyleIndex))
	require.Equal(t, 1, len(xl.styleSheet.borderIndex))
	require.Equal(t, 2, len(xl.styleSheet.fillIndex))
	require.Equal(t, 1, len(xl.styleSheet.fontIndex))
	require.Equal(t, 0, len(xl.styleSheet.numberIndex))

	//add font
	style := format.NewStyles(
		format.Font.Size(8),
		format.Font.Color("#FF1122"),
	)

	styleRef := xl.AddFormatting(style)
	require.Equal(t, format.DirectStyleID(1), styleRef)
	require.Equal(t, 2, len(xl.styleSheet.directStyleIndex))
	require.Equal(t, 1, len(xl.styleSheet.borderIndex))
	require.Equal(t, 2, len(xl.styleSheet.fillIndex))
	require.Equal(t, 2, len(xl.styleSheet.fontIndex))
	require.Equal(t, 0, len(xl.styleSheet.numberIndex))

	//add fill
	style.Set(
		format.Fill.Type(format.PatternTypeLightGrid),
		format.Fill.Background("#EFF142"),
	)

	styleRef = xl.AddFormatting(style)
	require.Equal(t, format.DirectStyleID(2), styleRef)
	require.Equal(t, 3, len(xl.styleSheet.directStyleIndex))
	require.Equal(t, 1, len(xl.styleSheet.borderIndex))
	require.Equal(t, 3, len(xl.styleSheet.fillIndex))
	require.Equal(t, 2, len(xl.styleSheet.fontIndex))
	require.Equal(t, 0, len(xl.styleSheet.numberIndex))

	//add built-in number
	style.Set(
		format.NumberFormatID(45),
	)

	styleRef = xl.AddFormatting(style)
	require.Equal(t, format.DirectStyleID(3), styleRef)
	require.Equal(t, 4, len(xl.styleSheet.directStyleIndex))
	require.Equal(t, 1, len(xl.styleSheet.borderIndex))
	require.Equal(t, 3, len(xl.styleSheet.fillIndex))
	require.Equal(t, 2, len(xl.styleSheet.fontIndex))
	require.Equal(t, 0, len(xl.styleSheet.numberIndex))

	//add custom number
	style.Set(
		format.NumberFormat(`$0.00" usd"`),
	)

	styleRef = xl.AddFormatting(style)
	require.Equal(t, format.DirectStyleID(4), styleRef)
	require.Equal(t, 5, len(xl.styleSheet.directStyleIndex))
	require.Equal(t, 1, len(xl.styleSheet.borderIndex))
	require.Equal(t, 3, len(xl.styleSheet.fillIndex))
	require.Equal(t, 2, len(xl.styleSheet.fontIndex))
	require.Equal(t, 1, len(xl.styleSheet.numberIndex))

	//add border
	style.Set(
		format.Border.Color("#1122FF"),
		format.Border.Type(format.BorderStyleDashDot),
	)

	styleRef = xl.AddFormatting(style)
	require.Equal(t, format.DirectStyleID(5), styleRef)
	require.Equal(t, 6, len(xl.styleSheet.directStyleIndex))
	require.Equal(t, 2, len(xl.styleSheet.borderIndex))
	require.Equal(t, 3, len(xl.styleSheet.fillIndex))
	require.Equal(t, 2, len(xl.styleSheet.fontIndex))
	require.Equal(t, 1, len(xl.styleSheet.numberIndex))

	//add alignment
	style.Set(
		format.Alignment.VAlign(format.VAlignBottom),
		format.Alignment.HAlign(format.HAlignFill),
	)

	styleRef = xl.AddFormatting(style)
	require.Equal(t, format.DirectStyleID(6), styleRef)
	require.Equal(t, 7, len(xl.styleSheet.directStyleIndex))
	require.Equal(t, 2, len(xl.styleSheet.borderIndex))
	require.Equal(t, 3, len(xl.styleSheet.fillIndex))
	require.Equal(t, 2, len(xl.styleSheet.fontIndex))
	require.Equal(t, 1, len(xl.styleSheet.numberIndex))

	//add protection
	style.Set(
		format.Protection.Hidden,
		format.Protection.Locked,
	)

	styleRef = xl.AddFormatting(style)
	require.Equal(t, format.DirectStyleID(7), styleRef)
	require.Equal(t, 8, len(xl.styleSheet.directStyleIndex))
	require.Equal(t, 2, len(xl.styleSheet.borderIndex))
	require.Equal(t, 3, len(xl.styleSheet.fillIndex))
	require.Equal(t, 2, len(xl.styleSheet.fontIndex))
	require.Equal(t, 1, len(xl.styleSheet.numberIndex))

	return styleRef
}

func addExistingStyles(xl *Spreadsheet, t *testing.T) {
	require.NotNil(t, xl)

	style := format.NewStyles(
		format.Font.Size(8),
		format.Font.Color("#FF1122"),
		format.Fill.Type(format.PatternTypeLightGrid),
		format.Fill.Background("#EFF142"),
		format.NumberFormat(`$0.00" usd"`),
		format.Border.Color("#1122FF"),
		format.Border.Type(format.BorderStyleDashDot),
		format.Alignment.VAlign(format.VAlignBottom),
		format.Alignment.HAlign(format.HAlignFill),
		format.Protection.Hidden,
		format.Protection.Locked,
	)

	styleRef := xl.AddFormatting(style)
	require.Equal(t, format.DirectStyleID(7), styleRef)
	require.Equal(t, 14, len(xl.styleSheet.directStyleIndex))
	require.Equal(t, 2, len(xl.styleSheet.borderIndex))
	require.Equal(t, 3, len(xl.styleSheet.fillIndex))
	require.Equal(t, 2, len(xl.styleSheet.fontIndex))
	require.Equal(t, 1, len(xl.styleSheet.numberIndex))
}

func checkStyles(xl *Spreadsheet, t *testing.T) {
	require.NotNil(t, xl)

	//validate stored fonts
	require.Equal(t, []*ml.Font{
		//default font
		{
			Name:   "Calibri",
			Family: format.FontFamilySwiss,
			Size:   11,
			Scheme: format.FontSchemeMinor,
		},
		//new font
		{
			Size:  8,
			Color: &ml.Color{RGB: "FFFF1122"},
		},
	}, xl.styleSheet.ml.Fonts.Items)

	//validate stored fills
	require.Equal(t, []*ml.Fill{
		//default fill
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
		//new fill
		{
			Pattern: &ml.PatternFill{
				Type:       format.PatternTypeLightGrid,
				Background: &ml.Color{RGB: "FFEFF142"},
			},
		},
	}, xl.styleSheet.ml.Fills.Items)

	//validate stored number
	require.Equal(t, []*ml.NumberFormat{
		{
			ID:   164,
			Code: `$0.00" usd"`,
		},
	}, xl.styleSheet.ml.NumberFormats.Items)

	//validate stored border
	require.Equal(t, []*ml.Border{
		//default border
		{
			Left:   &ml.BorderSegment{},
			Right:  &ml.BorderSegment{},
			Top:    &ml.BorderSegment{},
			Bottom: &ml.BorderSegment{},
		},
		//new border
		{
			Left:   &ml.BorderSegment{Type: format.BorderStyleDashDot, Color: &ml.Color{RGB: "FF1122FF"}},
			Right:  &ml.BorderSegment{Type: format.BorderStyleDashDot, Color: &ml.Color{RGB: "FF1122FF"}},
			Top:    &ml.BorderSegment{Type: format.BorderStyleDashDot, Color: &ml.Color{RGB: "FF1122FF"}},
			Bottom: &ml.BorderSegment{Type: format.BorderStyleDashDot, Color: &ml.Color{RGB: "FF1122FF"}},
		},
	}, xl.styleSheet.ml.Borders.Items)

	//validate stored Xf
	require.Equal(t, []*ml.DirectStyle{
		//default xf
		{
			XfId: 0,
			Style: ml.Style{
				FontId:   0,
				FillId:   0,
				BorderId: 0,
				NumFmtId: 0,
			},
		},
		//new xf
		{
			XfId: 0,
			Style: ml.Style{
				FontId:    1,
				FillId:    0,
				BorderId:  0,
				NumFmtId:  0,
				ApplyFont: true,
			},
		},
		{
			XfId: 0,
			Style: ml.Style{
				FontId:    1,
				FillId:    2,
				BorderId:  0,
				NumFmtId:  0,
				ApplyFont: true,
				ApplyFill: true,
			},
		},
		{
			XfId: 0,
			Style: ml.Style{
				FontId:            1,
				FillId:            2,
				BorderId:          0,
				NumFmtId:          45,
				ApplyFont:         true,
				ApplyFill:         true,
				ApplyNumberFormat: true,
			},
		},
		{
			XfId: 0,
			Style: ml.Style{
				FontId:            1,
				FillId:            2,
				BorderId:          0,
				NumFmtId:          164,
				ApplyFont:         true,
				ApplyFill:         true,
				ApplyNumberFormat: true,
			},
		},
		{
			XfId: 0,
			Style: ml.Style{
				FontId:            1,
				FillId:            2,
				BorderId:          1,
				NumFmtId:          164,
				ApplyFont:         true,
				ApplyFill:         true,
				ApplyNumberFormat: true,
				ApplyBorder:       true,
			},
		},
		{
			XfId: 0,
			Style: ml.Style{
				FontId:            1,
				FillId:            2,
				BorderId:          1,
				NumFmtId:          164,
				ApplyFont:         true,
				ApplyFill:         true,
				ApplyNumberFormat: true,
				ApplyBorder:       true,
				ApplyAlignment:    true,
				Alignment: &ml.CellAlignment{
					Vertical:   format.VAlignBottom,
					Horizontal: format.HAlignFill,
				},
			},
		},
		{
			XfId: 0,
			Style: ml.Style{
				FontId:            1,
				FillId:            2,
				BorderId:          1,
				NumFmtId:          164,
				ApplyFont:         true,
				ApplyFill:         true,
				ApplyNumberFormat: true,
				ApplyBorder:       true,
				ApplyAlignment:    true,
				ApplyProtection:   true,
				Alignment: &ml.CellAlignment{
					Vertical:   format.VAlignBottom,
					Horizontal: format.HAlignFill,
				},
				Protection: &ml.CellProtection{
					Hidden: true,
					Locked: true,
				},
			},
		},
		//types styles for number format
		{
			XfId: 0,
			Style: ml.Style{
				FontId:            0,
				FillId:            0,
				BorderId:          0,
				NumFmtId:          0x01,
				ApplyNumberFormat: true,
			},
		},
		{
			XfId: 0,
			Style: ml.Style{
				FontId:            0,
				FillId:            0,
				BorderId:          0,
				NumFmtId:          0x02,
				ApplyNumberFormat: true,
			},
		},
		{
			XfId: 0,
			Style: ml.Style{
				FontId:            0,
				FillId:            0,
				BorderId:          0,
				NumFmtId:          0x0e,
				ApplyNumberFormat: true,
			},
		},
		{
			XfId: 0,
			Style: ml.Style{
				FontId:            0,
				FillId:            0,
				BorderId:          0,
				NumFmtId:          0x14,
				ApplyNumberFormat: true,
			},
		},
		{
			XfId: 0,
			Style: ml.Style{
				FontId:            0,
				FillId:            0,
				BorderId:          0,
				NumFmtId:          0x16,
				ApplyNumberFormat: true,
			},
		},
		{
			XfId: 0,
			Style: ml.Style{
				FontId:            0,
				FillId:            0,
				BorderId:          0,
				NumFmtId:          0x2d,
				ApplyNumberFormat: true,
			},
		},
	}, xl.styleSheet.ml.CellXfs.Items)
}

func TestStyleSheets_create(t *testing.T) {
	xl := New()
	//after creating XLSX we must have only only default styles and new
	styleRef := addNewStyles(xl, t)

	//after creating sheet we must have also 'typed' styles
	sheet := xl.AddSheet("test sheet")
	sheet.Row(0).SetFormatting(styleRef)

	//try to add already existing styles
	addExistingStyles(xl, t)

	//check stored information, styles must be: default + new + typed
	checkStyles(xl, t)
	xl.SaveAs("./test_files/test_styles.xlsx")
	xl.Close()
}

func TestStyleSheets_reopen(t *testing.T) {
	xl, err := Open("./test_files/test_styles.xlsx")
	require.Nil(t, err)

	//after opening XLSX we must have only only saved styles
	addExistingStyles(xl, t)

	//after opening sheet we must have also 'typed' styles
	sheet := xl.Sheet(0) //we need to add types styles
	_ = sheet

	//check stored information, styles must be: default + new + typed
	checkStyles(xl, t)
	xl.Close()
}
