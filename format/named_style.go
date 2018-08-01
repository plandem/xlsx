package format

import (
	"github.com/plandem/ooxml/ml"
)

type namedStyleType int

//List of all possible types for NamedStyle
//
//N.B.:
// NamedStyleRowLevel and NamedStyleColLevel has default name only for first row and col.
// If you want to add styles for other level, then you have to provide a name. E.g.: RowLevel_6
const (
	NamedStyleCustom namedStyleType = -1
	NamedStyleNormal namedStyleType = iota
	NamedStyleRowLevel
	NamedStyleColLevel
	NamedStyleComma
	NamedStyleCurrency
	NamedStylePercent
	NamedStyleComma0
	NamedStyleCurrency0
	NamedStyleHyperlink
	NamedStyleHyperlinkFollowed
	NamedStyleNote
	NamedStyleWarning
	_
	_
	_
	NamedStyleTitle
	NamedStyleHeading1
	NamedStyleHeading2
	NamedStyleHeading3
	NamedStyleHeading4
	NamedStyleInput
	NamedStyleOutput
	NamedStyleCalculation
	NamedStyleCheckCell
	NamedStyleLinkedCell
	NamedStyleTotal
	NamedStyleGood
	NamedStyleBad
	NamedStyleNeutral
	NamedStyleAccent1
	NamedStyleAccent1_20
	NamedStyleAccent1_40
	NamedStyleAccent1_60
	NamedStyleAccent2
	NamedStyleAccent2_20
	NamedStyleAccent2_40
	NamedStyleAccent2_60
	NamedStyleAccent3
	NamedStyleAccent3_20
	NamedStyleAccent3_40
	NamedStyleAccent3_60
	NamedStyleAccent4
	NamedStyleAccent4_20
	NamedStyleAccent4_40
	NamedStyleAccent4_60
	NamedStyleAccent5
	NamedStyleAccent5_20
	NamedStyleAccent5_40
	NamedStyleAccent5_60
	NamedStyleAccent6
	NamedStyleAccent6_20
	NamedStyleAccent6_40
	NamedStyleAccent6_60
	NamedStyleExplanatory
)

var (
	namedStyleNames map[namedStyleType]string
)

func init() {
	namedStyleNames = map[namedStyleType]string{
		NamedStyleNormal:            "Normal",
		NamedStyleRowLevel:          "RowLevel_1",
		NamedStyleColLevel:          "ColLevel_1",
		NamedStyleComma:             "Comma",
		NamedStyleCurrency:          "Currency",
		NamedStylePercent:           "Percent",
		NamedStyleComma0:            "Comma[0]",
		NamedStyleCurrency0:         "Currency[0]",
		NamedStyleHyperlink:         "Hyperlink",
		NamedStyleHyperlinkFollowed: "Followed Hyperlink",
		NamedStyleNote:              "Note",
		NamedStyleWarning:           "Warning Text",
		NamedStyleTitle:             "Title",
		NamedStyleHeading1:          "Heading 1",
		NamedStyleHeading2:          "Heading 2",
		NamedStyleHeading3:          "Heading 3",
		NamedStyleHeading4:          "Heading 4",
		NamedStyleInput:             "Input",
		NamedStyleOutput:            "Output",
		NamedStyleCalculation:       "Calculation",
		NamedStyleCheckCell:         "CheckCell",
		NamedStyleLinkedCell:        "LinkedCell",
		NamedStyleTotal:             "Total",
		NamedStyleGood:              "Good",
		NamedStyleBad:               "Bad",
		NamedStyleNeutral:           "Neutral",
		NamedStyleAccent1:           "Accent1",
		NamedStyleAccent1_20:        "20% - Accent1",
		NamedStyleAccent1_40:        "40% - Accent1",
		NamedStyleAccent1_60:        "60% - Accent1",
		NamedStyleAccent2:           "Accent2",
		NamedStyleAccent2_20:        "20% - Accent2",
		NamedStyleAccent2_40:        "40% - Accent2",
		NamedStyleAccent2_60:        "60% - Accent2",
		NamedStyleAccent3:           "Accent3",
		NamedStyleAccent3_20:        "20% - Accent3",
		NamedStyleAccent3_40:        "40% - Accent3",
		NamedStyleAccent3_60:        "60% - Accent3",
		NamedStyleAccent4:           "Accent4",
		NamedStyleAccent4_20:        "20% - Accent4",
		NamedStyleAccent4_40:        "40% - Accent4",
		NamedStyleAccent4_60:        "60% - Accent4",
		NamedStyleAccent5:           "Accent5",
		NamedStyleAccent5_20:        "20% - Accent5",
		NamedStyleAccent5_40:        "40% - Accent5",
		NamedStyleAccent5_60:        "60% - Accent5",
		NamedStyleAccent6:           "Accent6",
		NamedStyleAccent6_20:        "20% - Accent6",
		NamedStyleAccent6_40:        "40% - Accent6",
		NamedStyleAccent6_60:        "60% - Accent6",
		NamedStyleExplanatory:       "Explanatory Text",
	}
}

//NamedStyle is option to update StyleFormat with provided settings for NamedStyleInfo
func NamedStyle(name string, t namedStyleType) func(*StyleFormat) {
	return func(s *StyleFormat) {
		defaultName, ok := namedStyleNames[t]

		if len(name) == 0 {
			//non built-in styles must have name
			if !ok {
				panic("you must provide a name for named style")
			}

			name = defaultName
		}

		s.namedInfo.Name = name

		if ok {
			builtInID := int(t)
			s.namedInfo.BuiltinId = ml.OptionalIndex(&builtInID)
		} else {
			s.namedInfo.BuiltinId = nil
		}
	}
}
