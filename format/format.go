package format

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
	"github.com/plandem/xlsx/internal/ml"
)

//StyleRefID is helper type do forbid usage of integers directly and getting valid ID for StyleFormat via style-sheet
//type StyleRefID int

//StyleFormat is objects that holds combined information about cell styling
type StyleFormat struct {
	key string

	Font       ml.Font
	Fill       fill
	Alignment  alignment
	NumFormat  numberFormat
	Protection protection
	Border     border
}

type option func(o *StyleFormat)

//New creates and returns StyleFormat object with requested options
func New(options ...option) *StyleFormat {
	s := &StyleFormat{}
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
		//string(s.Font.Name),
		//strconv.FormatInt(int64(s.Font.Family), 10),
		//strconv.FormatBool(bool(s.Font.Bold)),
		//strconv.FormatBool(bool(s.Font.Italic)),
		//strconv.FormatBool(bool(s.Font.Strike)),
		//strconv.FormatBool(bool(s.Font.Shadow)),
		//strconv.FormatBool(bool(s.Font.Condense)),
		//strconv.FormatBool(bool(s.Font.Extend)),
		//string(s.Font.Color),
		//strconv.FormatFloat(float64(s.Font.Size), 'f', -1, 64),
		//strconv.FormatInt(int64(s.Font.Underline), 10),
		//strconv.FormatInt(int64(s.Font.VAlign), 10),
		//strconv.FormatInt(int64(s.Font.Scheme), 10),
	}, ":")
}

func (s *StyleFormat) getKeyForFill() string {
	return strings.Join([]string{
		strconv.FormatInt(int64(s.Fill.Type), 10),
		//string(s.Fill.Color),
		//string(s.Fill.Background),
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
		//string(s.Border.Top.Color),
		strconv.FormatInt(int64(s.Border.Top.Type), 10),

		//string(s.Border.Bottom.Color),
		strconv.FormatInt(int64(s.Border.Bottom.Type), 10),

		//string(s.Border.Left.Color),
		strconv.FormatInt(int64(s.Border.Left.Type), 10),

		//string(s.Border.Right.Color),
		strconv.FormatInt(int64(s.Border.Right.Type), 10),
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
