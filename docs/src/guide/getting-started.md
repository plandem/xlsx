# Getting Started

::: warning COMPATIBILITY NOTE
Xlsx2Go requires GoLang >= 1.9
:::

## Installation
```
go get github.com/plandem/xlsx
```

## Create a new file
While creation of a new XLSX file is dead simple, valid XLSX file requires a bit more steps 
```go
package main

import (
	"fmt"
	"github.com/plandem/xlsx"
)

func main() {
	xl := xlsx.New()
	
	sheet := xl.AddSheet("sheet name")
	
	//
	// add new content to sheet
	//
	
	if err := xl.SaveAs("./foo.xlsx"); err != nil {
		fmt.Println(err)
	}
}
```

## Open the existing file
To open XLSX file using filename or file handler
```go
	xl, err := xlsx.Open("./foo.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	defer xl.Close()
````

## Save file
To update the existing XLSX file
```go
	if err := xl.Save(); err != nil {
		fmt.Println(err)
	}
```

To save a new XLSX file or under different name
```go
	if err := xl.SaveAs("./foo.xlsx"); err != nil {
		fmt.Println(err)
	}
```
