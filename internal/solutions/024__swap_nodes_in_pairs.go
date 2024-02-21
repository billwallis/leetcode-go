package solutions

// swapPairs is the twenty-fourth LeetCode problem:
// - https://leetcode.com/problems/swap-nodes-in-pairs/
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

var SwapPairs = swapPairs /* For testing */
