package hash_test

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/format/internal/hash"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAlignment(t *testing.T) {
	require.Equal(t, "0:0:0:false:0:0:false:false:0", hash.Alignment(nil))
	require.Equal(t, "0:3:0:false:0:0:false:false:0", hash.Alignment(&ml.CellAlignment{Vertical: format.VAlignBottom}))
	require.Equal(t, "8:0:0:false:0:0:false:false:0", hash.Alignment(&ml.CellAlignment{Horizontal: format.HAlignDistributed}))
	require.Equal(t, "0:0:90:false:0:0:false:false:0", hash.Alignment(&ml.CellAlignment{TextRotation: 90}))
	require.Equal(t, "0:0:0:true:0:0:false:false:0", hash.Alignment(&ml.CellAlignment{WrapText: true}))
	require.Equal(t, "0:0:0:false:10:0:false:false:0", hash.Alignment(&ml.CellAlignment{Indent: 10}))
	require.Equal(t, "0:0:0:false:0:12:false:false:0", hash.Alignment(&ml.CellAlignment{RelativeIndent: 12}))
	require.Equal(t, "0:0:0:false:0:0:true:false:0", hash.Alignment(&ml.CellAlignment{JustifyLastLine: true}))
	require.Equal(t, "0:0:0:false:0:0:false:true:0", hash.Alignment(&ml.CellAlignment{ShrinkToFit: true}))
	require.Equal(t, "0:0:0:false:0:0:false:false:13", hash.Alignment(&ml.CellAlignment{ReadingOrder: 13}))
	require.Equal(t, "8:3:90:true:10:12:true:true:13", hash.Alignment(&ml.CellAlignment{
		Horizontal:      format.HAlignDistributed,
		Vertical:        format.VAlignBottom,
		ReadingOrder:    13,
		TextRotation:    90,
		ShrinkToFit:     true,
		JustifyLastLine: true,
		WrapText:        true,
		Indent:          10,
		RelativeIndent:  12,
	}))
}

func TestBorder(t *testing.T) {
	b := ml.Border{}
	require.Equal(t, "false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false", hash.Border(nil))
	require.Equal(t, "false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false", hash.Border(&b))
	require.Equal(t, &ml.Border{}, &b)

	require.Equal(t, "false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false", hash.Border(&ml.Border{Left: &ml.BorderSegment{}}))
	require.Equal(t, "false:112233:0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false", hash.Border(&ml.Border{Left: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}))
	require.Equal(t, "false::0:::mediumDashDot:false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false", hash.Border(&ml.Border{Left: &ml.BorderSegment{Type: format.BorderStyleMediumDashDot}}))

	require.Equal(t, "false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false", hash.Border(&ml.Border{Right: &ml.BorderSegment{}}))
	require.Equal(t, "false::0::::false:112233:0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false", hash.Border(&ml.Border{Right: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}))
	require.Equal(t, "false::0::::false::0:::mediumDashDot:false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false", hash.Border(&ml.Border{Right: &ml.BorderSegment{Type: format.BorderStyleMediumDashDot}}))

	require.Equal(t, "false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false", hash.Border(&ml.Border{Top: &ml.BorderSegment{}}))
	require.Equal(t, "false::0::::false::0::::false:112233:0::::false::0::::false::0::::false::0::::false::0::::false:false:false", hash.Border(&ml.Border{Top: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}))
	require.Equal(t, "false::0::::false::0::::false::0:::mediumDashDot:false::0::::false::0::::false::0::::false::0::::false:false:false", hash.Border(&ml.Border{Top: &ml.BorderSegment{Type: format.BorderStyleMediumDashDot}}))

	require.Equal(t, "false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false", hash.Border(&ml.Border{Bottom: &ml.BorderSegment{}}))
	require.Equal(t, "false::0::::false::0::::false::0::::false:112233:0::::false::0::::false::0::::false::0::::false:false:false", hash.Border(&ml.Border{Bottom: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}))
	require.Equal(t, "false::0::::false::0::::false::0::::false::0:::mediumDashDot:false::0::::false::0::::false::0::::false:false:false", hash.Border(&ml.Border{Bottom: &ml.BorderSegment{Type: format.BorderStyleMediumDashDot}}))

	require.Equal(t, "false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false", hash.Border(&ml.Border{Diagonal: &ml.BorderSegment{}}))
	require.Equal(t, "false::0::::false::0::::false::0::::false::0::::false:112233:0::::false::0::::false::0::::false:false:false", hash.Border(&ml.Border{Diagonal: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}))
	require.Equal(t, "false::0::::false::0::::false::0::::false::0::::false::0:::mediumDashDot:false::0::::false::0::::false:false:false", hash.Border(&ml.Border{Diagonal: &ml.BorderSegment{Type: format.BorderStyleMediumDashDot}}))

	require.Equal(t, "false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false", hash.Border(&ml.Border{Vertical: &ml.BorderSegment{}}))
	require.Equal(t, "false::0::::false::0::::false::0::::false::0::::false::0::::false:112233:0::::false::0::::false:false:false", hash.Border(&ml.Border{Vertical: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}))
	require.Equal(t, "false::0::::false::0::::false::0::::false::0::::false::0::::false::0:::mediumDashDot:false::0::::false:false:false", hash.Border(&ml.Border{Vertical: &ml.BorderSegment{Type: format.BorderStyleMediumDashDot}}))

	require.Equal(t, "false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:false", hash.Border(&ml.Border{Horizontal: &ml.BorderSegment{}}))
	require.Equal(t, "false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:112233:0::::false:false:false", hash.Border(&ml.Border{Horizontal: &ml.BorderSegment{Color: &ml.Color{RGB: "112233"}}}))
	require.Equal(t, "false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0:::mediumDashDot:false:false:false", hash.Border(&ml.Border{Horizontal: &ml.BorderSegment{Type: format.BorderStyleMediumDashDot}}))

	require.Equal(t, "false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::true:false:false", hash.Border(&ml.Border{DiagonalUp: true}))
	require.Equal(t, "false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:true:false", hash.Border(&ml.Border{DiagonalDown: true}))
	require.Equal(t, "false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false::0::::false:false:true", hash.Border(&ml.Border{Outline: true}))

	require.Equal(t, "false:777777:0:::slantDashDot:false:666666:0:::thin:false:555555:0:::thick:false:444444:0:::hair:false:333333:0:::dotted:false:222222:0:::dashDot:false:111111:0:::medium:true:true:true", hash.Border(&ml.Border{
		Outline: true,
		DiagonalDown: true,
		DiagonalUp: true,
		Horizontal: &ml.BorderSegment{Color: &ml.Color{RGB: "111111"}, Type: format.BorderStyleMedium},
		Vertical: &ml.BorderSegment{Color: &ml.Color{RGB: "222222"}, Type: format.BorderStyleDashDot},
		Diagonal: &ml.BorderSegment{Color: &ml.Color{RGB: "333333"}, Type: format.BorderStyleDotted},
		Bottom: &ml.BorderSegment{Color: &ml.Color{RGB: "444444"}, Type: format.BorderStyleHair},
		Top: &ml.BorderSegment{Color: &ml.Color{RGB: "555555"}, Type: format.BorderStyleThick},
		Right: &ml.BorderSegment{Color: &ml.Color{RGB: "666666"}, Type: format.BorderStyleThin},
		Left: &ml.BorderSegment{Color: &ml.Color{RGB: "777777"}, Type: format.BorderStyleSlantDashDot},
	}))
}

func TestColor(t *testing.T) {
	index1 := 1
	index2 := 2

	require.Equal(t, "false::0::", hash.Color(nil))
	require.Equal(t, "false::0::", hash.Color(&ml.Color{}))
	require.Equal(t, "true::0::", hash.Color(&ml.Color{Auto: true}))
	require.Equal(t, "false:112233:0::", hash.Color(&ml.Color{RGB: "112233"}))
	require.Equal(t, "false::1::", hash.Color(&ml.Color{Tint: 1}))
	require.Equal(t, "false::0:1:", hash.Color(&ml.Color{Indexed: &index1}))
	require.Equal(t, "false::0::2", hash.Color(&ml.Color{Theme: &index2}))
	require.Equal(t, "true:112233:1:1:2", hash.Color(&ml.Color{
		Indexed: &index1,
		RGB:     "112233",
		Auto:    true,
		Theme:   &index2,
		Tint:    1,
	}))
}

func TestFill(t *testing.T) {

}

func TestFont(t *testing.T) {
	require.Equal(t, ":0:0:false:false:false:false:false:false:false::0:::0:::", hash.Font(nil))
	require.Equal(t, ":0:0:false:false:false:false:false:false:false::0:::0:::", hash.Font(&ml.Font{}))
	require.Equal(t, "calibri:0:0:false:false:false:false:false:false:false::0:::0:::", hash.Font(&ml.Font{Name: "calibri"}))
	require.Equal(t, ":1:0:false:false:false:false:false:false:false::0:::0:::", hash.Font(&ml.Font{Charset: 1}))
	require.Equal(t, ":0:1:false:false:false:false:false:false:false::0:::0:::", hash.Font(&ml.Font{Family: format.FontFamilyRoman}))
	require.Equal(t, ":0:0:false:false:false:false:false:false:false::0:::0:::", hash.Font(&ml.Font{Bold: false}))
	require.Equal(t, ":0:0:true:false:false:false:false:false:false::0:::0:::", hash.Font(&ml.Font{Bold: true}))
	require.Equal(t, ":0:0:false:true:false:false:false:false:false::0:::0:::", hash.Font(&ml.Font{Italic: true}))
	require.Equal(t, ":0:0:false:false:true:false:false:false:false::0:::0:::", hash.Font(&ml.Font{Strike: true}))
	require.Equal(t, ":0:0:false:false:false:true:false:false:false::0:::0:::", hash.Font(&ml.Font{Shadow: true}))
	require.Equal(t, ":0:0:false:false:false:false:true:false:false::0:::0:::", hash.Font(&ml.Font{Condense: true}))
	require.Equal(t, ":0:0:false:false:false:false:false:true:false::0:::0:::", hash.Font(&ml.Font{Extend: true}))
	require.Equal(t, ":0:0:false:false:false:false:false:false:false:112233:0:::0:::", hash.Font(&ml.Font{Color: &ml.Color{RGB: "112233"}}))
	require.Equal(t, ":0:0:false:false:false:false:false:false:false::0:::2.2:::", hash.Font(&ml.Font{Size: 2.2}))
	require.Equal(t, ":0:0:false:false:false:false:false:false:false::0:::0:doubleAccounting::", hash.Font(&ml.Font{Underline: format.UnderlineTypeDoubleAccounting}))
	require.Equal(t, ":0:0:false:false:false:false:false:false:false::0:::0::subscript:", hash.Font(&ml.Font{VAlign: format.FontVAlignSubscript}))
	require.Equal(t, ":0:0:false:false:false:false:false:false:false::0:::0:::major", hash.Font(&ml.Font{Scheme: format.FontSchemeMajor}))
	require.Equal(t, "calibri:1:1:true:true:true:true:true:true:false:112233:0:::2.2:doubleAccounting:subscript:major", hash.Font(&ml.Font{
		Scheme: format.FontSchemeMajor,
		VAlign: format.FontVAlignSubscript,
		Underline: format.UnderlineTypeDoubleAccounting,
		Size: 2.2,
		Color: &ml.Color{RGB: "112233"},
		Extend: true,
		Condense: true,
		Shadow: true,
		Strike: true,
		Italic: true,
		Bold: true,
		Family: format.FontFamilyRoman,
		Charset: 1,
		Name: "calibri",
	}))
}

func TestNumber(t *testing.T) {
	require.Equal(t, "0:", hash.NumberFormat(nil))
	require.Equal(t, "0:", hash.NumberFormat(&ml.NumberFormat{}))
	require.Equal(t, "1:", hash.NumberFormat(&ml.NumberFormat{ID: 1}))
	require.Equal(t, "0:aaa", hash.NumberFormat(&ml.NumberFormat{Code: "aaa"}))
	require.Equal(t, "1:aaa", hash.NumberFormat(&ml.NumberFormat{ID: 1, Code: "aaa"}))
}

func TestProtection(t *testing.T) {
	require.Equal(t, "false:false", hash.Protection(nil))
	require.Equal(t, "false:false", hash.Protection(&ml.CellProtection{}))
	require.Equal(t, "true:false", hash.Protection(&ml.CellProtection{Locked: true}))
	require.Equal(t, "false:true", hash.Protection(&ml.CellProtection{Hidden: true}))
	require.Equal(t, "true:true", hash.Protection(&ml.CellProtection{Locked: true, Hidden: true}))
}

func TestStyle(t *testing.T) {
}
