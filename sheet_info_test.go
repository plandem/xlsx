package xlsx

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/options"
	"github.com/plandem/xlsx/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSheetInfo(t *testing.T) {
	require.Equal(t, true, isCellEmpty(nil))
	require.Equal(t, true, isCellEmpty(&ml.Cell{}))
	require.Equal(t, true, isCellEmpty(&ml.Cell{Ref: "A10"}))
	require.Equal(t, false, isCellEmpty(&ml.Cell{Ref: "A10", Value: "1"}))

	require.Equal(t, true, isRowEmpty(nil))
	require.Equal(t, true, isRowEmpty(&ml.Row{}))
	require.Equal(t, false, isRowEmpty(&ml.Row{Cells: []*ml.Cell{{}}}))
	require.Equal(t, false, isRowEmpty(&ml.Row{CustomHeight: true}))

	xl, err := Open("./test_files/example_simple.xlsx")
	if err != nil {
		panic(err)
	}

	defer xl.Close()
	sheet := xl.Sheet(0)

	//test options
	o := options.NewSheetOptions(
		options.Sheet.Visibility(types.VisibilityTypeVeryHidden),
	)

	require.Equal(t, types.VisibilityType(0), xl.workbook.ml.Sheets[0].State)
	sheet.Set(o)
	require.Equal(t, types.VisibilityTypeVeryHidden, xl.workbook.ml.Sheets[0].State)

	//test set active
	require.Equal(t, 0, (*xl.workbook.ml.BookViews)[0].ActiveTab)
	sheet = xl.AddSheet("test")
	sheet.SetActive()
	require.Equal(t, 1, (*xl.workbook.ml.BookViews)[0].ActiveTab)
}
