package solutions_test

import (
	"github.com/billwallis/leetcode-go/internal/solutions"
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	var tests = []struct {
		nums []int
		want [][]int
	}{
		{nums: []int{-1, 0, 1, 2, -1, -4}, want: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{nums: []int{0, 1, 1}, want: [][]int{}},
		{nums: []int{0, 0, 0}, want: [][]int{{0, 0, 0}}},
	}

	for _, test := range tests {
		got := solutions.ThreeSum(test.nums)

		sort.Slice(got, func(i, j int) bool {
			// The arrays always have 3 elements
			position1 := got[i][0] < got[j][0]
			position2 := (got[i][0] == got[j][0]) && got[i][1] < got[j][1]
			position3 := (got[i][0] == got[j][0] && got[i][1] == got[j][1]) && got[i][2] < got[j][2]

			return position1 || position2 || position3
		})

		if len(got) == 0 && len(test.want) == 0 {
			// Pass
		} else if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("ThreeSum(%v) == %v, want %v", test.nums, got, test.want)
		}
	}
}
