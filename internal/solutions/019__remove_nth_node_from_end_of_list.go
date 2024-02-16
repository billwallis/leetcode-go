package solutions

func NodeLength(node *ListNode) int {
	if node == nil {
		return 0
	} else {
		return 1 + NodeLength(node.Next)
	}
}

func skipIfN(node *ListNode, i int, n int) (*ListNode, int) {
	var nextNode *ListNode
	i--

	if i == 0 {
		return nil, i
	} else if i == n {
		nextNode = node.Next.Next
	} else {
		nextNode = node.Next
	}

	newNode, _ := skipIfN(nextNode, i, n)
	return &ListNode{Val: node.Val, Next: newNode}, i
}

// removeNthFromEnd is the nineteenth LeetCode problem:
// - https://leetcode.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := NodeLength(head)
	if n == length {
		return head.Next
	}

	newNode, _ := skipIfN(head, length, n)
	return newNode
}

var RemoveNthFromEnd = removeNthFromEnd /* For testing */
