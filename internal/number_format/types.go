// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package number

import (
	"strings"

	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/plandem/xlsx/internal/number_format/convert"
)

//Type of underlying value of built-in number format
type Type byte

//builtInFormat is wrapper around ml.NumberFormat that adds type of number
type builtInFormat struct {
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

const dateFormatChars = "dmyhH"

type dateFormatMapping struct {
	excelFormat string
	goFormat    string
}

// the order of the mapping matters
var dateFormatParts = []dateFormatMapping{
	{"d",    "02"},
	{"mmm",  "Jan"},
	{"mm",  "04"},
	{"m",   "01"},
	{"yyyy", "2006"},
	{"yy",   "06"},
	{"h",    "03"},
	{"H",    "15"},
	{"ss",   "05"},
	{"AM/PM",   "PM"},
}

//New create and return ml.NumberFormat type for provided values, respecting built-in number formats
func New(id int, code string) ml.NumberFormat {
	return Normalize(ml.NumberFormat{ID: id, Code: code})
}

//IsBuiltIn returns true if id is one of built-in
func IsBuiltIn(id int) bool {
	return id >= 0 && id <= LastReservedID
}

//Resolve looks through built-in number formats and return related if there is any
func Resolve(nf ml.NumberFormat) *builtInFormat {
	//if id is one of built-in format, then try to resolve it via ID
	if IsBuiltIn(nf.ID) {
		if knownFormat, ok := builtIn[nf.ID]; ok {
			//known built-in format?
			return knownFormat
		}

		//general built-in format?
		return &builtInFormat{ml.NumberFormat{ID: nf.ID, Code: builtIn[0x00].Code}, General}
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

//Default returns default ID and code of number format for type
func Default(t Type) (int, string) {
	if id, ok := typeDefault[t]; ok {
		if number, ok := builtIn[id]; ok {
			return number.NumberFormat.ID, number.NumberFormat.Code
		}
	}

	number := builtIn[typeDefault[General]]
	return number.NumberFormat.ID, number.NumberFormat.Code
}

//Format tries to format value into required format code
func Format(value, code string, t primitives.CellType) string {
	if strings.ContainsAny(code, dateFormatChars) {
		return formatDate(value, code)
	}

	//TODO: implement formatting based on code and type
	return value
}

func formatDate(value, code string) string {
	dateFormat := code
	for _, mapping := range dateFormatParts {
		dateFormat = strings.ReplaceAll(dateFormat, mapping.excelFormat, mapping.goFormat)
	}

	dateValue, err := convert.ToDate(value)
	if err != nil {
		return value
	}

	return dateValue.UTC().Format(dateFormat)
}
