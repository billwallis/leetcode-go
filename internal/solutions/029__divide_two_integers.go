package solutions

import (
	"math"
)

//func Abs(x int) int {
//	if x < 0 {
//		return -x
//	} else {
//		return x
//	}
//}

//func Within(number int, min int, max int) int {
//	// Validation excluded for performance
//
//	// Validation option 1
//	//if max > min {
//	//	panic("min is greater than max")
//	//}
//
//	// Validation option 2
//	//min, max := math.Min(min, max), math.Max(min, max)
//
//	if number <= min {
//		return min
//	} else if number >= max {
//		return max
//	} else {
//		return number
//	}
//}

// divide is the twenty-ninth LeetCode problem:
// - https://leetcode.com/problems/divide-two-integers/
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	var quotient int
	absDividend, absDivisor := Abs(dividend), Abs(divisor)
	for quotient = 0; absDividend >= absDivisor; quotient++ {
		absDividend -= absDivisor
	}

	if (dividend < 0) != (divisor < 0) {
		quotient = -quotient
	}
	return Within(quotient, math.MinInt32, math.MaxInt32)
}

var Divide = divide /* For testing */
