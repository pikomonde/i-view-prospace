package transnum

// RomanToInt is used to translate roman numeric in string to integer
func (s *ServiceTransnum) RomanToInt(str string) int {
	prevCharVal := 0
	total := 0
	for i := len(str) - 1; i >= 0; i-- {
		curCharVal := roman(str[i]).Int()
		if curCharVal >= prevCharVal {
			total += curCharVal
		} else {
			total -= curCharVal
		}
		prevCharVal = curCharVal
	}
	return total
}

// GalaticToInt is used to translate galactic unit in array to integer
func (s *ServiceTransnum) GalaticToInt(words []string) (int, error) {
	prevCharVal := 0
	total := 0
	for i := len(words) - 1; i >= 0; i-- {
		curChar, ok := s.Dict[words[i]]
		if !ok {
			return 0, ErrInvalidGalacticUnit
		}
		curCharVal := roman(curChar).Int()
		if curCharVal >= prevCharVal {
			total += curCharVal
		} else {
			total -= curCharVal
		}
		prevCharVal = curCharVal
	}
	return total, nil
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
		return ErrInvalidRoman
	}
	s.Dict[galacticUnit] = roman(r)
	return nil
}
