package solutions

// longestValidParentheses is the thirty-second LeetCode problem:
// - https://leetcode.com/problems/longest-valid-parentheses/
func longestValidParentheses(s string) int {
	// Approach nicked from:
	// - https://leetcode.com/problems/longest-valid-parentheses/solutions/1139974/python-c-go-o-n-by-stack-w-comment
	if s == "" || len(s) == 1 {
		return 0
	}

	var longest int
	stack := []int{-1}

	for index, char := range s {
		if char == '(' {
			stack = append(stack, index)
		} else {
			// Pop the stack
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				stack = append(stack, index)
			} else {
				if newLen := index - stack[len(stack)-1]; longest < newLen {
					longest = newLen
				}
			}
		}
	}

	return longest
}

var LongestValidParentheses = longestValidParentheses /* For testing */
