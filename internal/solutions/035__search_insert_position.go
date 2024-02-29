package solutions

// searchInsert is the thirty-fifth LeetCode problem:
// - https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	leftIndex, rightIndex := 0, len(nums)-1
	middleIndex, middle := 0, 0

	for leftIndex < rightIndex {
		middleIndex = (leftIndex + rightIndex) / 2
		middle = nums[middleIndex]
		if middle >= target {
			rightIndex = middleIndex
		} else if middle < target {
			leftIndex = middleIndex + 1
		}
	}

	if nums[leftIndex] == target {
		return leftIndex
	} else if target > nums[rightIndex] {
		return rightIndex + 1
	} else {
		return leftIndex
	}
}

var SearchInsert = searchInsert /* For testing */
