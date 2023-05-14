package main

import "fmt"

func main() {
	// second smallest element
	var arr = []int{12, 35, 60, 10, 34, 2, 60, 3, 2, 5, 15, 20}
	var second_smallest int
	var smallest int
	if arr[0] > arr[1] {
		fmt.Println("heyyy", arr[0], arr[1])
		smallest = arr[1]
		second_smallest = arr[0]
	} else {
		smallest = arr[0]
		second_smallest = arr[1]
	}
	for i := 1; i < len(arr); i++ {
		if smallest > arr[i] {
			smallest = arr[i]
		} else if arr[i] != smallest && second_smallest > arr[i] {
			second_smallest = arr[i]
		}
	}
	fmt.Printf("Second Smallest element in given array is : %d", second_smallest)
}
