package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	var tests = []struct {
		head []int
		want []int
	}{
		{head: []int{}, want: []int{}},
		{head: []int{1}, want: []int{1}},
		{head: []int{1, 2, 3, 4}, want: []int{2, 1, 4, 3}},
		{head: []int{1, 2, 3, 4, 5}, want: []int{2, 1, 4, 3, 5}},
	}

	for _, test := range tests {
		head := solutions.NewListNode(test.head)
		want := solutions.NewListNode(test.want)

		got := solutions.SwapPairs(head)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("SwapPairs(%v) == %v, want %v", test.head, got, test.want)
		}
	}
}
