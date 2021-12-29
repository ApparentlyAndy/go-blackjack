package main

import "fmt"

type Dealer struct {
	deck          Deck
	cards         []Card
	totalValue    int
	altTotalValue int
	reveal        bool
}

func createDealer() Dealer {
	dealer := Dealer{
		deck:          createDeck(),
		cards:         []Card{},
		totalValue:    0,
		altTotalValue: 0,
		reveal:        false,
	}
	return dealer
}

func (dealer Dealer) getStat() {
	fmt.Println("\r\n===== DEALER =====")
	if !dealer.reveal {
		fmt.Printf("[ %s %s ]", dealer.cards[0].suit, dealer.cards[0].value)
		fmt.Printf("[ %s %s ]\r\n", "?", "?")
	} else {
		for _, card := range dealer.cards {
			fmt.Printf("[ %s %s ]", card.suit, card.value)
		}

		fmt.Printf("\r\nTotal: %d\r\n", dealer.totalValue)

		if dealer.altTotalValue > 0 {
			fmt.Printf("Alternative Total: %d\r\n", dealer.altTotalValue)
		}
	}
}

func (dealer *Dealer) dealSelf() {
	card := dealer.deck.dealCard()
	dealer.cards = append(dealer.cards, card)
	if card.getValue() == 1 && dealer.totalValue+11 <= 21 {
		dealer.altTotalValue = dealer.totalValue + 11
	} else {
		dealer.altTotalValue = 0
	}
	dealer.totalValue += card.getValue()
}

func (dealer *Dealer) deal() Card {
	card := dealer.deck.dealCard()
	return card
}

func (dealer *Dealer) revealDealer() {
	dealer.reveal = true
	if dealer.totalValue > 16 {
		dealer.getStat()
	} else {
		dealer.dealSelf()
		dealer.revealDealer()
	}
}

func (dealer *Dealer) resetScore() {
	dealer.cards = []Card{}
	dealer.totalValue = 0
	dealer.altTotalValue = 0
	dealer.reveal = false
}
