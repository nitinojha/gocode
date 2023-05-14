package main

import "fmt"

func main() {
	//Arr : [3, 1, 7, 9, 11, 1, 8, 1,  2, 8]

	//Sum : 20
	var arr []int = []int{3, 2, 7, 9, 12, 1, 8, 1, 2, 8}
	index := 0
	var sum = arr[index]
	var temp_arr []int

	for i := index + 1; i < len(arr); i++ {
		fmt.Println(index, arr[index])
		temp_arr = append(temp_arr, arr[index])
		sum = sum + arr[index]
		if sum < 20 {
			fmt.Println(i, arr[i])
			temp_arr = append(temp_arr, arr[i])
			sum = sum + arr[i]
			if sum == 20 {
				fmt.Println(temp_arr)
				break
			}
			for j := index + 2; j < len(arr); j++ {
				sum += arr[j]
				temp_arr = append(temp_arr, arr[j])
				if sum == 20 {
					fmt.Println(temp_arr)
					break
				} else if sum > 20 {
					fmt.Println("hello", temp_arr)
					break
				}
			}
		}
		if sum == 20 {
			fmt.Println(temp_arr)
			break
		} else {
			index++
		}
		fmt.Println("heyy", temp_arr, sum)
		newSlice := make([]int, 0, len(temp_arr)) // create a new empty slice with the same underlying type and capacity
		copy(newSlice, temp_arr)                  // copy the elements from the original slice to the new slice
		fmt.Println(temp_arr, newSlice)
		temp_arr = newSlice             // assign the new slice to the original variable
		fmt.Println(temp_arr, newSlice) // Output: []
		sum = 0

	}

	fmt.Println("Hello, 世界")
}
