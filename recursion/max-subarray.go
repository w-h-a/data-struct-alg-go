package recursion

import "math"

/*
Assum:
- Arguments: int slice (xs)
- Returns: int
- Find the contiguous subarray which has the largest sum and return the sum

Examples:

Input: xs = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Input: xs = [1]
Output: 1

Input: xs = [5,4,-1,7,8]
Output: 23
*/

func MaxSubArray(xs []int) int {
	var max func([]int, int, int) int
	max = func(xs []int, maxEndingHere, maxSoFar int) int {
		if len(xs) == 0 {
			return maxSoFar
		}
		maxEndingHere = int(math.Max(float64(maxEndingHere+xs[0]), float64(xs[0])))
		maxSoFar = int(math.Max(float64(maxEndingHere), float64(maxSoFar)))
		return max(xs[1:], maxEndingHere, maxSoFar)
	}
	return max(xs[1:], xs[0], xs[0])
}

type Tally struct {
	TotalSum  int
	MaxSum    int
	MaxPrefix int
	MaxSuffix int
}

func DivideAndConqMaxSubArray(xs []int) int {
	return divide(xs, 0, len(xs)-1).MaxSum
}

func divide(xs []int, low, high int) *Tally {
	if low == high {
		return &Tally{
			TotalSum:  xs[low],
			MaxSum:    xs[low],
			MaxPrefix: xs[low],
			MaxSuffix: xs[low],
		}
	}
	mid := (low + high) / 2
	return conquer(divide(xs, low, mid), divide(xs, mid+1, high))
}

func conquer(L, R *Tally) *Tally {
	return &Tally{
		TotalSum:  L.TotalSum + R.TotalSum,
		MaxPrefix: int(math.Max(float64(L.MaxPrefix), float64(L.TotalSum+R.MaxPrefix))),
		MaxSuffix: int(math.Max(float64(R.MaxSuffix), float64(R.TotalSum+L.MaxSuffix))),
		MaxSum:    int(math.Max(math.Max(float64(L.MaxSum), float64(R.MaxSum)), float64(L.MaxSuffix+R.MaxPrefix))),
	}
}
