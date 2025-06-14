package solutions_test

import (
	"github.com/billwallis/leetcode-go/internal/solutions"
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	var tests = []struct {
		head []int
		k    int
		want []int
	}{
		{head: []int{1, 2, 3, 4, 5}, k: 2, want: []int{2, 1, 4, 3, 5}},
		{head: []int{1, 2, 3, 4, 5}, k: 3, want: []int{3, 2, 1, 4, 5}},
		{head: []int{1, 2, 3, 4, 5}, k: 4, want: []int{4, 3, 2, 1, 5}},
	}

	for _, test := range tests {
		head := solutions.NewListNode(test.head)
		want := solutions.NewListNode(test.want)

		got := solutions.ReverseKGroup(head, test.k)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("ReverseKGroup(%v, %d) == %v, want %v", test.head, test.k, got.AsString(), test.want)
		}
	}
}
