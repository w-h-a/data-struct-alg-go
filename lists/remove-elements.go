package lists

/*
Assumptions:
- Arguments: pointer to ListNode (xs) and int (val)
- Return: pointer to ListNode
- Remove all the nodes fo the linked list that has val

Examples:

Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]

Input: head = [], val = 1
Output: []

Input: head = [7,7,7,7], val = 7
Output: []
*/

func RemoveElements(xs *ListNode, val int) *ListNode {
	if xs == nil {
		return nil
	}
	if xs.Val == val {
		return RemoveElements(xs.Next, val)
	}
	return &ListNode{
		Val:  xs.Val,
		Next: RemoveElements(xs.Next, val),
	}
}
