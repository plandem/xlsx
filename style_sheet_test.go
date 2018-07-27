package xlsx

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
)

func addNewStyles(xl *Spreadsheet, t *testing.T) format.StyleRefID {
	require.NotNil(t, xl)

	require.Equal(t, 1, len(xl.styleSheet.xfIndex))
	require.Equal(t, 1, len(xl.styleSheet.borderIndex))
	require.Equal(t, 2, len(xl.styleSheet.fillIndex))
	require.Equal(t, 1, len(xl.styleSheet.fontIndex))
	require.Equal(t, 0, len(xl.styleSheet.numberIndex))

	//add font
	style := format.New(
		format.Font.Size(8),
		format.Font.Color("#FF1122"),
	)

	styleRef := xl.AddFormatting(style)
	require.Equal(t, format.StyleRefID(1), styleRef)
	require.Equal(t, 2, len(xl.styleSheet.xfIndex))
	require.Equal(t, 1, len(xl.styleSheet.borderIndex))
	require.Equal(t, 2, len(xl.styleSheet.fillIndex))
	require.Equal(t, 2, len(xl.styleSheet.fontIndex))
	require.Equal(t, 0, len(xl.styleSheet.numberIndex))

	//add fill
	style.Set(
		format.Fill.Type(format.PatternTypeLightGrid),
		//format.Fill.Color("#FF1122"),
		format.Fill.Background("#EFF142"),
	)

	styleRef = xl.AddFormatting(style)
	require.Equal(t, format.StyleRefID(2), styleRef)
	require.Equal(t, 3, len(xl.styleSheet.xfIndex))
	require.Equal(t, 1, len(xl.styleSheet.borderIndex))
	require.Equal(t, 3, len(xl.styleSheet.fillIndex))
	require.Equal(t, 2, len(xl.styleSheet.fontIndex))
	require.Equal(t, 0, len(xl.styleSheet.numberIndex))

	//add built-in number
	style.Set(
		format.NumberFormatID(45),
	)

	styleRef = xl.AddFormatting(style)
	require.Equal(t, format.StyleRefID(3), styleRef)
	require.Equal(t, 4, len(xl.styleSheet.xfIndex))
	require.Equal(t, 1, len(xl.styleSheet.borderIndex))
	require.Equal(t, 3, len(xl.styleSheet.fillIndex))
	require.Equal(t, 2, len(xl.styleSheet.fontIndex))
	require.Equal(t, 0, len(xl.styleSheet.numberIndex))

	//add custom number
	style.Set(
		format.NumberFormat(`$0.00" usd"`),
	)

	styleRef = xl.AddFormatting(style)
	require.Equal(t, format.StyleRefID(4), styleRef)
	require.Equal(t, 5, len(xl.styleSheet.xfIndex))
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
	require.Equal(t, format.StyleRefID(5), styleRef)
	require.Equal(t, 6, len(xl.styleSheet.xfIndex))
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
	require.Equal(t, format.StyleRefID(6), styleRef)
	require.Equal(t, 7, len(xl.styleSheet.xfIndex))
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
	require.Equal(t, format.StyleRefID(7), styleRef)
	require.Equal(t, 8, len(xl.styleSheet.xfIndex))
	require.Equal(t, 2, len(xl.styleSheet.borderIndex))
	require.Equal(t, 3, len(xl.styleSheet.fillIndex))
	require.Equal(t, 2, len(xl.styleSheet.fontIndex))
	require.Equal(t, 1, len(xl.styleSheet.numberIndex))

	return styleRef
}

func addExistingStyles(xl *Spreadsheet, t *testing.T) {
	require.NotNil(t, xl)

	style := format.New(
		format.Font.Size(8),
		format.Font.Color("#FF1122"), //font is red
		format.Fill.Type(format.PatternTypeLightGrid),
		//format.Fill.Color("#FF1122"), //color is red
		format.Fill.Background("#EFF142"), //background is green
		format.NumberFormat(`$0.00" usd"`),
		format.Border.Color("#1122FF"),                //border is blue
		format.Border.Type(format.BorderStyleDashDot), //
		format.Alignment.VAlign(format.VAlignBottom),  //
		format.Alignment.HAlign(format.HAlignFill),    //
		format.Protection.Hidden,
		format.Protection.Locked,
	)

	styleRef := xl.AddFormatting(style)
	require.Equal(t, format.StyleRefID(7), styleRef)
	require.Equal(t, 8, len(xl.styleSheet.xfIndex))
	require.Equal(t, 2, len(xl.styleSheet.borderIndex))
	require.Equal(t, 3, len(xl.styleSheet.fillIndex))
	require.Equal(t, 2, len(xl.styleSheet.fontIndex))
	require.Equal(t, 1, len(xl.styleSheet.numberIndex))
}

func checkStyles(xl *Spreadsheet, t *testing.T) {
	require.NotNil(t, xl)

	//validate stored fonts
	require.Equal(t, &[]*ml.Font{
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
	}, xl.styleSheet.ml.Fonts)

	//validate stored fills
	require.Equal(t, &[]*ml.Fill{
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
				Type: format.PatternTypeLightGrid,
				//Color:      &ml.Color{RGB: "FFEFF142"},
				Background: &ml.Color{RGB: "FFEFF142"},
				//Background: &ml.Color{Indexed: 4},
			},
		},
	}, xl.styleSheet.ml.Fills)

	//validate stored number
	require.Equal(t, &[]*ml.NumberFormat{
		{
			ID:   164,
			Code: `$0.00" usd"`,
		},
	}, xl.styleSheet.ml.NumberFormats)

	//validate stored border
	require.Equal(t, &[]*ml.Border{
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
	}, xl.styleSheet.ml.Borders)

	//validate stored Xf
	require.Equal(t, &[]*ml.StyleRef{
		//default xf
		{
			XfId:     0,
			FontId:   0,
			FillId:   0,
			BorderId: 0,
			NumFmtId: 0,
		},
		//new xf
		{
			XfId:      0,
			FontId:    1,
			FillId:    0,
			BorderId:  0,
			NumFmtId:  0,
			ApplyFont: true,
		},
		{
			XfId:      0,
			FontId:    1,
			FillId:    2,
			BorderId:  0,
			NumFmtId:  0,
			ApplyFont: true,
			ApplyFill: true,
		},
		{
			XfId:              0,
			FontId:            1,
			FillId:            2,
			BorderId:          0,
			NumFmtId:          45,
			ApplyFont:         true,
			ApplyFill:         true,
			ApplyNumberFormat: true,
		},
		{
			XfId:              0,
			FontId:            1,
			FillId:            2,
			BorderId:          0,
			NumFmtId:          164,
			ApplyFont:         true,
			ApplyFill:         true,
			ApplyNumberFormat: true,
		},
		{
			XfId:              0,
			FontId:            1,
			FillId:            2,
			BorderId:          1,
			NumFmtId:          164,
			ApplyFont:         true,
			ApplyFill:         true,
			ApplyNumberFormat: true,
			ApplyBorder:       true,
		},
		{
			XfId:              0,
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
		{
			XfId:              0,
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
	}, xl.styleSheet.ml.CellXfs)
}

func TestStyleSheets(t *testing.T) {
	testList := []struct {
		name     string
		onBefore func(fileName string, tt *testing.T) *Spreadsheet
		onAfter  func(fileName string, xl *Spreadsheet, tt *testing.T)
	}{
		{
			"create",
			func(fileName string, tt *testing.T) *Spreadsheet {
				xl := New()
				styleRef := addNewStyles(xl, tt)
				sheet := xl.AddSheet("test sheet")
				sheet.Row(0).SetFormatting(styleRef)
				//sheet.Cell(10, 10).SetFormatting(styleRef)
				return xl
			},
			func(fileName string, xl *Spreadsheet, tt *testing.T) {
				xl.SaveAs(fileName)
			},
		},
		{
			"reopen",
			func(fileName string, tt *testing.T) *Spreadsheet {
				xl, err := Open(fileName)
				require.Nil(tt, err)
				return xl
			},
			func(fileName string, xl *Spreadsheet, tt *testing.T) {},
		},
	}

	for _, info := range testList {
		t.Run(info.name, func(tt *testing.T) {
			//create/load
			xl := info.onBefore("./test_files/test_styles.xlsx", tt)

			//try to add already existing styles
			addExistingStyles(xl, tt)

			//check stored information
			checkStyles(xl, tt)

			//safe
			info.onAfter("./test_files/test_styles.xlsx", xl, tt)
			xl.Close()
		})
	}
}
