package format

type numberFormat struct {
	ID   int
	Code string
}

//NumberFormat is option to update StyleFormat with provided id and code
func NumberFormat(id int, code string) func(*StyleFormat) {
	return func(s *StyleFormat) {
		s.NumFormat.ID = id
		s.NumFormat.Code = code
	}
}
