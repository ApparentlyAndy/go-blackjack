package main

import "strconv"

type Card struct {
	suit  string
	value string
}

func (card Card) getValue() int {
	value := 0
	switch card.value {
	case "Ace":
		value = 1
	case "Jack":
		value = 10
	case "Queen":
		value = 10
	case "King":
		value = 10
	default:
		if i, err := strconv.Atoi(card.value); err == nil {
			value = i
		}
	}
	return value
}

func (card Card) getSuit() string {
	return card.suit
}

func newCard(suit, value string) Card {
	return Card{
		suit:  suit,
		value: value,
	}
}
