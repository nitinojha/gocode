package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	fmt.Printf("a = %v\n", a)
	// for i, b := range a {
	// 	fmt.Println("%v\n", b[i])
	// }
	changevalue(&a)
	// for i, b := range a {
	// 	fmt.Println("%v\n", b[i])
	// }
	fmt.Printf("a = %v\n", a)

}

func changevalue(p *[3]int) {
	p[0] *= 2
	p[1] *= 3
	p[2] *= 4
}
