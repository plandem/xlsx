package styles

import (
	"github.com/plandem/ooxml/ml"
)

type namedStyleType int

//List of all possible types for NamedStyle
const (
	NamedStyleNormal namedStyleType = iota
	_NamedStyleRowLevel
	_NamedStyleColLevel
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
	//pseudo styles for RowLevel 1 - 7
	NamedStyleRowLevel1 = iota + (100 + _NamedStyleRowLevel)
	NamedStyleRowLevel2
	NamedStyleRowLevel3
	NamedStyleRowLevel4
	NamedStyleRowLevel5
	NamedStyleRowLevel6
	NamedStyleRowLevel7
	//pseudo styles for ColLevel 1 - 7
	NamedStyleColLevel1 = iota + (200 + _NamedStyleColLevel)
	NamedStyleColLevel2
	NamedStyleColLevel3
	NamedStyleColLevel4
	NamedStyleColLevel5
	NamedStyleColLevel6
	NamedStyleColLevel7
)

var (
	namedStyleNames map[namedStyleType]string
)

func init() {
	namedStyleNames = map[namedStyleType]string{
		NamedStyleNormal:            "Normal",
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

		//pseudo types
		NamedStyleRowLevel1: "RowLevel_1",
		NamedStyleRowLevel2: "RowLevel_2",
		NamedStyleRowLevel3: "RowLevel_3",
		NamedStyleRowLevel4: "RowLevel_4",
		NamedStyleRowLevel5: "RowLevel_5",
		NamedStyleRowLevel6: "RowLevel_6",
		NamedStyleRowLevel7: "RowLevel_7",
		NamedStyleColLevel1: "ColLevel_1",
		NamedStyleColLevel2: "ColLevel_2",
		NamedStyleColLevel3: "ColLevel_3",
		NamedStyleColLevel4: "ColLevel_4",
		NamedStyleColLevel5: "ColLevel_5",
		NamedStyleColLevel6: "ColLevel_6",
		NamedStyleColLevel7: "ColLevel_7",
	}
}

//NamedStyle is option to update Info with provided settings for NamedStyleInfo
func NamedStyle(name interface{}) func(*Info) {
	return func(s *Info) {
		if n, ok := name.(string); ok {
			if len(n) == 0 {
				panic("you must provide a name for custom named style")
			}

			s.namedInfo.BuiltinId = nil
			s.namedInfo.Name = n
		} else if t, ok := name.(namedStyleType); ok {
			//is known built-in style?
			defaultName, known := namedStyleNames[t]
			if !known {
				panic("unknown ID of built-in named style")
			}

			//fix our pseudo styles
			if t >= NamedStyleRowLevel1 && t <= NamedStyleRowLevel7 {
				t = _NamedStyleRowLevel
			}

			if t >= NamedStyleColLevel1 && t <= NamedStyleColLevel7 {
				t = _NamedStyleColLevel
			}

			builtInID := int(t)
			s.namedInfo.BuiltinId = ml.OptionalIndex(builtInID)
			s.namedInfo.Name = defaultName
		} else {
			panic("unsupported format of NamedStyle (can'be name for custom style or built-in type only)")
		}
	}
}
