package main

import (
	"fmt"

	"github.com/w-h-a/more-dsa/pointers"
)

func main() {
	fmt.Println(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334"))

	fmt.Println(twoSum([]int{3, 1, 2, 3, 1}, 6))

	arr := []int{0, 1, 0, 3, 12}

	pointers.MoveZeroes(arr)

	fmt.Println(arr)

	arr1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	result := pointers.RemoveDuplicates(arr1)

	fmt.Println(result)

	fmt.Println(arr1)
}
