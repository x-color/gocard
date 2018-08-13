package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/x-color/gocard"
)

type Player struct {
	deck *Deck
	hand Cards
}

func (player *Player) draw() (card Card, err error) {
	card, err = player.deck.Draw()
	if err != nil {
		return card, err
	}
	player.hand = append((*player).hand, card)
	return card, nil
}

func culcTotalOfCards(cards Cards) (total int) {
	cards.SortByRank()
	for _, card := range cards {
		if int(card.Rank) > 10 || (card.Rank == ACE && total <= 11) {
			total += 10
		} else {
			total += int(card.Rank)
		}
	}
	return total
}

func printHand(hand Cards) {
	fmt.Println("Hand is")
	for _, card := range hand {
		fmt.Println("-", card)
	}
}

func isContinue() (yes bool) {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := strings.TrimSpace(stdin.Text())
	switch text {
	case "y":
		return true
	default:
		return false
	}
}

func turnOfPlayer(player Player) (total int, burst bool) {
	fmt.Println("## Start your turn ##")
	for {
		hand := player.hand
		printHand(hand)
		total := culcTotalOfCards(hand)
		fmt.Println("Total number of cards is", total)
		if total > 21 {
			fmt.Println("Your hand is burst!!")
			return total, true
		}

		fmt.Printf("Do you draw a card from the deck? [y/n]: ")
		if isContinue() {
			card, _ := player.draw()
			fmt.Println("You drew", card)
			fmt.Println("")
		} else {
			fmt.Println("Your turn is finished")
			return total, burst
		}
	}
}

func turnOfDealer(player Player) (total int, burst bool) {
	fmt.Println("## Start dealer's turn ##")
	for {
		hand := player.hand
		printHand(hand)
		total := culcTotalOfCards(hand)
		fmt.Println("Total number of cards is", total)
		if total > 21 {
			fmt.Println("Dealer's hand is burst!!")
			return total, true
		}
		if total >= 17 {
			fmt.Println("Dealer's turn is finished")
			return total, burst
		}
		card, _ := player.draw()
		fmt.Println("Dealer drew", card)
		fmt.Println("")
	}
}

func startGame(player, dealer Player) (result int) {
	fmt.Println("## Start Blackjack!! ##")
	fmt.Println("## Both Player draw two cards from the deck ##")

	fmt.Println("Your drawn cards are")
	player.draw()
	player.draw()
	fmt.Println("-", player.hand[0])
	fmt.Println("-", player.hand[1])
	fmt.Println("")

	fmt.Println("Dealer's drawn cards are")
	dealer.draw()
	dealer.draw()
	fmt.Println("-", dealer.hand[0])
	fmt.Println("- Unknown")
	fmt.Println("")

	totalOfPlayer, burstOfPlayer := turnOfPlayer(player)
	fmt.Println("")
	if burstOfPlayer {
		return -1
	}
	totalOfDealer, burstOfDealer := turnOfDealer(dealer)
	fmt.Println("")
	if burstOfDealer {
		return 1
	}
	return totalOfPlayer - totalOfDealer
}

func main() {
	deck := NewDeck()
	deck.Shuffle()
	player := Player{deck: &deck}
	dealer := Player{deck: &deck}

	result := startGame(player, dealer)
	switch {
	case result > 0:
		fmt.Println("## You won!! ##")
	case result == 0:
		fmt.Println("## Drew the game ##")
	case result < 0:
		fmt.Println("## You lost.. ##")
	}
	fmt.Println("## Finished game. Bye!! ##")
}
