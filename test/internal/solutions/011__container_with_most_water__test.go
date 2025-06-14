package solutions_test

import (
	"github.com/billwallis/leetcode-go/internal/solutions"
	"testing"
)

func TestMaxArea(t *testing.T) {
	var tests = []struct {
		height []int
		want   int
	}{
		{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, want: 49},
		{height: []int{1, 1}, want: 1},
	}

	for _, test := range tests {
		got := solutions.MaxArea(test.height)
		if got != test.want {
			t.Fatalf("MaxArea(%v) == %d, want %d", test.height, got, test.want)
		}
	}
}
