package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var tests = []struct {
		text string
		want int
	}{
		{text: "abcabcbb", want: 3},
		{text: "bbbbb", want: 1},
		{text: "pwwkew", want: 3},
	}

	for _, test := range tests {
		if got := solutions.LengthOfLongestSubstring(test.text); got != test.want {
			t.Fatalf(`LengthOfLongestSubstring(%s) = %d, want %d`, test.text, got, test.want)
		}
	}
}
