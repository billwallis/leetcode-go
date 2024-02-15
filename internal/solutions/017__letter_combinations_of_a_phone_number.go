package solutions

var digitToLetters = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

// letterCombinations is the seventeenth LeetCode problem:
// - https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	results := []string{""}
	for _, char := range digits {
		letters := digitToLetters[string(char)]
		var newResults []string
		for _, result := range results {
			for _, letter := range letters {
				newResults = append(newResults, result+letter)
			}
		}
		results = newResults
	}
	return results
}

var LetterCombinations = letterCombinations /* For testing */
