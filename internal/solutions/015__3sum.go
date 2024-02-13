package solutions

import (
	"reflect"
	"sort"
)

type ThreeSumResults struct {
	arrays [][]int
}

func (results *ThreeSumResults) Contains(arr []int) bool {
	for _, res := range results.arrays {
		if reflect.DeepEqual(res, arr) {
			return true
		}
	}
	return false
}

func (results *ThreeSumResults) Append(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	if !results.Contains(arr) {
		results.arrays = append(results.arrays, arr)
	}
}

// threeSum is the fifteenth LeetCode problem:
// - https://leetcode.com/problems/3sum/
func threeSum(nums []int) [][]int {
	results := ThreeSumResults{}

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if n1, n2, n3 := nums[i], nums[j], nums[k]; n1+n2+n3 == 0 {
					results.Append([]int{n1, n2, n3})
				}
			}
		}
	}

	return results.arrays
}

var ThreeSum = threeSum /* For testing */
