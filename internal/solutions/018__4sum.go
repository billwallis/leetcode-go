package solutions

import "sort"

// fourSum is the eighteenth LeetCode problem:
// - https://leetcode.com/problems/4sum/
func fourSum(nums []int, target int) [][]int {
	var results [][]int

	if len(nums) < 4 {
		return [][]int{}
	} else if len(nums) == 4 {
		if target == (nums[0] + nums[1] + nums[2] + nums[3]) {
			return [][]int{nums}
		}
	}

	sort.Ints(nums)
	for leftmostIndex := 0; leftmostIndex < len(nums)-3; leftmostIndex++ {
		if leftmostIndex > 0 && nums[leftmostIndex-1] == nums[leftmostIndex] {
			continue
		}
		for rightmostIndex := len(nums) - 1; rightmostIndex > (leftmostIndex + 2); rightmostIndex-- {
			if rightmostIndex < len(nums)-1 && nums[rightmostIndex] == nums[rightmostIndex+1] {
				continue
			}
			leftmostNumber, rightmostNumber := nums[leftmostIndex], nums[rightmostIndex]
			leftIndex, rightIndex := leftmostIndex+1, rightmostIndex-1
			for leftIndex < rightIndex {
				leftNumber, rightNumber := nums[leftIndex], nums[rightIndex]
				sum := leftmostNumber + leftNumber + rightNumber + rightmostNumber

				if sum == target {
					results = append(results, []int{leftmostNumber, leftNumber, rightNumber, rightmostNumber})
				}

				// If the sum is smaller than the target, then we need to _increase_
				// our sum by moving the left index up; otherwise, we need to
				// _decrease_ our sum by moving the right index down.
				//
				// If the sum equals the target, we advance both indexes.
				if sum <= target {
					for leftIndex < rightIndex && nums[leftIndex] == leftNumber {
						leftIndex++
					}
				}
				if sum >= target {
					for leftIndex < rightIndex && nums[rightIndex] == rightNumber {
						rightIndex--
					}
				}
			}
		}
	}

	return results
}

var FourSum = fourSum /* For testing */
