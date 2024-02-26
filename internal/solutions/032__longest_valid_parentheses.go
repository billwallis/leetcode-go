package solutions

import "fmt"

//func opposite(bracket uint8) uint8 {
//	switch bracket {
//	case "("[0]:
//		return ")"[0]
//	case ")"[0]:
//		return "("[0]
//	case "{"[0]:
//		return "}"[0]
//	case "}"[0]:
//		return "{"[0]
//	case "["[0]:
//		return "]"[0]
//	case "]"[0]:
//		return "["[0]
//	default:
//		panic("Unknown bracket")
//	}
//}
//
//// isValid is the twentieth LeetCode problem:
//// - https://leetcode.com/problems/valid-parentheses/
//func isValid(s string) bool {
//	openBrackets := ""
//	for i := 0; i < len(s); i++ {
//		char := s[i]
//		if char == "("[0] || char == "["[0] || char == "{"[0] {
//			openBrackets += string(char)
//		} else if length := len(openBrackets); length > 0 && opposite(char) == openBrackets[length-1] {
//			openBrackets = openBrackets[:length-1]
//		} else {
//			return false
//		}
//	}
//	return openBrackets == ""
//}

// longestValidParentheses is the thirty-second LeetCode problem:
// - https://leetcode.com/problems/longest-valid-parentheses/
func longestValidParentheses(s string) int {
	fmt.Printf("longestValidParentheses(%s)\n", s)
	if s == "" || len(s) == 1 {
		return 0
	}

	var longest int
	var skip bool
	for left := 0; left < len(s)-1; left++ {
		if len(s)-longest < left {
			break
		}

		skip = false
		for right := left + 1; right < len(s); right++ {
			fmt.Printf("  left, right: (%d, %d)\n", left, right)
			if substring := s[left : right+1]; isValid(substring) {
				fmt.Printf("    substring: %s\n", substring)
				if longest < len(substring) {
					longest = len(substring)

					// If it's a valid string, then skip left to the end
					skip = true
				}
			}
		}
		if skip {
			left += longest
		}
	}
	return longest
}

var LongestValidParentheses = longestValidParentheses /* For testing */
