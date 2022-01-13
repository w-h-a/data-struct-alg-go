package binarysearch

/*
Assumptions:
- Arguments: sorted int slice (numbers), int (target)
- Returns: slice of two indices whose values add up to target
- The two indices should be in ascending order
- The first index is 1 rather than 0
- There is a unique solution
- You may not use the same element twice

Examples:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].

Algorithm:
- binary search
*/

func TwoSum(numbers []int, target int) []int {
	u := len(numbers) - 1
	for i, x := range numbers {
		l := i + 1
		j := BinarySearch(numbers, target-x, l, u, (l+u)/2)
		if j != -1 {
			return []int{i + 1, j + 1}
		}
	}
	return nil
}
