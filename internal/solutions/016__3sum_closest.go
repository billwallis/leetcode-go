package solutions

import (
	"sort"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

type ThreeSumClosestResults struct {
	target        int
	closestTarget int
}

func (results *ThreeSumClosestResults) UpdateIfCloser(newTarget int) int {
	diff := newTarget - results.target
	if Abs(diff) < Abs(results.target-results.closestTarget) {
		results.closestTarget = newTarget
	}
	return diff
}

// threeSum is the sixteenth LeetCode problem:
// - https://leetcode.com/problems/3sum-closest/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	results := ThreeSumClosestResults{
		target:        target,
		closestTarget: 1000000,
	}

	for i := 0; i < len(nums)-2; {
		var leftNumber, rightNumber int
		leftIndex, rightIndex := i+1, len(nums)-1
		currentNumber := nums[i]

		for leftIndex < rightIndex {
			leftNumber, rightNumber = nums[leftIndex], nums[rightIndex]
			sum := currentNumber + leftNumber + rightNumber
			diff := results.UpdateIfCloser(sum)

			// Increasing the left index increases our diff; decreasing the
			// right index decreases our diff. Since we want the diff to be
			// as close to `0` as possible, adjust the index that will get
			// us closer to `0`!
			//
			// If the diff _is_ `0`, we want the next pair of numbers.
			//
			// We always want to go to the next _different_ number to avoid
			// duplicates.
			if diff <= 0 {
				for leftIndex < rightIndex && leftNumber == nums[leftIndex] {
					leftIndex++
				}
			}
			if diff >= 0 {
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

	return results.closestTarget
}

var ThreeSumClosest = threeSumClosest /* For testing */
