package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"testing"
)

func TestIsPalindrome_(t *testing.T) {
	var tests = []struct {
		s    string
		want bool
	}{
		{s: "babab", want: true},
		{s: "cbbda", want: false},
	}

	for _, test := range tests {
		if got := solutions.IsPalindrome_(test.s); got != test.want {
			t.Fatalf(`IsPalindrome(%s) = %v, want %v`, test.s, got, test.want)
		}
	}
}

func TestLongestPalindrome(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{s: "babad", want: "bab"}, // "aba" is also a valid answer
		{s: "cbbd", want: "bb"},
	}

	for _, test := range tests {
		if got := solutions.LongestPalindrome(test.s); got != test.want {
			t.Fatalf(`LongestPalindrome(%s) = %s, want %s`, test.s, got, test.want)
		}
	}
}
