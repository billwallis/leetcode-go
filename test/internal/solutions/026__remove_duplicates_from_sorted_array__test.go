package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	// The actual test in LeetCode is more complex that this, but it's
	// a crappy question, so I can't be bothered to implement the test
	// correctly
	var tests = []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 1, 2}, want: 2},
		{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, want: 5},
	}

	for _, test := range tests {
		got := solutions.RemoveDuplicates(test.nums)
		if got != test.want {
			t.Fatalf("RemoveDuplicates(%v) == %d, want %d", test.nums, got, test.want)
		}
	}
}
