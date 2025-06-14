package solutions_test

import (
	"github.com/billwallis/leetcode-go/internal/solutions"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{s: "III", want: 3},
		{s: "LVIII", want: 58},
		{s: "MCMXCIV", want: 1994},
		{s: "DCXXI", want: 621},
	}

	for _, test := range tests {
		got := solutions.RomanToInt(test.s)
		if got != test.want {
			t.Fatalf("RomanToInt(%s) == %d, want %d", test.s, got, test.want)
		}
	}
}
