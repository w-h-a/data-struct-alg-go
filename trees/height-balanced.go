package trees

import (
	"math"
)

/*
Assumptions:
- Arguments: pointer to TreeNode (tn)
- Returns: bool
- Return true if and only if the tree is height-balanced
- A tree is height-balanced just in case the left and right subtrees of every node differ in height by no more than 1.

Examples:

Input: root = [3,9,20,null,null,15,7]
Output: true

Input: root = [1,2,2,3,3,null,null,4,4]
Output: false

Input: root = []
Output: true
*/

func IsBalanced(tn *TreeNode) bool {
	if tn == nil {
		return true
	}
	return IsBalanced(tn.Left) && IsBalanced(tn.Right) && math.Abs(float64(Depth(tn.Left)-Depth(tn.Right))) <= 1
}
