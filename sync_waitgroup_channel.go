package main

import (
	"fmt"
	"sync"
	"time"
)

var Wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	ch1 := make(chan int)
	Wg.Add(3)
	time.Sleep(5)
	fmt.Println("wait")
	go func() {
		j := <-ch
		ch1 <- j + 1
		fmt.Printf("%v", j)
		Wg.Done()
	}()
	go func() {
		i := 32
		ch <- i
		fmt.Println("done 2")
		Wg.Done()
	}()
	go func() {
		k := <-ch1
		fmt.Printf("%v", k)
		Wg.Done()
	}()
	Wg.Wait()
}
