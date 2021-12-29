package main

import (
	"fmt"
)

type Player struct {
	cards         []Card
	totalValue    int
	altTotalValue int
}

func createPlayer() Player {
	player := Player{
		cards:         []Card{},
		totalValue:    0,
		altTotalValue: 0,
	}

	return player
}

func (player Player) getStat() {
	fmt.Println("\r\n===== PLAYER =====")
	for _, card := range player.cards {
		fmt.Printf("[ %s %s ]", card.suit, card.value)
	}

	fmt.Printf("\r\nTotal: %d\r\n", player.totalValue)

	if player.altTotalValue > 0 {
		fmt.Printf("Alternative Total: %d\r\n", player.altTotalValue)
	}
}

func (player *Player) hit(card Card) {
	player.cards = append(player.cards, card)
	if card.getValue() == 1 && player.totalValue+11 < 21 {
		player.altTotalValue = player.totalValue + 11
	} else {
		player.altTotalValue = 0
	}
	player.totalValue += card.getValue()
}

func (player *Player) resetScore() {
	player.cards = []Card{}
	player.totalValue = 0
	player.altTotalValue = 0
}
