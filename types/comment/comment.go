package comment

type Comment struct {
}

type commentOption func(co *Comment)

//NewComment create and returns option set for comment
func New(options ...commentOption) *Comment {
	s := &Comment{}
	s.Set(options...)
	return s
}

//Set sets new options for option set
func (co *Comment) Set(options ...commentOption) {
	for _, o := range options {
		o(co)
	}
}

func (o *commentOption) Author(author string) commentOption {
	return func(co *Comment) {

	}
}

func (o *commentOption) Visible(visibility bool) commentOption {
	return func(co *Comment) {

	}
}

func (o *commentOption) XScale(scale float32) commentOption {
	return func(co *Comment) {

	}
}

func (o *commentOption) YScale(scale float32) commentOption {
	return func(co *Comment) {

	}
}

func (o *commentOption) XOffset(offset float32) commentOption {
	return func(co *Comment) {

	}
}

func (o *commentOption) YOffset(offset float32) commentOption {
	return func(co *Comment) {

	}
}

func (o *commentOption) Width(width float32) commentOption {
	return func(co *Comment) {

	}
}

func (o *commentOption) Height(height float32) commentOption {
	return func(co *Comment) {

	}
}

func (o *commentOption) StartCell(start int) commentOption {
	return func(co *Comment) {

	}
}

func (o *commentOption) StartRow(start int) commentOption {
	return func(co *Comment) {

	}
}

func (o *commentOption) StartCol(start int) commentOption {
	return func(co *Comment) {

	}
}

func (o *commentOption) Font(font interface{}) commentOption {
	return func(co *Comment) {

	}
}

func (o *commentOption) FontSize(size interface{}) commentOption {
	return func(co *Comment) {

	}
}

func (o *commentOption) Color(color interface{}) commentOption {
	return func(co *Comment) {

	}
}
