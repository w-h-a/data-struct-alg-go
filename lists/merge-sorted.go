package lists

/*
Assumptions:
- Arguments: two ListNode pointers
- Returns: ListNode pointer
- Merge the two lists in to one sorted list
- The return list should be made by splicing together the nodes of the argument lists

Examples:

Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Input: list1 = [], list2 = []
Output: []

Input: list1 = [], list2 = [0]
Output: [0]
*/

func MergeTwoLists(xs, ys *ListNode) *ListNode {
	if xs == nil {
		return ys
	}
	if ys == nil {
		return xs
	}
	if xs.Val <= ys.Val {
		return &ListNode{
			Val:  xs.Val,
			Next: MergeTwoLists(xs.Next, ys),
		}
	}
	return &ListNode{
		Val:  ys.Val,
		Next: MergeTwoLists(xs, ys.Next),
	}
}
