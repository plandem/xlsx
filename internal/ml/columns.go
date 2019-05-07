package ml

import (
	"encoding/xml"
	"sort"
)

//ColList is a direct mapping of XSD CT_Cols
type ColList struct {
	Items []*Col `xml:"col,omitempty"`
}

func (cols *ColList) pack() {
	//moving grouped column ahead
	packed := cols.Items
	sort.Slice(packed, func(i, j int) bool { return packed[i].Min != packed[i].Max })

	//unpack columns
	unpacked := make(map[int]*Col, len(packed))
	unpackedKeys := make([]int, 0, len(packed))

	for _, c := range packed {
		//add columns with settings only
		if *c != (Col{Min: c.Min, Max: c.Max}) {
			if c.Min == c.Max {
				//we need unique indexes, but unpacked has mix of grouped/non-grouped columns with intersection of indexes
				if _, ok := unpacked[c.Min]; !ok {
					unpackedKeys = append(unpackedKeys, c.Min)
				}

				unpacked[c.Min] = c
			} else {
				for i := c.Min; i <= c.Max; i++ {
					unpacked[i] = c
					unpackedKeys = append(unpackedKeys, i)
				}
			}
		}
	}

	//pack columns
	var prevCol *Col
	packed = make([]*Col, 0, len(packed))
	sort.Ints(unpackedKeys)

	//cases:
	//same data: 1-10, 2, 3, 11 => 1-11
	//diff data: 1-10, 2, 3, 11 => 1, 2, 3, 4-10, 11
	for _, idx := range unpackedKeys {
		col := unpacked[idx]

		if prevCol == nil {
			prevCol = &Col{}
			*prevCol = *col
			prevCol.Min = idx
			prevCol.Max = idx
			packed = append(packed, prevCol)
		} else {
			colA := *prevCol
			colB := *col
			colA.Min = 0
			colA.Max = 0
			colB.Min = 0
			colB.Max = 0

			if colA == colB && idx == prevCol.Max+1 {
				prevCol.Max++
			} else {
				prevCol = &Col{}
				*prevCol = *col
				prevCol.Min = idx
				prevCol.Max = idx
				packed = append(packed, prevCol)
			}
		}
	}

	cols.Items = packed
}

func (cols *ColList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	//after using columns manager ColList can contain mix of grouped/non-grouped columns with intersection of indexes
	cols.pack()

	if len(cols.Items) > 0 {
		return e.EncodeElement(*cols, start)
	}

	return nil
}
