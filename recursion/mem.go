package recursion

import (
	"fmt"
	"strings"
)

type Cache map[string]int

func Memoize(f func(...int) int, cache Cache) func(...int) int {
	return func(args ...int) int {
		key := strOfIntSlice(args, ",")
		v, ok := cache[key]
		if !ok {
			cache[key] = f(args...)
			return cache[key]
		}
		return v
	}
}

func strOfIntSlice(xs []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(xs), " ", delim, -1), "[]")
}

// e.g.,

var MemFib = Memoize(Fib, Cache{})
var MemUniquePaths = Memoize(UniquePaths, Cache{})

func Fib(n ...int) int {
	if n[0] == 0 || n[0] == 1 {
		return n[0]
	}
	return Fib(n[0]-1) + Fib(n[0]-2)
}

func UniquePaths(rc ...int) int {
	if rc[0] == 1 || rc[1] == 1 {
		return 1
	}
	return UniquePaths(rc[0]-1, rc[1]) + UniquePaths(rc[0], rc[1]-1)
}
