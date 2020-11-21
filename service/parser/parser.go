package parser

import (
	"strings"
)

// Parse parsing user's input and returns system's output
func (s *ServiceParser) Parse(in string) string {
	inSplit := strings.Split(in, " "+defaultAssignmentOperator+" ")
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
	}

	return ""
}
