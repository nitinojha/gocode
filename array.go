package main

import (
	"fmt"
)

func main() {
	t := "level"
	for i, l := range t {
		fmt.Println(l[i])
	}
	// var k int = 3
	// sort.Ints(t)
	// //p := sort.IntSlice(t)
	// //fmt.Println(p)
	// //sort.Sort(p)
	// //fmt.Println("p", p)

	// fmt.Println("Ints:   ", t)
	// siz := len(t)
	// fmt.Printf("%v th largest number %v", k, t[siz-k])
}
