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

func UniquePathsWithObstaclesIterative(grid [][]int) int {
	cache := map[[2]int]int{}
	m, n := len(grid)-1, len(grid[0])-1
	if grid[0][0] == 1 {
		return 0
	}
	cache[[2]int{0, 0}] = 1
	for r := 0; r <= m; r++ {
		for c := 0; c <= n; c++ {
			if r == 0 && c == 0 {
				continue
			}
			key := [2]int{r, c}
			if grid[r][c] == 1 {
				cache[key] = 0
			} else {
				cache[key] = cache[[2]int{r - 1, c}] + cache[[2]int{r, c - 1}]
			}
		}
	}
	return cache[[2]int{m, n}]
}
