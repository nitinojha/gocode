package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{23, 32, 45, 53, 45, 53, 67}

	sort.Ints(arr)
	fmt.Println("Second Largest number by sort method is ", arr[len(arr)-2])
	var largest, second_largest int
	if arr[0] > arr[1] {
		largest = arr[0]
		second_largest = arr[1]
	} else {
		largest = arr[1]
		second_largest = arr[0]
	}
	for i := 2; i < len(arr); i++ {
		if largest < arr[i] {
			second_largest = largest
			largest = arr[i]

		} else if arr[i] > second_largest && second_largest != largest {
			second_largest = arr[i]
		}
	}
	fmt.Println("Second Largest Number is ", second_largest)
}
