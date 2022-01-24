package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/mgaza/goTools"
)

func ReadFilePaths(ericexportfilepath string, remarkPtr *bool) AllLegalInfo {
	fileInfoToPass := AllLegalInfo{}
	fileInfoToPass.exportContent = make(map[string][][]string)
	fileInfoToPass.remarkPtr = *remarkPtr

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
			openReadFile(s, outdirectory, &fileInfoToPass)

			// fmt.Println("writing: ", s)
		}
	}

	return fileInfoToPass
}

func openReadFile(path string, outdirectory string, fileInfoToPass *AllLegalInfo) {
	yearMonth := goTools.GetExportYearMonth(path)

	sourcefile, err := os.Open(path)
	goTools.CheckErrorFatal("could not open: ", err)
	defer goTools.CloseFile(sourcefile)

	r := csv.NewReader(bufio.NewReader(sourcefile))
	// fileReader(r)
	records, _ := r.ReadAll()
	fileInfoToPass.exportContent[yearMonth] = records

	// for _, i := range records {
	// 	for _, j := range i {
	// 		fmt.Println(j)
	// 	}
	// }
	// goTools.OpenAndWriteCSVFile("thisisatest.csv", outdirectory, records)

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
