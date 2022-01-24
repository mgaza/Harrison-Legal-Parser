package main

import (
	"fmt"
)

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

type AllLegalInfo struct {
	ExportContent map[string][][]string
	ExportKey     []string
	IndexContent  map[string][][]string
	IndexKey      []string
	OutputFiles   []string
	RemarkPtr     bool
}

func HarrisonParser(AllInfo AllLegalInfo) {

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
