package solutions

// Results is a helper struct for LongestPalindrome.
// It stores the current longest palindrome and its length.
type Results struct {
	length int
	word   string
}

// UpdateIfLarger updates the results if the new palindrome is longer
// than the current one.
func (results *Results) UpdateIfLarger(str string) {
	if newLen := len(str); newLen > results.length {
		results.length = newLen
		results.word = str
	}
}

// isPalindrome checks whether a string is a palindrome.
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

var IsPalindrome_ = isPalindrome /* For testing */

// LongestPalindrome is the fifth LeetCode problem:
// - https://leetcode.com/problems/longest-palindromic-substring/
func LongestPalindrome(s string) string {
	results := &Results{}
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			word := s[i : j+1]

			if isPalindrome(word) {
				results.UpdateIfLarger(word)
			}
		}
	}

	return results.word
}
