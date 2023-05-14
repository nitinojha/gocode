package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}
	unique_arr_element := remove_dup(arr)
	fmt.Println(unique_arr_element)
}

func remove_dup(arr []int) []int {
	var tmp_arr []int
	my_map := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		my_map[arr[i]] += 1
		if my_map[arr[i]] > 1 {
			continue
		} else {
			tmp_arr = append(tmp_arr, arr[i])
		}
	}
	return tmp_arr
}
