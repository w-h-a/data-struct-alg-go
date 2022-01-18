package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/w-h-a/more-dsa/binarysearch"
	"github.com/w-h-a/more-dsa/recursion"
	"github.com/w-h-a/more-dsa/sorting"
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
}

func randFloats(min, max float64, n int, result []float64) []float64 {
	if n == 0 {
		return result
	}
	return randFloats(min, max, n-1, append(result, min+rand.Float64()*(max-min)))
}
