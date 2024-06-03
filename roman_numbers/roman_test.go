package roman

import (
	"hello/assertions"
	"testing"
)

func TestDecToRoman(t *testing.T) {
	testCases := []struct {
		name          string
		decimal       int
		expectedRoman string
	}{
		{"1 converts to I", 1, "I"},
		{"2 converts to II", 2, "II"},
		{"3 converts to III", 3, "III"},
		{"4 converts to IV", 4, "IV"},
		{"9 converts to IX", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
		{"14 gets converted to XIV", 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
		{"40 gets converted to XL", 40, "XL"},
		{"47 gets converted to XLVII", 47, "XLVII"},
		{"49 gets converted to XLIX", 49, "XLIX"},
		{"50 gets converted to L", 50, "L"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := DecToRoman(testCase.decimal)
			want := testCase.expectedRoman
			assertions.AssertString(t, got, want)
		})
	}
}
