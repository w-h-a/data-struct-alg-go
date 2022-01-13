package binarysearch

func BinarySearch(xs []int, y, l, u, m int) int {
	if l > u {
		return -1
	}
	if xs[m] < y {
		l = m + 1
		return BinarySearch(xs, y, l, u, (l+u)/2)
	}
	if xs[m] > y {
		u = m - 1
		return BinarySearch(xs, y, l, u, (l+u)/2)
	}
	return m
}
