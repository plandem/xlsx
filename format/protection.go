package format

type protection struct {
	Locked bool
	Hidden bool
}

type protectionOption byte

//Protection is a 'namespace' for all possible settings for protection
var Protection protectionOption

func (p *protectionOption) Hidden(s *StyleFormat) {
	s.Protection.Hidden = true
}

func (p *protectionOption) Locked(s *StyleFormat) {
	s.Protection.Locked = true
}
