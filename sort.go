package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sorting integers
	numbers := []int{9, 3, 6, 1, 8, 2, 4, 7, 5}
	sort.Ints(numbers)
	fmt.Println(numbers)

	//Reverse Sorting Data
	sort.Slice(numbers, func(x, y int) bool {
		return numbers[y] < numbers[x]
	})
	fmt.Println(numbers)

	// Sorting strings
	names := []string{"Alice", "Charlie", "Bob", "Eve"}
	sort.Strings(names)
	fmt.Println(names)

	// Sorting custom types
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Charlie", Age: 30},
		{Name: "Bob", Age: 20},
	}

	for key, val := range people {
		fmt.Println(key, val)
		fmt.Printf("Type of val is %T with val %v \n", val.Name, val.Name)
		fmt.Printf("Type of val is %T with val %v \n", val.Age, val.Age)
		fmt.Printf(" pesron %d name is %s whose  age is %d \n", key+1, val.Name, val.Age)
	}
	// Implementing the sort.Interface for custom sorting
	// type ByAge []Person

	// func (a ByAge) Len() int {
	// 	 return len(a)
	// }
	// func (a ByAge) Swap(i, j int) {
	// 	a[i], a[j] = a[j], a[i]
	// }
	// func (a ByAge) Less(i, j int) bool {
	// 	return a[i].Age < a[j].Age
	// }

	// sort.Sort(ByAge(people))
	fmt.Println(people)
}
