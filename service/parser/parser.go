package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// Parse parsing user's input and returns system's output
func (s *ServiceParser) Parse(in string) string {
	inSplit := strings.Split(in, space(defaultAssignmentKeyword))
	if len(inSplit) != 2 {
		return defaultUnexpectedInput
	}

	// Getting left and right side of sentence
	leftIn := strings.Split(inSplit[0], " ")
	rightIn := strings.Split(inSplit[1], " ")
	if (len(leftIn) == 0) || (len(rightIn) == 0) {
		return defaultUnexpectedInput
	}

	// Parse logic
	// TODO: better code writing?
	if (len(rightIn[0]) > 0) && s.ServiceTransnum.IsRomanChar(rune(rightIn[0][0])) {
		if len(leftIn) != 1 {
			return defaultUnexpectedInput
		}
		s.ServiceTransnum.AddGalacticUnit(leftIn[0], rune(rightIn[0][0]))
		return ""
	} else if rightIn[len(rightIn)-1] == defaultCreditsKeyword {
		// Left
		if len(leftIn) < 2 {
			return defaultUnexpectedInput
		}
		unit, err := s.ServiceTransnum.GalaticToInt(leftIn[:len(leftIn)-1])
		if err != nil {
			return defaultUnexpectedInput
		}

		// Right
		if len(rightIn) != 2 {
			return defaultUnexpectedInput
		}
		price, err := strconv.ParseInt(rightIn[0], 10, 64)
		if err != nil {
			return defaultUnexpectedInput
		}
		err = s.ServiceResource.AddResourcePrice(leftIn[len(leftIn)-1], unit, int(price))
		if err != nil {
			return defaultUnexpectedInput
		}
		return ""
	} else if (inSplit[0] == defaultQueryTransnumKeyword) && (rightIn[len(rightIn)-1] == defaultQuestionMarkKeyword) {
		words := rightIn[:len(rightIn)-1]

		// Translate galactic unit to decimal
		result, err := s.ServiceTransnum.GalaticToInt(words)
		if err != nil {
			return defaultUnexpectedInput
		}

		return strings.Join(append(words, []string{defaultAssignmentKeyword, fmt.Sprint(result)}...), " ")
	} else if (inSplit[0] == defaultQueryPriceKeyword) && (rightIn[len(rightIn)-1] == defaultQuestionMarkKeyword) {
		rightIn = rightIn[:len(rightIn)-1]
		if len(rightIn) < 2 {
			return defaultUnexpectedInput
		}

		// Translate galactic unit to decimal
		unit, err := s.ServiceTransnum.GalaticToInt(rightIn[:len(rightIn)-1])
		if err != nil {
			return defaultUnexpectedInput
		}

		totalPrice, err := s.ServiceResource.GetResourcePrice(unit, rightIn[len(rightIn)-1])
		if err != nil {
			return defaultUnexpectedInput
		}

		return strings.Join(append(rightIn, []string{defaultAssignmentKeyword, fmt.Sprint(totalPrice), defaultCreditsKeyword}...), " ")
	}

	return defaultUnexpectedInput
}
