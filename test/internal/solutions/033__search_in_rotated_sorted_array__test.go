package solutions_test

import (
	"github.com/billwallis/leetcode-go/internal/solutions"
	"testing"
)

func TestSearch(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   int
	}{
		{nums: []int{0, 1, 2, 3}, target: 0, want: 0},
		{nums: []int{0, 1, 2, 4, 5, 6, 7}, target: 0, want: 0},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, want: 4},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3, want: -1},
		{nums: []int{1}, target: 0, want: -1},
	}

	for _, test := range tests {
		got := solutions.Search(test.nums, test.target)
		if got != test.want {
			t.Fatalf("Search(%v, %d) == %d, want %d", test.nums, test.target, got, test.want)
		}
	}
}
