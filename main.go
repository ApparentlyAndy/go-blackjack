package main

import "fmt"

var player Player
var dealer Dealer

func askUserContinue() {
	var input string
	fmt.Println("\r\nDo you wish to play again? ( Y/N )")
	fmt.Scanln(&input)

	switch input {
	case "N":
		fmt.Println("Goodbye!")
	case "Y":
		player.resetScore()
		dealer.resetScore()
		fmt.Println("\r\n================================================")
		startGame()
	default:
		fmt.Println("Unrecognized character...goodbye!")
	}
}

func checkWinner() bool {

	if player.totalValue > 21 {
		fmt.Println("\r\nDealer wins!")
		return false
	}

	if dealer.reveal {
		if player.altTotalValue == 21 || dealer.totalValue > 21 || player.totalValue > dealer.totalValue {
			fmt.Println("\r\nPlayer wins!")
		} else if dealer.totalValue == player.totalValue {
			fmt.Println("\r\nPlayer and Dealer tie!")
		} else {
			fmt.Println("\r\nDealer wins!")
		}
		return false
	}

	return true
}

func askUser() {
	fmt.Println("\r\n[ S ] to stand.\r\n[ H ] to hit.\r\n[ Q ] to quit.")
	fmt.Println("\r\n================================================")
	var input string
	fmt.Printf("Your selection: ")
	fmt.Scanln(&input)

	switch input {
	case "S":
		player.getStat()
		dealer.revealDealer()
	case "H":
		player.hit(dealer.deal())
		player.getStat()
		dealer.getStat()

	case "Q":
		fmt.Println("Goodbye!")
	}

	shouldContinue := checkWinner()
	if shouldContinue {
		askUser()
	} else {
		askUserContinue()
	}
}

func startGame() {
	player.hit(dealer.deal())
	dealer.dealSelf()
	player.hit(dealer.deal())
	dealer.dealSelf()

	player.getStat()
	dealer.getStat()

	askUser()
}

func main() {
	dealer = createDealer()
	player = createPlayer()
	startGame()
}
