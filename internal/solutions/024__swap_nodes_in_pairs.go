package solutions

// This the first solution, before realising that this swaps _values_
// when the problem asked to swap _nodes_.
/*
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	return &ListNode{
		Val: head.Next.Val,
		Next: &ListNode{
			Val:  head.Val,
			Next: swapPairs(head.Next.Next),
		},
	}
}
*/

// swapPairs is the twenty-fourth LeetCode problem:
// - https://leetcode.com/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	swap := head.Next
	head.Next = swapPairs(head.Next.Next)
	swap.Next = head

	return swap
}

var SwapPairs = swapPairs /* For testing */
