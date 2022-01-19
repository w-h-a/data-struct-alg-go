package recursion

import (
	"math"
)

/*
Assumptions:
- Arguments: (non-negative) int matrix (grid)
- Returns: int
- Find a path from top left to bottom right that minimimizes the sum along the path
- You can only move either down or right at any point in time.

Examples:

Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
Output: 7
Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.

Input: grid = [[1,2,3],[4,5,6]]
Output: 12
*/

func MinPathSum(grid [][]int) int {
	cache := map[[2]int]int{}
	m, n := len(grid)-1, len(grid[0])-1
	var memoized func(int, int) int
	memoized = func(m, n int) int {
		if m < 0 || n < 0 {
			return math.MaxInt32
		}
		if m == 0 && n == 0 {
			return grid[m][n]
		}
		key := [2]int{m, n}
		_, ok := cache[key]
		if !ok {
			cache[key] = grid[m][n] + min(memoized(m-1, n), memoized(m, n-1))
		}
		return cache[key]
	}
	return memoized(m, n)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MinPathSumIter(grid [][]int) int {
	m, n := len(grid)-1, len(grid[0])-1
	for r := m; r >= 0; r-- {
		for c := n; c >= 0; c-- {
			if r == m && c == n {
				continue
			}
			if r == m {
				grid[r][c] = grid[r][c] + grid[r][c+1]
				continue
			}
			if c == n {
				grid[r][c] = grid[r][c] + grid[r+1][c]
				continue
			}
			grid[r][c] = grid[r][c] + min(grid[r][c+1], grid[r+1][c])
		}
	}
	return grid[0][0]
}
