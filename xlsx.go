package xlsx

import (
	"github.com/plandem/ooxml"
	"github.com/plandem/xlsx/internal"
	"io"
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

//Open opens a XLSX file with name
func Open(fileName string) (*Spreadsheet, error) {
	doc, err := shared.OpenFile(fileName, newSpreadsheet)
	if err != nil {
		return nil, err
	}

	if xlDoc, ok := doc.(*Spreadsheet); ok {
		return xlDoc, nil
	}

	return nil, shared.ErrorUnknownPackage(Spreadsheet{})
}

//OpenStream opens a XLSX stream
func OpenStream(stream io.Reader) (*Spreadsheet, error) {
	doc, err := shared.OpenStream(stream, newSpreadsheet)
	if err != nil {
		return nil, err
	}

	if xlDoc, ok := doc.(*Spreadsheet); ok {
		return xlDoc, nil
	}

	return nil, shared.ErrorUnknownPackage(Spreadsheet{})
}

//New creates and returns a new XLSX document
func New() *Spreadsheet {
	if doc, err := newSpreadsheet(shared.NewPackage(nil)); err == nil {
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
