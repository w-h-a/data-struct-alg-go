package pointers

/*
Assumptions:
- Arguments: slice of ints
- Returns: n/a
- In pace, move all 0s to the end while maintaining the relative order of the non-zero elements

Examples:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Input: nums = [0]
Output: [0]

Algorithm:
- Reader/Writer
*/

func MoveZeroes(nums []int) {
	w, r := 0, 0
	for ; r < len(nums); r++ {
		if nums[r] != 0 {
			nums[w] = nums[r]
			w++
		}
	}
	for ; w < len(nums); w++ {
		nums[w] = 0
	}
}
