package solutions

import (
	"sort"
)

func isEven(n int) bool {
	return n%2 == 0
}

func get(nums []int, index int) int {
	if index >= len(nums) {
		return 1_000_000 // Default value, max value as per the problem statement
	}
	return nums[index]
}

// FindMedianSortedArrays is the fourth LeetCode problem:
// - https://leetcode.com/problems/median-of-two-sorted-arrays/
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Walk through the arrays until we reach the "middle".
	//
	// If the combined length of the arrays is even, we need to average
	// the middle two numbers; otherwise, we just need to return the
	// middle number
	//
	// For example, a combined length of 5 would have a middle index of
	// 2 (0-based), and a combined length of 6 would have middle indexes
	// of 2 and 3 (0-based)
	length := len(nums1) + len(nums2)

	// `stopIndex` is the index at which we stop walking through the arrays
	stopIndex := int(float64(length-1) / 2.0)
	values := []int{0, 0}
	var ret float64

	// `mergedIndex` is the index for the merged array (conceptually)
	n1Index, n2Index, mergedIndex := 0, 0, 0

	for {
		if mergedIndex == stopIndex {
			// The smallest two of the next four numbers
			nextFourValues := []int{
				get(nums1, n1Index),
				get(nums1, n1Index+1),
				get(nums2, n2Index),
				get(nums2, n2Index+1),
			}
			sort.Slice(
				nextFourValues,
				func(i, j int) bool {
					return nextFourValues[i] < nextFourValues[j]
				},
			)

			values = nextFourValues[:2]
			break
		}

		if get(nums1, n1Index) < get(nums2, n2Index) {
			n1Index++
		} else {
			n2Index++
		}
		mergedIndex++
	}

	if isEven(length) {
		ret = float64(values[0]+values[1]) / 2.0
	} else {
		ret = float64(values[0])
	}

	return ret
}
