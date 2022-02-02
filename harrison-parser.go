package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/mgaza/goTools"
)

var containsSurveyNumber = regexp.MustCompile(`\d+$`).MatchString

var surveyNameRegList = []string{
	`\s+SUR[\s+]*$`,
}

func HarrisonRemarksParser(remark string, InstrumentInfo *InstrumentInfo) {
	splitRemark := strings.Split(remark, ",")
	for _, property := range splitRemark {
		locLegalAtt := LegalAttributes{}
		property += "}"

		labelRemarksPropertyString(&property)
		fillRemarksLegalAttributes(&locLegalAtt, property)

		InstrumentInfo.LegalAttributes = append(InstrumentInfo.LegalAttributes, locLegalAtt)
		// fmt.Println(locLegalAtt)
	}

}

func HarrisonIndexParser(indexes []string, indexMap map[string]InstrumentInfo) map[string]InstrumentInfo {
	for _, line := range indexes {
		lineElem := strings.Split(line, "|")

		mapKey := makeMapKey(lineElem[0], lineElem[1])

		indexMap[mapKey] = InstrumentInfo{
			ID:                  lineElem[0],
			LegalAttributes:     fillIndexLegalAttributes(lineElem[9]),
			AmountAttributes:    append(indexMap[lineElem[0]].AmountAttributes, AmountAttributes{}),
			ReferenceAttributes: append(indexMap[lineElem[0]].ReferenceAttributes, fillIndexReferenceAttributes(lineElem[6])),
		}

		// fmt.Println(indexMap[mapKey].ID, indexMap[mapKey].ReferenceAttributes)
	}
	return indexMap
}

func makeMapKey(f string, s string) string {
	ye := regexp.MustCompile(`^19[78]\d-00`)
	matchedYe := ye.MatchString(f)

	if matchedYe {
		f = f[5:]
	}

	f = strings.ReplaceAll(f, "-", "")

	re := regexp.MustCompile(`B:\s+(.+?)\s+V:\s+(.+?)\s+P:\s+(.+?)$`)
	matched := re.MatchString(s)

	if matched {
		matches := re.FindStringSubmatch(s)
		return f + matches[1] + matches[2] + matches[3]
	} else {
		return f
	}
}

func labelRemarksPropertyString(property *string) {
	if strings.Contains(*property, "Survey Name:") {
		*property = strings.ReplaceAll(*property, "Survey Name:", "}Survey Name:")
	}
	if strings.Contains(*property, "Acres:") {
		*property = strings.ReplaceAll(*property, "Acres:", "}Acres:")
	}
	if strings.Contains(*property, "Abst #:") {
		*property = strings.ReplaceAll(*property, "Abst #:", "}Abst #:")
	}
	if strings.Contains(*property, "Subdivision:") {
		*property = strings.ReplaceAll(*property, "Subdivision:", "}Subdivision:")
	}
	if strings.Contains(*property, "Lot:") {
		*property = strings.ReplaceAll(*property, "Lot:", "}Lot:")
	}
	if strings.Contains(*property, "Block:") {
		*property = strings.ReplaceAll(*property, "Block:", "}Block:")
	}
}

func fillRemarksLegalAttributes(locLegalAtt *LegalAttributes, property string) {
	locLegalAtt.AbstractNumber = strings.TrimSpace(findRemarksLegalProperty(`Abst \#:(?P<abstractnumber>.+?)}`, property))
	locLegalAtt.Acreage = strings.TrimSpace(findRemarksLegalProperty(`Acres:(?P<acres>.+?)}`, property))
	locLegalAtt.Block = strings.TrimSpace(findRemarksLegalProperty(`Block:(?P<block>.+?)}`, property))
	locLegalAtt.Lot = strings.TrimSpace(findRemarksLegalProperty(`Lot:(?P<lot>.+?)}`, property))
	locLegalAtt.Subdivision = strings.TrimSpace(findRemarksLegalProperty(`Subdivision:(?P<subdivision>.+?)}`, property))
	locLegalAtt.SurveyName = strings.TrimSpace(findRemarksLegalProperty(`Survey Name:(?P<surveyname>.+?)}`, property))

	if containsSurveyNumber(locLegalAtt.SurveyName) {
		locLegalAtt.SurveyNumber = extractSurveyNumber(&locLegalAtt.SurveyName)
	}

	cleanSurveyName(&locLegalAtt.SurveyName)
	cleanAcres(&locLegalAtt.Acreage)
	expSubdivision(&locLegalAtt.Subdivision)

	locLegalAtt.SubBlockKey = locLegalAtt.Subdivision + locLegalAtt.Block
	locLegalAtt.SurNameNumKey = locLegalAtt.SurveyName + locLegalAtt.SurveyNumber
}

func findRemarksLegalProperty(regVar string, property string) string {
	re := regexp.MustCompile(regVar)
	matched, err := regexp.MatchString(regVar, property)

	goTools.CheckErrorFatal("There's a problem with the legal regex: ", err)

	if matched {
		matches := re.FindStringSubmatch(property)
		return matches[1]
	} else {
		return ""
	}
}

func fillIndexLegalAttributes(legal string) []LegalAttributes {
	toFill := []LegalAttributes{}
	props := strings.Split(legal, ";")

	sub := ""
	lot := ""
	block := ""
	sur := ""
	surNum := ""
	abs := ""
	sec := ""
	acre := ""

	reAcre := regexp.MustCompile(`ACS$`)
	reAbst := regexp.MustCompile(`A-\d+`)
	reInt := regexp.MustCompile(`\d+`)
	reSur := regexp.MustCompile(`SUR$?`)

	for _, p := range props {
		elem := strings.Split(p, ":")

		for key, e := range elem {
			switch key {
			case 0:
				sub = strings.TrimSpace(e)

				if reSur.MatchString(e) {
					sub = ""
					sur = e

					if reAbst.MatchString(e) {
						sur = strings.Split(e, "A-")[0]
						abs = strings.Split(e, "A-")[1]
					}
				}
			case 1:
				if e != "P" {
					lot = strings.TrimSpace(e)
				}

				if reAcre.MatchString(e) {
					lot = ""
					acre = strings.ReplaceAll(e, "ACS", "")
				}
			case 2:
				block = strings.TrimSpace(e)
			case 5:
				if sur == "" {
					sur = strings.TrimSpace(e)
				}

				if reAbst.MatchString(e) {
					sur = strings.Split(e, "A-")[0]
					abs = strings.Split(e, "A-")[1]
				}
			case 6:
				if abs == "" {
					abs = strings.TrimSpace(e)
				}
			case 9:
				sec = strings.TrimSpace(e)
			case 11:
				if acre == "" {
					acre = strings.TrimSpace(e)
				}

				if !reInt.MatchString(e) && e != "" {
					acre = ""
					sub = e
				}
			}
		}

		if containsSurveyNumber(sur) {
			// fmt.Println(toFill.SurveyName)
			surNum = extractSurveyNumber(&sur)
		}

		cleanSurveyName(&sur)
		cleanAcres(&acre)
		expSubdivision(&sub)

		SubBlkKey := sub + block
		SurNNKey := sur + surNum

		toFill = append(toFill,
			LegalAttributes{
				Subdivision:    sub,
				Lot:            lot,
				Block:          block,
				SurveyName:     sur,
				SurveyNumber:   surNum,
				AbstractNumber: abs,
				Section:        sec,
				Acreage:        acre,
				SubBlockKey:    SubBlkKey,
				SurNameNumKey:  SurNNKey,
			})

		sub = ""
		lot = ""
		block = ""
		sur = ""
		surNum = ""
		abs = ""
		sec = ""
		acre = ""
	}

	return toFill
}

func fillIndexReferenceAttributes(ref string) ReferenceAttributes {
	toFill := ReferenceAttributes{}
	elem := strings.Split(ref, ":")

	for key, e := range elem {
		switch key {
		case 0:
			toFill.Number = strings.TrimSpace(e)
		case 2:
			toFill.Volume = strings.TrimSpace(e)
		case 3:
			toFill.Page = strings.TrimSpace(e)
		}
	}

	if toFill.Volume == "" || toFill.Page == "" {
		toFill.Volume = ""
		toFill.Page = ""
	}

	return toFill
}

func extractSurveyNumber(surveyName *string) string {
	re := regexp.MustCompile(`\s*#?(?P<surveynumber>\d+)[\s+]*$`)
	matches := re.FindStringSubmatch(*surveyName)

	*surveyName = strings.ReplaceAll(*surveyName, matches[1], "")

	return matches[1]
}

func cleanSurveyName(surveyName *string) {
	if *surveyName != "" {
		for _, elem := range surveyNameRegList {
			re := regexp.MustCompile(elem)
			match := re.FindString(*surveyName)

			*surveyName = strings.ReplaceAll(*surveyName, match, "")
		}
	}
}

func cleanAcres(acres *string) {
	if *acres != "" {
		*acres = strings.ReplaceAll(*acres, "-", " ")
		*acres = strings.ReplaceAll(*acres, "PT", "")

		acreSlice := strings.Split(*acres, " ")

		for i := 0; i < len(acreSlice); i++ {
			convertFractionToDecimalString(&acreSlice[i])
		}

		*acres = strings.Join(acreSlice, "")
	}
}

func convertFractionToDecimalString(fraction *string) {

	if strings.Contains(*fraction, "/") {
		fractionPieces := strings.Split(*fraction, "/")
		first, err := strconv.ParseFloat(fractionPieces[0], 64)
		goTools.CheckErrorFatal("not a number: ", err)

		second, err := strconv.ParseFloat(fractionPieces[1], 64)
		goTools.CheckErrorFatal("not a number: ", err)

		decFloatSum := first / second
		decStringSum := strconv.FormatFloat(decFloatSum, 'f', -1, 64)
		strLen := len(decStringSum)

		switch {
		case strLen == 3:
			*fraction = decStringSum[1:3]
		case strLen > 3:
			*fraction = decStringSum[1:4]
		default:
			*fraction = decStringSum
		}

		// fmt.Println(first, "/", second, "=", *fraction)
	}
}

func expSubdivision(sub *string) {
	*sub = strings.ReplaceAll(*sub, "SUBD", "SUBDIVISION")
	*sub = strings.ReplaceAll(*sub, "ADDN", "ADDITION")
}
