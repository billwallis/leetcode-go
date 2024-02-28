package solutions

// searchRange is the thirty-fourth LeetCode problem:
// - https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	// Two binary searches are probably more correct -- if the values in
	// nums are all the same, this is O(n), but two binary searches
	// would still be O(2 log n) ~ O(log n)
	if len(nums) == 0 {
		return []int{-1, -1}
	} else if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}

	leftIndex, middleIndex, rightIndex := 0, 0, len(nums)-1
	for leftIndex < rightIndex {
		if nums[leftIndex] == target {
			middleIndex = leftIndex
		} else if nums[rightIndex] == target {
			middleIndex = rightIndex
		} else {
			middleIndex = leftIndex + ((rightIndex - leftIndex) / 2)
		}
		if nums[middleIndex] < target {
			leftIndex = middleIndex + 1
		} else if nums[middleIndex] > target {
			rightIndex = middleIndex - 1
		} else {
			// Loop to find the bounds
			leftIndex, rightIndex = middleIndex, middleIndex
			for leftIndex > 0 && nums[leftIndex-1] == target {
				leftIndex--
			}
			for rightIndex < len(nums)-1 && nums[rightIndex+1] == target {
				rightIndex++
			}
			return []int{leftIndex, rightIndex}
		}
	}

	return []int{-1, -1}
}

var SearchRange = searchRange /* For testing */
