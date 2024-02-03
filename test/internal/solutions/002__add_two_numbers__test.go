package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"testing"
)

func TestNewListNode(t *testing.T) {
	got := solutions.NewListNode([]int{1, 2, 3})
	want := &solutions.ListNode{
		Val: 1,
		Next: &solutions.ListNode{
			Val: 2,
			Next: &solutions.ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	if !got.Equals(want) {
		t.Fatalf(`NewListNode(%v) = %v, want %v`, []int{1, 2, 3}, got.AsString(), want.AsString())
	}
}

func TestAddTwoNumbers(t *testing.T) {
	var tests = []struct {
		l1   []int
		l2   []int
		want []int
	}{
		{l1: []int{2, 4, 3}, l2: []int{5, 6, 4}, want: []int{7, 0, 8}},
		{l1: []int{0}, l2: []int{0}, want: []int{0}},
		{l1: []int{9, 9, 9, 9, 9, 9, 9}, l2: []int{9, 9, 9, 9}, want: []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for _, test := range tests {
		number1 := solutions.NewListNode(test.l1)
		number2 := solutions.NewListNode(test.l2)
		want := solutions.NewListNode(test.want)
		got := solutions.AddTwoNumbers(number1, number2)

		if !got.Equals(want) {
			t.Fatalf(`AddTwoNumbers(%v, %v) = %v, want %v`, test.l1, test.l2, got.AsString(), want.AsString())
		}
	}
}
