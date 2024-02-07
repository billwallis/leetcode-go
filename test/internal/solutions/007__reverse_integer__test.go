package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		x    int
		want int
	}{
		{x: 123, want: 321},
		{x: -123, want: -321},
		{x: 120, want: 21},
		{x: 1534236469, want: 0},
	}

	for _, test := range tests {
		got := solutions.Reverse(test.x)
		if got != test.want {
			t.Fatalf("Reverse(%d) == %d, want %d", test.x, got, test.want)
		}
	}
}
