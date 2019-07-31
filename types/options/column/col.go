// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package options

//Info hold advanced settings of column.
// N.B.: You should NOT mutate any value directly.
type Info struct {
	OutlineLevel uint8
	Collapsed    bool
	Phonetic     bool
	Hidden       bool
	Width        float32
	Format       interface{}
}

//Option is helper type to set options for comment
type Option func(co *Info)

//New create and returns option set for column
func New(settings ...Option) *Info {
	i := &Info{}
	i.Set(settings...)
	return i
}

//Set sets new options for option set
func (i *Info) Set(settings ...Option) {
	for _, o := range settings {
		o(i)
	}
}

//OutlineLevel sets outline level of affected column. Range is 0 to 7.
func OutlineLevel(level uint8) Option {
	return func(i *Info) {
		if level < 8 {
			i.OutlineLevel = level
		}
	}
}

//Collapsed sets flag indicating if the outlining of the affected column is in the collapsed state.
func Collapsed(collapsed bool) Option {
	return func(i *Info) {
		i.Collapsed = collapsed
	}
}

//Phonetic sets flag indicating if the phonetic information should be displayed by default for the affected column of the worksheet.
func Phonetic(phonetic bool) Option {
	return func(i *Info) {
		i.Phonetic = phonetic
	}
}

//Hidden sets flag indicating if the affected column are hidden on this worksheet.
func Hidden(hidden bool) Option {
	return func(i *Info) {
		i.Hidden = hidden
	}
}

//Width sets the column width in the same units used by Excel which is: the number of characters in the default font. For more details: https://support.microsoft.com/en-us/kb/214123
func Width(width float32) Option {
	return func(i *Info) {
		i.Width = width
	}
}

//Styles sets style format to requested DirectStyleID or styles.Info
func Styles(s interface{}) Option {
	return func(i *Info) {
		i.Format = s
	}
}
