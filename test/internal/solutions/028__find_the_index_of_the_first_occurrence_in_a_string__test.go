package solutions_test

import (
	"github.com/billwallis/leetcode-go/internal/solutions"
	"testing"
)

func TestStrStr(t *testing.T) {
	var tests = []struct {
		haystack string
		needle   string
		want     int
	}{
		{haystack: "sadbutsad", needle: "sad", want: 0},
		{haystack: "leetcode", needle: "leeto", want: -1},
		{haystack: "abcd", needle: "bc", want: 1},
	}

	for _, test := range tests {
		got := solutions.StrStr(test.haystack, test.needle)
		if got != test.want {
			t.Fatalf("StrStr(%s, %s) == %d, want %d", test.haystack, test.needle, got, test.want)
		}
	}
}
