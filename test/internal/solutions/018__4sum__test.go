package solutions_test

import (
	"github.com/billwallis/leetcode-go/internal/solutions"
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   [][]int
	}{
		{nums: []int{1, 0, -1, 0, -2, 2}, target: 0, want: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{nums: []int{2, 2, 2, 2, 2}, target: 8, want: [][]int{{2, 2, 2, 2}}},
		{nums: []int{0, 0, 0, 0}, target: 0, want: [][]int{{0, 0, 0, 0}}},
	}

	for _, test := range tests {
		got := solutions.FourSum(test.nums, test.target)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("FourSum(%v, %d) == %v, want %v", test.nums, test.target, got, test.want)
		}
	}
}
