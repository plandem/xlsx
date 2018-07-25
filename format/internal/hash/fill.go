package hash

import (
	"github.com/plandem/xlsx/internal/ml"
	"strconv"
	"strings"
)

//Fill returns string with all values of fill
func Fill(fill *ml.Fill) Key {
	var f ml.Fill

	if fill == nil {
		f = ml.Fill{}
	} else {
		//we don't want to mutate original fill
		f = *fill
	}

	if f.Pattern == nil {
		f.Pattern = &ml.PatternFill{}
	}

	if f.Gradient == nil {
		f.Gradient = &ml.GradientFill{}
	}

	result := []string{
		strconv.FormatInt(int64(f.Pattern.Type), 10),
		string(Color(f.Pattern.Color)),
		string(Color(f.Pattern.Background)),
		strconv.FormatInt(int64(f.Gradient.Type), 10),
		strconv.FormatFloat(float64(f.Gradient.Degree), 'f', -1, 64),
		strconv.FormatFloat(float64(f.Gradient.Left), 'f', -1, 64),
		strconv.FormatFloat(float64(f.Gradient.Right), 'f', -1, 64),
		strconv.FormatFloat(float64(f.Gradient.Top), 'f', -1, 64),
		strconv.FormatFloat(float64(f.Gradient.Bottom), 'f', -1, 64),
	}

	for _, stop := range f.Gradient.Stop {
		result = append(result,
			strconv.FormatFloat(float64(stop.Position), 'f', -1, 64),
			string(Color(stop.Color)),
		)
	}

	return Key(strings.Join(result, ":"))
}
