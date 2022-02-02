/*
	Group Properties

	At this point, whatever county legal instrument that needed parsing should
	have populated `InstrumentInfo.LegalAttributes`. Which means there's a slice of
	Legal Attribute structures that we can group together for final legal line output.

	Line should look something like `subdivision:NAME;lot:1,2,3,4,5;block:20`
*/

package main

import (
	"regexp"
	"strings"

	"github.com/mgaza/goTools"
)

func GroupLegals(InstrumentInfo InstrumentInfo) string {
	var formattedLegalInfo = []string{}
	var legalInfoKeys = []string{}
	completedLegalLine := ""
	// stophere := ""

	for _, elem := range InstrumentInfo.LegalAttributes {
		switch {
		case elem.SubBlockKey != "" && elem.SurNameNumKey != "":
			indexes := grabKeyIndexes(legalInfoKeys, elem.SubBlockKey, elem.SurNameNumKey)

			checkIfKeyExists(&legalInfoKeys, elem.SubBlockKey)
			checkIfSliceLengthExists(&formattedLegalInfo, indexes)
			addInstInfoElem(&formattedLegalInfo[indexes[0]], elem, 0)

			checkIfKeyExists(&legalInfoKeys, elem.SurNameNumKey)
			checkIfSliceLengthExists(&formattedLegalInfo, indexes)
			addInstInfoElem(&formattedLegalInfo[indexes[1]], elem, 1)

			// fmt.Println(indexes, elem, formattedLegalInfo[indexes[0]], formattedLegalInfo[indexes[1]])
			// fmt.Scanln(&stophere)
		case elem.SubBlockKey != "" && elem.SurNameNumKey == "":
			indexes := grabKeyIndexes(legalInfoKeys, elem.SubBlockKey, elem.SurNameNumKey)

			checkIfKeyExists(&legalInfoKeys, elem.SubBlockKey)
			checkIfSliceLengthExists(&formattedLegalInfo, indexes)
			addInstInfoElem(&formattedLegalInfo[indexes[0]], elem, 0)

			// fmt.Println(indexes, elem, formattedLegalInfo[indexes[0]])
			// fmt.Scanln(&stophere)
		case elem.SubBlockKey == "" && elem.SurNameNumKey != "":
			indexes := grabKeyIndexes(legalInfoKeys, elem.SubBlockKey, elem.SurNameNumKey)

			checkIfKeyExists(&legalInfoKeys, elem.SurNameNumKey)
			checkIfSliceLengthExists(&formattedLegalInfo, indexes)
			addInstInfoElem(&formattedLegalInfo[indexes[1]], elem, 1)

			// fmt.Println(indexes, elem, formattedLegalInfo[indexes[1]])
			// fmt.Scanln(&stophere)
		default:
			indexes := grabKeyIndexes(legalInfoKeys, elem.SubBlockKey, elem.SurNameNumKey)

			checkIfKeyExists(&legalInfoKeys, "")
			checkIfSliceLengthExists(&formattedLegalInfo, indexes)
			addInstInfoElem(&formattedLegalInfo[indexes[0]], elem, 2)

			// fmt.Println(indexes, elem, formattedLegalInfo[indexes[0]])
			// fmt.Scanln(&stophere)
		}
	}

	completedLegalLine = strings.Join(formattedLegalInfo, "|")
	completedLegalLine = strings.TrimSuffix(completedLegalLine, "|")

	return completedLegalLine
}

func GroupAttributes(InstrumentInfo InstrumentInfo) string {
	completedAttributeLine := ""

	for _, elem := range InstrumentInfo.AmountAttributes {
		if elem.Amount != "" {
			if completedAttributeLine != "" {
				completedAttributeLine = completedAttributeLine + "|amount:" + elem.Amount
			} else {
				completedAttributeLine = "amount:" + elem.Amount
			}
		}
	}

	return completedAttributeLine
}

func GroupReferences(InstrumentInfo InstrumentInfo) string {
	completedReferenceLine := ""

	for _, elem := range InstrumentInfo.ReferenceAttributes {
		if elem.Number != "" {
			if elem.Volume != "" {
				if completedReferenceLine != "" {
					completedReferenceLine = completedReferenceLine + "|number:" + elem.Number + ";volume:" + elem.Volume + ";page:" + elem.Page
				} else {
					completedReferenceLine = "number:" + elem.Number + ";volume:" + elem.Volume + ";page:" + elem.Page
				}
			} else {
				if completedReferenceLine != "" {
					completedReferenceLine = completedReferenceLine + "|number:" + elem.Number
				} else {
					completedReferenceLine = "number:" + elem.Number
				}
			}
		} else if elem.Volume != "" {
			if completedReferenceLine != "" {
				completedReferenceLine = completedReferenceLine + "|volume:" + elem.Volume + ";page:" + elem.Page
			} else {
				completedReferenceLine = "volume:" + elem.Volume + ";page:" + elem.Page
			}
		}
	}

	return completedReferenceLine
}

func grabKeyIndexes(keySlc []string, sbKey string, snKey string) [2]int {
	positions := [2]int{len(keySlc), len(keySlc) + 1}

	if sbKey == "" || snKey == "" {
		positions[1] = len(keySlc)
	}

	for key, a := range keySlc {
		switch a {
		case sbKey:
			positions[0] = key
			if positions[1] == len(keySlc)+1 {
				positions[1] = len(keySlc)
			}
		case snKey:
			positions[1] = key
		}
	}
	return positions
}

func checkIfKeyExists(infoKeys *[]string, keyName string) {
	if len(*infoKeys) == 0 {
		*infoKeys = append(*infoKeys, keyName)
	} else {
		notfound := true

		for _, elem := range *infoKeys {
			if elem == keyName {
				notfound = false
			}
		}

		if notfound {
			*infoKeys = append(*infoKeys, keyName)
		}
	}
}

func checkIfSliceLengthExists(formattedLegal *[]string, indexes [2]int) {
	for _, elem := range indexes {
		if len(*formattedLegal) == elem {
			*formattedLegal = append(*formattedLegal, "")
		}
	}
}

func addInstInfoElem(formattedLegal *string, elem LegalAttributes, indexPos int) {
	if *formattedLegal == "" {
		formatNewProperty(formattedLegal, elem, indexPos)
	} else {
		addNewProperty(formattedLegal, elem)
	}
}

func addNewProperty(formattedLegal *string, elem LegalAttributes) {
	checkForSubElem(formattedLegal, elem.Phase, `phase:`, `phase:(.+?)(;|$)`)
	checkForSubElem(formattedLegal, elem.Lot, `lot:`, `lot:(.+?)(;|$)`)
	checkForSubElem(formattedLegal, elem.Outlot, `outlot:`, `outlot:(.+?)(;|$)`)
	checkForSubElem(formattedLegal, elem.Section, `section:`, `section:(.+?)(;|$)`)
	checkForSubElem(formattedLegal, elem.Township, `township:`, `township:(.+?)(;|$)`)
	checkForSubElem(formattedLegal, elem.Tract, `tract:`, `tract:(.+?)(;|$)`)
	checkForSubElem(formattedLegal, elem.Unit, `unit:`, `unit:(.+?)(;|$)`)
	checkForSubElem(formattedLegal, elem.Parcel, `parcel:`, `parcel:(.+?)(;|$)`)
	checkForSubElem(formattedLegal, elem.Porcion, `porcion:`, `porcion:(.+?)(;|$)`)
	checkForSubElem(formattedLegal, elem.Portion, `portion:`, `portion:(.+?)(;|$)`)
	checkForSubElem(formattedLegal, elem.AbstractName, `abstract name:`, `abstract name:(.+?)(;|$)`)
	checkForSubElem(formattedLegal, elem.AbstractNumber, `abstract number:`, `abstract number:(.+?)(;|$)`)
	checkForSubElem(formattedLegal, elem.Acreage, `acreage:`, `acreage:(.+?)(;|$)`)
	checkForSubElem(formattedLegal, elem.Addition, `addition:`, `addition:(.+?)(;|$)`)
}

func checkForSubElem(formattedLegal *string, subElem string, propLabel string, propRegLabel string) {
	if subElem != "" {
		re := regexp.MustCompile(propRegLabel)
		matched, err := regexp.MatchString(propRegLabel, *formattedLegal)

		goTools.CheckErrorFatal("There's a problem with the property label regex: ", err)

		if matched {
			matches := re.FindStringSubmatch(*formattedLegal)
			newRe := regexp.MustCompile(`,?` + subElem + `,?`)
			foundMatch := newRe.MatchString(*formattedLegal)

			if !foundMatch {
				oldInfo := propLabel + matches[1]
				newInfo := propLabel + matches[1] + "," + subElem
				*formattedLegal = strings.ReplaceAll(*formattedLegal, oldInfo, newInfo)
			}

		} else {
			*formattedLegal += ";" + propLabel + subElem
		}
	}
}

func formatNewProperty(formattedLegal *string, elem LegalAttributes, indexPos int) {

	switch indexPos {
	case 0:
		if elem.Subdivision != "" {
			if *formattedLegal == "" {
				*formattedLegal = "subdivision:" + elem.Subdivision
			} else {
				*formattedLegal += ";subdivision:" + elem.Subdivision
			}
		}
		if elem.SubdivisionAddition != "" {
			if *formattedLegal == "" {
				*formattedLegal = "subdivision addition:" + elem.SubdivisionAddition
			} else {
				*formattedLegal += ";subdivision addition:" + elem.SubdivisionAddition
			}
		}
		if elem.SubdivisionNotes != "" {
			if *formattedLegal == "" {
				*formattedLegal = "subdivision notes:" + elem.SubdivisionNotes
			} else {
				*formattedLegal += ";subdivision notes:" + elem.SubdivisionNotes
			}
		}
	case 1:
		if elem.SurveyName != "" {
			if *formattedLegal == "" {
				*formattedLegal = "survey name:" + elem.SurveyName
			} else {
				*formattedLegal += ";survey name:" + elem.SurveyName
			}
		}
		if elem.SurveyNumber != "" {
			if *formattedLegal == "" {
				*formattedLegal = "survey number:" + elem.SurveyNumber
			} else {
				*formattedLegal += ";survey number:" + elem.SurveyNumber
			}
		}
		if elem.SurveyBlock != "" {
			if *formattedLegal == "" {
				*formattedLegal = "survey block:" + elem.SurveyBlock
			} else {
				*formattedLegal += ";survey block:" + elem.SurveyBlock
			}
		}
	}

	if elem.City != "" {
		if *formattedLegal == "" {
			*formattedLegal = "city:" + elem.City
		} else {
			*formattedLegal += ";city:" + elem.City
		}
	}
	if elem.Phase != "" {
		if *formattedLegal == "" {
			*formattedLegal = "phase:" + elem.Phase
		} else {
			*formattedLegal += ";phase:" + elem.Phase
		}
	}
	if elem.Lot != "" {
		if *formattedLegal == "" {
			*formattedLegal = "lot:" + elem.Lot
		} else {
			*formattedLegal += ";lot:" + elem.Lot
		}
	}
	if elem.Outlot != "" {
		if *formattedLegal == "" {
			*formattedLegal = "outlot:" + elem.Outlot
		} else {
			*formattedLegal += ";outlot:" + elem.Outlot
		}
	}
	if elem.Block != "" {
		if *formattedLegal == "" {
			*formattedLegal = "block:" + elem.Block
		} else {
			*formattedLegal += ";block:" + elem.Block
		}
	}
	if elem.Section != "" {
		if *formattedLegal == "" {
			*formattedLegal = "section:" + elem.Section
		} else {
			*formattedLegal += ";section:" + elem.Section
		}
	}
	if elem.Township != "" {
		if *formattedLegal == "" {
			*formattedLegal = "township:" + elem.Township
		} else {
			*formattedLegal += ";township:" + elem.Township
		}
	}
	if elem.Tract != "" {
		if *formattedLegal == "" {
			*formattedLegal = "tract:" + elem.Tract
		} else {
			*formattedLegal += ";tract:" + elem.Tract
		}
	}
	if elem.Unit != "" {
		if *formattedLegal == "" {
			*formattedLegal = "unit:" + elem.Unit
		} else {
			*formattedLegal += ";unit:" + elem.Unit
		}
	}
	if elem.PlatBookType != "" {
		if *formattedLegal == "" {
			*formattedLegal = "plat book type:" + elem.PlatBookType
		} else {
			*formattedLegal += ";plat book type:" + elem.PlatBookType
		}
	}
	if elem.PlatBookNumber != "" {
		if *formattedLegal == "" {
			*formattedLegal = "plat book number:" + elem.PlatBookNumber
		} else {
			*formattedLegal += ";plat book number:" + elem.PlatBookNumber
		}
	}
	if elem.PlatVolume != "" {
		if *formattedLegal == "" {
			*formattedLegal = "plat volume:" + elem.PlatVolume
		} else {
			*formattedLegal += ";plat volume:" + elem.PlatVolume
		}
	}
	if elem.PlatPageNumber != "" {
		if *formattedLegal == "" {
			*formattedLegal = "plat page number:" + elem.PlatPageNumber
		} else {
			*formattedLegal += ";plat page number:" + elem.PlatPageNumber
		}
	}
	if elem.Parcel != "" {
		if *formattedLegal == "" {
			*formattedLegal = "parcel:" + elem.Parcel
		} else {
			*formattedLegal += ";parcel:" + elem.Parcel
		}
	}
	if elem.Porcion != "" {
		if *formattedLegal == "" {
			*formattedLegal = "porcion:" + elem.Porcion
		} else {
			*formattedLegal += ";porcion:" + elem.Porcion
		}
	}
	if elem.Portion != "" {
		if *formattedLegal == "" {
			*formattedLegal = "portion:" + elem.Portion
		} else {
			*formattedLegal += ";portion:" + elem.Portion
		}
	}
	if elem.AbstractName != "" {
		if *formattedLegal == "" {
			*formattedLegal = "abstract name:" + elem.AbstractName
		} else {
			*formattedLegal += ";abstract name:" + elem.AbstractName
		}
	}
	if elem.AbstractNumber != "" {
		if *formattedLegal == "" {
			*formattedLegal = "abstract number:" + elem.AbstractNumber
		} else {
			*formattedLegal += ";abstract number:" + elem.AbstractNumber
		}
	}
	if elem.Acreage != "" {
		if *formattedLegal == "" {
			*formattedLegal = "acreage:" + elem.Acreage
		} else {
			*formattedLegal += ";acreage:" + elem.Acreage
		}
	}
	if elem.Addition != "" {
		if *formattedLegal == "" {
			*formattedLegal = "addition:" + elem.Addition
		} else {
			*formattedLegal += ";addition:" + elem.Addition
		}
	}
	if elem.AppraisalReviewBoard != "" {
		if *formattedLegal == "" {
			*formattedLegal = "appraisal review board:" + elem.AppraisalReviewBoard
		} else {
			*formattedLegal += ";appraisal review board:" + elem.AppraisalReviewBoard
		}
	}
	if elem.Comment != "" {
		if *formattedLegal == "" {
			*formattedLegal = "comment:" + elem.Comment
		} else {
			*formattedLegal += ";comment:" + elem.Comment
		}
	}
	if elem.LegalRemarks != "" {
		if *formattedLegal == "" {
			*formattedLegal = "legal remarks:" + elem.LegalRemarks
		} else {
			*formattedLegal += ";legal remarks:" + elem.LegalRemarks
		}
	}
	if elem.PreviousArbitration != "" {
		if *formattedLegal == "" {
			*formattedLegal = "previous arbitration:" + elem.PreviousArbitration
		} else {
			*formattedLegal += ";previous arbitration:" + elem.PreviousArbitration
		}
	}
	if elem.PrevParcel != "" {
		if *formattedLegal == "" {
			*formattedLegal = "prev parcel:" + elem.PrevParcel
		} else {
			*formattedLegal += ";prev parcel:" + elem.PrevParcel
		}
	}
	if elem.PublicImprovementDistrict != "" {
		if *formattedLegal == "" {
			*formattedLegal = "public improvement district:" + elem.PublicImprovementDistrict
		} else {
			*formattedLegal += ";public improvement district:" + elem.PublicImprovementDistrict
		}
	}
	if elem.Shares != "" {
		if *formattedLegal == "" {
			*formattedLegal = "shares:" + elem.Shares
		} else {
			*formattedLegal += ";shares:" + elem.Shares
		}
	}
}
