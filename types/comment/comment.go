// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package comment

//Info hold advanced settings of comment
type Info struct {
	Width      float32
	Height     float32
	XScale     float32
	YScale     float32
	XOffset    float32
	YOffset    float32
	Author     string
	Background string
	Shadow     string
	Visible    bool
	Text       []interface{}
}

//Option is helper type to set options for comment
type Option func(o *Info)

//New create and returns option set for comment
func New(options ...Option) *Info {
	s := &Info{
		Background: "#ffffe1",
		Width:      128,
		Height:     74,
		XScale:     1,
		YScale:     1,
	}

	s.Set(options...)
	return s
}

//Set sets new options for option set
func (i *Info) Set(options ...Option) {
	for _, o := range options {
		o(i)
	}
}

//Author adds author information to comment
func Author(author string) Option {
	return func(i *Info) {
		i.Author = author
	}
}

//Visible sets visibility of comment
func Visible(visible bool) Option {
	return func(i *Info) {
		i.Visible = visible
	}
}

//XScale sets x-scaling value for comment
func XScale(scale float32) Option {
	return func(i *Info) {
		i.XScale = scale
	}
}

//YScale sets y-scaling value for comment
func YScale(scale float32) Option {
	return func(i *Info) {
		i.YScale = scale
	}
}

//XOffset sets x-offset value for comment
func XOffset(offset float32) Option {
	return func(i *Info) {
		i.XOffset = offset
	}
}

//YOffset sets y-offset value for comment
func YOffset(offset float32) Option {
	return func(i *Info) {
		i.YOffset = offset
	}
}

//Width sets width of comment
func Width(width float32) Option {
	return func(i *Info) {
		i.Width = width
	}
}

//Height sets height of comment
func Height(height float32) Option {
	return func(i *Info) {
		i.Height = height
	}
}

//Background sets background color for comment
func Background(rgb string) Option {
	return func(i *Info) {
		i.Background = rgb
	}
}

//Shadow sets shadow color for comment
func Shadow(rgb string) Option {
	return func(i *Info) {
		i.Shadow = rgb
	}
}

//Text sets rich text of comment
func Text(parts ...interface{}) Option {
	return func(i *Info) {
		i.Text = parts
	}
}
