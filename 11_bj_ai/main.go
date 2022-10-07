package main

import (
	"blackjack/cards"
	blackjack_players "blackjack/player"
	"fmt"
)

func main() {
	deck := cards.Shuffle(cards.New(cards.Deal(1)))

	// https://www.geeksforgeeks.org/inheritance-in-golang/
	pl := blackjack_players.HumanPlayer{
		Player: blackjack_players.Player{Cards: deck.DealFromDeck(2)},
	}
	dl := blackjack_players.DealerPlayer{
		Player: blackjack_players.Player{Cards: deck.DealFromDeck(2)},
	}
	pc := blackjack_players.ComputerPlayer{
		Player: blackjack_players.Player{Cards: deck.DealFromDeck(2)},
	}

	pl.SetHandVal()
	dl.SetHandVal()
	pc.SetHandVal()

	fmt.Println(dl.Cards[0])

	players := []blackjack_players.PlayerTurn{&pl, &pc}
	for _, p := range players {
		// here's why pointers here https://stackoverflow.com/questions/46306888/implementing-interface-in-golang-gives-method-has-pointer-receiver
		fmt.Println("Going to next player")
		blackjack_players.RunTurn(p, &deck)

	}
	fmt.Println("Going to dealer")
	blackjack_players.RunTurn(&dl, &deck)
	nonDealers := []blackjack_players.PlayerTurn{&pl, &pc}
	for i, p := range nonDealers {
		if (p.CheckBust() && dl.CheckBust()) || (p.GetHandValue() == dl.HandValue) {
			fmt.Printf("Player: %v Ties\n", i)
		} else if ((p.GetHandValue() > dl.HandValue) && !p.CheckBust()) || dl.CheckBust() {
			fmt.Printf("Player: %v Wins\n", i)
		} else {
			fmt.Printf("Player: %v Loses\n", i)
		}

	}

}
