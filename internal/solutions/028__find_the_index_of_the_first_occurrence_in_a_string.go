package solutions

// strStr is the twenty-eighth LeetCode problem:
// - https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
func strStr(haystack string, needle string) int {
	lenHaystack, lenNeedle := len(haystack), len(needle)
	for left := 0; left < len(haystack); left++ {
		if lenHaystack-left < lenNeedle {
			continue
		}
		if haystack[left:left+lenNeedle] == needle {
			return left
		}
	}
	return -1
}

var StrStr = strStr /* For testing */
