package main

import (
	"flag"
	"fmt"

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

	for i, s := range importfilepaths {
		fmt.Println(i, s)
	}

}

// func filePathWalker(filePath string, extfile string) []string {
// 	re := regexp.MustCompile(extfile)
// 	var paths []string

// 	err := filepath.WalkDir(filePath, func(path string, d fs.DirEntry, err error) error {
// 		if err != nil {
// 			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
// 			return err
// 		}

// 		if re.MatchString(d.Name()) {
// 			paths = append(paths, path)
// 		}

// 		return nil
// 	})

// 	goTools.CheckErrorFatal("error walking the path: ", err)
// 	return paths
// }

// func openReadFile(path string) {
// 	sourcefile, err := os.Open(path)
// 	goTools.CheckErrorFatal("could not open: ", err)
// 	defer closeFile(sourcefile)

// 	r := csv.NewReader(bufio.NewReader(sourcefile))
// 	fileReader(r)
// }

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

// func closeFile(f *os.File) {
// 	err := f.Close()
// 	goTools.CheckErrorFatal("could not close file: ", err)
// }
