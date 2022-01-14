package lists

/*
Assumptions:
- Arguments: pointer to ListNode (xs)
- Return: pointer to ListNode
- Delete all duplicates such that each element appears only once and the elements are sorted.
- xs is already sorted

Examples:

Input: head = [1,1,2]
Output: [1,2]

Input: head = [1,1,2,3,3]
Output: [1,2,3]
*/

func DeleteDuplicates(xs *ListNode) *ListNode {
	hash := map[int]bool{}
	var keepGoing func(*ListNode) *ListNode
	keepGoing = func(xs *ListNode) *ListNode {
		if xs == nil {
			return nil
		}
		if hash[xs.Val] {
			return keepGoing(xs.Next)
		}
		hash[xs.Val] = true
		return &ListNode{
			Val:  xs.Val,
			Next: keepGoing(xs.Next),
		}
	}
	return keepGoing(xs)
}
