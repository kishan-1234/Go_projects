package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { //t i the test handler
	d := newDeck()
	if len(d) != 16 {
		t.Error("Expected Deck length of 16,but got", len(d)) //formatted string
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.savetofile("_decktesting")
	loadeddeck := newDeckFromfile("_decktesting")
	if len(loadeddeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadeddeck))

	}
	os.Remove("_decktesting")
}
