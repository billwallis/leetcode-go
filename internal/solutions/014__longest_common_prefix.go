package solutions

// MinLengthOfStrings returns the length of the shortest string given a
// slice of strings.
func MinLengthOfStrings(strings []string) int {
	if len(strings) < 1 {
		return 0
	}

	minLen := len(strings[0])
	for i := 1; i < len(strings); i++ {
		if sLen := len(strings[i]); sLen < minLen {
			minLen = sLen
		}
	}
	return minLen
}

// AreCharsAtPositionEqual checks whether the character at `position` is
// the same for each string in the given slice of strings.
func AreCharsAtPositionEqual(strings []string, position int) bool {
	match := true
	for i := 1; match && i < len(strings); i++ {
		match = match && (strings[i-1][position] == strings[i][position])
	}
	return match
}

// longestCommonPrefix is the fourteenth LeetCode problem:
// - https://leetcode.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	prefix := ""
	for i := 0; i < MinLengthOfStrings(strs); i++ {
		if AreCharsAtPositionEqual(strs, i) {
			prefix += string(strs[0][i]) // If they're equal, we can take any -- hence `0`
		} else {
			break
		}
	}
	return prefix
}

var LongestCommonPrefix = longestCommonPrefix /* For testing */
