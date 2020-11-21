package transnum

import "errors"

type roman rune

// Int is used to translate roman digit to decimal.
func (r roman) Int() int {
	switch r {
	case romanI:
		return 1
	case romanV:
		return 5
	case romanX:
		return 10
	case romanL:
		return 50
	case romanC:
		return 100
	case romanD:
		return 500
	case romanM:
		return 1000
	}
	return 0
}

const (
	romanI roman = 'I'
	romanV roman = 'V'
	romanX roman = 'X'
	romanL roman = 'L'
	romanC roman = 'C'
	romanD roman = 'D'
	romanM roman = 'M'
)

type dictionary map[string]roman

// ErrInvalidRoman is returned when there is invalid roman character found.
var ErrInvalidRoman = errors.New("Invalid roman character found")
