package comment

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

func Author(author string) Option {
	return func(i *Info) {
		i.Author = author
	}
}

func Visible(visible bool) Option {
	return func(i *Info) {
		i.Visible = visible
	}
}

func XScale(scale float32) Option {
	return func(co *Info) {
		co.XScale = scale
	}
}

func YScale(scale float32) Option {
	return func(i *Info) {
		i.YScale = scale
	}
}

func XOffset(offset float32) Option {
	return func(i *Info) {
		i.XOffset = offset
	}
}

func YOffset(offset float32) Option {
	return func(i *Info) {
		i.YOffset = offset
	}
}

func Width(width float32) Option {
	return func(i *Info) {
		i.Width = width
	}
}

func Height(height float32) Option {
	return func(i *Info) {
		i.Height = height
	}
}

func Background(rgb string) Option {
	return func(i *Info) {
		i.Background = rgb
	}
}

func Shadow(rgb string) Option {
	return func(i *Info) {
		i.Shadow = rgb
	}
}

func Text(parts ...interface{}) Option {
	return func(i *Info) {
		i.Text = parts
	}
}
