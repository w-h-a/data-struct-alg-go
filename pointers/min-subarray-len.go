package pointers

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
	minLen := len(nums) + 1
	s := 0
	e := 0
	sum := 0
	for e < len(nums) {
		sum += nums[e]
		for sum >= target {
			minLen = min(minLen, e-s+1)
			leftNum := nums[s]
			sum -= leftNum
			s++
		}
		e++
	}
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
