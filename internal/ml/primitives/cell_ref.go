package primitives

import (
	"fmt"
	"github.com/plandem/ooxml"
	"math"
	"strconv"
	"strings"
)

//CellRef is a type to encode XSD ST_CellRef, an A1 style reference to the location of this cell
type CellRef string

//ToIndexes returns 0-based indexes of reference
func (cr CellRef) ToIndexes() (int, int) {
	colPart := strings.Map(ooxml.GetLettersFn, string(cr))
	rowPart := strings.Map(ooxml.GetNumbersFn, string(cr))

	var colIndex, rowIndex int
	for i, j := len(colPart)-1, 0; i >= 0; i, j = i-1, j+1 {
		colIndex += (int(colPart[i]) - int('A') + 1) * int(math.Pow(26, float64(j)))
	}

	rowIndex, _ = strconv.Atoi(rowPart)

	rowIndex--
	colIndex--

	return colIndex, rowIndex
}

//CellRefFromIndexes returns a CellRef for 0-based indexes
func CellRefFromIndexes(colIndex, rowIndex int) CellRef {
	if colIndex < 0 || rowIndex < 0 {
		return ""
	}

	var colName string
	i := colIndex + 1
	for i > 0 {
		colName = string((i-1)%26+65) + colName
		i = (i - 1) / 26
	}

	return CellRef(fmt.Sprintf("%s%d", colName, rowIndex+1))
}
