package options_test

import (
	"github.com/plandem/xlsx/options"
	"github.com/plandem/xlsx/types"
)

func ExampleNewColumnOptions() {
	o := options.NewColumnOptions(
		options.Column.OutlineLevel(5),
		options.Column.Hidden(true),
		options.Column.Phonetic(true),
		options.Column.Width(45.5),
	)

	_ = o
}

func ExampleNewRowOptions() {
	o := options.NewRowOptions(
		options.Row.OutlineLevel(5),
		options.Row.Hidden(true),
		options.Row.Phonetic(true),
		options.Row.Height(45.5),
	)

	_ = o
}

func ExampleNewSheetOptions() {
	o := options.NewSheetOptions(
		options.Sheet.Visibility(types.VisibilityTypeVisible),
	)

	_ = o
}
