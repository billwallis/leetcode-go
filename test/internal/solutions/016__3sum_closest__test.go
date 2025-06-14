package solutions_test

import (
	"github.com/billwallis/leetcode-go/internal/solutions"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   int
	}{
		{nums: []int{-1, 2, 1, -4}, target: 1, want: 2},
		{nums: []int{0, 0, 0}, target: 1, want: 0},
		{nums: []int{4, 0, 5, -5, 3, 3, 0, -4, -5}, target: -2, want: -2},
	}

	for _, test := range tests {
		got := solutions.ThreeSumClosest(test.nums, test.target)
		if got != test.want {
			t.Fatalf("ThreeSumClosest(%v, %d) == %d, want %d", test.nums, test.target, got, test.want)
		}
	}
}
