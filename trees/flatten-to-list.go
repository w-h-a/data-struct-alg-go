package trees

/*
Assumptions:
- Arguments: pointer to TreeNode (tn)
- Returns: n/a
- Flatten the tree into a linked list
- The linked list should use the same TreeNode 'class',
    where the right child pointer points to the next node and the left child pointer is always nil
- The linked list should be in the same order as a pre-order traversal of the binary tree

Examples:

Input: root = [1,2,5,3,4,null,6]
Output: [1,null,2,null,3,null,4,null,5,null,6]

Input: root = []
Output: []

Input: root = [0]
Output: [0]
*/

func Flatten(tn *TreeNode) {
	if tn == nil {
		return
	}
	xs := SlicePreOrder(tn)
	tn.Left = nil
	for _, x := range xs[1:] {
		tn.Right = &TreeNode{Val: x}
		tn = tn.Right
	}
}
