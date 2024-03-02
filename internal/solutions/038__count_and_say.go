package solutions

import "strconv"

func Say(word string) string {
	if len(word) == 1 {
		return "1" + word
	}

	var count int
	var newWord string
	for index := 0; index < len(word); index++ {
		count++
		if index == len(word)-1 || word[index] != word[index+1] {
			newWord += strconv.Itoa(count) + string(word[index])
			count = 0
		}
	}

	return newWord
}

// countAndSay is the thirty-eighth LeetCode problem:
// - https://leetcode.com/problems/count-and-say/
func countAndSay(n int) string {
	word := "1"
	for count := 2; count <= n; count++ {
		word = Say(word)
	}
	return word
}

var CountAndSay = countAndSay /* For testing */
