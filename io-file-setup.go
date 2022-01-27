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
	fileInfoToPass.ExportContent = make(map[string][][]string)
	fileInfoToPass.RemarkPtr = *remarkPtr

	switch *remarkPtr {
	case true:
		importfilepaths := goTools.FilePathWalker(ericexportfilepath, `([a-z]+|[a-z]+_[a-z]+)_\d{4}-\d{2}-\d{2}_\d{4}-\d{2}-\d{2}`)

		outdirectory := ericexportfilepath + "\\output"
		err := os.Mkdir(outdirectory, 0755)
		goTools.CheckErrorNonFatal("Could not make directory: ", err)

		for _, s := range importfilepaths {
			openReadFile(s, outdirectory, &fileInfoToPass) // Do Not Read All Exports at Once. Instead Process One at a Time
		}

	default:
		// Remember to write for 1930-1980
		fmt.Println("No function exists yet for index reading")
	}

	return fileInfoToPass
}

// Refactor to read in Contents, Write to Path, and Dump Export Info to save memory
func openReadFile(path string, outdirectory string, fileInfoToPass *AllLegalInfo) {
	countyname, yearMonth := goTools.GetExportCountyYearMonth(path)
	fileInfoToPass.CountyName = countyname

	writeFileNamePath := outdirectory + "\\" + yearMonth + ".csv"
	fileInfoToPass.OutputFiles = append(fileInfoToPass.OutputFiles, writeFileNamePath)

	sourcefile, err := os.Open(path)
	goTools.CheckErrorFatal("could not open: ", err)
	defer goTools.CloseFile(sourcefile)

	r := csv.NewReader(bufio.NewReader(sourcefile))
	// fileReader(r)
	records, _ := r.ReadAll()
	fileInfoToPass.ExportContent[yearMonth] = records
	fileInfoToPass.ExportKey = append(fileInfoToPass.ExportKey, yearMonth)

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
