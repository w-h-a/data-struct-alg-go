package lists

/*
Assum:
- Arguments: pointer to ListNode (xs), int (n)
- Returns: pointer to ListNode
- Rremove the nth node from the end of the list and return the new head.

Examples:

Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Input: head = [1], n = 1
Output: []

Input: head = [1,2], n = 1
Output: [1]
*/

func RemoveNthFromEnd(xs *ListNode, n int) *ListNode {
	return Delete(xs, Length(xs)-n)
}
