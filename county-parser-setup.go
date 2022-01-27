package main

import (
	"fmt"
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
	CountyName    string
	ExportContent map[string][][]string // Change to Not Map
	ExportKey     []string
	IndexContent  map[string][][]string
	IndexKey      []string
	OutputFiles   []string
	RemarkPtr     bool
}

func CountyParser(AllInfo AllLegalInfo) {

	switch AllInfo.RemarkPtr {
	case true:
		for _, key_year := range AllInfo.ExportKey {
			for row_num, info := range AllInfo.ExportContent[key_year] {
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
						AllInfo.ExportContent[key_year][row_num][15] = GroupLegals(InstrumentInfo)
						// fmt.Println(AllInfo.ExportContent[key_year][row_num][13], "|", AllInfo.ExportContent[key_year][row_num][15])
					}
				}

				// fmt.Println("Row_Num:", row_num, "=>", "Info:", info)
			}
			// fmt.Println("Key_Year:", key_year)
		}
	default:
		// Remember to write for 1930-1980
		fmt.Println("No function exists yet for index reading")
	}
}

func callCountyRemarksParser(countyName string, remark string, InstrumentInfo *InstrumentInfo) {
	switch countyName {
	case "harrison":
		HarrisonRemarksParser(remark, InstrumentInfo)
	}
}