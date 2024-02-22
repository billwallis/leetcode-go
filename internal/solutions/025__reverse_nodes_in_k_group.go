package solutions

// reverseKGroup is the twenty-fifth LeetCode problem:
// - https://leetcode.com/problems/reverse-nodes-in-k-group/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || NodeLength(head) < k {
		return head
	}

	var prev, last *ListNode
	// Reverse the k nodes on the left
	// Nicked from https://stackoverflow.com/q/76377587/8213085
	for i := 0; i < k; i++ {
		if i == 0 {
			last = head
		}
		prev, head, head.Next = head, head.Next, prev
	}
	last.Next = reverseKGroup(head, k)

	return prev
}

var ReverseKGroup = reverseKGroup /* For testing */
