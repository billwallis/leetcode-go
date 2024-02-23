package solutions

// removeDuplicates is the twenty-sixth LeetCode problem:
// - https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	left := 0
	for i := 1; i < len(nums); i++ {
		if nums[left] != nums[i] {
			left++
			nums[left] = nums[i]
		}
	}
	return left + 1
}

var RemoveDuplicates = removeDuplicates /* For testing */
