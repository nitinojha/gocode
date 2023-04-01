package main

import "fmt"

func main() {
	fmt.Println("this is tets code !!")
	cards := newdeck()
	cards.print()
	deck1, deck2 := deal(cards, 5)
	deck1.print()
	deck2.print()
}
