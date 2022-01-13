package binarysearch

/*
Assumptions:
- Arguments: sorted but rotated int slice (xs), int (target)
- Returns: int (index)
- xs has pairwise distinct values
- If the target is in xs, return the index; else, return -1
- You must write an algorithm with O(log n) runtime complexity.

Examples:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

Input: nums = [1], target = 0
Output: -1
*/

func Search(xs []int, target int) int {
	var bSearch func(int, int, int) int
	bSearch = func(l, u, m int) int {
		if l > u {
			return -1
		}
		if xs[m] == target {
			return m
		}
		if xs[l] <= xs[m] {
			if xs[l] <= target && target < xs[m] {
				u = m - 1
				return bSearch(l, u, (l+u)/2)
			}
			l = m + 1
			return bSearch(l, u, (l+u)/2)
		}
		if xs[m] <= xs[u] {
			if xs[m] < target && target <= xs[u] {
				l = m + 1
				return bSearch(l, u, (l+u)/2)
			}
			u = m - 1
			return bSearch(l, u, (l+u)/2)
		}
		return -1
	}
	return bSearch(0, len(xs)-1, len(xs)/2)
}
