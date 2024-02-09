package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		x    int
		want bool
	}{
		{x: 121, want: true},
		{x: -121, want: false},
		{x: 10, want: false},
	}

	for _, test := range tests {
		got := solutions.IsPalindrome(test.x)
		if got != test.want {
			t.Fatalf("IsPalindrome(%d) == %v, want %v", test.x, got, test.want)
		}
	}
}
