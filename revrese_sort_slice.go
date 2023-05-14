package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 3, 2, 4, 5}
	strings1 := []string{"a", "b", "c", "A", "B", "C"}
	fmt.Printf("Type of arr is %T with value %v", arr, arr)
	fmt.Printf("Type of arr is %T with value %v", strings1, strings1)
	//sort.Ints(arr)
	fmt.Println(arr)
	fmt.Println(strings1)
	// sort.Slice(arr , func (a int , b int) bool {
	// 	return arr[a] > arr[b]
	// })
	sort.Slice(sort.Reverse(sort.Ints(arr)))
	sort.Slice(sort.Reverse(sort.Strings(strings1)))
	fmt.Println(arr)
	fmt.Println(strings1)
}
