package xlsx_test

import (
	"github.com/plandem/xlsx"
	"log"
	"os"
)

func ExampleNew() {
	xl := xlsx.New()

	//... add a new content

	xl.SaveAs("new_file.xlsx")
}

func ExampleOpen_filename() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()
}

func ExampleOpen_file() {
	zipFile, err := os.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	xl, err := xlsx.Open(zipFile)
	if err != nil {
		log.Fatal(err)
	}

	_ = xl
}

func ExampleSave() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//... change content

	err = xl.Save()
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleSaveAs() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//... change content

	err = xl.SaveAs("new_file.xlsx")
	if err != nil {
		log.Fatal(err)
	}
}
