package lists

/*
Assumptions:
- Arguments: pointer to ListNode (xs)
- Returns: bool
- Return true iff the linked list has a cycle in it.
- There is a cycle just in case
    some node in the list can be reached by _continuously_ following the next pointer

*/

func HasCycle(xs *ListNode) bool {
	visited := map[*ListNode]bool{}
	var keepGoing func(*ListNode) bool
	keepGoing = func(xs *ListNode) bool {
		if xs == nil {
			return false
		}
		if visited[xs] {
			return true
		}
		visited[xs] = true
		return keepGoing(xs.Next)
	}
	return keepGoing(xs)
}

func HasCycleFastSlow(xs *ListNode) bool {
	if xs == nil || xs.Next == nil || xs.Next.Next == nil {
		return false
	}
	s := xs
	f := xs.Next.Next
	for f.Next != nil && f.Next.Next != nil {
		if s == f {
			return true
		}
		s = s.Next
		f = f.Next.Next
	}
	return false
}
