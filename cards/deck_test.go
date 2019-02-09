package main

import (
	"os"
	"testing"
)

func TestNewdeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck of 16, but got %d", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Looking for Ace of spade but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Last card is a %v instead of a Four of Clubs", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {

	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Error("The loaded deck don have exactly 16 cards")
	}
	os.Remove("_decktesting")
}
