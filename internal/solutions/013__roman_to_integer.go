package solutions

// romanToInt is the thirteenth LeetCode problem:
// - https://leetcode.com/problems/roman-to-integer/
func romanToInt(s string) int {
	decomposer := Decomposer() /* Nicked from solution 12 */
	number, symbolIndex := 0, len(decomposer)-1
	for s > "" {
		symbol := decomposer[symbolIndex]

		if symbolLen := len(symbol.numeral); len(s) >= symbolLen && symbol.numeral == s[:symbolLen] {
			number += symbol.decimal
			s = s[symbolLen:]
		} else {
			symbolIndex--
		}
	}

	return number
}

var RomanToInt = romanToInt /* For testing */
