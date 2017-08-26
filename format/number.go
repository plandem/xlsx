package format

type numberFormat struct {
	ID   int
	Code string
}

func NumberFormat(id int, code string) func(*StyleFormat) {
	return func(s *StyleFormat) {
		s.NumFormat.ID = id
		s.NumFormat.Code = code
	}
}
