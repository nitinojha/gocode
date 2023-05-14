package main

import (
	"fmt"
)

func main() {
	mystr := "hello alexa" //create string
	fmt.Println("The string given here is:", mystr)
	frequency := make(map[rune]int) //create frequency map using make function

	for _, val := range mystr {
		if _, ok := frequency[val]; ok {
			frequency[val]++
		} else {
			frequency[val] = 1 //assign frequencies
		}
	}
	fmt.Println("The frequency of characters in the string is:")

	fmt.Println(frequency) //print frequency map

	type Record struct {
		UID  int
		Type string
		Year string
	}

	type SumRecord struct {
		Type string
		Year string
	}

	m := make(map[int]Record)

	// e.g. [{"1996","non-fiction"}:4], representing 4 occurrences of {"1996","non-fiction"}
	srMap := make(map[SumRecord]int)

	// add records
	m[1] = { "Type": "fiction", "Year": 1996},
	m[2] = {"Type": "non-fiction", "Year": 1996}
	// loop over records
	for key := range m {
		sr := SumRecord{
			Type: m[key].Type,
			Year: m[key].Year,
		}
		// creates new counter or increments existing pair counter by 1
		srMap[sr] += 1
	}
	// print all mappings
	fmt.Println(srMap)

	// specific example
	fmt.Println(srMap[SumRecord{
		Year: "1996",
		Type: "non-fiction",
	}])
}
