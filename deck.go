package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// A (Recievers) function which works like methods exactly
func newDeck() deck {
	cards := deck{}
	cardSuits := []string {"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string {"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[handSize:], d[:handSize]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}