package main

import (
	"os"
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card Four of Clubs, got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktestfile")

	deck := newDeck()
	deck.saveToFile("_decktestfile")

	loadedDeck := newDeckFromFile("_decktestfile")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, got %v", len(loadedDeck))
	}

	os.Remove("_decktestfile")
}

func TestShuffle(t *testing.T) {
	deck, origDeck := newDeck(), newDeck()
	deck.shuffle()
	if reflect.DeepEqual(deck, origDeck) == true {
		t.Errorf("Expeced original deck to be different from shuffled, got equal")
	}
}
