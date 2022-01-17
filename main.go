package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	"github.com/mgaza/goTools"
)

func main() {

	// path to source files
	// ericexportfilepath := "C:\\Users\\mgaza\\Desktop\\temp\\Harrison\\1903-1929"

	// path to source files using flags
	var ericexportfilepath string
	flag.StringVar(&ericexportfilepath, "source", "None", "full path to location of source files")
	flag.Parse()

	importfilepaths := goTools.FilePathWalker(ericexportfilepath, "csv")

	outdirectory := ericexportfilepath + "\\output"
	err := os.Mkdir(outdirectory, 0755)
	goTools.CheckErrorNonFatal("Could not make directory: ", err)

	for _, s := range importfilepaths {
		openReadFile(s, outdirectory)

		fmt.Println("writing: ", s)
	}

}

func openReadFile(path string, outdirectory string) {
	sourcefile, err := os.Open(path)
	goTools.CheckErrorFatal("could not open: ", err)
	defer goTools.CloseFile(sourcefile)

	r := csv.NewReader(bufio.NewReader(sourcefile))
	//fileReader(r)
	records, _ := r.ReadAll()
	goTools.OpenAndWriteCSVFile("thisisatest.csv", outdirectory, records)

}

// func fileReader(newFile *csv.Reader) {

// 	// Iterate through the records
// 	for {
// 		// Read each record from csv
// 		record, err := newFile.Read()

// 		if err == io.EOF {
// 			break
// 		}
// 		goTools.CheckErrorFatal("Found an error: ", err)
// 		fmt.Printf("remark: %s\n", record[13])
// 	}
// }
