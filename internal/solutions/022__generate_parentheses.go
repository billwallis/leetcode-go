package solutions

import (
	"slices"
	"strings"
)

type Stack struct {
	stack []string
}

func (p *Stack) Add(s string) {
	if !slices.Contains(p.stack, s) {
		p.stack = append(p.stack, s)
	}
}

func SplitParentheses(s string) []string {
	if s == "" {
		return []string{}
	}

	var splitString []string
	counter, lastCut := 0, 0
	for i := 0; i < len(s); i++ {
		counter += map[uint8]int{
			"("[0]: 1,
			")"[0]: -1,
		}[s[i]]

		if counter == 0 {
			splitString = append(splitString, s[lastCut:i+1])
			lastCut = i + 1
		}
	}
	return splitString
}

// generateParenthesis is the twenty-second LeetCode problem:
// - https://leetcode.com/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	parens := []string{"()"}
	for i := 2; i <= n; i++ {
		var newParens Stack
		for _, char := range parens {
			newParens.Add("(" + char + ")")
			newParens.Add(char + "()")
			newParens.Add("()" + char)

			split := SplitParentheses(char)
			length := len(split)
			for j := 0; j < length; j++ {
				if j == 0 {
					newParens.Add("(" + split[0] + ")" + strings.Join(split[1:], ""))
				} else if j == length-1 {
					newParens.Add(strings.Join(split[:length-1], "") + "(" + split[length-1] + ")")
				} else {
					newParens.Add(
						strings.Join(split[:j], "") + "(" + split[j] + ")" + strings.Join(split[j+1:], ""))
				}
			}
		}
		parens = newParens.stack
	}
	return parens
}

var GenerateParenthesis = generateParenthesis /* For testing */
