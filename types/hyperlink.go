package types

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
)

type HyperlinkInfo struct {
	ml.Hyperlink
	*format.StyleFormat
}

type option func(o *HyperlinkInfo)

func NewHyperlink(options ...option) *HyperlinkInfo {
	s := &HyperlinkInfo{}
	s.Set(options...)
	return s
}

//Set sets new options for hyperlink
func (s *HyperlinkInfo) Set(options ...option) {
	for _, o := range options {
		o(s)
	}
}

//
//import (
//	"github.com/plandem/xlsx/format"
//)
//
//type hyperlink struct {
//	//location string
//	File string
//	Url string
//	Tooltip  string
//	Display  string
//	Style *format.StyleFormat
//}
//
//type hyperlinkOption func(o *hyperlink)
//
////Hyperlink is a 'namespace' for all possible settings for hyperlink
//var Hyperlink hyperlinkOption
//
////NewHyperlink creates and returns hyperlink object with requested options
//func NewHyperlink(options ...hyperlinkOption) *hyperlink {
//	h := &hyperlink{}
//	h.Set(options...)
//	return h
//}
//
////Set sets new options for hyperlink
//func (h *hyperlink) Set(options ...hyperlinkOption) {
//	for _, o := range options {
//		o(h)
//	}
//}

//
//func (ho *hyperlinkOption) A(v int) hyperlinkOption {
//	return func(h *hyperlink) {
//		h.Horizontal = v
//	}
//}
