package solutions_test

import (
	"github.com/billwallis/leetcode-go/internal/solutions"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	var tests = []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{nums1: []int{1, 3}, nums2: []int{2}, want: 2.00000},
		{nums1: []int{1, 2}, nums2: []int{3, 4}, want: 2.50000},
		{nums1: []int{1, 3, 5}, nums2: []int{2, 4, 6}, want: 3.50000},
	}

	for _, test := range tests {
		if got := solutions.FindMedianSortedArrays(test.nums1, test.nums2); got != test.want {
			t.Fatalf(`FindMedianSortedArrays(%v, %v) = %2f, want %2f`, test.nums1, test.nums2, got, test.want)
		}
	}
}
