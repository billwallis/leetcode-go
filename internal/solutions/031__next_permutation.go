package solutions

import "sort"

// nextPermutation is the thirty-first LeetCode problem:
// - https://leetcode.com/problems/next-permutation/
func nextPermutation(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	for left := len(nums) - 2; left >= 0; left-- {
		for right := len(nums) - 1; left < right; right-- {
			if nums[left] < nums[right] {
				nums[left], nums[right] = nums[right], nums[left]
				sort.Ints(nums[left+1:])
				return nums
			}
		}
	}

	// If no swapping, then it's the "largest" permutation, so we reset
	sort.Ints(nums)
	return nums
}

var NextPermutation = nextPermutation /* For testing */
