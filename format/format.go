package format

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
	"github.com/plandem/xlsx/internal/ml"
)

//StyleRefID is alias of original ml.StyleRefID type to:
// 1) make it public
// 2) forbid usage of integers directly
// 3) getting valid ID for StyleFormat via style-sheet
// 4) put everything related to stylesheet to same package
type StyleRefID ml.StyleRefID

//StyleFormat is objects that holds combined information about cell styling
type StyleFormat struct {
	key string

	Font       ml.Font
	Fill       ml.Fill
	Alignment  ml.CellAlignment
	NumFormat  ml.NumberFormat
	Protection ml.CellProtection
	Border     ml.Border
}

type option func(o *StyleFormat)

//New creates and returns StyleFormat object with requested options
func New(options ...option) *StyleFormat {
	s := &StyleFormat{
		Fill: ml.Fill {
			Pattern: &ml.PatternFill{},
		},
		Border: ml.Border{
			Left: &ml.BorderSegment {},
			Right: &ml.BorderSegment {},
			Top: &ml.BorderSegment {},
			Bottom: &ml.BorderSegment {},
		},
	}
	s.Set(options...)
	return s
}

//Key returns unique hash for style settings
func (s *StyleFormat) Key() string {
	return s.key
}

//Set sets new options for style
func (s *StyleFormat) Set(options ...option) {
	for _, o := range options {
		o(s)
	}

	h := md5.New()
	io.WriteString(h, strings.Join([]string{
		s.getKeyForFont(),
		s.getKeyForFill(),
		s.getKeyForAlignment(),
		s.getKeyForNumFormat(),
		s.getKeyForProtection(),
		s.getKeyForBorder(),
	}, ":"))

	s.key = fmt.Sprintf("%x", h.Sum(nil))
}

func (s *StyleFormat) getKeyForFont() string {
	return strings.Join([]string{
		string(s.Font.Name),
		strconv.FormatInt(int64(s.Font.Family), 10),
		strconv.FormatBool(bool(s.Font.Bold)),
		strconv.FormatBool(bool(s.Font.Italic)),
		strconv.FormatBool(bool(s.Font.Strike)),
		strconv.FormatBool(bool(s.Font.Shadow)),
		strconv.FormatBool(bool(s.Font.Condense)),
		strconv.FormatBool(bool(s.Font.Extend)),
		s.getKeyForColor(s.Font.Color),
		strconv.FormatFloat(float64(s.Font.Size), 'f', -1, 64),
		string(s.Font.Underline),
		string(s.Font.VAlign),
		string(s.Font.Scheme),
	}, ":")
}

func (s *StyleFormat) getKeyForColor(color *ml.Color) string {
	if color == nil {
		return ""
	}

	result := []string {
		strconv.FormatBool(color.Auto),
		color.RGB,
	}

	if color.Indexed != nil {
		result = append(result, strconv.FormatInt(int64(*color.Indexed), 10))
	} else {
		result = append(result, "")
	}

	if color.Theme != nil {
		result = append(result, strconv.FormatInt(int64(*color.Theme), 10))
	} else {
		result = append(result, "")
	}

	if color.Tint != nil {
		result = append(result, strconv.FormatFloat(*color.Tint, 'f', -1, 64))
	} else {
		result = append(result, "")
	}

	return strings.Join(result, ":")
}

func (s *StyleFormat) getKeyForFill() string {
	return strings.Join([]string{
		strconv.FormatInt(int64(s.Fill.Pattern.Type), 10),
		s.getKeyForColor(s.Fill.Pattern.Color),
		s.getKeyForColor(s.Fill.Pattern.Background),
	}, ":")
}

func (s *StyleFormat) getKeyForNumFormat() string {
	return strings.Join([]string{
		strconv.FormatInt(int64(s.NumFormat.ID), 10),
		s.NumFormat.Code,
	}, ":")
}

func (s *StyleFormat) getKeyForProtection() string {
	return strings.Join([]string{
		strconv.FormatBool(s.Protection.Locked),
		strconv.FormatBool(s.Protection.Hidden),
	}, ":")
}

func (s *StyleFormat) getKeyForBorder() string {
	return strings.Join([]string{
		s.getKeyForColor(s.Border.Top.Color),
		s.Border.Top.Type.String(),

		s.getKeyForColor(s.Border.Bottom.Color),
		s.Border.Bottom.Type.String(),

		s.getKeyForColor(s.Border.Left.Color),
		s.Border.Left.Type.String(),

		s.getKeyForColor(s.Border.Right.Color),
		s.Border.Right.Type.String(),
	}, ":")
}

func (s *StyleFormat) getKeyForAlignment() string {
	return strings.Join([]string{
		strconv.FormatInt(int64(s.Alignment.Horizontal), 10),
		strconv.FormatInt(int64(s.Alignment.Vertical), 10),
		strconv.FormatInt(int64(s.Alignment.TextRotation), 10),
		strconv.FormatBool(s.Alignment.WrapText),
		strconv.FormatInt(int64(s.Alignment.Indent), 10),
		strconv.FormatInt(int64(s.Alignment.RelativeIndent), 10),
		strconv.FormatBool(s.Alignment.JustifyLastLine),
		strconv.FormatBool(s.Alignment.ShrinkToFit),
		strconv.FormatInt(int64(s.Alignment.ReadingOrder), 10),
	}, ":")
}
