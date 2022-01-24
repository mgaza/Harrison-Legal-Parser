package main

import "fmt"

type AllLegalInfo struct {
	exportContent map[string][][]string
	indexContent  map[string][][]string
	outputFiles   []string
	remarkPtr     bool
}

func HarrisonParser(AllInfo AllLegalInfo) {
	fmt.Println(AllInfo.exportContent["1903-01-01_1903-12-31"][1])
	fmt.Println(AllInfo.indexContent)
	fmt.Println(AllInfo.outputFiles)
	fmt.Println(AllInfo.remarkPtr)
}
