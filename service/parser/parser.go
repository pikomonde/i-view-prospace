package parser

import (
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
	if (len(rightIn[0]) > 0) && s.ServiceTransnum.IsRomanChar(rune(rightIn[0][0])) {
		if len(leftIn) != 1 {
			return defaultUnexpectedInput
		}
		s.ServiceTransnum.AddGalacticUnit(leftIn[0], rune(rightIn[0][0]))
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
	}

	return ""
}
