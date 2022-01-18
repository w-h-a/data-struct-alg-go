package recursion

import "math"

/*
Assumptions
- Arguments: triangle
- Returns: int
- Return the min path sum from top to bottom.
- If you are index i on the current row, you may move to either index i or index i + 1 on the next row

Examples:

Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
Output: 11
Explanation: The triangle looks like:
   2
  3 4
 6 5 7
4 1 8 3
The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).

Input: triangle = [[-10]]
Output: -10
*/

func MinimumTotal(triangle [][]int) int {
	cache := map[[2]int]int{}
	m, n := 0, 0
	var memoized func(int, int) int
	memoized = func(m, n int) int {
		if m < 0 || n < 0 {
			return math.MaxInt32
		}
		if m == len(triangle)-1 {
			return triangle[m][n]
		}
		key := [2]int{m, n}
		_, ok := cache[key]
		if !ok {
			cache[key] = triangle[m][n] + min(memoized(m+1, n), memoized(m+1, n+1))
		}
		return cache[key]
	}
	return memoized(m, n)
}
