package solutions

// indexOf is a helper function for LengthOfLongestSubstring.
// It returns the index of an element in a slice.
func indexOf(element any, data []any) int {
	// Nicked from https://stackoverflow.com/q/8307478/8213085
	for index, value := range data {
		if element == value {
			return index
		}
	}
	return -1
}

// LengthOfLongestSubstring is the third LeetCode problem:
// - https://leetcode.com/problems/longest-substring-without-repeating-characters/
func LengthOfLongestSubstring(s string) int {
	var chars []any
	var maxLength = 0

	for _, char := range s {
		index := indexOf(char, chars)
		chars = append(chars, char)

		if index == -1 {
			maxLength = max(maxLength, len(chars))
		} else {
			chars = chars[index+1:]
		}
	}

	return maxLength
}
