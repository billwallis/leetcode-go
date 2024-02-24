package solutions

// findSubstring is the thirtieth LeetCode problem:
// - https://leetcode.com/problems/substring-with-concatenation-of-all-words/
func findSubstring(s string, words []string) []int {
	wordLength := len(words[0]) // They all have the same length
	concatenatedSubstringLength := wordLength * len(words)
	if concatenatedSubstringLength > len(s) {
		return []int{}
	}

	startingIndexes := []int{}
	for index := 0; index <= len(s)-concatenatedSubstringLength; index++ {
		remainingWords := make([]string, len(words))
		_ = copy(remainingWords, words)
		// At the index, check whether all words are present
		for wordIndex := index; len(remainingWords) > 0; wordIndex += wordLength {
			wordAtIndex := s[wordIndex : wordIndex+wordLength]
			found := false
			for i, w := range remainingWords {
				if wordAtIndex == w {
					found = true
					// Remove the element
					remainingWords[i] = remainingWords[len(remainingWords)-1]
					remainingWords = remainingWords[:len(remainingWords)-1]
					break
				}
			}
			if !found {
				break
			}
		}
		// If all words are present, append this index to the solution slice
		if len(remainingWords) == 0 {
			startingIndexes = append(startingIndexes, index)
		}
	}

	return startingIndexes
}

var FindSubstring = findSubstring /* For testing */
