package misc

/*
Assumptions:
- Arguments: slice of ints and int target
- Return: indices of the two numbers that add up to target
- Each input has a unique solution
- You may not use the same element twice
- You can return the answer in any order

Examples:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Input: nums = [3,2,4], target = 6
Output: [1,2]

Input: nums = [3,3], target = 6
Output: [0,1]

Data Structure:
- Hash table that stores index of integers ("indexOfValue")

Algorithm:
- For each index, value pair in numbers slice
    - check if indexOfValue contains target - value
    - if so,
        - return slice with current index of loop paired with return value of indexOfValue hash
    - if not,
        - update indexOfValue hash so that the current value of loop is associated with the current index
*/

func TwoSum(numbers []int, target int) []int {
	indexOfValue := map[int]int{}
	for i, v := range numbers {
		j, ok := indexOfValue[target-v]
		if ok {
			return []int{i, j}
		}
		indexOfValue[v] = i
	}
	return []int{}
}
