package blackjack_players

import "fmt"

type HumanPlayer struct {
	Player
}

func (p HumanPlayer) ShouldHit() bool {
	var text string
	fmt.Printf("Hand is at %v, hit?: ", p.HandValue)
	fmt.Scanf("%s\n", &text)
	if text == "s" {
		fmt.Println("Staying")
		return false
	}
	return true
}
