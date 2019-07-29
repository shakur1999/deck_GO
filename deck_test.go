package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clunbes" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

//what to test?
// ensure there are X ard in the deck
// last card should be four of spades
