package solutions

/*
	Symbol    Value
	I         1
	V         5
	X         10
	L         50
	C         100
	D         500
	M         1000

Subtraction rules:

- I can be placed before V (5) and X (10) to make 4 and 9.
- X can be placed before L (50) and C (100) to make 40 and 90.
- C can be placed before D (500) and M (1000) to make 400 and 900.
*/

type Symbol struct {
	decimal int
	numeral string
}

func Decomposer() []Symbol {
	return []Symbol{
		{decimal: 1, numeral: "I"},
		{decimal: 4, numeral: "IV"},
		{decimal: 5, numeral: "V"},
		{decimal: 9, numeral: "IX"},
		{decimal: 10, numeral: "X"},
		{decimal: 40, numeral: "XL"},
		{decimal: 50, numeral: "L"},
		{decimal: 90, numeral: "XC"},
		{decimal: 100, numeral: "C"},
		{decimal: 400, numeral: "CD"},
		{decimal: 500, numeral: "D"},
		{decimal: 900, numeral: "CM"},
		{decimal: 1000, numeral: "M"},
	}
}

// intToRoman is the twelfth LeetCode problem:
// - https://leetcode.com/problems/integer-to-roman/
func intToRoman(num int) string {
	decomposer := Decomposer()
	roman, symbolIndex := "", len(decomposer)-1
	for num > 0 {
		symbol := decomposer[symbolIndex]
		if symbol.decimal <= num {
			roman += symbol.numeral
			num -= symbol.decimal
		} else {
			symbolIndex--
		}
	}

	return roman
}

var IntToRoman = intToRoman /* For testing */
