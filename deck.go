package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func toString(list []string) string {
	return strings.Join(list, "\n")
}

func (d deck) saveToFile(path string) error {
	// 0666 => is a basic permission means anyone can read and write from the file system
	return ioutil.WriteFile(path, []byte(toString(d)), 0666)
}