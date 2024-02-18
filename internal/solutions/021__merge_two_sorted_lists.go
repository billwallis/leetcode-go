package solutions

// mergeTwoLists is the twenty-first LeetCode problem:
// - https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list2 == nil {
		return list1
	} else if list1 == nil {
		return list2
	}

	if list1.Val < list2.Val {
		return &ListNode{
			Val:  list1.Val,
			Next: mergeTwoLists(list1.Next, list2),
		}
	} else {
		return &ListNode{
			Val:  list2.Val,
			Next: mergeTwoLists(list1, list2.Next),
		}
	}
}

var MergeTwoLists = mergeTwoLists /* For testing */
