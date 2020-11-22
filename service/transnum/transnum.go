package transnum

// RomanToInt is used to translate roman numeric in string to integer
func (s *ServiceTransnum) RomanToInt(str string) (int, error) {
	// Validate roman structure
	if !s.IsValidRoman(str) {
		return 0, ErrInvalidRomanStructure
	}

	prevCharVal := 0
	total := 0
	for i := len(str) - 1; i >= 0; i-- {
		// Validate roman character
		if !s.IsRomanChar(rune(str[i])) {
			return 0, ErrInvalidRomanFound
		}

		curCharVal := roman(str[i]).Int()
		if curCharVal >= prevCharVal {
			total += curCharVal
		} else {
			total -= curCharVal
		}
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

// IsValidRoman is used to check whether a roman numeral valid or not
func (s *ServiceTransnum) IsValidRoman(str string) bool {
	prevChar := ' '
	prevCharVal := 0
	curLargestVal := 0
	countSameConsecutiveChar := 0
	for i := len(str) - 1; i >= 0; i-- {
		// Initialize curChar and curCharVal
		curChar := rune(str[i])
		curCharVal := roman(curChar).Int()

		if curCharVal > prevCharVal {
			curLargestVal = curCharVal
			countSameConsecutiveChar = 1
		} else if curCharVal == prevCharVal {
			// Rule: Same char cannot be used as addition after substraction, ex: IIV
			if curLargestVal != curCharVal {
				return false
			}

			// Rule: 'D', 'L', 'V' can not be repeated, ex: DD
			if (curChar == 'V') || (curChar == 'L') || (curChar == 'D') {
				return false
			}

			// Rule: 'I', 'X', 'C', 'M' can not be repeated more than 3 times, ex: XIIII
			countSameConsecutiveChar++
			if countSameConsecutiveChar > 3 {
				return false
			}
		} else {
			// 'I' can not be substracted from character other than 'V' and 'X', ex: IM
			// 'X' can not be substracted from character other than 'L' and 'C', ex: XM
			// 'C' can not be substracted from character other than 'D' and 'M', ex: C? (we can remove this validation)
			// 'V', 'L', and 'D' can not be used to substract other character, ex: DM
			if (curChar == 'I') && ((prevChar != 'V') && (prevChar != 'X')) ||
				(curChar == 'X') && ((prevChar != 'L') && (prevChar != 'C')) ||
				(curChar == 'C') && ((prevChar != 'D') && (prevChar != 'M')) ||
				(curChar == 'V') || (curChar == 'L') || (curChar == 'D') {
				return false
			}
			countSameConsecutiveChar = 1
		}
		prevChar = curChar
		prevCharVal = curCharVal
	}
	return true
}
