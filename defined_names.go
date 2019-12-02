package xlsx

type defineNames struct {
	doc  *Spreadsheet
}

//newDefinedNames creates an object that implements defined names functionality
func newDefinedNames(doc  *Spreadsheet) *defineNames {
	return &defineNames{doc: doc}
}

//Add adds a new defined name for value with sheetID scope or global scope for -1
func (n* defineNames) Add(name, value string, sheetID int) error {
	panic(errorNotSupported)
}
