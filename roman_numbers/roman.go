package roman

import (
	"errors"
	"strings"
)

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

// must be in decreasing order
var allRomanNumerals = []RomanNumeral{
	{1_000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

var ErrValueTooLarge = errors.New("the value trying to be represented is too large")

const RomanMaxNumber = 3999

func DecToRoman(number uint16) (string, error) {
	var result string
	if number > RomanMaxNumber {
		return "", ErrValueTooLarge
	}
	for _, roman := range allRomanNumerals {
		for number >= roman.Value {
			result += roman.Symbol
			number -= roman.Value
		}
	}
	return result, nil
}

var ErrTooManyRepetions = errors.New("the roman numeral has more than 3 repeated values")

func RomanToDec(roman string) (uint16, error) {
	var result uint16

	for _, translation := range allRomanNumerals {
		var repetitionCount uint8
		for strings.HasPrefix(roman, translation.Symbol) {
			if repetitionCount == 3 {
				return 0, ErrTooManyRepetions
			}
			result += translation.Value
			roman = strings.TrimPrefix(roman, translation.Symbol)
			repetitionCount++
		}
	}
	return result, nil
}
