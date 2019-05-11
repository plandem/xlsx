package styles

type protectionOption byte

//Protection is a 'namespace' for all possible settings for protection
var Protection protectionOption

func (p *protectionOption) Hidden(s *Info) {
	s.styleInfo.Protection.Hidden = true
}

func (p *protectionOption) Locked(s *Info) {
	s.styleInfo.Protection.Locked = true
}
