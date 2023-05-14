package main

import (
	"fmt"
	"time"
)

func sender(ch chan string, message string, waitTime time.Duration) {
	time.Sleep(waitTime)
	ch <- message

}

func receiver(ch chan string) {
	for {
		message := <-ch
		fmt.Println("Received:", message)
	}
}

func main() {
	ch := make(chan string)

	go sender(ch, "Hello, channels!", 1*time.Second)
	go sender(ch, "Another message", 2*time.Second)

	go receiver(ch)

	// Wait for the goroutines to finish
	time.Sleep(3 * time.Second)
	close(ch)
}
