package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	// The real solution doesn't return anyway, but I don't know how to
	// test that, so Imma return stuff anyway
	var tests = []struct {
		nums []int
		want []int
	}{
		{nums: []int{1, 2, 3}, want: []int{1, 3, 2}},
		{nums: []int{3, 2, 1}, want: []int{1, 2, 3}},
		{nums: []int{1, 1, 5}, want: []int{1, 5, 1}},
		{nums: []int{1, 3, 4, 2}, want: []int{1, 4, 2, 3}},
		{nums: []int{1, 4, 2, 3}, want: []int{1, 4, 3, 2}},
		{nums: []int{1, 4, 3, 2}, want: []int{2, 1, 3, 4}},
	}

	for _, test := range tests {
		got := solutions.NextPermutation(test.nums)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("NextPermutation(%v) == %v, want %v", test.nums, got, test.want)
		}
	}
}
