package transnum_test

import (
	"testing"

	"github.com/pikomonde/i-view-prospace/service/transnum"
	"github.com/stretchr/testify/assert"
)

func testCasesRoman() map[string]int {
	return map[string]int{
		"I":         1,
		"II":        2,
		"III":       3,
		"IV":        4,
		"V":         5,
		"VI":        6,
		"VII":       7,
		"VIII":      8,
		"IX":        9,
		"X":         10,
		"XI":        11,
		"XIV":       14,
		"XV":        15,
		"XIX":       19,
		"XXIV":      24,
		"XLI":       41,
		"XLIV":      44,
		"XLV":       45,
		"XLIX":      49,
		"LX":        60,
		"LXI":       61,
		"LXIV":      64,
		"LXV":       65,
		"LXIX":      69,
		"LXXXVII":   87,
		"XC":        90,
		"XCI":       91,
		"XCIV":      94,
		"XCV":       95,
		"XCIX":      99,
		"CDX":       410,
		"CDXLIV":    444,
		"CDXCVI":    496,
		"MDCCXCIV":  1794,
		"MMDCXV":    2615,
		"MMMCMXCIX": 3999,
	}
}

func testCasesErrorRoman() map[string]int {
	return map[string]int{
		"-": 0,
	}
}

func getDict() map[string]rune {
	return map[string]rune{
		"glob": 'I',
		"prok": 'V',
		"pish": 'X',
		"tegj": 'L',
		"adsw": 'C',
		"qqq":  'D',
		"qaz":  'M',
	}
}

func getErrorDict() map[string]rune {
	return map[string]rune{
		"gacob": 'Q',
		"pacok": 'S',
		"pacsh": 'F',
		"tacgj": 'P',
		"aacsw": 'q',
		"qacq":  'a',
		"qacz":  '1',
	}
}

func TestRomanToInt(t *testing.T) {
	sNum := transnum.ServiceTransnum{}

	// Positive testing
	for k, v := range testCasesRoman() {
		assert.Equal(t, v, sNum.RomanToInt(k))
	}

	// Negative testing
	for k, v := range testCasesErrorRoman() {
		assert.Equal(t, v, sNum.RomanToInt(k))
	}
}

func TestGalacticToInt(t *testing.T) {
	dict := getDict()

	// Populate galacticUnit to service
	dictReverse := make(map[rune]string)
	sNum := transnum.ServiceTransnum{}
	for k, v := range dict {
		sNum.AddGalacticUnit(k, v)
		dictReverse[v] = k
	}

	// Positive testing
	for k, v := range testCasesRoman() {
		words := make([]string, 0)
		for _, w := range k {
			words = append(words, dictReverse[w])
		}
		res, err := sNum.GalaticToInt(words)
		assert.Equal(t, v, res)
		assert.Equal(t, nil, err)
	}

	// Negative testing
	for k, v := range testCasesErrorRoman() {
		words := make([]string, 0)
		for _, w := range k {
			words = append(words, dictReverse[w])
		}
		res, err := sNum.GalaticToInt(words)
		assert.Equal(t, v, res)
		assert.Equal(t, transnum.ErrInvalidGalacticUnit, err)
	}
}

func TestMustGalaticToInt(t *testing.T) {
	dict := getDict()

	// Populate galacticUnit to service
	dictReverse := make(map[rune]string)
	sNum := transnum.ServiceTransnum{}
	for k, v := range dict {
		sNum.AddGalacticUnit(k, v)
		dictReverse[v] = k
	}

	// Positive testing
	for k, v := range testCasesRoman() {
		words := make([]string, 0)
		for _, w := range k {
			words = append(words, dictReverse[w])
		}
		res := sNum.MustGalaticToInt(words)
		assert.Equal(t, v, res)
	}

	// Negative testing
	for k, v := range testCasesErrorRoman() {
		words := make([]string, 0)
		for _, w := range k {
			words = append(words, dictReverse[w])
		}
		res := sNum.MustGalaticToInt(words)
		assert.Equal(t, v, res)
	}
}

func TestIsRomanChar(t *testing.T) {
	testCases := map[rune]bool{
		'I': true,
		'V': true,
		'X': true,
		'L': true,
		'C': true,
		'D': true,
		'M': true,
		'i': false,
		'v': false,
		'x': false,
		'l': false,
		'c': false,
		'd': false,
		'm': false,
		'Q': false,
	}

	sNum := transnum.ServiceTransnum{}
	for k, v := range testCases {
		assert.Equal(t, v, sNum.IsRomanChar(k))
	}
}

func TestAddGalacticUnitPositive(t *testing.T) {
	dict := getDict()

	// Positive testing
	sNum := transnum.ServiceTransnum{}
	for k, v := range dict {
		err := sNum.AddGalacticUnit(k, v)
		assert.Equal(t, nil, err)
	}
	assert.Equal(t, len(dict), len(sNum.Dict))
}

func TestAddGalacticUnitNegative(t *testing.T) {
	errorDict := getErrorDict()

	// Positive testing
	sNum := transnum.ServiceTransnum{}
	for k, v := range errorDict {
		err := sNum.AddGalacticUnit(k, v)
		assert.Equal(t, transnum.ErrInvalidRoman, err)
	}
	assert.Equal(t, 0, len(sNum.Dict))

}
