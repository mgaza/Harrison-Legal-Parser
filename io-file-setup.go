package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/mgaza/goTools"
)

func ReadFilePaths(ericexportfilepath string, remarkPtr *bool) {

	switch *remarkPtr {
	case false:
		// Remember to write for 1930-1980
		fmt.Println("No function exists yet for index reading")
	default:
		importfilepaths := goTools.FilePathWalker(ericexportfilepath, "csv")

		outdirectory := ericexportfilepath + "\\output"
		err := os.Mkdir(outdirectory, 0755)
		goTools.CheckErrorNonFatal("Could not make directory: ", err)

		for _, s := range importfilepaths {
			openReadFile(s, outdirectory)

			fmt.Println("writing: ", s)
		}
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
