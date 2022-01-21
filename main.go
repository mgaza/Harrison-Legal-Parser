package main

import (
	"flag"
)

func main() {

	// path to source files
	// ericexportfilepath := "C:\\Users\\mgaza\\Desktop\\temp\\Harrison\\1903-1929"

	// path to source files using flags
	var ericexportfilepath string
	flag.StringVar(&ericexportfilepath, "source", "None", "full path to location of source files")
	remarkPtr := flag.Bool("remarkRead", true, "bool to show whether or not to parse from remarks or index files")
	flag.Parse()

	ReadFilePaths(ericexportfilepath, remarkPtr)

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
