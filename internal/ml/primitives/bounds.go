package primitives

import (
	"encoding/xml"
)

//Bounds is implementation of Ref
type Bounds struct {
	FromCol     int
	FromRow     int
	ToCol       int
	ToRow       int
	initialized bool
}

//BoundsFromIndexes returns a Bounds information for provided 0-based indexes
func BoundsFromIndexes(fromCol, fromRow, toCol, toRow int) Bounds {
	//rebound cols if required
	if fromCol > toCol {
		toCol, fromCol = fromCol, toCol
	}

	//rebound rows if required
	if fromRow > toRow {
		toRow, fromRow = fromRow, toRow
	}

	return Bounds{
		fromCol,
		fromRow,
		toCol,
		toRow,
		true,
	}
}

//ContainsRef checks if celRef is inside of bounds
func (b *Bounds) ContainsRef(celRef CellRef) bool {
	return b.Contains(celRef.ToIndexes())
}

//Contains checks if indexes cIdx and rIdx are inside of bounds
func (b *Bounds) Contains(cIdx, rIdx int) bool {
	return (cIdx >= b.FromCol && cIdx <= b.ToCol) && (rIdx >= b.FromRow && rIdx <= b.ToRow)
}

//Dimension returns total number of cols and rows in bounds
func (b *Bounds) Dimension() (width int, height int) {
	width = b.ToCol - b.FromCol + 1
	height = b.ToRow - b.FromRow + 1
	return
}

//ToRef returns reference of bounds. Alias of String() method
func (b *Bounds) ToRef() Ref {
	return Ref(b.String())
}

//String return textual version of bounds
func (b Bounds) String() string {
	return string(RefFromCellRefs(
		CellRefFromIndexes(b.FromCol, b.FromRow),
		CellRefFromIndexes(b.ToCol, b.ToRow)),
	)
}

//IsEmpty return true if type was not initialized
func (b Bounds) IsEmpty() bool {
	return b == Bounds{}
}

//MarshalXMLAttr marshal Bounds
func (b *Bounds) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if b.IsEmpty() {
		attr = xml.Attr{}
	} else {
		attr.Value = b.String()
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal Bounds
func (b *Bounds) UnmarshalXMLAttr(attr xml.Attr) error {
	*b = Ref(attr.Value).ToBounds()
	return nil
}
