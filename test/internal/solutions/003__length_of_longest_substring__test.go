package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"testing"
)

func TestIndexOf(t *testing.T) {
	var tests = []struct {
		element any
		data    []any
		want    int
	}{
		{element: "b", data: []any{"a", "b", "c"}, want: 1},
		{element: "x", data: []any{"a", "b", "c"}, want: -1},
		{element: 0, data: []any{0, 1, 2}, want: 0},
		{element: 4, data: []any{0, 1, 2}, want: -1},
	}

	for _, test := range tests {
		got := solutions.IndexOf(test.element, test.data)
		if got != test.want {
			t.Fatalf(`IndexOf(%v, %v) == %d, want %d`, test.element, test.data, got, test.want)
		}
	}
}

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
