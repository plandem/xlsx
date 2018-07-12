package options

type columnOption func(co *ColumnOptions)

//ColumnOptions is a helper type to simplify process of settings options for column
type ColumnOptions struct {
	OutlineLevel uint8
	Collapsed    bool
	Phonetic     bool
	Hidden       bool
	Width        float32
}

//Column is a 'namespace' for all possible options for column
//
// Possible options are:
// OutlineLevel
// Collapsed
// Phonetic
// Hidden
// Width
var Column columnOption

//NewColumnOptions create and returns option set for column
func NewColumnOptions(options ...columnOption) *ColumnOptions {
	s := &ColumnOptions{}
	s.Set(options...)
	return s
}

//Set sets new options for option set
func (co *ColumnOptions) Set(options ...columnOption) {
	for _, o := range options {
		o(co)
	}
}

//OutlineLevel sets outline level of affected column. Range is 0 to 7.
func (o *columnOption) OutlineLevel(level uint8) columnOption {
	return func(co *ColumnOptions) {
		if level < 8 {
			co.OutlineLevel = level
		}
	}
}

//Collapsed sets flag indicating if the outlining of the affected column is in the collapsed state.
func (o *columnOption) Collapsed(collapsed bool) columnOption {
	return func(co *ColumnOptions) {
		co.Collapsed = collapsed
	}
}

//Phonetic sets flag indicating if the phonetic information should be displayed by default for the affected column of the worksheet.
func (o *columnOption) Phonetic(phonetic bool) columnOption {
	return func(co *ColumnOptions) {
		co.Phonetic = phonetic
	}
}

//Hidden sets flag indicating if the affected column are hidden on this worksheet.
func (o *columnOption) Hidden(hidden bool) columnOption {
	return func(co *ColumnOptions) {
		co.Hidden = hidden
	}
}

//Width sets the column width in the same units used by Excel which is: the number of characters in the default font. For more details: https://support.microsoft.com/en-us/kb/214123
func (o *columnOption) Width(width float32) columnOption {
	return func(co *ColumnOptions) {
		co.Width = width
	}
}
