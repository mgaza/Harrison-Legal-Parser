package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"

	"github.com/mgaza/goTools"
)

func ReadFilePaths(ericexportfilepath string, indexfilespath string, remarkPtr *bool) {

	err := os.Mkdir(ericexportfilepath+`\legals-import`, 0755)
	goTools.CheckErrorNonFatal("Could not make directory: ", err)

	importfilepaths := goTools.FilePathWalker(ericexportfilepath, `([a-z]+|[a-z]+_[a-z]+)_\d{4}-\d{2}-\d{2}_\d{4}-\d{2}-\d{2}`)

	for _, s := range importfilepaths {
		fileInfoToPass := AllLegalInfo{}
		fileInfoToPass.RemarkPtr = *remarkPtr

		openReadFile(s, ericexportfilepath, &fileInfoToPass, remarkPtr)

		err := os.Mkdir(fileInfoToPass.OutputFilePath, 0755)
		goTools.CheckErrorNonFatal("Could not make directory: ", err)

		if !*remarkPtr {
			fileName := fileInfoToPass.OutputFileName[0:4] + `\.txt`
			indexes := goTools.FilePathWalker(indexfilespath, fileName)

			dat, err := os.Open(indexes[0])
			goTools.CheckErrorFatal("Could not read index: ", err)
			defer dat.Close()

			r := bufio.NewScanner(dat)

			for r.Scan() {
				fileInfoToPass.IndexContent = append(fileInfoToPass.IndexContent, r.Text())
			}
		}

		CountyParser(fileInfoToPass)
	}

}

func openReadFile(path string, ericexportfilepath string, fileInfoToPass *AllLegalInfo, remarkPtr *bool) {
	countyname, yearMonth := goTools.GetExportCountyYearMonth(path)
	fileInfoToPass.CountyName = countyname

	fileInfoToPass.OutputFilePath = ericexportfilepath + "\\legals-import\\" + yearMonth
	fileInfoToPass.OutputFileName = yearMonth + ".csv"

	sourcefile, err := os.Open(path)
	goTools.CheckErrorFatal("could not open: ", err)
	defer goTools.CloseFile(sourcefile)

	r := csv.NewReader(bufio.NewReader(sourcefile))
	appendExport(r, fileInfoToPass, remarkPtr)

}

func appendExport(newFile *csv.Reader, fileInfoToPass *AllLegalInfo, remarkPtr *bool) {

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := newFile.Read()

		if err == io.EOF {
			break
		}
		goTools.CheckErrorFatal("Found an error: ", err)

		if record[13] != "" && *remarkPtr {
			fileInfoToPass.ExportContent = append(fileInfoToPass.ExportContent, record)
		} else if !*remarkPtr {
			fileInfoToPass.ExportContent = append(fileInfoToPass.ExportContent, record)
		}
	}
}
