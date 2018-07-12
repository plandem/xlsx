package options

type rowOption func(co *RowOptions)

//RowOptions is a helper type to simplify process of settings options for row
type RowOptions struct {
	OutlineLevel uint8
	Collapsed    bool
	Phonetic     bool
	Hidden       bool
	Height       float32
}

//Row is a 'namespace' for all possible options for row
//
// Possible options are:
// OutlineLevel
// Collapsed
// Phonetic
// Hidden
// Height
var Row rowOption

//NewRowOptions create and returns option set for row
func NewRowOptions(options ...rowOption) *RowOptions {
	s := &RowOptions{}
	s.Set(options...)
	return s
}

//Set sets new options for option set
func (ro *RowOptions) Set(options ...rowOption) {
	for _, o := range options {
		o(ro)
	}
}

//OutlineLevel set outlining level of the row, when outlining is on.
func (o *rowOption) OutlineLevel(level uint8) rowOption {
	return func(ro *RowOptions) {
		if level < 8 {
			ro.OutlineLevel = level
		}
	}
}

//Collapsed sets flag indicating if the outlining of the affected row is in the collapsed state.
func (o *rowOption) Collapsed(collapsed bool) rowOption {
	return func(ro *RowOptions) {
		ro.Collapsed = collapsed
	}
}

//Phonetic sets flag indicating if the phonetic information should be displayed by default for the affected row of the worksheet.
func (o *rowOption) Phonetic(phonetic bool) rowOption {
	return func(ro *RowOptions) {
		ro.Phonetic = phonetic
	}
}

//Hidden sets flag indicating if the affected row are hidden on this worksheet.
func (o *rowOption) Hidden(hidden bool) rowOption {
	return func(ro *RowOptions) {
		ro.Hidden = hidden
	}
}

//Height sets the row height in point size. There is no margin padding on row height.
func (o *rowOption) Height(height float32) rowOption {
	return func(ro *RowOptions) {
		ro.Height = height
	}
}
