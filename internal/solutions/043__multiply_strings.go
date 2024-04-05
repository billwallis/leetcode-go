package solutions

import (
	"fmt"
	"strconv"
)

// multiply is the forty-third LeetCode problem:
// - https://leetcode.com/problems/multiply-strings/
//
// We implement the way that we'd multiply numbers using pen and paper:
//
//	    1 2
//	  x 3 4
//	  -----
//	      8
//	    4 0
//	    6 0
//	+ 3 0 0
//	  -----
//	  4 0 8
func multiply(num1 string, num2 string) string {
	fmt.Printf("\nmultiply(%v, %v)\n", num1, num2)
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	unit2, result := 1, 0
	for i := len(num1) - 1; i >= 0; i-- {
		unit1 := 1
		for j := len(num2) - 1; j >= 0; j-- {
			digit1 := num1[i] - '0' // Adjust to the actual integer, not the ASCII decimal
			digit2 := num2[j] - '0' // Adjust to the actual integer, not the ASCII decimal
			fmt.Printf("  digit1, digit2: %v, %v\n", digit1, digit2)
			result += int(digit1) * unit1 * int(digit2) * unit2
			fmt.Printf("  result: %v\n", result)
			unit1 *= 10
		}
		unit2 *= 10
	}

	fmt.Printf("result: %v\n", result)
	return strconv.Itoa(result)
}

var Multiply = multiply /* For testing */
