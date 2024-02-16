package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"testing"
)

func TestNodeLength(t *testing.T) {
	var tests = []struct {
		node []int
		want int
	}{
		{node: []int{0, 1, 2, 3, 4, 5}, want: 6},
		{node: []int{}, want: 0},
		{node: []int{1}, want: 1},
	}

	for _, test := range tests {
		node := solutions.NewListNode(test.node)
		got := solutions.NodeLength(node)
		if got != test.want {
			t.Fatalf("NodeLength(%v) == %d, want %d", test.node, got, test.want)
		}
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	var tests = []struct {
		head []int
		n    int
		want []int
	}{
		{head: []int{1, 2, 3, 4, 5}, n: 2, want: []int{1, 2, 3, 5}},
		{head: []int{1}, n: 1, want: []int{}},
		{head: []int{1, 2}, n: 1, want: []int{1}},
		{head: []int{1, 2}, n: 2, want: []int{2}},
		{head: []int{1, 2, 3, 4}, n: 0, want: []int{1, 2, 3, 4}},
	}

	for _, test := range tests {
		input := solutions.NewListNode(test.head)
		want := solutions.NewListNode(test.want)

		got := solutions.RemoveNthFromEnd(input, test.n)
		if !got.Equals(want) {
			//t.Fatalf("RemoveNthFromEnd(%v, %d) == %v, want %v", test.head, test.n, got, test.want)
			t.Fatalf("RemoveNthFromEnd(%v, %d) == %v, want %v", input.AsString(), test.n, got.AsString(), want.AsString())
		}
	}
}
