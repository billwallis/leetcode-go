package solutions

// mergeKLists is the twenty-third LeetCode problem:
// - https://leetcode.com/problems/merge-k-sorted-lists/
func mergeKLists(lists []*ListNode) *ListNode {
	var notNilLists []*ListNode
	for _, list := range lists {
		if list != nil {
			notNilLists = append(notNilLists, list)
		}
	}
	if len(notNilLists) == 0 {
		return nil
	}

	shortestNodeIndex := 0
	shortestNodeValue := 10_000
	for index, node := range notNilLists {
		if node.Val < shortestNodeValue {
			shortestNodeIndex = index
			shortestNodeValue = node.Val
		}
	}

	var adjustedLists []*ListNode
	for index, node := range notNilLists {
		if index == shortestNodeIndex {
			adjustedLists = append(adjustedLists, node.Next)
		} else {
			adjustedLists = append(adjustedLists, node)
		}
	}

	return &ListNode{
		Val:  shortestNodeValue,
		Next: mergeKLists(adjustedLists),
	}
}

var MergeKLists = mergeKLists /* For testing */
