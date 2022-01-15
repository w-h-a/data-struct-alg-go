package lists

/*
Assum:
- Arguments: pointer to ListNode (xs)
- Returns: pointer to ListNode
- Delete all nodes that have duplicate values, leaving only distinct numbers
- The input list is sorted
- the output list is sorted


Examples:

Input: head = [1,2,3,3,4,4,5]
Output: [1,2,5]

Input: head = [1,1,1,2,3]
Output: [2,3]
*/

func DeleteAllDuplicates(xs *ListNode) *ListNode {
	var keepGoing func(*ListNode, *ListNode) *ListNode
	keepGoing = func(prev, curr *ListNode) *ListNode {
		if curr == nil {
			return nil
		}
		if (prev != nil && curr != nil && prev.Val == curr.Val) || (curr != nil && curr.Next != nil && curr.Val == curr.Next.Val) {
			return keepGoing(curr, curr.Next)
		}
		return &ListNode{
			Val:  curr.Val,
			Next: keepGoing(curr, curr.Next),
		}
	}
	return keepGoing(nil, xs)
}
