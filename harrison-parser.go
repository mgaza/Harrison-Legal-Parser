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

func extractSurveyNumber(surveyName *string) string {
	re := regexp.MustCompile(`\s+(?P<surveynumber>\d+)[\s+]*$`)
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
