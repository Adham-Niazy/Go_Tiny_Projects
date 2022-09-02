package main

import "fmt"

// We can declare variable outside a function but without intialization
var deckSize int

func main() {
	cards := newDeck()
	// cards.print()

	hand, remaingCards := cards.deal(5)

	fmt.Println(hand)
	fmt.Println(remaingCards)
	fmt.Println(cards)

}