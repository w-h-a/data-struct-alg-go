package binarysearch

/*
Assumptions:
- Arguments: int (target), int slice (nums)
- Returns: int
- The return integer is the minimal length of a contiguous subarray of which
    the sum of the elements >= target.
    So, if there is no such subarray, return 0.

Examples:

Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.

Input: target = 4, nums = [1,4,4]
Output: 1

Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0

*/

func MinSubArrayLen(target int, nums []int) int {
	var bSearch func(int, int, int, int) int
	bSearch = func(l, u, m, result int) int {
		if l > u {
			return result
		}
		if isSolution(m, target, nums) {
			result = m
			u = m - 1
			return bSearch(l, u, (l+u)/2, result)
		}
		l = m + 1
		return bSearch(l, u, (l+u)/2, result)
	}
	return bSearch(0, len(nums), len(nums)/2, 0)
}

func isSolution(m, target int, nums []int) bool {
	a, r := 0, 0
	sum := 0
	for r < len(nums) {
		if r < m {
			sum += nums[r]
			r++
		} else {
			sum -= nums[a]
			a++
			sum += nums[r]
			r++
		}
		if sum >= target {
			return true
		}
	}
	return false
}
