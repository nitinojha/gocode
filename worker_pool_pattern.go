package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

//test
func main() {
	maxWorkers := 3
	ctx := context.Background()

	// initialise the worker pool setup
	h := NewHotelStaff(maxWorkers)

	for i := 0; i < maxWorkers; i++ {
		// start workers
		go h.Run(ctx, i+1)
	}

	// listen to errors
	go h.ListerToComplaints(ctx)

	// trigger jobs
	for i := 1; i <= 10; i++ {
		k := i
		task := func() (int, error) {
			if k%3 == 0 {
				return k, errors.New("error..!!")
			}

			return k, nil
		}

		h.Task <- task
	}
	time.Sleep(time.Second * 5)
	fmt.Println("hotel is closed...!! see you tomorrow  ...!!!")

}

type HotelStaff struct {
	NoWorkers int
	Task      chan func() (int, error)
	err       chan error
}

func NewHotelStaff(noWorkers int) *HotelStaff {
	return &HotelStaff{NoWorkers: noWorkers,
		Task: make(chan func() (int, error), noWorkers), err: make(chan error, noWorkers)}
}

func (h *HotelStaff) Run(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			return

		case x, ok := <-h.Task:
			if !ok {
				return
			}
			tableNo, err := x()
			fmt.Printf("%d table order served by %d\n", tableNo, id)
			if err != nil {
				h.err <- fmt.Errorf("A compliant raised from table %d which is served by %d", tableNo, id)
			}
		}
	}
}

func (h *HotelStaff) ListerToComplaints(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case x, ok := <-h.err:
			if !ok {
				return
			}
			fmt.Println(x)
		}
	}
}
