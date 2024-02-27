package solutions

// search is the thirty-third LeetCode problem:
// - https://leetcode.com/problems/search-in-rotated-sorted-array/
func search(nums []int, target int) int {
	var leftIndex, rightIndex, rotationPoint int
	length := len(nums)

	// Binary search for rotation point
	leftIndex, rightIndex = 0, length-1
	for leftIndex < rightIndex {
		middleIndex := leftIndex + ((rightIndex - leftIndex) / 2)
		if nums[middleIndex] > nums[rightIndex] {
			leftIndex = middleIndex + 1
		} else {
			rightIndex = middleIndex
		}
	}
	rotationPoint = leftIndex

	// Binary search for target
	// We'll account for the rotation by starting the search at the rotation
	leftIndex, rightIndex = rotationPoint, rotationPoint-1+length
	for leftIndex <= rightIndex {
		middleIndex := leftIndex + ((rightIndex - leftIndex) / 2)
		middle := nums[middleIndex%length] // Use mod to account for rotation

		if middle > target {
			rightIndex = middleIndex - 1
		} else if middle < target {
			leftIndex = middleIndex + 1
		} else {
			return middleIndex % length
		}
	}

	return -1
}

var Search = search /* For testing */
