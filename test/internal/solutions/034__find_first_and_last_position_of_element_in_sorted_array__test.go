package solutions_test

import (
	"github.com/billwallis/leetcode-go/internal/solutions"
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   []int
	}{
		{nums: []int{}, target: 0, want: []int{-1, -1}},
		{nums: []int{1}, target: 1, want: []int{0, 0}},
		{nums: []int{2, 2}, target: 2, want: []int{0, 1}},
		{nums: []int{1, 3}, target: 1, want: []int{0, 0}},
		{nums: []int{1, 3}, target: 3, want: []int{1, 1}},
		{nums: []int{2, 2, 2, 2, 2, 2, 2, 2}, target: 2, want: []int{0, 7}},
		{nums: []int{5, 7, 7, 8, 8, 10}, target: 8, want: []int{3, 4}},
		{nums: []int{5, 7, 7, 8, 8, 10}, target: 6, want: []int{-1, -1}},
	}

	for _, test := range tests {
		got := solutions.SearchRange(test.nums, test.target)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("SearchRange(%v, %d) == %v, want %v", test.nums, test.target, got, test.want)
		}
	}
}
