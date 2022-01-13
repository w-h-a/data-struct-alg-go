package binarysearch

/*
Assumptions:
- Arguments: sorted in slice (xs), int (target)
- Returns: int slice (indices)
- Return value contains starting and ending position of target value
- If target is not found, return [-1, -1].
- What if there is only one instance of the target?
    Answer: the first and ending postions are the same, in that case
- How many duplicates of the target are possible?
    Answer: no limit
- You must write an algorithm with O(log n) runtime complexity.

Examples:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

Input: nums = [], target = 0
Output: [-1,-1]
*/

func SearchRange(xs []int, target int) []int {
	leftResult, rightResult := -1, -1
	var findLeftMost func(int, int, int) int
	findLeftMost = func(l, u, m int) int {
		if l > u {
			return leftResult
		}
		if xs[m] < target {
			l = m + 1
			return findLeftMost(l, u, (l+u)/2)
		}
		if xs[m] > target {
			u = m - 1
			return findLeftMost(l, u, (l+u)/2)
		}
		leftResult = m
		u = m - 1
		return findLeftMost(l, u, (l+u)/2)
	}
	var findRightMost func(int, int, int) int
	findRightMost = func(l, u, m int) int {
		if l > u {
			return rightResult
		}
		if xs[m] < target {
			l = m + 1
			return findRightMost(l, u, (l+u)/2)
		}
		if xs[m] > target {
			u = m - 1
			return findRightMost(l, u, (l+u)/2)
		}
		rightResult = m
		l = m + 1
		return findRightMost(l, u, (l+u)/2)
	}
	leftMost, rightMost := findLeftMost(0, len(xs)-1, len(xs)/2), findRightMost(0, len(xs)-1, len(xs)/2)
	return []int{leftMost, rightMost}
}
