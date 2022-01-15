package sorting

// MergeSort O(n log n)
func MergeSort(xs []float64) []float64 {
	if len(xs) == 0 || len(xs) == 1 {
		return xs
	}
	k := len(xs) / 2
	return merge(MergeSort(xs[0:k]), MergeSort(xs[k:]))
}

func merge(xs, ys []float64) []float64 {
	if len(xs) == 0 {
		return ys
	}
	if len(ys) == 0 {
		return xs
	}
	if xs[0] <= ys[0] {
		return append([]float64{xs[0]}, merge(xs[1:], ys)...)
	}
	return append([]float64{ys[0]}, merge(xs, ys[1:])...)
}
