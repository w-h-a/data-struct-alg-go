package binarysearch

/*
Assumptions:
- Arguments: int (num)
- Returns: bool
- Returns true iff num is a perfect square
- num is a perfect square just in case num is the product of some integer with itself
- Don't use built-ins

Examples:

Input: num = 16
Output: true

Input: num = 14
Output: false

Algorithm:
- binary search
*/

func IsPerfectSquare(num int) bool {
	var bSearch func(int, int, int) bool
	bSearch = func(l, u, m int) bool {
		if l > u {
			return false
		}
		if m*m < num {
			l = m + 1
			return bSearch(l, u, (l+u)/2)
		}
		if m*m > num {
			u = m - 1
			return bSearch(l, u, (l+u)/2)
		}
		return true
	}
	return bSearch(0, num, num/2)
}
