package trees

/*
Assumptions:
- Arguments: pointer to root of binary tree
- Returns: bool
- Return true if and only if the tree is a valid binary search tree
- A tree is a valid binary search true just in case
    - the left subtree of a node contains only nodes with values less than the node's
    - the right subtree of a node contains only nodes with values greater than the node's
    - both the left and right subtrees are binary search trees.

Examples:

Input: root = [2,1,3]
Output: true

Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
*/

func IsValidBST(tn *TreeNode) bool {
	if tn == nil {
		return true
	}
	return isSubtreeLessThan(tn.Left, tn.Val) && isSubtreeGreaterThan(tn.Right, tn.Val) && IsValidBST(tn.Left) && IsValidBST(tn.Right)
}

func isSubtreeLessThan(tn *TreeNode, x int) bool {
	if tn == nil {
		return true
	}
	return tn.Val < x && isSubtreeLessThan(tn.Left, x) && isSubtreeLessThan(tn.Right, x)
}

func isSubtreeGreaterThan(tn *TreeNode, x int) bool {
	if tn == nil {
		return true
	}
	return tn.Val > x && isSubtreeGreaterThan(tn.Left, x) && isSubtreeGreaterThan(tn.Right, x)
}

func IsValidBSTII(root *TreeNode) bool {
	var prev *TreeNode
	var isIncreasing func(*TreeNode) bool
	isIncreasing = func(p *TreeNode) bool {
		if p == nil {
			return true
		}
		if isIncreasing(p.Left) {
			if prev != nil && p.Val <= prev.Val {
				return false
			}
			prev = p
			return isIncreasing(p.Right)
		}
		return false
	}
	return isIncreasing(root)
}
