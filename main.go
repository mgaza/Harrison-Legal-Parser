package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"

	"github.com/mgaza/goTools"
)

func main() {

	// path to source files
	ericexportfilepath := "C:\\Users\\mgaza\\Desktop\\temp\\Harrison\\1903-1929"
	extfile := "csv"

	err := filepath.WalkDir(ericexportfilepath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		re := regexp.MustCompile(extfile)

		if re.MatchString(d.Name()) {
			sourcefile, err := os.Open(path)
			goTools.CheckErrorFatal("could not open: ", err)

			r := csv.NewReader(bufio.NewReader(sourcefile))
			fileReader(r)

			sourcefile.Close()
		}

		return nil
	})

	goTools.CheckErrorFatal("error walking the path: ", err)
}

func fileReader(newFile *csv.Reader) {

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := newFile.Read()

		if err == io.EOF {
			break
		}
		goTools.CheckErrorFatal("Found an error: ", err)
		fmt.Printf("remark: %s\n", record[13])
	}
}
