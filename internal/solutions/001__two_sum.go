package solutions

// TwoSum is the first LeetCode problem:
// - https://leetcode.com/problems/two-sum/
func TwoSum(nums []int, target int) []int {
	for index1, element1 := range nums {
		for index2, element2 := range nums[index1+1:] {
			// The second index will be offset by `index_1 + 1` so we
			// need to add it back on in the return.
			if (element1 + element2) == target {
				return []int{index1, index2 + index1 + 1}
			}
		}
	}
	return []int{}
}
