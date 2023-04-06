package main

import "fmt"

func main() {
	var str = "abcdbabd"

	s := longuniqueletter(str)
	fmt.Println(s)
}

func longuniqueletter(str1 string) int {
	//var cnt int
	mymap := make(map[rune]int)

	for _, char := range str1 {
		mymap[char]++
	}

	// for char, count := range charCount {
	// 	fmt.Printf("%c: %d\n", char, count)
	// }
	fmt.Println(mymap)

	return len(mymap)
}
