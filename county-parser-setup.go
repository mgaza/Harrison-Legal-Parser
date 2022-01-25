package main

import (
	"fmt"
)

type InstrumentInfo struct {
	ID                  string
	LegalAttributes     struct{}
	AmountAttributes    struct{}
	ReferenceAttributes struct{}
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
	ExportContent map[string][][]string
	ExportKey     []string
	IndexContent  map[string][][]string
	IndexKey      []string
	OutputFiles   []string
	RemarkPtr     bool
}

func CountyParser(AllInfo AllLegalInfo) {

	switch AllInfo.RemarkPtr {
	case true:
		for _, key := range AllInfo.ExportKey {
			fmt.Println("Key:", key, "=>", "Element:", AllInfo.ExportContent[key][0][0])
		}
	default:
		// Remember to write for 1930-1980
		fmt.Println("No function exists yet for index reading")
	}
}
