package lists

/*
Assum:
- Arguments: pointer to ListNode (xs)
- Returns: bool
- return true iff the list is a palindrome
- What's a palindrom here?

Examples:

Input: head = [1,2,2,1]
Output: true

Input: head = [1,2]
Output: false
*/

func IsPalindrome(xs *ListNode) bool {
	ys := Reverse(xs)
	var check func(*ListNode, *ListNode) bool
	check = func(xs, ys *ListNode) bool {
		if xs == nil {
			return true
		}
		if xs.Val != ys.Val {
			return false
		}
		return check(xs.Next, ys.Next)
	}
	return check(xs, ys)
}
