package trees

/*
Assumptions:
- Arguments: pointer to TreeNode (tn)
- Returns: int
- The tree contains digits from 0 to 9 only
- Return the total sum of all root-to-leaf numbers (see below)

Examples:

Input: root = [1,2,3]
Output: 25
Explanation:
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.
Therefore, sum = 12 + 13 = 25.

Input: root = [4,9,0,5,1]
Output: 1026
Explanation:
The root-to-leaf path 4->9->5 represents the number 495.
The root-to-leaf path 4->9->1 represents the number 491.
The root-to-leaf path 4->0 represents the number 40.
Therefore, sum = 495 + 491 + 40 = 1026.
*/

func SumNumbers(tn *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(tn *TreeNode, result int) int {
		if tn == nil {
			return 0
		}
		result = 10*result + tn.Val
		if tn.Left == nil && tn.Right == nil {
			return result
		}
		return dfs(tn.Left, result) + dfs(tn.Right, result)
	}
	return dfs(tn, 0)
}
