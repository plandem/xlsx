// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package options

//Info hold advanced settings of row.
// N.B.: You should NOT mutate any value directly.
type Info struct {
	OutlineLevel uint8
	Collapsed    bool
	Phonetic     bool
	Hidden       bool
	Height       float32
	Format       interface{}
}

//Option is helper type to set options for row
type Option func(co *Info)

//New create and returns option set for row
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

//OutlineLevel set outlining level of the row, when outlining is on.
func OutlineLevel(level uint8) Option {
	return func(i *Info) {
		if level < 8 {
			i.OutlineLevel = level
		}
	}
}

//Collapsed sets flag indicating if the outlining of the affected row is in the collapsed state.
func Collapsed(collapsed bool) Option {
	return func(i *Info) {
		i.Collapsed = collapsed
	}
}

//Phonetic sets flag indicating if the phonetic information should be displayed by default for the affected row of the worksheet.
func Phonetic(phonetic bool) Option {
	return func(i *Info) {
		i.Phonetic = phonetic
	}
}

//Hidden sets flag indicating if the affected row are hidden on this worksheet.
func Hidden(hidden bool) Option {
	return func(i *Info) {
		i.Hidden = hidden
	}
}

//Height sets the row height in point size. There is no margin padding on row height.
func Height(height float32) Option {
	return func(i *Info) {
		i.Height = height
	}
}

//Styles sets style format to requested DirectStyleID or styles.Info
func Styles(s interface{}) Option {
	return func(i *Info) {
		i.Format = s
	}
}
