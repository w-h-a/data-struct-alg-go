package lists

import (
	"errors"
)

var (
	ErrNilList     = errors.New("nil list")
	ErrValNotFound = errors.New("value not found")
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func UpTo(m, n int) *ListNode {
	if m > n {
		return nil
	}
	return &ListNode{
		Val:  m,
		Next: UpTo(m+1, n),
	}
}

func SliceOfList(xs *ListNode) []int {
	if xs == nil {
		return []int{}
	}
	return append([]int{xs.Val}, SliceOfList(xs.Next)...)
}

func Head(xs *ListNode) (int, error) {
	if xs == nil {
		return 0, ErrNilList
	}
	return xs.Val, nil
}

func Tail(xs *ListNode) *ListNode {
	if xs == nil || xs.Next == nil {
		return nil
	}
	return xs.Next
}

func Last(xs *ListNode) (int, error) {
	if xs == nil {
		return 0, ErrNilList
	}
	if xs.Next == nil {
		return xs.Val, nil
	}
	return Last(xs.Next)
}

func Length(xs *ListNode) int {
	if xs == nil {
		return 0
	}
	return 1 + Length(xs.Next)
}

func Take(xs *ListNode, idx int) *ListNode {
	if xs == nil || idx < 1 {
		return nil
	}
	return &ListNode{
		Val:  xs.Val,
		Next: Take(xs.Next, idx-1),
	}
}

func Drop(xs *ListNode, idx int) *ListNode {
	if xs == nil || idx < 0 {
		return nil
	}
	if idx > 0 {
		return Drop(xs.Next, idx-1)
	}
	return xs
}

func Read(xs *ListNode, idx int) (int, error) {
	return Head(Drop(xs, idx))
}

func Search(xs *ListNode, val int) (int, error) {
	var keepSearching func(*ListNode, int) (int, error)
	keepSearching = func(xs *ListNode, idx int) (int, error) {
		if xs == nil {
			return 0, ErrValNotFound
		}
		if xs.Val == val {
			return idx, nil
		}
		return keepSearching(xs.Next, idx+1)
	}
	return keepSearching(xs, 0)
}

func Prepend(xs, ys *ListNode) *ListNode {
	if xs == nil {
		return ys
	}
	if ys == nil {
		return xs
	}
	return &ListNode{
		Val:  xs.Val,
		Next: Prepend(xs.Next, ys),
	}
}

func Insert(xs *ListNode, idx, val int) *ListNode {
	return Prepend(Take(xs, idx), Prepend(&ListNode{Val: val}, Drop(xs, idx)))
}

func Delete(xs *ListNode, idx int) *ListNode {
	return Prepend(Take(xs, idx), Drop(xs, idx+1))
}

func Reverse(xs *ListNode) *ListNode {
	return ReversePrepend(xs, nil)
}

func ReversePrepend(xs, ys *ListNode) *ListNode {
	switch {
	case ys == nil:
		if xs.Next == nil {
			return &ListNode{
				Val: xs.Val,
			}
		}
		return ReversePrepend(xs.Next, &ListNode{Val: xs.Val})
	default:
		if xs.Next == nil {
			return &ListNode{
				Val:  xs.Val,
				Next: ys,
			}
		}
		return ReversePrepend(xs.Next, &ListNode{Val: xs.Val, Next: ys})
	}
}

//
