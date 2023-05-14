package main

import "fmt"

type Human struct {
	name string
}

type Student struct {
	Human
	job string
}

type Worker struct {
	Human
	job    string
	salary int
}

func (h Human) Info() {
	fmt.Printf("I am %s\n", h.name)
}

func (s Student) Info() {
	fmt.Printf("I am %s. I am a %s.\n", s.name, s.job)
}
func (w Worker) Info() {
	fmt.Printf("I am %s. I am a %s. I have %d salary\n", w.name, w.job, w.salary)
}

type Person interface {
	Info()
}

func main() {
	Tom := Student{Human{"Tom"}, "student"}
	Ben := Worker{Human{"Ben"}, "worker", 1000}
	Susan := Human{"Susan"}

	Tom.Info()
	Ben.Info()
	Susan.Info()

	var i Person

	// Becaus Student and Worker and Human all have implemented Info method
	i = Tom
	i.Info()

	i = Ben
	i.Info()

	i = Susan
	i.Info()
}
