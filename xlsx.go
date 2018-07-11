package xlsx

import (
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx/internal"
)

var (
	//ColumnOption is a 'namespace' for all possible options of ColumnOptions object
	//
	// Possible options are:
	// OutlineLevel
	// Collapsed
	// Phonetic
	// Hidden
	// Formatting
	// Width
	ColumnOption internal.ColumnOption

	//RowOption is a 'namespace' for all possible options of RowOptions object
	//
	// Possible options are:
	// OutlineLevel
	// Collapsed
	// Phonetic
	// Hidden
	// Formatting
	// Height
	RowOption internal.RowOption
)

//Open opens a XLSX file with name or io.Reader
func Open(f interface{}) (*Spreadsheet, error) {
	doc, err := ooxml.Open(f, newSpreadsheet)
	if err != nil {
		return nil, err
	}

	if xlDoc, ok := doc.(*Spreadsheet); ok {
		return xlDoc, nil
	}

	return nil, ooxml.ErrorUnknownPackage(Spreadsheet{})
}

//New creates and returns a new XLSX document
func New() *Spreadsheet {
	if doc, err := newSpreadsheet(ooxml.NewPackage(nil)); err == nil {
		if xlDoc, ok := doc.(*Spreadsheet); ok {
			return xlDoc
		}
	}

	panic("Could not create a new XLSX document.")
}

//NewColumnOptions create and returns option set for column
func NewColumnOptions(options ...internal.ColumnOption) *internal.ColumnOptions {
	s := &internal.ColumnOptions{}
	s.Set(options...)
	return s
}

//NewRowOptions create and returns option set for row
func NewRowOptions(options ...internal.RowOption) *internal.RowOptions {
	s := &internal.RowOptions{}
	s.Set(options...)
	return s
}
