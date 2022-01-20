package trees

/*
Assumptions:
- Arguments: pointer to TreeNode (tn), int (low), int (high)
- Returns: int
- Return the sum of values of all nodes with a value in the inclusive range [low, high]

Examples:

Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
Output: 32
Explanation: Nodes 7, 10, and 15 are in the range [7, 15]. 7 + 10 + 15 = 32.

Input: root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
Output: 23
Explanation: Nodes 6, 7, and 10 are in the range [6, 10]. 6 + 7 + 10 = 23.
*/

func RangeSumBST(tn *TreeNode, low int, high int) int {
	if tn == nil {
		return 0
	}
	if tn.Val >= low && tn.Val <= high {
		return tn.Val + RangeSumBST(tn.Left, low, high) + RangeSumBST(tn.Right, low, high)
	}
	return RangeSumBST(tn.Left, low, high) + RangeSumBST(tn.Right, low, high)
}
