package main

import (
	"fmt"
	"time"
	//"time"
)

func main() {
	ch1 := make(chan int)
	//var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		go func() {
			//defer wg.Done()
			time.Sleep(time.Second * 1)
			//fmt.Println(i)
			ch1 <- i
		}()
		fmt.Println(<-ch1)

		// 	go func() {
		// 		defer wg.Done()
		// 		time.Sleep(time.Second * 1)
		// 		wg.Add(1)
		// 		v := <-ch1
		// 		fmt.Println(v)
		// 		close(ch1)

		// 	}()
	}
	close(ch1)

}

// func write(ch1 chan int) {
// 	for v := range ch1 {
// 		time.Sleep(time.Second * 1)
// 		fmt.Println(v)
// 	}

// }

// 1 1 2 2 3 3 4 4 5 5 6 6 7 7 8 8 9 9 10 10
