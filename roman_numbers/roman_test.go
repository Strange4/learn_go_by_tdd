package roman

import (
	"fmt"
	"hello/assertions"
	"testing"
	"testing/quick"
)

var testCases = []struct {
	decimal uint16
	roman   string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{100, "C"},
	{90, "XC"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestDecToRoman(t *testing.T) {
	for _, testCase := range testCases {
		name := fmt.Sprintf("%d converts to %q", testCase.decimal, testCase.roman)
		t.Run(name, func(t *testing.T) {
			got, err := DecToRoman(testCase.decimal)
			if err != nil {
				t.Errorf("expected no error but got %v", err)
			}
			want := testCase.roman
			assertions.AssertString(t, got, want)
		})
	}
}

func TestRomanToDec(t *testing.T) {
	for _, testCase := range testCases {
		name := fmt.Sprintf("%q converts to %d", testCase.roman, testCase.decimal)
		t.Run(name, func(t *testing.T) {
			got, err := RomanToDec(testCase.roman)
			if err != nil {
				t.Errorf("expected no error but got %v", err)
			}
			want := testCase.decimal
			assertions.AssertEqual(t, got, want)
		})
	}
}

func TestEncodingAndDecoding(t *testing.T) {

	assertToAndFromRoman := func(number uint16) bool {
		roman, err := DecToRoman(number)
		if err != nil {
			if number > RomanMaxNumber {
				// expected error
				return true
			} else {
				t.Errorf("expected no error but got %v", err)
				return false
			}
		}

		fromRoman, _ := RomanToDec(roman)
		return number == fromRoman
	}

	if err := quick.Check(assertToAndFromRoman, nil); err != nil {
		t.Errorf("failed the test with error: %v", err)
	}
}
