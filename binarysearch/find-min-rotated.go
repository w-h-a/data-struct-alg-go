package binarysearch

/*
Assumptions:
- Arguments: sorted but rotated int slice (xs)
- Returns: int
- Return value is the minimum element
- You must write an algorithm that runs in O(log n) time.

Examples:

Input: nums = [3,4,5,1,2]
Output: 1
Explanation: The original array was [1,2,3,4,5] rotated 3 times.

Input: nums = [4,5,6,7,0,1,2]
Output: 0
Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.

Input: nums = [11,13,15,17]
Output: 11
Explanation: The original array was [11,13,15,17] and it was rotated 4 times.

Main idea:
Carefully move lower bound up until the element at lower bound is finally
    less than the element at upper bound
When do we move the lower bound?
    Whenever the mid point is still greater than the upper bound.
On the other hand, if the mid point is less than the upper bound,
    we just need to reset the upper bound
What about in the case where, we have something like: [7, 0]?
    m needs to be lower bound to ensure l moves in which case we have our answer when l == u
    even when xs[l] == xs[u]
*/

func FindMin(xs []int) int {
	var bSearch func(int, int, int) int
	bSearch = func(l, u, m int) int {
		if l == u || xs[l] < xs[u] {
			return l
		}
		if xs[m] > xs[u] {
			l = m + 1
			return bSearch(l, u, (l+u)/2)
		}
		u = m
		return bSearch(l, u, (l+u)/2)
	}
	return xs[bSearch(0, len(xs)-1, (len(xs)-1)/2)]
}
