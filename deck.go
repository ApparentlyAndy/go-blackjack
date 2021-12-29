package main

import (
	"math/rand"
	"time"
)

type Deck []Card

func (deck Deck) shuffleDeck() Deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}

func createDeck() Deck {
	suits := []string{"♠", "♥", "♦", "♣"}
	values := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	var deck Deck = []Card{}

	for _, value := range values {
		for _, suit := range suits {
			deck = append(deck, newCard(suit, value))
		}
	}
	return deck.shuffleDeck()
}

func (deck *Deck) dealCard() Card {
	card := (*deck)[0]
	*deck = (*deck)[1:]
	if len(*deck) == 1 {
		*deck = append(*deck, createDeck()...)
	}
	return card
}
