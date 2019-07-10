// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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
