package solutions_test

import (
	"github.com/billwallis/leetcode-go/internal/solutions"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   int
	}{
		{nums: []int{1, 3, 5, 6}, target: 5, want: 2},
		{nums: []int{1, 3, 5, 6}, target: 2, want: 1},
		{nums: []int{1, 3, 5, 6}, target: 7, want: 4},
		{nums: []int{1, 3, 5, 6}, target: 0, want: 0},
	}

	for _, test := range tests {
		got := solutions.SearchInsert(test.nums, test.target)
		if got != test.want {
			t.Fatalf("SearchInsert(%v, %d) == %d, want %d", test.nums, test.target, got, test.want)
		}
	}
}
