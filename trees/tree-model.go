package trees

import (
	"errors"
)

var (
	ErrTreeIsNil     = errors.New("tree is nil")
	ErrValPresent    = errors.New("value is already present")
	ErrValNotPresent = errors.New("value is not present")
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Size(tn *TreeNode) int {
	if tn == nil {
		return 0
	}
	return 1 + Size(tn.Left) + Size(tn.Right)
}

func Max(tn *TreeNode) (int, error) {
	if tn == nil {
		return 0, ErrTreeIsNil
	}
	if tn.Right == nil {
		return tn.Val, nil
	}
	return Max(tn.Right)
}

// Depth calculates the length of the longest path from root to leaf
func Depth(tn *TreeNode) int {
	if tn == nil {
		return 0
	}
	return 1 + max(Depth(tn.Left), Depth(tn.Right))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// CompleteTree creates a complete binary tree of depth n and
// labels the nodes from k to (2^n)-1
func CompleteTree(k, n int) *TreeNode {
	if n <= 0 {
		return nil
	}
	return &TreeNode{
		Val:   k,
		Left:  CompleteTree(2*k, n-1),
		Right: CompleteTree(2*k+1, n-1),
	}
}

// Reflect forms the mirror image of a tree by exchanging left and right subtrees
// all the way down
func Reflect(tn *TreeNode) *TreeNode {
	if tn == nil {
		return nil
	}
	return &TreeNode{
		Val:   tn.Val,
		Left:  Reflect(tn.Right),
		Right: Reflect(tn.Left),
	}
}

func IsReflection(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val && IsReflection(t1.Left, t2.Right) && IsReflection(t1.Right, t2.Left)
}

func IsSameTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val && IsSameTree(t1.Left, t2.Left) && IsSameTree(t1.Right, t2.Right)
}

func SliceInOrder(tn *TreeNode) []int {
	if tn == nil {
		return []int{}
	}
	return append(append(SliceInOrder(tn.Left), tn.Val), SliceInOrder(tn.Right)...)
}

func SlicePreOrder(tn *TreeNode) []int {
	if tn == nil {
		return []int{}
	}
	return append(append([]int{tn.Val}, SlicePreOrder(tn.Left)...), SlicePreOrder(tn.Right)...)
}

func SlicePostOrder(tn *TreeNode) []int {
	if tn == nil {
		return []int{}
	}
	return append(append(SlicePostOrder(tn.Left), SlicePostOrder(tn.Right)...), tn.Val)
}

func SliceLevelOrder(tn *TreeNode) []int {
	var helper func([]*TreeNode, []int) []int
	helper = func(q []*TreeNode, result []int) []int {
		if len(q) == 0 {
			return result
		}
		if q[0].Left != nil {
			q = append(q, q[0].Left)
		}
		if q[0].Right != nil {
			q = append(q, q[0].Right)
		}
		return helper(q[1:], append(result, q[0].Val))
	}
	return helper([]*TreeNode{tn}, []int{})
}

func BalancedTreeInOrder(xs []int) *TreeNode {
	if len(xs) == 0 {
		return nil
	}
	k := len(xs) / 2
	ys := xs[k:]
	return &TreeNode{
		Val:   ys[0],
		Left:  BalancedTreeInOrder(xs[:k]),
		Right: BalancedTreeInOrder(ys[1:]),
	}
}

func BalancedTreePreOrder(xs []int) *TreeNode {
	if len(xs) == 0 {
		return nil
	}
	x := xs[0]
	xs = xs[1:]
	k := len(xs) / 2
	return &TreeNode{
		Val:   x,
		Left:  BalancedTreePreOrder(xs[:k]),
		Right: BalancedTreePreOrder(xs[k:]),
	}
}

func BalancedTreePostOrder(xs []int) *TreeNode {
	return Reflect(BalancedTreePreOrder(rev(xs)))
}

func rev(xs []int) []int {
	if len(xs) == 0 {
		return []int{}
	}
	return append(rev(xs[1:]), xs[0])
}

func BalancedTreeLevelOrder(xs []int) *TreeNode {
	var root *TreeNode
	for _, x := range xs {
		root = levelOrder(root, x)
	}
	return root
}

func levelOrder(tn *TreeNode, x int) *TreeNode {
	if tn == nil {
		return &TreeNode{
			Val: x,
		}
	}
	if x <= tn.Val {
		tn.Left = levelOrder(tn.Left, x)
	} else {
		tn.Right = levelOrder(tn.Right, x)
	}
	return tn
}

func Search(tn *TreeNode, b int) (*TreeNode, error) {
	if tn == nil {
		return nil, ErrValNotPresent
	}
	switch {
	case tn.Val > b:
		return Search(tn.Left, b)
	case tn.Val < b:
		return Search(tn.Right, b)
	default:
		return tn, nil
	}
}

func Insert(tn *TreeNode, b int) (*TreeNode, error) {
	if tn == nil {
		return &TreeNode{
			Val: b,
		}, nil
	}
	switch {
	case tn.Val > b:
		res, err := Insert(tn.Left, b)
		if err != nil {
			panic(err)
		}
		return &TreeNode{
			Val:   tn.Val,
			Left:  res,
			Right: tn.Right,
		}, nil
	case tn.Val < b:
		res, err := Insert(tn.Right, b)
		if err != nil {
			panic(err)
		}
		return &TreeNode{
			Val:   tn.Val,
			Left:  tn.Left,
			Right: res,
		}, nil
	default:
		return nil, ErrValPresent
	}
}

func Delete(tn *TreeNode, b int) (*TreeNode, error) {
	if _, err := Search(tn, b); err != nil {
		return nil, err
	}
	var delete func(*TreeNode, int) *TreeNode
	delete = func(tn *TreeNode, b int) *TreeNode {
		switch {
		case tn.Val > b:
			return &TreeNode{
				Val:   tn.Val,
				Left:  delete(tn.Left, b),
				Right: tn.Right,
			}
		case tn.Val < b:
			return &TreeNode{
				Val:   tn.Val,
				Left:  tn.Left,
				Right: delete(tn.Right, b),
			}
		default:
			switch {
			case tn.Left == nil && tn.Right == nil:
				return nil
			case tn.Left == nil:
				return tn.Right
			case tn.Right == nil:
				return tn.Left
			default:
				return &TreeNode{
					Val:   lift(tn.Right),
					Left:  tn.Left,
					Right: rightOf(tn.Right),
				}
			}
		}
	}
	return delete(tn, b), nil
}

func lift(tn *TreeNode) int {
	if tn.Left == nil {
		return tn.Val
	}
	return lift(tn.Left)
}

func rightOf(tn *TreeNode) *TreeNode {
	if tn.Left == nil {
		return tn.Right
	}
	return &TreeNode{
		Val:   tn.Val,
		Left:  rightOf(tn.Left),
		Right: tn.Right,
	}
}
