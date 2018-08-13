package card

import (
	"fmt"
	"math"
	"testing"
)

// #################################
// Test Card.String()
// #################################

func TestStringOfCards(t *testing.T) {
	testCases := map[Card]string{
		Card{Rank: ACE, Suit: SPADES}:     "Ace of Spades",
		Card{Rank: TWO, Suit: SPADES}:     "Two of Spades",
		Card{Rank: THREE, Suit: SPADES}:   "Three of Spades",
		Card{Rank: FOUR, Suit: HEARTS}:    "Four of Hearts",
		Card{Rank: FIVE, Suit: HEARTS}:    "Five of Hearts",
		Card{Rank: SIX, Suit: HEARTS}:     "Six of Hearts",
		Card{Rank: SEVEN, Suit: DIAMONDS}: "Seven of Diamonds",
		Card{Rank: EIGHT, Suit: DIAMONDS}: "Eight of Diamonds",
		Card{Rank: NINE, Suit: DIAMONDS}:  "Nine of Diamonds",
		Card{Rank: TEN, Suit: CLUBS}:      "Ten of Clubs",
		Card{Rank: JACK, Suit: CLUBS}:     "Jack of Clubs",
		Card{Rank: QUEEN, Suit: CLUBS}:    "Queen of Clubs",
		Card{Rank: KING, Suit: CLUBS}:     "King of Clubs",
	}
	for card, expected := range testCases {
		actual := fmt.Sprintf("%s", card)
		if actual != expected {
			msg := "String of card is not expected string"
			t.Fatalf("%s\nExpected: %s\nActual  : %s", msg, expected, actual)
		}
	}
}

// #################################
// Test Cards.SortByXXX()
// #################################

// Setup for test
func setupCards() (cards Cards) {
	cards = append(cards, Card{Suit: HEARTS, Rank: SEVEN})
	cards = append(cards, Card{Suit: SPADES, Rank: ACE})
	cards = append(cards, Card{Suit: SPADES, Rank: JACK})
	cards = append(cards, Card{Suit: CLUBS, Rank: NINE})
	cards = append(cards, Card{Suit: CLUBS, Rank: TEN})
	cards = append(cards, Card{Suit: HEARTS, Rank: KING})
	cards = append(cards, Card{Suit: DIAMONDS, Rank: ACE})
	cards = append(cards, Card{Suit: CLUBS, Rank: JACK})
	cards = append(cards, Card{Suit: DIAMONDS, Rank: FIVE})
	cards = append(cards, Card{Suit: DIAMONDS, Rank: SEVEN})
	cards = append(cards, Card{Suit: SPADES, Rank: FIVE})
	cards = append(cards, Card{Suit: HEARTS, Rank: QUEEN})
	return cards
}

// For test
func searchLowestSuit() (lowestSuit Suit) {
	lowestRankingOfSuit := math.MaxInt64
	for suit, rankingOfSuit := range rankingOfSuits {
		if lowestRankingOfSuit < rankingOfSuit {
			lowestRankingOfSuit = rankingOfSuit
			lowestSuit = suit
		}
	}
	return lowestSuit
}

// For test
func searchLowestRank() (lowestRank Rank) {
	lowestRankingOfRank := math.MaxInt64
	for rank, rankingOfRank := range rankingOfRanks {
		if lowestRankingOfRank < rankingOfRank {
			lowestRankingOfRank = rankingOfRank
			lowestRank = rank
		}
	}
	return lowestRank
}

// #################################
// Test Cards.SortBySuit()
// #################################

func TestSortBySuitCheckSuitsOrder(t *testing.T) {
	cards := setupCards()
	cards.SortBySuit()
	lastSuit := searchLowestSuit()
	for i, card := range cards {
		if rankingOfSuits[card.Suit] < rankingOfSuits[lastSuit] {
			expected := fmt.Sprintf("cards[%d]=%s < cards[%d]=%s", i-1, card.Suit, i, lastSuit)
			actual := fmt.Sprintf("cards[%d]=%s > cards[%d]=%s", i-1, lastSuit, i, card.Suit)
			msg := fmt.Sprintf("Expected cards[%d] is not higher than cards[%d]", i-1, i)
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
		}
		lastSuit = card.Suit
	}
}

func TestSortBySuitCheckRanksOrder(t *testing.T) {
	cards := setupCards()
	cards.SortBySuit()
	lastRank := searchLowestRank()
	lastSuit := searchLowestSuit()
	for i, card := range cards {
		if rankingOfSuits[card.Suit] == rankingOfSuits[lastSuit] {
			if rankingOfRanks[card.Rank] < rankingOfRanks[lastRank] {
				expected := fmt.Sprintf("cards[%d]=%s < cards[%d]=%s", i-1, card.Rank, i, lastRank)
				actual := fmt.Sprintf("cards[%d]=%s > cards[%d]=%s", i-1, lastRank, i, card.Rank)
				msg := fmt.Sprintf("Expected %s is not higher than %s", lastRank, card.Rank)
				t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
			}
		}
		lastRank = card.Rank
		lastSuit = card.Suit
	}
}

// #################################
// Test Cards.SortByRank()
// #################################

func TestSortByRankCheckSuitsOrder(t *testing.T) {
	cards := setupCards()
	cards.SortByRank()
	lastRank := searchLowestRank()
	lastSuit := searchLowestSuit()
	for i, card := range cards {
		if rankingOfRanks[card.Rank] == rankingOfRanks[lastRank] {
			if rankingOfSuits[card.Suit] < rankingOfSuits[lastSuit] {
				expected := fmt.Sprintf("cards[%d]=%s < cards[%d]=%s", i-1, card.Suit, i, lastSuit)
				actual := fmt.Sprintf("cards[%d]=%s > cards[%d]=%s", i-1, lastSuit, i, card.Suit)
				msg := fmt.Sprintf("Expected cards[%d] is not higher than cards[%d]", i-1, i)
				t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
			}
		}
		lastRank = card.Rank
		lastSuit = card.Suit
	}
}

func TestSortByRankCheckRanksOrder(t *testing.T) {
	cards := setupCards()
	cards.SortByRank()
	lastRank := searchLowestRank()
	for i, card := range cards {
		if rankingOfRanks[card.Rank] < rankingOfRanks[lastRank] {
			expected := fmt.Sprintf("cards[%d]=%s < cards[%d]=%s", i-1, card.Rank, i, lastRank)
			actual := fmt.Sprintf("cards[%d]=%s > cards[%d]=%s", i-1, lastRank, i, card.Rank)
			msg := fmt.Sprintf("Expected %s is not higher than %s", lastRank, card.Rank)
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
		}
		lastRank = card.Rank
	}
}

// #################################
// Test Cards.SetRankingOfXXX()
// #################################

// #################################
// Test Cards.SetRankingOfSuits()
// #################################

func TestSetRankingOfSuits(t *testing.T) {
	suits := []Suit{
		CLUBS,
		DIAMONDS,
		HEARTS,
		SPADES,
	}
	SetRankingOfSuits(suits)

	lastSuit := searchLowestSuit()
	for i, suit := range suits {
		if rankingOfSuits[suit] < rankingOfSuits[lastSuit] {
			expected := i
			actual := rankingOfSuits[suit]
			msg := fmt.Sprintf("Expected %s is higher than %s, but not", suit, lastSuit)
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
		}
		lastSuit = suit
	}
}

// #################################
// Test Cards.SetRankingOfRanks()
// #################################

func TestSetRankingOfRanks(t *testing.T) {
	ranks := []Rank{
		THREE,
		FOUR,
		FIVE,
		SIX,
		SEVEN,
		EIGHT,
		NINE,
		TEN,
		JACK,
		QUEEN,
		KING,
		ACE,
		TWO,
	}
	SetRankingOfRanks(ranks)

	lastRank := searchLowestRank()
	for i, rank := range ranks {
		if rankingOfRanks[rank] < rankingOfRanks[lastRank] {
			expected := i
			actual := rankingOfRanks[rank]
			msg := fmt.Sprintf("Expected %s is higher than %s, but not", rank, lastRank)
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
		}
		lastRank = rank
	}
}

// #################################
// Test Cards.CompareByXXX()
// #################################

// #################################
// Test Cards.CompareBySuit()
// #################################

func TestCompareBySuitDifferentCards(t *testing.T) {
	cardD1 := Card{
		Suit: DIAMONDS,
		Rank: ACE,
	}
	cardS6 := Card{
		Suit: SPADES,
		Rank: SIX,
	}

	if r := CompareBySuit(cardD1, cardS6); r >= 0 {
		expected := "less than 0"
		actual := r
		msg := fmt.Sprintf("Expected %s is lower than %s, but not", cardD1, cardS6)
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
	}
}

func TestCompareBySuitSameSuitCards(t *testing.T) {
	cardD1 := Card{
		Suit: DIAMONDS,
		Rank: ACE,
	}
	cardD6 := Card{
		Suit: DIAMONDS,
		Rank: SIX,
	}

	if r := CompareBySuit(cardD1, cardD6); r <= 0 {
		expected := "larger than 0"
		actual := r
		msg := fmt.Sprintf("Expected %s is higher than %s, but not", cardD1, cardD6)
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
	}
}

func TestCompareBySuitSameRankCards(t *testing.T) {
	cardS6 := Card{
		Suit: SPADES,
		Rank: SIX,
	}
	cardD6 := Card{
		Suit: DIAMONDS,
		Rank: SIX,
	}

	if r := CompareBySuit(cardS6, cardD6); r <= 0 {
		expected := "larger than 0"
		actual := r
		msg := fmt.Sprintf("Expected %s is higher than %s, but not", cardS6, cardD6)
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
	}
}

func TestCompareBySuitSameCards(t *testing.T) {
	cardH6 := Card{
		Suit: HEARTS,
		Rank: SIX,
	}

	if r := CompareBySuit(cardH6, cardH6); r != 0 {
		expected := 0
		actual := r
		msg := fmt.Sprintf("Expected %s is draw %s, but not", cardH6, cardH6)
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
	}
}

// #################################
// Test Cards.CompareByRank()
// #################################

func TestCompareByRankDifferentCards(t *testing.T) {
	cardS6 := Card{
		Suit: SPADES,
		Rank: SIX,
	}
	cardD1 := Card{
		Suit: DIAMONDS,
		Rank: ACE,
	}

	if r := CompareByRank(cardS6, cardD1); r >= 0 {
		expected := "less than 0"
		actual := r
		msg := fmt.Sprintf("Expected %s is lower than %s, but not", cardD1, cardS6)
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
	}
}

func TestCompareByRankSameSuitCards(t *testing.T) {
	cardD1 := Card{
		Suit: DIAMONDS,
		Rank: ACE,
	}
	cardD6 := Card{
		Suit: DIAMONDS,
		Rank: SIX,
	}

	if r := CompareByRank(cardD1, cardD6); r <= 0 {
		expected := "larger than 0"
		actual := r
		msg := fmt.Sprintf("Expected %s is higher than %s, but not", cardD1, cardD6)
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
	}
}

func TestCompareByRankSameRankCards(t *testing.T) {
	cardS6 := Card{
		Suit: SPADES,
		Rank: SIX,
	}
	cardD6 := Card{
		Suit: DIAMONDS,
		Rank: SIX,
	}

	if r := CompareByRank(cardS6, cardD6); r <= 0 {
		expected := "larger than 0"
		actual := r
		msg := fmt.Sprintf("Expected %s is higher than %s, but not", cardS6, cardD6)
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
	}
}

func TestCompareByRankSameCards(t *testing.T) {
	cardH6 := Card{
		Suit: HEARTS,
		Rank: SIX,
	}

	if r := CompareByRank(cardH6, cardH6); r != 0 {
		expected := 0
		actual := r
		msg := fmt.Sprintf("Expected %s is draw %s, but not", cardH6, cardH6)
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
	}
}
