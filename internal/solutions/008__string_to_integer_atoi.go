package solutions

import (
	"math"
	"slices"
	"strconv"
)

// LeftTrim removes _all_ the leading characters from a string.
func LeftTrim(s string, c string) string {
	// Validation excluded for performance

	//if len(c) != 1 {
	//	panic("c must be a single character")
	//}
	for len(s) > 0 && s[0] == c[0] {
		s = s[1:]
	}
	return s
}

// Within returns the number if it is within the range of min and max,
// otherwise it returns the min or max value which is closest to the
// number.
func Within(number int, min int, max int) int {
	// Validation excluded for performance

	// Validation option 1
	//if max > min {
	//	panic("min is greater than max")
	//}

	// Validation option 2
	//min, max := math.Min(min, max), math.Max(min, max)

	if number <= min {
		return min
	} else if number >= max {
		return max
	} else {
		return number
	}
}

// MyAtoi is the eighth LeetCode problem:
// - https://leetcode.com/problems/string-to-integer-atoi/
func MyAtoi(s string) int {
	sign, result := "", ""
	str := LeftTrim(s, " ")
	minimum := int(math.Pow(float64(-2), float64(31)))
	maximum := int(-1 + math.Pow(float64(2), float64(31)))

	if len(str) != 0 {
		if firstChar := string(str[0]); slices.Contains([]string{"-", "+"}, firstChar) {
			sign = firstChar
			str = str[1:]
		}
	}

	str = LeftTrim(str, "0")
	if str == "" {
		return 0
	}

	for _, char := range str {
		// 0 to 9 (integers) == 48 to 57 (bytes)
		if char >= 48 && char <= 57 {
			result += string(char)
		} else {
			if result == "" {
				return 0
			}
			break
		}
	}

	if len(result) > 10 {
		if sign == "-" {
			return minimum
		} else {
			return maximum
		}
	}

	ret, err := strconv.Atoi(sign + result)
	if err != nil {
		panic(err)
	} else {
		return Within(ret, minimum, maximum)
	}
}
