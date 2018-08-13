/*
Package card implements simple functions for playing card.

Default ranking of cards is
 Suits: Clubs < Diamonds < Hearts < Spades
 Ranks: Two < Three < Four < Five < Six < Seven < Eight < Nine < Ten < Jack < Queen < King < Ace
*/
package card

import (
	"fmt"
	"sort"
)

// Rank is rank of card. (ACE ~ KING)
type Rank int

// String returns string of rank of card. (e.g. Ace)
func (rank Rank) String() (msg string) {
	switch rank {
	case 1:
		return "Ace"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	case 10:
		return "Ten"
	case 11:
		return "Jack"
	case 12:
		return "Queen"
	case 13:
		return "King"
	default:
		return "Unknown"
	}
}

// Suit is suit of card. (SPADES, HEARTS, DIAMONDS, CLUBS)
type Suit int

// String returns string of suit of card. (e.g. Spades)
func (suit Suit) String() (msg string) {
	switch suit {
	case 1:
		return "Spades"
	case 2:
		return "Hearts"
	case 3:
		return "Diamonds"
	case 4:
		return "Clubs"
	default:
		return "Unknown"
	}
}

// Card is a card has rank and suit.
type Card struct {
	Rank Rank
	Suit Suit
}

// String returns string of card. (e.g. Ace of Spades)
func (card Card) String() (msg string) {
	msg = fmt.Sprintf("%s of %s", card.Rank, card.Suit)
	return msg
}

// Cards is a slice of Card.
type Cards []Card

// Len returns length of Cards for SortByXXX.
func (cards Cards) Len() (length int) {
	return len(cards)
}

// Swap swaps two cards in Cards for SortByXXX.
func (cards Cards) Swap(i, j int) {
	cards[i], cards[j] = cards[j], cards[i]
}

// BySuit is wrapper of Cards for SortBySuit.
type BySuit struct {
	Cards
}

// Less returns result first card is less than second card in Cards for SortBySuit.
func (b BySuit) Less(i, j int) (less bool) {
	return CompareBySuit(b.Cards[i], b.Cards[j]) < 0
}

// ByRank is wrapper of Cards for SortByRank.
type ByRank struct {
	Cards
}

// Less returns result first card is less than second card in Cards for SortByRank.
func (b ByRank) Less(i, j int) bool {
	return CompareByRank(b.Cards[i], b.Cards[j]) < 0
}

// SortBySuit sorts cards by suit of cards in ascending order.
// If cards have same suit, it sorts cards by rank of cards.
func (cards Cards) SortBySuit() {
	sort.Sort(BySuit{cards})
}

// SortByRank sorts cards by rank of cards in ascending order.
// If cards have same rank, it sorts cards by suit of cards.
func (cards Cards) SortByRank() {
	sort.Sort(ByRank{cards})
}

// These constant values are suits of card.
const (
	SPADES Suit = iota + 1
	HEARTS
	DIAMONDS
	CLUBS
)

// These constant values are ranks of card.
const (
	ACE Rank = iota + 1
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

var rankingOfSuits = map[Suit]int{
	CLUBS:    1,
	DIAMONDS: 2,
	HEARTS:   3,
	SPADES:   4,
}

var rankingOfRanks = map[Rank]int{
	TWO:   1,
	THREE: 2,
	FOUR:  3,
	FIVE:  4,
	SIX:   5,
	SEVEN: 6,
	EIGHT: 7,
	NINE:  8,
	TEN:   9,
	JACK:  10,
	QUEEN: 11,
	KING:  12,
	ACE:   13,
}

// SetRankingOfSuits sets ranking of suits of card.
// Ranking is set last suit is higher than first suit.
func SetRankingOfSuits(suits []Suit) {
	for i, suit := range suits {
		rankingOfSuits[suit] = i
	}
}

// SetRankingOfRanks sets ranking of ranks of card.
// Ranking is set last rank is higher than first sudrankit.
func SetRankingOfRanks(ranks []Rank) {
	for i, rank := range ranks {
		rankingOfRanks[rank] = i
	}
}

// CompareBySuit compares two cards in Cards by suit of cards and returns diff of cards.
// If two cards have same suit, it compares cards by rank of cards.
// Return diff > 0 (card1 > card2), diff = 0 (card1 == card2), diff < 0 (card1 < card2)
func CompareBySuit(card1 Card, card2 Card) (diff int) {
	diff = rankingOfSuits[card1.Suit] - rankingOfSuits[card2.Suit]
	if diff == 0 {
		diff = rankingOfRanks[card1.Rank] - rankingOfRanks[card2.Rank]
	}
	return diff
}

// CompareByRank compares two cards in Cards by rank of cards and returns diff of cards.
// If two cards have same rank, it compares cards by suit of cards.
// Return diff > 0 (card1 > card2), diff = 0 (card1 == card2), diff < 0 (card1 < card2)
func CompareByRank(card1 Card, card2 Card) (diff int) {
	diff = rankingOfRanks[card1.Rank] - rankingOfRanks[card2.Rank]
	if diff == 0 {
		diff = rankingOfSuits[card1.Suit] - rankingOfSuits[card2.Suit]
	}
	return diff
}
