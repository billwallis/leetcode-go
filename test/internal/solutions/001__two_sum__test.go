package solutions_test

import (
	"github.com/billwallis/leetcode-go/internal/solutions"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, want: []int{0, 1}},
	}

	for _, test := range tests {
		if got := solutions.TwoSum(test.nums, test.target); !reflect.DeepEqual(got, test.want) {
			t.Fatalf("TwoSum(%v, %v) = %v, want %v", test.nums, test.target, got, test.want)
		}
	}
}
