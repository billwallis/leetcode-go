package solutions_test

import (
	"github.com/billwallis/leetcode-go/internal/solutions"
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	t.Skip("Not implemented yet...")

	var tests = []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{candidates: []int{2}, target: 1, want: [][]int{}},
		{candidates: []int{2, 3, 6, 7}, target: 7, want: [][]int{{2, 2, 3}, {7}}},
		{candidates: []int{2, 3, 5}, target: 8, want: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
	}

	for _, test := range tests {
		got := solutions.CombinationSum(test.candidates, test.target)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("CombinationSum(%v, %d) == %v, want %v", test.candidates, test.target, got, test.want)
		}
	}
}
