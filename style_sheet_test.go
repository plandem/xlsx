// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/plandem/xlsx/internal/ml"
)

func addNewStyles(xl *Spreadsheet, t *testing.T) styles.DirectStyleID {
	require.NotNil(t, xl)

	require.Equal(t, 1, xl.styleSheet.directStyleIndex.Count())
	require.Equal(t, 1, xl.styleSheet.borderIndex.Count())
	require.Equal(t, 2, xl.styleSheet.fillIndex.Count())
	require.Equal(t, 1, xl.styleSheet.fontIndex.Count())
	require.Equal(t, 0, xl.styleSheet.numberIndex.Count())

	//add font
	style := styles.New(
		styles.Font.Size(8),
		styles.Font.Color("#FF1122"),
	)

	styleRef := xl.AddStyles(style)
	require.Equal(t, styles.DirectStyleID(1), styleRef)
	require.Equal(t, 2, xl.styleSheet.directStyleIndex.Count())
	require.Equal(t, 1, xl.styleSheet.borderIndex.Count())
	require.Equal(t, 2, xl.styleSheet.fillIndex.Count())
	require.Equal(t, 2, xl.styleSheet.fontIndex.Count())
	require.Equal(t, 0, xl.styleSheet.numberIndex.Count())

	//add fill
	style.Set(
		styles.Fill.Type(styles.PatternTypeLightGrid),
		styles.Fill.Background("#EFF142"),
	)

	styleRef = xl.AddStyles(style)
	require.Equal(t, styles.DirectStyleID(2), styleRef)
	require.Equal(t, 3, xl.styleSheet.directStyleIndex.Count())
	require.Equal(t, 1, xl.styleSheet.borderIndex.Count())
	require.Equal(t, 3, xl.styleSheet.fillIndex.Count())
	require.Equal(t, 2, xl.styleSheet.fontIndex.Count())
	require.Equal(t, 0, xl.styleSheet.numberIndex.Count())

	//add built-in number
	style.Set(
		styles.NumberFormatID(45),
	)

	styleRef = xl.AddStyles(style)
	require.Equal(t, styles.DirectStyleID(3), styleRef)
	require.Equal(t, 4, xl.styleSheet.directStyleIndex.Count())
	require.Equal(t, 1, xl.styleSheet.borderIndex.Count())
	require.Equal(t, 3, xl.styleSheet.fillIndex.Count())
	require.Equal(t, 2, xl.styleSheet.fontIndex.Count())
	require.Equal(t, 0, xl.styleSheet.numberIndex.Count())

	//add custom number
	style.Set(
		styles.NumberFormat(`$0.00" usd"`),
	)

	styleRef = xl.AddStyles(style)
	require.Equal(t, styles.DirectStyleID(4), styleRef)
	require.Equal(t, 5, xl.styleSheet.directStyleIndex.Count())
	require.Equal(t, 1, xl.styleSheet.borderIndex.Count())
	require.Equal(t, 3, xl.styleSheet.fillIndex.Count())
	require.Equal(t, 2, xl.styleSheet.fontIndex.Count())
	require.Equal(t, 1, xl.styleSheet.numberIndex.Count())

	//add border
	style.Set(
		styles.Border.Color("#1122FF"),
		styles.Border.Type(styles.BorderStyleDashDot),
	)

	styleRef = xl.AddStyles(style)
	require.Equal(t, styles.DirectStyleID(5), styleRef)
	require.Equal(t, 6, xl.styleSheet.directStyleIndex.Count())
	require.Equal(t, 2, xl.styleSheet.borderIndex.Count())
	require.Equal(t, 3, xl.styleSheet.fillIndex.Count())
	require.Equal(t, 2, xl.styleSheet.fontIndex.Count())
	require.Equal(t, 1, xl.styleSheet.numberIndex.Count())

	//add alignment
	style.Set(
		styles.Alignment.VAlign(styles.VAlignBottom),
		styles.Alignment.HAlign(styles.HAlignFill),
	)

	styleRef = xl.AddStyles(style)
	require.Equal(t, styles.DirectStyleID(6), styleRef)
	require.Equal(t, 7, xl.styleSheet.directStyleIndex.Count())
	require.Equal(t, 2, xl.styleSheet.borderIndex.Count())
	require.Equal(t, 3, xl.styleSheet.fillIndex.Count())
	require.Equal(t, 2, xl.styleSheet.fontIndex.Count())
	require.Equal(t, 1, xl.styleSheet.numberIndex.Count())

	//add protection
	style.Set(
		styles.Protection.Hidden,
		styles.Protection.Locked,
	)

	styleRef = xl.AddStyles(style)
	require.Equal(t, styles.DirectStyleID(7), styleRef)
	require.Equal(t, 8, xl.styleSheet.directStyleIndex.Count())
	require.Equal(t, 2, xl.styleSheet.borderIndex.Count())
	require.Equal(t, 3, xl.styleSheet.fillIndex.Count())
	require.Equal(t, 2, xl.styleSheet.fontIndex.Count())
	require.Equal(t, 1, xl.styleSheet.numberIndex.Count())

	return styleRef
}

func addExistingStyles(xl *Spreadsheet, t *testing.T) {
	require.NotNil(t, xl)

	style := styles.New(
		styles.Font.Size(8),
		styles.Font.Color("#FF1122"),
		styles.Fill.Type(styles.PatternTypeLightGrid),
		styles.Fill.Background("#EFF142"),
		styles.NumberFormat(`$0.00" usd"`),
		styles.Border.Color("#1122FF"),
		styles.Border.Type(styles.BorderStyleDashDot),
		styles.Alignment.VAlign(styles.VAlignBottom),
		styles.Alignment.HAlign(styles.HAlignFill),
		styles.Protection.Hidden,
		styles.Protection.Locked,
	)

	styleRef := xl.AddStyles(style)
	require.Equal(t, styles.DirectStyleID(7), styleRef)
	require.Equal(t, 14, xl.styleSheet.directStyleIndex.Count())
	require.Equal(t, 2, xl.styleSheet.borderIndex.Count())
	require.Equal(t, 3, xl.styleSheet.fillIndex.Count())
	require.Equal(t, 2, xl.styleSheet.fontIndex.Count())
	require.Equal(t, 1, xl.styleSheet.numberIndex.Count())
}

func checkStyles(xl *Spreadsheet, t *testing.T) {
	require.NotNil(t, xl)

	//validate stored fonts
	require.Equal(t, []*ml.Font{
		//default font
		{
			Name:   "Calibri",
			Family: styles.FontFamilySwiss,
			Size:   11,
			Scheme: styles.FontSchemeMinor,
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
				Type: styles.PatternTypeNone,
			},
		},
		{
			Pattern: &ml.PatternFill{
				Type: styles.PatternTypeGray125,
			},
		},
		//new fill
		{
			Pattern: &ml.PatternFill{
				Type:       styles.PatternTypeLightGrid,
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
			Left:   &ml.BorderSegment{Type: styles.BorderStyleDashDot, Color: &ml.Color{RGB: "FF1122FF"}},
			Right:  &ml.BorderSegment{Type: styles.BorderStyleDashDot, Color: &ml.Color{RGB: "FF1122FF"}},
			Top:    &ml.BorderSegment{Type: styles.BorderStyleDashDot, Color: &ml.Color{RGB: "FF1122FF"}},
			Bottom: &ml.BorderSegment{Type: styles.BorderStyleDashDot, Color: &ml.Color{RGB: "FF1122FF"}},
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
					Vertical:   styles.VAlignBottom,
					Horizontal: styles.HAlignFill,
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
					Vertical:   styles.VAlignBottom,
					Horizontal: styles.HAlignFill,
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
	sheet.Row(0).SetStyles(styleRef)

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
