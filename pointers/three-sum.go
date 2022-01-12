package pointers

import (
	"sort"
)

/*
Assumptions:
- Arguments: int slice (nums)
- Returns: [][]int
- In particular: the return value ought to be all the _unique_ triplets ([]int{num[i], num[j], num[k]})
    s.t. i, j, and k are pairwise distinct _and_ the values sum to 0.

Examples:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]

Input: nums = []
Output: []

Input: nums = [0]
Output: []

*/

func ThreeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	for i, v := range nums {
		if i != 0 && v == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			x, y := nums[l], nums[r]
			if v+x+y == 0 {
				result = append(result, []int{v, x, y})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
			if v+x+y < 0 {
				l++
			}
			if v+x+y > 0 {
				r--
			}
		}
	}
	return result
}
