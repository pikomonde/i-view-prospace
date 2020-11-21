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

	// Parse logic
	// TODO: better code writing? Maybe using regeluar expression?
	if (len(leftIn) == 1) && (len(rightIn) == 1) && (len(rightIn[0]) == 1) {
		// Add Galactic Unit to dictionary
		err := s.ServiceTransnum.AddGalacticUnit(leftIn[0], rune(rightIn[0][0]))
		if err != nil {
			return defaultUnexpectedInput
		}

		return ""
	} else if (len(leftIn) >= 1) && (len(rightIn) == 2) && (rightIn[len(rightIn)-1] == defaultCreditsKeyword) {
		// Translate Galactic unit to decimal
		unit, err := s.ServiceTransnum.GalaticToInt(leftIn[:len(leftIn)-1])
		if err != nil {
			return defaultUnexpectedInput
		}

		// Parse price from string to integer
		price, err := strconv.ParseInt(rightIn[0], 10, 64)
		if err != nil {
			return defaultUnexpectedInput
		}

		// Add price to dictionary
		err = s.ServiceResource.AddResourcePrice(leftIn[len(leftIn)-1], unit, int(price))
		if err != nil {
			return defaultUnexpectedInput
		}

		return ""
	} else if (strings.Join(leftIn, " ") == defaultQueryTransnumKeyword) && (rightIn[len(rightIn)-1] == defaultQuestionMarkKeyword) {
		words := rightIn[:len(rightIn)-1]

		// Translate galactic unit to decimal
		result, err := s.ServiceTransnum.GalaticToInt(words)
		if err != nil {
			return defaultUnexpectedInput
		}

		return strings.Join(append(words, []string{defaultAssignmentKeyword, fmt.Sprint(result)}...), " ")
	} else if (strings.Join(leftIn, " ") == defaultQueryPriceKeyword) && (len(rightIn) >= 3) && (rightIn[len(rightIn)-1] == defaultQuestionMarkKeyword) {
		rightIn = rightIn[:len(rightIn)-1]

		// Translate galactic unit to decimal
		unit, err := s.ServiceTransnum.GalaticToInt(rightIn[:len(rightIn)-1])
		if err != nil {
			return defaultUnexpectedInput
		}

		// Get and calculate total price of a resource
		totalPrice, err := s.ServiceResource.GetResourcePrice(unit, rightIn[len(rightIn)-1])
		if err != nil {
			return defaultUnexpectedInput
		}

		return strings.Join(append(rightIn, []string{defaultAssignmentKeyword, fmt.Sprint(totalPrice), defaultCreditsKeyword}...), " ")
	}

	return defaultUnexpectedInput
}
