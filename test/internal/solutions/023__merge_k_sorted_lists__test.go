package solutions_test

import (
	"github.com/billwallis/leetcode-go/internal/solutions"
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	var tests = []struct {
		lists [][]int
		want  []int
	}{
		{lists: [][]int{}, want: []int{}},
		{lists: [][]int{{}}, want: []int{}},
		{lists: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, want: []int{1, 1, 2, 3, 4, 4, 5, 6}},
	}

	for _, test := range tests {
		var lists []*solutions.ListNode
		for _, list := range test.lists {
			lists = append(lists, solutions.NewListNode(list))
		}
		want := solutions.NewListNode(test.want)

		got := solutions.MergeKLists(lists)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("MergeKLists(%v) == %v, want %v", test.lists, got, test.want)
		}
	}
}
