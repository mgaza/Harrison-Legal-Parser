package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	// Open the file
	ericexportfile, err := os.Open("C:\\Users\\mgaza\\Desktop\\temp\\Harrison\\1903-1929\\harrison_1903-01-01_1903-12-31.csv")
	if err != nil {
		log.Fatal(err)
	}

	// goTools.checkError("Found an error: ", err)

	// Parse the file
	r := csv.NewReader(bufio.NewReader(ericexportfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		// goTools.checkError("Found an error: ", err)
		fmt.Printf("remark: %s\n", record[13])
	}
}

// func checkError(message string, err error) {
// 	if err != nil {
// 		log.Fatal(message, err)
// 	}
// }
