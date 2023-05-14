package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, tasks chan int, wg *sync.WaitGroup) {
	for task := range tasks {
		fmt.Printf("Worker  %d is executing task %d", id, task)
		fmt.Println(time.Now())
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Printf("Worker  %d finished task %d", id, task)
		fmt.Println(time.Now())
		wg.Done()
	}
}

func main() {
	const noofworkers = 3
	const taskqueues = 5
	const task = 10

	tasks := make(chan int, taskqueues)
	var wg sync.WaitGroup
	fmt.Println(tasks)
	for i := 1; i <= task; i++ {
		go worker(i, tasks, &wg)
	}

	// Add tasks to the task queue
	for i := 1; i <= task; i++ {
		wg.Add(1)
		tasks <- i
	}

	// Wait for all tasks to be completed
	wg.Wait()
	close(tasks)

}
