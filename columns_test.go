package xlsx

import (
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestColumns_Delete(t *testing.T) {
	xl := New()
	defer xl.Close()

	xl.AddSheet("The first sheet")
	cols := xl.sheets[0].columns

	//non grouped columns
	require.NotNil(t, cols.sheet.ml.Cols)
	cols.Resolve(0)
	require.EqualValues(t, &[]*ml.Col{
		{
			Min: 1,
			Max: 1,
		},
	}, cols.sheet.ml.Cols)
	cols.Delete(0)
	require.EqualValues(t, &[]*ml.Col{}, cols.sheet.ml.Cols)

	cols.Resolve(0)
	cols.Resolve(5)
	require.EqualValues(t, &[]*ml.Col{
		{
			Min: 1,
			Max: 1,
		},
		{
			Min: 6,
			Max: 6,
		},
	}, cols.sheet.ml.Cols)
	cols.Delete(0)
	require.EqualValues(t, &[]*ml.Col{
		{
			Min: 6,
			Max: 6,
		},
	}, cols.sheet.ml.Cols)

	//grouped columns
	cols.sheet.ml.Cols = &[]*ml.Col{
		{
			Min:   1,
			Max:   100,
			Width: 32,
		},
	}
	cols.Delete(0)
	require.EqualValues(t, &[]*ml.Col{
		{
			Min:     1,
			Max:     99,
			Width:   32,
		},
	}, cols.sheet.ml.Cols)
	cols.Resolve(0)
	require.EqualValues(t, &[]*ml.Col{
		{
			Min:     1,
			Max:     99,
			Width:   32,
		},
		{
			Min:   1,
			Max:   1,
			Width: 32,
		},
	}, cols.sheet.ml.Cols)
	cols.Delete(0)
	require.EqualValues(t, &[]*ml.Col{
		{
			Min:     1,
			Max:     98,
			Width:   32,
		},
	}, cols.sheet.ml.Cols)
	cols.Resolve(0)
	cols.Resolve(5)
	require.EqualValues(t, &[]*ml.Col{
		{
			Min:     1,
			Max:     98,
			Width:   32,
		},
		{
			Min:   1,
			Max:   1,
			Width: 32,
		},
		{
			Min:   6,
			Max:   6,
			Width: 32,
		},
	}, cols.sheet.ml.Cols)
	cols.Delete(5)
	require.EqualValues(t, &[]*ml.Col{
		{
			Min:     1,
			Max:     97,
			Width:   32,
		},
		{
			Min:   1,
			Max:   1,
			Width: 32,
		},
	}, cols.sheet.ml.Cols)
}

func TestColumns_Resolve(t *testing.T) {
	xl := New()
	defer xl.Close()

	xl.AddSheet("The first sheet")
	cols := xl.sheets[0].columns

	//non grouped columns
	require.NotNil(t, cols.sheet.ml.Cols)
	cols.Resolve(0)
	require.EqualValues(t, &[]*ml.Col{
		{
			Min: 1,
			Max: 1,
		},
	}, cols.sheet.ml.Cols)
	cols.Resolve(0)
	require.EqualValues(t, &[]*ml.Col{
		{
			Min: 1,
			Max: 1,
		},
	}, cols.sheet.ml.Cols)
	cols.Resolve(5)
	require.EqualValues(t, &[]*ml.Col{
		{
			Min: 1,
			Max: 1,
		},
		{
			Min: 6,
			Max: 6,
		},
	}, cols.sheet.ml.Cols)

	//grouped columns
	cols.sheet.ml.Cols = &[]*ml.Col{
		{
			Min:   1,
			Max:   100,
			Width: 32,
		},
	}
	cols.Resolve(0)
	require.EqualValues(t, &[]*ml.Col{
		{
			Min:     1,
			Max:     100,
			Width:   32,
		},
		{
			Min:   1,
			Max:   1,
			Width: 32,
		},
	}, cols.sheet.ml.Cols)

	cols.Resolve(5)
	require.EqualValues(t, &[]*ml.Col{
		{
			Min:     1,
			Max:     100,
			Width:   32,
		},
		{
			Min:   1,
			Max:   1,
			Width: 32,
		},
		{
			Min:   6,
			Max:   6,
			Width: 32,
		},
	}, cols.sheet.ml.Cols)
}

func TestColumns_pack(t *testing.T) {
	xl := New()
	defer xl.Close()

	xl.AddSheet("The first sheet")
	cols := xl.sheets[0].columns

	//empty cols should be removed
	cols.Resolve(0)
	cols.Resolve(5)
	cols.Resolve(10)
	require.Nil(t, cols.pack())

	//serial cols with same settings should be merged
	colsIdx := []int{10, 2, 0, 5, 1}
	for _, idx := range colsIdx {
		c := cols.Resolve(idx)
		c.Width = 100
	}

	require.NotNil(t, cols.pack())
	require.EqualValues(t, &[]*ml.Col{
		{
			Min:   1,
			Max:   3,
			Width: 100,
		},
		{
			Min:   6,
			Max:   6,
			Width: 100,
		},
		{
			Min:   11,
			Max:   11,
			Width: 100,
		},
	}, cols.sheet.ml.Cols)

	//serial and grouped, same data
	cols.sheet.ml.Cols = &[]*ml.Col{
		{
			Min:   1,
			Max:   10,
			Width: 100,
		},
	}
	colsIdx = []int{2, 3, 10, 11, 15, 14}
	for _, idx := range colsIdx {
		c := cols.Resolve(idx)
		c.Width = 100
	}
	require.NotNil(t, cols.pack())
	require.EqualValues(t, &[]*ml.Col{
		{
			Min:   1,
			Max:   12,
			Width: 100,
		},
		{
			Min:   15,
			Max:   16,
			Width: 100,
		},
	}, cols.sheet.ml.Cols)

	//serial and grouped, different data
	cols.sheet.ml.Cols = &[]*ml.Col{
		{
			Min:   1,
			Max:   10,
			Width: 100,
		},
	}
	colsIdx = []int{2, 3, 10, 11, 15, 14}
	for _, idx := range colsIdx {
		c := cols.Resolve(idx)
		c.Width = 200
	}
	require.NotNil(t, cols.pack())
	require.EqualValues(t, &[]*ml.Col{
		{
			Min:   1,
			Max:   2,
			Width: 100,
		},
		{
			Min:   3,
			Max:   4,
			Width: 200,
		},
		{
			Min:   5,
			Max:   10,
			Width: 100,
		},
		{
			Min:   11,
			Max:   12,
			Width: 200,
		},
		{
			Min:   15,
			Max:   16,
			Width: 200,
		},
	}, cols.sheet.ml.Cols)
}
