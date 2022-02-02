package main

import (
	"fmt"

	"github.com/mgaza/goTools"
)

type InstrumentInfo struct {
	ID                  string
	LegalAttributes     []LegalAttributes
	AmountAttributes    []AmountAttributes
	ReferenceAttributes []ReferenceAttributes
}

type LegalAttributes struct {
	AbstractName              string
	AbstractNumber            string
	Acreage                   string
	Addition                  string
	AppraisalReviewBoard      string
	Block                     string
	City                      string
	Comment                   string
	LegalRemarks              string
	Lot                       string
	Outlot                    string
	Parcel                    string
	Phase                     string
	PlatBookNumber            string
	PlatBookType              string
	PlatPageNumber            string
	PlatVolume                string
	Porcion                   string
	Portion                   string
	PreviousArbitration       string
	PrevParcel                string
	PublicImprovementDistrict string
	Section                   string
	Shares                    string
	Subdivision               string
	SubdivisionAddition       string
	SubdivisionNotes          string
	SurveyBlock               string
	SurveyName                string
	SurveyNumber              string
	Township                  string
	Tract                     string
	Unit                      string

	// use these for combining properties
	SubBlockKey   string
	SurNameNumKey string
}

type AmountAttributes struct {
	Amount string
}

type ReferenceAttributes struct {
	Number string
	Volume string
	Page   string
}

type AllLegalInfo struct {
	CountyName     string
	ExportContent  [][]string
	IndexContent   []string
	OutputFilePath string
	OutputFileName string
	RemarkPtr      bool
}

func CountyParser(AllInfo AllLegalInfo) {
	// stophere := ""

	switch AllInfo.RemarkPtr {
	case true:
		for row_num, info := range AllInfo.ExportContent {
			if row_num == 0 {
				continue
			}

			InstrumentInfo := InstrumentInfo{}

			for elem_num, content := range info {
				switch elem_num {
				case 4:
					InstrumentInfo.ID = content
				case 13:
					callCountyRemarksParser(AllInfo.CountyName, content, &InstrumentInfo)
					AllInfo.ExportContent[row_num][15] = GroupLegals(InstrumentInfo)

					// fmt.Println(AllInfo.ExportContent[row_num][13], "|", AllInfo.ExportContent[row_num][15])
					// fmt.Scanln(&stophere)
				}
			}

			// fmt.Println("Row_Num:", row_num, "=>", "Info:", info)
		}

		goTools.OpenAndWriteCSVFile(AllInfo.OutputFileName, AllInfo.OutputFilePath, AllInfo.ExportContent)
		fmt.Println("Completed Writing to:", AllInfo.OutputFileName)
	default:
		indexMap := make(map[string]InstrumentInfo)
		callCountyIndexParser(AllInfo.CountyName, AllInfo.IndexContent, &indexMap)

		for row_num := range AllInfo.ExportContent {
			if row_num == 0 {
				continue
			}

			mapKey := AllInfo.ExportContent[row_num][4] + AllInfo.ExportContent[row_num][5] + AllInfo.ExportContent[row_num][6] + AllInfo.ExportContent[row_num][7]

			AllInfo.ExportContent[row_num][14] = GroupAttributes(indexMap[mapKey])
			AllInfo.ExportContent[row_num][15] = GroupLegals(indexMap[mapKey])
			AllInfo.ExportContent[row_num][16] = GroupReferences(indexMap[mapKey])

			// fmt.Println("Inst_Num:", AllInfo.ExportContent[row_num][4], "Row_Num:", row_num, "=>", "Info:", GroupReferences(indexMap[mapKey]))
		}

		AllInfo.ExportContent = removeEmptyInst(AllInfo.ExportContent)

		goTools.OpenAndWriteCSVFile(AllInfo.OutputFileName, AllInfo.OutputFilePath, AllInfo.ExportContent)
		fmt.Println("Completed Writing to:", AllInfo.OutputFileName)
	}
}

func callCountyRemarksParser(countyName string, remark string, InstrumentInfo *InstrumentInfo) {
	switch countyName {
	case "harrison":
		HarrisonRemarksParser(remark, InstrumentInfo)
	}
}

func callCountyIndexParser(countyName string, indexes []string, indexMap *map[string]InstrumentInfo) {
	switch countyName {
	case "harrison":
		*indexMap = HarrisonIndexParser(indexes, *indexMap)
	}
}

func removeEmptyInst(content [][]string) [][]string {
	newOrder := [][]string{}

	for row_num := range content {
		if row_num == 0 {
			newOrder = append(newOrder, content[row_num])
			continue
		}

		if content[row_num][14] != "" || content[row_num][15] != "" || content[row_num][16] != "" {
			newOrder = append(newOrder, content[row_num])
		}
	}

	return newOrder
}
