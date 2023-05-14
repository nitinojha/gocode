package main

import (
	"fmt"
	"strconv"
)

type Processor struct {
	core string
	arch string
}

type Memory struct {
	size     int
	capacity string
}

type Computer struct {
	Processor
	Memory
	price int
}

func main() {
	var a = "intiial"
	var b, c int = 10, 12
	var d = true
	e := "apple"
	fmt.Println(a, b, c, d, e)
	var h rune = 'b'
	fmt.Printf("%d %d %s %s %d %c", b, c, a, e, d, h)
	fmt.Println("")
	fmt.Printf("%v \t %T \n", h, h)
	fmt.Println("")

	fmt.Printf("%v \t %T \n", d, d)
	fmt.Println("")

	var i int = 12
	var j int8 = 127
	var k string = strconv.Itoa(i)
	var arr = make([]int, 3, 10)
	fmt.Printf("%v \t %T \n", i, i)
	fmt.Printf("%v \t %T \n", j, j)

	fmt.Printf("%v \t %T \n", k, k)
	fmt.Printf("%v \t %T \n", arr, arr)
	fmt.Printf("%v \t %v \n", len(arr), cap(arr))

	comp := Computer{}
	comp.core = "i3"
	comp.arch = "x86"
	comp.size = 256
	comp.capacity = "ddr3"
	comp.price = 400000

	fmt.Printf("%v ,%T", comp, comp)

	t := []int{1, 2, 3, 4, 5}

	for _, k := range t {
		switch {
		case k == 1, k == 3, k == 5:
			fmt.Println("odd")
		case k == 2, k == 4:
			fmt.Println("even")
		default:
			fmt.Println("default")
		}
		fmt.Println(k)
	}

}
