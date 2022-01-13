package main

import (
	"fmt"

	"github.com/w-h-a/more-dsa/binarysearch"
	"github.com/w-h-a/more-dsa/lists"
)

func main() {
	fmt.Println(binarysearch.FindMin([]int{0, 7}))
	fmt.Println(binarysearch.Search([]int{3, 1}, 3))

	list := lists.UpTo(0, 4)
	fmt.Println(list)
	fmt.Println(lists.SliceOfList(list))
	list2 := lists.Take(list, 3)
	fmt.Println(lists.SliceOfList(list2))
	list3 := lists.Drop(list, 3)
	fmt.Println(lists.SliceOfList(list3))
	res, _ := lists.Read(list, 3)
	fmt.Println(res)
	fmt.Println(lists.SliceOfList(list))
	list4 := lists.Insert(list, 3, 7)
	fmt.Println(lists.SliceOfList(list4))
	fmt.Println(lists.SliceOfList(list))
	list5 := lists.Delete(list4, 3)
	fmt.Println("5", lists.SliceOfList(list5))
	fmt.Println("4", lists.SliceOfList(list4))
	fmt.Println("0", lists.SliceOfList(list))
	list6 := lists.Reverse(list)
	fmt.Println("6", lists.SliceOfList(list6))
	fmt.Println("0", lists.SliceOfList(list))
}
