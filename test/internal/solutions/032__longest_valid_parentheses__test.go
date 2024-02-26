package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{s: "", want: 0},
		{s: "(()", want: 2},
		{s: ")()())", want: 4},
		{s: ")))()(()()))(((((((((", want: 8},
		{s: "(((((())))))", want: 12},
	}

	for _, test := range tests {
		got := solutions.LongestValidParentheses(test.s)
		if got != test.want {
			t.Fatalf("LongestValidParentheses(%s) == %d, want %d", test.s, got, test.want)
		}
	}
}
