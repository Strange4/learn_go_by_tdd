package roman

type RomanNumeral struct {
	Value  int
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

func DecToRoman(number int) string {
	var result string
	for _, roman := range allRomanNumerals {
		for number >= roman.Value {
			result += roman.Symbol
			number -= roman.Value
		}
	}
	return result
}

func RomanToDec(roman string) int {
	return 0
}
