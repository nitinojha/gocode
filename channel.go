package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	go sender1(ch)
	time.Sleep(time.Second * 5)
	for v := range ch {
		fmt.Println(v)
		time.Sleep(time.Second * 1)
	}

}

func sender1(ch chan int) {
	cap1 := cap(ch)
	fmt.Println("hi", cap1)
	for i := 0; i < cap1; i++ {
		ch <- i
		fmt.Println("Value of ch ", i, " sent")
	}
	close(ch)
}
