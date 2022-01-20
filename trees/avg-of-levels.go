package trees

/*
Assumptions:
- Arguments: tn
- Returns: float slice
- Return the average value of the nodes at each level

Examples:

Input: root = [3,9,20,null,null,15,7]
Output: [3.00000,14.50000,11.00000]
Explanation: The average value of nodes on level 0 is 3, on level 1 is 14.5, and on level 2 is 11.
Hence return [3, 14.5, 11].

Input: root = [3,9,20,15,7]
Output: [3.00000,14.50000,11.00000]
*/

func AverageOfLevels(tn *TreeNode) []float64 {
	result := []float64{}
	q := []*TreeNode{}
	q = append(q, tn)
	for len(q) > 0 {
		size := len(q)
		var sum float64
		for i := 0; i < size; i++ {
			if q[0].Left != nil {
				q = append(q, q[0].Left)
			}
			if q[0].Right != nil {
				q = append(q, q[0].Right)
			}
			sum = sum + float64(q[0].Val)
			q = q[1:]
		}
		result = append(result, sum/float64(size))
	}
	return result
}
