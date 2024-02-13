package solutions

import (
	"sort"
)

// threeSum is the fifteenth LeetCode problem:
// - https://leetcode.com/problems/3sum/
func threeSum(nums []int) [][]int {
	// Nicked the logic from an existing solution since my solution timed out:
	// - https://leetcode.com/problems/3sum/solutions/295224/go-768-ms-faster-than-100-00

	// If we sort the numbers, we can efficiently walk through them
	sort.Ints(nums)
	var results [][]int

	for i := 0; i < len(nums)-2; {
		var leftNumber, rightNumber int
		leftIndex, rightIndex := i+1, len(nums)-1
		currentNumber := nums[i]

		for leftIndex < rightIndex {
			leftNumber, rightNumber = nums[leftIndex], nums[rightIndex]
			sum := currentNumber + leftNumber + rightNumber
			if sum == 0 {
				results = append(results, []int{currentNumber, leftNumber, rightNumber})
			}

			// Increasing the left index increases our sum; decreasing the right
			// index decreases our sum. Since we want the sum to be `0`, adjust
			// the index that will get us closer to `0`!
			//
			// If the sum _is_ `0`, we want the next pair of numbers.
			//
			// We always want to go to the next _different_ number to avoid
			// duplicates.
			if sum <= 0 {
				for leftIndex < rightIndex && leftNumber == nums[leftIndex] {
					leftIndex++
				}
			}
			if sum >= 0 {
				for leftIndex < rightIndex && rightNumber == nums[rightIndex] {
					rightIndex--
				}
			}
		}
		// Similarly, we want the `i` index to skip any repeated numbers
		if currentNumber < nums[i+1] {
			i++
		} else {
			for i < len(nums)-2 && currentNumber == nums[i] {
				i++
			}
		}
	}

	return results
}

var ThreeSum = threeSum /* For testing */
