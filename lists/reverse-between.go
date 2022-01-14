package lists

/*
Assumptions:
- Arguments: pointer to ListNode (xs), int (left), int (right)
- Returns: pointer to ListNode
- left <= right
- Reverse the nodes of the list from position left to position right and return

Examples:

Input: head = [1,2,3,4,5], left = 2, right = 4
Output: [1,4,3,2,5]

Input: head = [5], left = 1, right = 1
Output: [5]

Input: head = [3,5], left = right = 2
Output: [3,5]
*/

func ReverseBetween(xs *ListNode, left int, right int) *ListNode {
	return Prepend(Take(xs, left-1), Prepend(Reverse(Drop(Take(xs, right), left-1)), Drop(xs, right)))
}
