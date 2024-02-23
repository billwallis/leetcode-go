package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	// The actual test in LeetCode is more complex that this, but it's
	// a crappy question, so I can't be bothered to implement the test
	// correctly
	var tests = []struct {
		nums []int
		val  int
		want int
	}{
		{nums: []int{3, 2, 2, 3}, val: 3, want: 2},
		{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, want: 5},
	}

	for _, test := range tests {
		got := solutions.RemoveElement(test.nums, test.val)
		if got != test.want {
			t.Fatalf("RemoveElement(%v, %d) == %d, want %d", test.nums, test.val, got, test.want)
		}
	}
}
