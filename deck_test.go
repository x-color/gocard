package card

import (
	"errors"
	"testing"
)

// #################################
// Test Deck.Draw()
// #################################

func TestDrawCardFromNotEmptyDeck(t *testing.T) {
	deck := NewDeck()
	numOfcards := len(deck)
	for i := 0; i < numOfcards; i++ {
		_, err := deck.Draw()
		if err != nil {
			expected := error(nil)
			actual := err
			msg := "Couldn't draw cards from deck at 52 times"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
		}
		if len(deck) != numOfcards-i-1 {
			expected := numOfcards - i - 1
			actual := len(deck)
			msg := "Deck is not decreased"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
		}
	}
}

func TestDrawCardFromEmptyDeck(t *testing.T) {
	deck := make(Deck, 0)
	_, err := deck.Draw()
	if err == nil {
		expected := errors.New("couldn't draw, deck is empty")
		actual := err
		msg := "Couldn't catch error as drawing a card from empty deck"
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
	}
}

// #################################
// Test Deck.PutTop()
// #################################

func TestPutTop(t *testing.T) {
	deck := make(Deck, 5)
	card := Card{
		Suit: DIAMONDS,
		Rank: ACE,
	}

	if deck.PutTop(card); card != deck[0] {
		expected := card
		actual := deck[0]
		msg := "Expected putted card on the first of deck, but not"
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
	}
}

// #################################
// Test Deck.Bottom()
// #################################

func TestPutBottom(t *testing.T) {
	deck := make(Deck, 52)
	card := Card{
		Suit: DIAMONDS,
		Rank: ACE,
	}

	if deck.PutBottom(card); card != deck[len(deck)-1] {
		expected := card
		actual := deck[len(deck)-1]
		msg := "Expected putted card on the last of deck, but not"
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
	}
}

// #################################
// Test Deck.Shuffle()
// #################################

func TestShuffle(t *testing.T) {
	deck := NewDeck()
	originDeck := map[Card]int{}
	for i, card := range deck {
		originDeck[card] = i
	}

	deck.Shuffle()
	checkShuffled := false
	for i, card := range deck {
		if j, ok := originDeck[card]; ok {
			if i != j {
				checkShuffled = true
			}
		} else {
			expected := true
			actual := ok
			msg := "Expected card is contained in the deck, but not"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
		}
	}

	if !checkShuffled {
		expected := true
		actual := false
		msg := "Expected deck is shuffled, but card position is not moved"
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
	}
}
