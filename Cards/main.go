package main

// We can declare variable outside a function but without intialization
var deckSize int

func main() {
	cards := newDeck()
	cards.saveToFile("dummys/Deck.txt")

	importedHand := importDeckFromFile("dummys/favorite_cards.txt")

	importedHand.shuffle()
	// Printing after shuffling
	importedHand.print()
}