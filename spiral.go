package main

/*
Problem:
- Argument: m*n matrix
- Return: slice of elements in spiral order

Assumptions:
- The length of the slices are greater than or equal to 1 and less than or equal to 10
- The elements are integers between -100 and 100
- a slice of elements is in spiral order just in case
    it begins with the first row,
    proceeds to the last elements of each row,
    proceeds to the final row in reverse order,
    proceeds with the first elements of each row except the first, etc

Examples:

Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]

*/

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	m := len(matrix)
	n := len(matrix[0])
	r := 0
	c := -1
	for {
		for i := 0; i < n; i++ {
			c++
			result = append(result, matrix[r][c])
		}
		m--
		if m == 0 {
			break
		}
		for i := 0; i < m; i++ {
			r++
			result = append(result, matrix[r][c])
		}
		n--
		if n == 0 {
			break
		}
		for i := 0; i < n; i++ {
			c--
			result = append(result, matrix[r][c])
		}
		m--
		if m == 0 {
			break
		}
		for i := 0; i < m; i++ {
			r--
			result = append(result, matrix[r][c])
		}
		n--
		if n == 0 {
			break
		}
	}
	return result
}
