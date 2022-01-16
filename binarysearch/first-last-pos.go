package binarysearch

/*
Assumptions:
- Arguments: sorted in slice (xs), int (target)
- Returns: int slice (indices)
- Return value contains starting and ending position of target value
- If target is not found, return [-1, -1].
- What if there is only one instance of the target?
    Answer: the first and ending postions are the same, in that case
- How many duplicates of the target are possible?
    Answer: no limit
- You must write an algorithm with O(log n) runtime complexity.

Examples:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

Input: nums = [], target = 0
Output: [-1,-1]
*/

func SearchRange(xs []int, target int) []int {
	leftResult, rightResult := -1, -1
	var findLeftMost func(int, int, int) int
	findLeftMost = func(l, u, m int) int {
		if l > u {
			return leftResult
		}
		if xs[m] < target {
			l = m + 1
			return findLeftMost(l, u, (l+u)/2)
		}
		if xs[m] > target {
			u = m - 1
			return findLeftMost(l, u, (l+u)/2)
		}
		leftResult = m
		u = m - 1
		return findLeftMost(l, u, (l+u)/2)
	}
	var findRightMost func(int, int, int) int
	findRightMost = func(l, u, m int) int {
		if l > u {
			return rightResult
		}
		if xs[m] < target {
			l = m + 1
			return findRightMost(l, u, (l+u)/2)
		}
		if xs[m] > target {
			u = m - 1
			return findRightMost(l, u, (l+u)/2)
		}
		rightResult = m
		l = m + 1
		return findRightMost(l, u, (l+u)/2)
	}
	leftMost, rightMost := findLeftMost(0, len(xs)-1, len(xs)/2), findRightMost(0, len(xs)-1, len(xs)/2)
	return []int{leftMost, rightMost}
}

type Tally struct {
	LeftMost  int
	RightMost int
	Size      int
}

func DivAndConqSearchRange(xs []int, target int) []int {
	if len(xs) == 0 {
		return []int{-1, -1}
	}
	result := divide(xs, target, 0, len(xs)-1)
	return []int{result.LeftMost, result.RightMost}
}

func divide(xs []int, target, low, high int) *Tally {
	if low == high {
		if xs[low] == target {
			return &Tally{
				LeftMost:  0,
				RightMost: 0,
				Size:      1,
			}
		}
		return &Tally{
			LeftMost:  -1,
			RightMost: -1,
			Size:      1,
		}
	}
	mid := (low + high) / 2
	return conquer(divide(xs, target, low, mid), divide(xs, target, mid+1, high))
}

func conquer(L, R *Tally) *Tally {
	if L.LeftMost != -1 && R.RightMost != -1 {
		return &Tally{
			LeftMost:  L.LeftMost,
			RightMost: R.RightMost + L.Size,
			Size:      L.Size + R.Size,
		}
	}
	if L.LeftMost != -1 {
		return &Tally{
			LeftMost:  L.LeftMost,
			RightMost: L.RightMost,
			Size:      L.Size + R.Size,
		}
	}
	if R.RightMost != -1 {
		return &Tally{
			LeftMost:  R.LeftMost + L.Size,
			RightMost: R.RightMost + L.Size,
			Size:      L.Size + R.Size,
		}
	}
	return &Tally{
		LeftMost:  -1,
		RightMost: -1,
		Size:      L.Size + R.Size,
	}
}
