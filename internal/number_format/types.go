package numberFormat

import (
	"github.com/plandem/xlsx/internal/ml"
)

//Type of underlying value of built-in number format
type Type byte

//BuiltInFormat is wrapper around ml.NumberFormat that adds type of number
type BuiltInFormat struct {
	ml.NumberFormat
	Type Type
}

//List of all possible types for Type
const (
	General Type = iota
	Integer
	Float
	Date
	Time
	DateTime
	DeltaTime
)

//LastReservedID is id of last built-in/reserved format
const LastReservedID = 163

//New create and return ml.NumberFormat type for provided values, respecting built-in number formats
func New(id int, code string) ml.NumberFormat {
	return Normalize(ml.NumberFormat{ID: id, Code: code})
}

//IsBuiltIn returns true if id is one of built-in
func IsBuiltIn(id int) bool {
	return id >= 0 && id <= LastReservedID
}

//Resolve looks through built-in number formats and return related if there is any
func Resolve(nf ml.NumberFormat) *BuiltInFormat {
	//if id is one of built-in format, then try to resolve it via ID
	if IsBuiltIn(nf.ID) {
		if knownFormat, ok := builtIn[nf.ID]; ok {
			//known built-in format?
			return knownFormat
		}

		//general built-in format?
		return &BuiltInFormat{ml.NumberFormat{ID: nf.ID, Code: builtIn[0x00].Code}, General}
	}

	//if there is a known format code, then use that built-in format
	for _, knownFormat := range builtIn {
		if knownFormat.Code == nf.Code {
			return knownFormat
		}
	}

	return nil
}

//Normalize tries resolve provided format via list of built-in formats and returns one of built-in or original format
func Normalize(nf ml.NumberFormat) ml.NumberFormat {
	if found := Resolve(nf); found != nil {
		return found.NumberFormat
	}

	//if there is valid code, then return original code and unknown ID
	if len(nf.Code) > 0 {
		return ml.NumberFormat{ID: -1, Code: nf.Code}
	}

	//looks like known ID was provided, so just use it as is
	return nf
}
