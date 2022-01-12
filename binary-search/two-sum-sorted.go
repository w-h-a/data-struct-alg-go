package binarysearch

/*

 */

func TwoSum(numbers []int, target int) []int {
	var bSearch func(int, int, int, int) int
	bSearch = func(k, l, u, m int) int {
		if l > u {
			return -1
		}
		if numbers[m] < k {
			l = m + 1
			return bSearch(k, l, u, (l+u)/2)
		}
		if numbers[m] > k {
			u = m - 1
			return bSearch(k, l, u, (l+u)/2)
		}
		return m
	}
	for i, x := range numbers {
		j := bSearch(target-x, i+1, len(numbers)-1, (i+1+len(numbers)-1)/2)
		if j != -1 {
			return []int{i + 1, j + 1}
		}
	}
	return []int{0, 0}
}
