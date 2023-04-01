package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newdeck() deck {
	cards := deck{}
	cardsuits := []string{"Diamond", "Spade", "Heart"}
	cardvalues := []string{"One", "Two", "Jack", "King", "Queen", "Ace"}

	for _, suit := range cardsuits {
		for _, value := range cardvalues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

func deal(z deck, handsize int) (deck, deck) {
	return z[:handsize], z[handsize:]
}
