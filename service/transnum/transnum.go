package transnum

// RomanToInt is used to translate roman numeric in string to integer
func (s *ServiceTransnum) RomanToInt(str string) (int, error) {
	prevChar := ' '
	prevCharVal := 0
	curLargestVal := 0
	total := 0
	// TODO: move roman validation to another function?
	countSameConsecutiveChar := 0
	for i := len(str) - 1; i >= 0; i-- {
		// Validate roman numeric
		curChar := rune(str[i])
		curCharVal := roman(curChar).Int()
		if !s.IsRomanChar(curChar) {
			return 0, ErrInvalidRomanFound
		}

		if curCharVal > prevCharVal {
			curLargestVal = curCharVal
			countSameConsecutiveChar = 1
			total += curCharVal
		} else if (curCharVal == prevCharVal) && (curLargestVal == curCharVal) {
			if (curChar == 'V') || (curChar == 'L') || (curChar == 'D') {
				return 0, ErrInvalidRomanStructure
			}
			countSameConsecutiveChar++
			if countSameConsecutiveChar > 3 {
				return 0, ErrInvalidRomanStructure
			}
			total += curCharVal
		} else {
			if (curChar == 'I') && ((prevChar != 'V') && (prevChar != 'X')) ||
				(curChar == 'X') && ((prevChar != 'L') && (prevChar != 'C')) ||
				(curChar == 'C') && ((prevChar != 'D') && (prevChar != 'M')) ||
				(curChar == 'V') || (curChar == 'L') || (curChar == 'D') {
				return 0, ErrInvalidRomanStructure
			}
			countSameConsecutiveChar = 1
			total -= curCharVal
		}
		prevChar = curChar
		prevCharVal = curCharVal
	}
	return total, nil
}

// GalaticToInt is used to translate galactic unit in array to integer
func (s *ServiceTransnum) GalaticToInt(words []string) (int, error) {
	romanNumeral, err := s.GalaticToRoman(words)
	if err != nil {
		return 0, err
	}

	decimal, err := s.RomanToInt(romanNumeral)
	if err != nil {
		return 0, err
	}
	return decimal, nil
}

// GalaticToRoman translate Galactic Unit to Roman Numerals
func (s *ServiceTransnum) GalaticToRoman(words []string) (string, error) {
	var res string
	for _, word := range words {
		curChar, ok := s.Dict[word]
		if !ok {
			return "", ErrInvalidGalacticUnitFound
		}
		res += string(curChar)
	}
	return res, nil
}

// MustGalaticToInt similar to GalaticToInt, but not returning error
func (s *ServiceTransnum) MustGalaticToInt(words []string) int {
	res, _ := s.GalaticToInt(words)
	return res
}

// IsRomanChar is used to check whether a character is roman or not
func (s *ServiceTransnum) IsRomanChar(r rune) bool {
	if (roman(r) == romanI) ||
		(roman(r) == romanV) ||
		(roman(r) == romanX) ||
		(roman(r) == romanL) ||
		(roman(r) == romanC) ||
		(roman(r) == romanD) ||
		(roman(r) == romanM) {
		return true
	}
	return false
}

// AddGalacticUnit is used to add galactic unit to database
func (s *ServiceTransnum) AddGalacticUnit(galacticUnit string, r rune) error {
	if !s.IsRomanChar(r) {
		return ErrInvalidRomanFound
	}
	s.Dict[galacticUnit] = roman(r)
	return nil
}
