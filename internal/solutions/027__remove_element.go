package solutions

// removeElement is the twenty-seventh LeetCode problem:
// - https://leetcode.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}

var RemoveElement = removeElement /* For testing */
