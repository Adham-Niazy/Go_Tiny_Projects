package main

import "fmt"

// We can declare variable outside a function but without intialization
var deckSize int

func main() {
	cards := newDeck()
	fmt.Println(toString(cards))
}