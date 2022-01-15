package sorting

// InsertSort O(n^2)
func InsertSort(list []float64) []float64 {
	if len(list) == 0 {
		return list
	}
	return ins(list[0], InsertSort(list[1:]))
}

func ins(x float64, list []float64) []float64 {
	if len(list) == 0 {
		return []float64{x}
	}
	if x <= list[0] {
		return append([]float64{x}, list...)
	}
	return append([]float64{list[0]}, ins(x, list[1:])...)
}
