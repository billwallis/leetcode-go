package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	var tests = []struct {
		list1 []int
		list2 []int
		want  []int
	}{
		{list1: []int{1, 2, 4}, list2: []int{1, 3, 4}, want: []int{1, 1, 2, 3, 4, 4}},
		{list1: []int{}, list2: []int{}, want: []int{}},
		{list1: []int{}, list2: []int{0}, want: []int{0}},
	}

	for _, test := range tests {
		node1 := solutions.NewListNode(test.list1)
		node2 := solutions.NewListNode(test.list2)
		want := solutions.NewListNode(test.want)
		got := solutions.MergeTwoLists(node1, node2)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("MergeTwoLists(%v, %v) == %v, want %v", test.list1, test.list2, got, test.want)
		}
	}
}
