package sorting

// QuickSort O(n^2)
// Choose the head
// Partition the remaining items into two parts: the items less than or equal to the head
// and the items greater than the head
// Sort each part recursively, then put the smaller part before the greater
func QuickSort(list []float64) []float64 {
	if len(list) == 0 || len(list) == 1 {
		return list
	}
	var partition func(left, right, xs []float64) []float64
	partition = func(left, right, xs []float64) []float64 {
		if len(xs) == 0 {
			return append(append(QuickSort(left), list[0]), QuickSort(right)...)
		}
		if xs[0] <= list[0] {
			return partition(append([]float64{xs[0]}, left...), right, xs[1:])
		}
		return partition(left, append([]float64{xs[0]}, right...), xs[1:])
	}
	return partition([]float64{}, []float64{}, list[1:])
}
