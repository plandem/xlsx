package format

type protectionOption byte

//Protection is a 'namespace' for all possible settings for protection
var Protection protectionOption

func (p *protectionOption) Hidden(s *StyleFormat) {
	s.styleInfo.Protection.Hidden = true
}

func (p *protectionOption) Locked(s *StyleFormat) {
	s.styleInfo.Protection.Locked = true
}
