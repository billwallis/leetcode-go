package solutions

import (
	"strconv"
)

// Reverse is the seventh LeetCode problem:
// - https://leetcode.com/problems/reverse-integer/
func Reverse(x int) int {
	var reverse string
	var sign string
	str := strconv.FormatInt(int64(x), 10)

	if string(str[0]) == "-" {
		sign = "-"
		str = str[1:]
	}

	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}

	ret, err := strconv.ParseInt(sign+reverse, 10, 32)
	if err != nil {
		return 0
	} else {
		return int(ret)
	}
}
