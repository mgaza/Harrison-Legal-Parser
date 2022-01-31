package main

import (
	"flag"
)

func main() {
	// path to source files using flags
	var ericexportfilepath string
	flag.StringVar(&ericexportfilepath, "source", "None", "full path to location of source files")

	// path to index files using flags
	var indexfilespath string
	flag.StringVar(&indexfilespath, "index", "None", "full path to location of index files")

	// flag to determine whether or not to read from remarks or indexes
	remarkPtr := flag.Bool("remarkRead", true, "bool to show whether or not to parse from remarks or index files")
	flag.Parse()

	ReadFilePaths(ericexportfilepath, indexfilespath, remarkPtr)

}
