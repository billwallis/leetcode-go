package solutions

import "math"

// O(n^2) solution, too slow
/*
func maxArea(height []int) int {
	maxVolume := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			width := j - i
			containerHeight := int(math.Min(
				float64(height[i]),
				float64(height[j]),
			))

			volume := width * containerHeight
			if volume > maxVolume {
				maxVolume = volume
			}
		}
	}

	return maxVolume
}
*/

// maxArea is the eleventh LeetCode problem:
// - https://leetcode.com/problems/container-with-most-water/
func maxArea(height []int) int {
	maxVolume := 0
	leftIndex, rightIndex := 0, len(height)-1

	for leftIndex < rightIndex {
		width := rightIndex - leftIndex
		leftHeight, rightHeight := height[leftIndex], height[rightIndex]
		containerHeight := int(math.Min(
			float64(leftHeight),
			float64(rightHeight),
		))

		if volume := width * containerHeight; volume > maxVolume {
			maxVolume = volume
		}

		if leftHeight < rightHeight {
			leftIndex++
		} else {
			rightIndex--
		}
	}

	return maxVolume
}

var MaxArea = maxArea /* For testing */
