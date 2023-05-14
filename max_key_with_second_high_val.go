// c1 5
// c2 3
// // c3 7
// c4 1
// ouput := highest t price point candidate should be winner but he has to pay the proce of second highest candidate
// c3 5

// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	//fmt.Println("Hello, 世界")

	mymap := make(map[rune]int)
	mymap['a'] = 4
	mymap['b'] = 5
	mymap['c'] = 6
	mymap['d'] = 3
	//fmt.Println(mymap)
	max := mymap['a']
	sec_max := mymap['b']
	var index rune
	for k, v := range mymap {
		if max < v {
			//fmt.Println(k, v)
			sec_max = max
			max = v
			index = k
		}

	}
	fmt.Printf("%c %d", index, sec_max)

}
