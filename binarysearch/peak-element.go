package binarysearch

/*
Assumptions:
- Arguments: int slice (xs)
- Returns: int (an index)
- Return the index of one peak (of possibly many)
- A peak is just an element that is strictly greater than both of its neighbors
- You may imagine that xs[-1] == xs[n] == -∞
- You are required to write an algorithm that runs in O(logN) time.
- What if there is no peak? There is always some peak.

Examples:

Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.

Input: nums = [1,2,1,3,5,6,4]
Output: 5
Explanation: Your function can return either index number 1 where the peak element is 2, or index number 5 where the peak element is 6.
*/

func FindPeakElement(xs []int) int {
	if len(xs) == 1 {
		return 0
	}
	var bSearch func(int, int, int) int
	bSearch = func(l, u, m int) int {
		switch m {
		case 0:
			if xs[m+1] > xs[m] {
				l = m + 1
				return bSearch(l, u, (l+u)/2)
			}
			return m
		case len(xs) - 1:
			if xs[m-1] > xs[m] {
				u = m - 1
				return bSearch(l, u, (l+u)/2)
			}
			return m
		default:
			if xs[m-1] > xs[m] {
				u = m - 1
				return bSearch(l, u, (l+u)/2)
			}
			if xs[m+1] > xs[m] {
				l = m + 1
				return bSearch(l, u, (l+u)/2)
			}
			return m
		}
	}
	return bSearch(0, len(xs), len(xs)/2)
}

type Peak struct {
	Idx  int
	Val  int
	Size int
}

func DivAndConqFindPeakElement(xs []int) int {
	return div(xs, 0, len(xs)-1).Idx
}

func div(xs []int, low, high int) *Peak {
	if low == high {
		return &Peak{
			Idx:  0,
			Val:  xs[low],
			Size: 1,
		}
	}
	k := (low + high) / 2
	return conq(div(xs, low, k), div(xs, k+1, high))
}

func conq(L, R *Peak) *Peak {
	if L.Val > R.Val {
		return &Peak{
			Idx:  L.Idx,
			Val:  L.Val,
			Size: L.Size + R.Size,
		}
	}
	return &Peak{
		Idx:  R.Idx + L.Size,
		Val:  R.Val,
		Size: L.Size + R.Size,
	}
}
