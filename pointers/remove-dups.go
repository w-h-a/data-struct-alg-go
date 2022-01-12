package pointers

/*
Assumptions:
- Arguments: int slice (nums)
- Returns: int
- In place, remove duplicates from nums without disturbing relative position of elements
- If there are k elements after removing duplicates,
    the first k slots of nums should hold the k elements, and
    it does not matter what you leave behind in the slots after the first k slots
- The return value is k.
- You are not allowed to allocate extra space for another array

Examples:

Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

Algorithm:
- Reader/Writer
*/

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	w, r := 0, 1
	for r < len(nums) {
		if nums[r] != nums[w] {
			w++
			nums[w] = nums[r]
		}
		r++
	}
	w++
	return w
}
