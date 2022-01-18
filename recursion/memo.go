package recursion

func ClimbStairs(n int) int {
	cache := map[int]int{}
	var memoized func(int) int
	memoized = func(n int) int {
		if n == 1 {
			return 1
		}
		if n == 2 {
			return 2
		}
		key := n
		_, ok := cache[key]
		if !ok {
			cache[key] = memoized(n-1) + memoized(n-2)
		}
		return cache[key]
	}
	return memoized(n)
}

func UniquePaths(m, n int) int {
	cache := map[[2]int]int{}
	var memoized func(int, int) int
	memoized = func(m, n int) int {
		if m < 1 || n < 1 {
			return 0
		}
		if m == 1 && n == 1 {
			return 1
		}
		key := [2]int{m, n}
		_, ok := cache[key]
		if !ok {
			cache[key] = memoized(m-1, n) + memoized(m, n-1)
		}
		return cache[key]
	}
	return memoized(m, n)
}

func UniquePathsWithObstacles(grid [][]int) int {
	cache := map[[2]int]int{}
	m, n := len(grid)-1, len(grid[0])-1
	var memoized func(int, int) int
	memoized = func(m, n int) int {
		if m < 0 || n < 0 || grid[m][n] == 1 {
			return 0
		}
		if m == 0 && n == 0 {
			return 1
		}
		key := [2]int{m, n}
		_, ok := cache[key]
		if !ok {
			cache[key] = memoized(m-1, n) + memoized(m, n-1)
		}
		return cache[key]
	}
	return memoized(m, n)
}
