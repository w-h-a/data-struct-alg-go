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
	if len(xs) == 0 {
		return []int{-1, -1}
	}
	var bSearch func(int, int, int) []int
	bSearch = func(l, u, m int) []int {
		if l > u {
			return []int{-1, -1}
		}
		if xs[m] < target {
			l = m + 1
			return bSearch(l, u, (l+u)/2)
		}
		if xs[m] > target {
			u = m - 1
			return bSearch(l, u, (l+u)/2)
		}
		n := m
		for n != -1 && xs[n] == target {
			n--
		}
		for m != len(xs) && xs[m] == target {
			m++
		}
		return []int{n + 1, m - 1}
	}
	return bSearch(0, len(xs)-1, len(xs)/2)
}
