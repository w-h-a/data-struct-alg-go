package binarysearch

/*
Assumptions:
- Arguments: sorted int slice (xs), int (target)
- Returns: int (index)
- If the target value is found, return its index;
    otherwise, return the index where it would be if it were inserted in order
- You must write an algorithm with O(log n) runtime complexity.



Examples:

Input: nums = [1,3,5,6], target = 5
Output: 2

Input: nums = [1,3,5,6], target = 2
Output: 1

Input: nums = [1,3,5,6], target = 7
Output: 4
*/

func SearchInsert(xs []int, target int) int {
	var bSearch func(int, int, int) int
	bSearch = func(l, u, m int) int {
		if l > u {
			return l
		}
		if target < xs[m] {
			u = m - 1
			return bSearch(l, u, (l+u)/2)
		}
		if target > xs[m] {
			l = m + 1
			return bSearch(l, u, (l+u)/2)
		}
		return m
	}
	return bSearch(0, len(xs)-1, len(xs)/2)
}
