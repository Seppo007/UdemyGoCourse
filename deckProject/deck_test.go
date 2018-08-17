package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("expected deck length of 16, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("expected first card to be Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("expected last card to be Four of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")
	d := newDeck();
	d.saveToFile("_decktesting");
	loadedDeck := newDeckFromFile("_decktesting");

	if len(loadedDeck) != 16 {
		t.Errorf("expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}