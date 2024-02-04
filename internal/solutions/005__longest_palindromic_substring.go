package solutions

type Results struct {
	length int
	word   string
}

func (results *Results) UpdateIfLarger(str string) {
	if newLen := len(str); newLen > results.length {
		results.length = newLen
		results.word = str
	}
}

func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func LongestPalindrome(s string) string {
	results := &Results{}
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			word := s[i : j+1]

			if IsPalindrome(word) {
				results.UpdateIfLarger(word)
			}
		}
	}

	return results.word
}
