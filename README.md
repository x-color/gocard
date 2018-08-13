# gocard

This package implements simple functions for playing card.

## Install

```bash
$ go get github.com/x-color/gocard
```

## Functions

### Make a card and cards

```go
// Make ace of spades card
card1 := gocard.Card{
  Rank: gocard.JACK,
  Suit: gocard.SPADES,
}
card2 := gocard.Card{
  Rank: gocard.ACE,
  Suit: gocard.DIAMONDS,
}

// Make card list
cards := gocard.Cards{
  card1,
  card2,
}
```

### Sort cards

```go
// Sort cards by suit of cards.
cards.SortBySuit() // cards = [card1, card2]
// Sort cards by rank of cards.
cards.SortByRank() // cards = [card2, card1]
```

### Compare cards

```go
// Compare two cards by suit of cards
r := gocard.CompareBySuit(card1, card2)
// Compare two cards by rank of cards
r := gocard.CompareByRank(card1, card2)
switch {
case r > 0:
  // card1 > card2
case r < 0:
  // card1 < card2
case r == 0:
  // card1 == card2  
}
```

### Set ranking of cards

Set it if You want to change order of sorted cards and result of comparing cards.

```go
// Set ranking of cards, King is highest rank, Ace is lowest rank.
ranks := []gocard.Rank{gocard.ACE, gocard.TWO, gocard.THREE, gocard.FOUR, gocard.FIVE, gocard.SIX, gocard.SEVEN, gocard.EIGHT, gocard.NINE, gocard.TEN, gocard.JACK, gocard.QUEEN, gocard.KING}
gocard.SetRankingOfRanks(ranks)

// Set ranking of cards, DIAMONDS is highest suit, SPADES is lowest suit.
suits := []gocard.Suit{gocard.SPADES, gocard.CLUBS, gocard.HEARTS, gocard.DIAMONDS}
gocard.SetRankingOfSuits(suits)
```

#### Default ranking

Ranks of card
```go
gocard.TWO < gocard.THREE < gocard.FOUR < gocard.FIVE < gocard.SIX < gocard.SEVEN < gocard.EIGHT < gocard.NINE < gocard.TEN < gocard.JACK < gocard.QUEEN < gocard.KING < gocard.ACE
```

Suits of card
```go
gocard.CLUBS < gocard.DIAMONDS < gocard.HEARTS < gocard.SPADES
```

### Make a deck

```go
// Generate new deck
deck := gocard.NewDeck()
```

### Shuffle a deck

```go
// Shuffle the deck
deck.Shuffle()
```

### Draw a card

```go
// Draw a card from the top of the deck
card, err := deck.Draw()
```

### Put a card

```go
// Put a card on the top of the deck
deck.PutTop(card)
// Put a card on the bottom of the deck
deck.PutBottom(card)
```

## Files

```bash
gocard/
├── card.go       # define Card, Cards
├── card_test.go  # test code
├── deck.go       # define Deck
├── deck_test.go  # test code
└── example
    └── main.go   # simple Blackjack
```
