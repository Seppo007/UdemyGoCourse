package main

func main() {
	cards := newDeck()
	cards.saveToFile("mydeck.deck")

	loadedDeck := newDeckFromFile("mydeck.deck")
	loadedDeck.print()
}
