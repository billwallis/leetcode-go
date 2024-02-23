package solutions

import (
	"math"
)

// divide is the twenty-ninth LeetCode problem:
// - https://leetcode.com/problems/divide-two-integers/
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	var quotient int
	absDividend, absDivisor := Abs(dividend), Abs(divisor)

	// Doing subtraction is slow and causing the solution to time out
	//for quotient = 0; absDividend >= absDivisor; quotient++ {
	//	absDividend -= absDivisor
	//}

	// This bitshift business is nicked from:
	// - https://leetcode.com/problems/divide-two-integers/solutions/1774409/go-succinct-intuitive-100-with-explanation
	for absDividend >= absDivisor {
		sub, add := absDivisor, 1
		// Skip increasing powers of 2
		for absDividend >= sub<<1 {
			sub <<= 1
			add <<= 1
		}
		absDividend -= sub
		quotient += add
	}

	if (dividend < 0) != (divisor < 0) {
		quotient = -quotient
	}
	return Within(quotient, math.MinInt32, math.MaxInt32)
}

var Divide = divide /* For testing */
