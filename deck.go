package card

import (
	"errors"
	"math/rand"
	"time"
)

// Deck is a set of cards.
type Deck Cards

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Shuffle shuffles the deck.
func (deck Deck) Shuffle() {
	for i := len(deck); i > 0; i-- {
		randIndex := rand.Intn(i)
		deck[i-1], deck[randIndex] = deck[randIndex], deck[i-1]
	}
}

// Draw draws card from the top of the deck and returns drawn card, error of empty deck.
func (deck *Deck) Draw() (card Card, err error) {
	if len(*deck) == 0 {
		err = errors.New("couldn't draw, deck is empty")
		return card, err
	}
	card, *deck = (*deck)[0], (*deck)[1:]
	return card, err
}

// PutTop puts a card on the top of the deck.
func (deck *Deck) PutTop(card Card) {
	*deck = append(Deck{card}, *deck...)
}

// PutBottom puts a card on the bottom of the deck.
func (deck *Deck) PutBottom(card Card) {
	*deck = append(*deck, card)
}

// NewDeck returns new deck sorted by suits is a set of 52 cards.
func NewDeck() (deck Deck) {
	for suitNumber := 1; suitNumber <= 4; suitNumber++ {
		for rankNumber := 1; rankNumber <= 13; rankNumber++ {
			deck = append(deck, Card{Rank: Rank(rankNumber), Suit: Suit(suitNumber)})
		}
	}
	return deck
}
