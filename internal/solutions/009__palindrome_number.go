package solutions

// IsPalindrome is the ninth LeetCode problem:
// - https://leetcode.com/problems/palindrome-number/
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// Reverse the number
	var reversed int
	for temp := x; temp != 0; temp /= 10 {
		reversed = reversed*10 + temp%10
	}

	return x == reversed
}
