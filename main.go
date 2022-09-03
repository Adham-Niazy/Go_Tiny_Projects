package main

// We can declare variable outside a function but without intialization
var deckSize int

func main() {
	cards := newDeck()
	cards.saveToFile("Deck.txt")

	importedHand := importDeckFromFile("favorite_cards.txt")
	importedHand.print()
}