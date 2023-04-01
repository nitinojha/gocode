package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func main() {
	p1 := person{}
	p1.firstname = "nitin"
	p1.lastname = "ojha"
	data := persondetails(p1)
	fmt.Println(data)
}

func persondetails(p person) string {
	return p.firstname + p.lastname
}
