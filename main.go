package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/w-h-a/more-dsa/binarysearch"
	"github.com/w-h-a/more-dsa/recursion"
	"github.com/w-h-a/more-dsa/sorting"
	"github.com/w-h-a/more-dsa/trees"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	initial := randFloats(50.0, 200.0, 10, []float64{})
	fmt.Println("orig", initial)
	fmt.Println("insert", sorting.InsertSort(initial))
	fmt.Println("quick", sorting.QuickSort(initial))
	fmt.Println("merge", sorting.MergeSort(initial))

	fmt.Println(recursion.MyPow(3.0, 40))

	binarysearch.DivAndConqFindPeakElement([]int{1, 2, 3, 1})

	fmt.Println(recursion.ClimbStairs(4))

	fmt.Println(recursion.MinPathSum([][]int{[]int{1, 2, 5}, []int{3, 2, 1}}))

	fmt.Println(recursion.CoinChange([]int{1, 2, 5}, 11))

	fmt.Println("1", recursion.LongestCommonSubsequenceRec("psnw", "vozsh"))
	fmt.Println("2", recursion.LongestCommonSubsequenceRec("ezupkr", "ubmrapg"))
	fmt.Println("3", recursion.LongestCommonSubsequenceRec("abcde", "ace"))
	fmt.Println("5", recursion.LongestCommonSubsequenceRec("hofubmnylkra", "pqhgxgdofcvmr"))

	fmt.Println(" ")
	fmt.Println(trees.CompleteTree(1, 3))
	tree1 := &trees.TreeNode{
		Val: 4,
		Left: &trees.TreeNode{
			Val:   2,
			Left:  &trees.TreeNode{Val: 1},
			Right: &trees.TreeNode{Val: 3},
		},
		Right: &trees.TreeNode{
			Val:   5,
			Left:  &trees.TreeNode{Val: 6},
			Right: &trees.TreeNode{Val: 7},
		},
	}
	fmt.Println(" ")
	fmt.Println(trees.Reflect(tree1))
	fmt.Println(" ")
	inOrder := trees.SliceInOrder(tree1)
	preOrder := trees.SlicePreOrder(tree1)
	postOrder := trees.SlicePostOrder(tree1)
	fmt.Println("InOrder ", trees.BalancedTreeInOrder(inOrder))
	fmt.Println("PreOrder ", trees.BalancedTreePreOrder(preOrder))
	fmt.Println("PostOrder ", trees.BalancedTreePostOrder(postOrder))
	fmt.Println(" ")
	tree2 := &trees.TreeNode{
		Val: 50,
		Left: &trees.TreeNode{
			Val: 25,
			Left: &trees.TreeNode{
				Val:   10,
				Left:  &trees.TreeNode{Val: 4},
				Right: &trees.TreeNode{Val: 11},
			},
			Right: &trees.TreeNode{
				Val:   33,
				Left:  &trees.TreeNode{Val: 30},
				Right: &trees.TreeNode{Val: 40},
			},
		},
		Right: &trees.TreeNode{
			Val: 75,
			Left: &trees.TreeNode{
				Val:   56,
				Left:  &trees.TreeNode{Val: 52},
				Right: &trees.TreeNode{Val: 61},
			},
			Right: &trees.TreeNode{
				Val:   89,
				Left:  &trees.TreeNode{Val: 82},
				Right: &trees.TreeNode{Val: 95},
			},
		},
	}
	tree3, _ := trees.Insert(tree2, 45)
	fmt.Println("before insert: ", tree2.Left.Right.Right)
	fmt.Println("after insert: ", tree3.Left.Right.Right)
	fmt.Println(" ")
	fmt.Println("testing Delete for Trees")
	tree4, _ := trees.Delete(tree2, 4)
	fmt.Println("before delete: ", tree2.Left.Left)
	fmt.Println("after delete: ", tree4.Left.Left)
	tree5, _ := trees.Delete(tree4, 10)
	fmt.Println("before delete: ", tree4.Left.Left)
	fmt.Println("after delete: ", tree5.Left.Left)
	tree6, _ := trees.Delete(tree5, 56)
	fmt.Println("before: ", tree5.Right.Left)
	fmt.Println("after: ", tree6.Right.Left)
	tree9, _ := trees.Delete(tree6, 50)
	fmt.Println("EXTRA")
	fmt.Println("before: ", tree6.Right.Left)
	fmt.Println("after: ", tree9.Right.Left)
	fmt.Println("EXTRA")
	tree7, _ := trees.Insert(tree6, 55)
	fmt.Println("before insert: ", tree6.Right.Left.Left)
	fmt.Println("after insert: ", tree7.Right.Left.Left)
	tree8, _ := trees.Delete(tree7, 50)
	fmt.Println("before delete: ", tree7.Right.Left.Left)
	fmt.Println("after delete: ", tree8.Right.Left.Left)
}

func randFloats(min, max float64, n int, result []float64) []float64 {
	if n == 0 {
		return result
	}
	return randFloats(min, max, n-1, append(result, min+rand.Float64()*(max-min)))
}
