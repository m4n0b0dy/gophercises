package blackjack_players

import (
	"blackjack/cards"
	"fmt"
)

type Player struct {
	HandValue uint8
	Cards     []cards.Card
}

// A big improvement could be instead of slice of cards make a hand type
// then re write a lot of these hand receivers with that type

func (p Player) GetHandValue() uint8 {
	var v uint8
	containsAce := false
	for _, c := range p.Cards {
		v = v + uint8(c.Rank)
		if c.Rank == cards.Ace {
			containsAce = true
		}
	}
	if containsAce && v <= 11 {
		v = v + 10
	}
	return v
}

func (p Player) PrintCards() {
	fmt.Println(p.Cards)
}

func (p *Player) SetHandVal() {
	p.HandValue = p.GetHandValue()
}

func (p *Player) Hit(c cards.Card) {
	fmt.Printf("Hit Card: %v-%v\n", c.Suit, c.Rank)
	p.Cards = append(p.Cards, c)
	p.SetHandVal()
}

func (p Player) CheckBust() bool {
	if p.HandValue > 21 {
		fmt.Println("busted")
		return true
	}
	return false
}

func (p Player) ShowHand() {
	fmt.Println(p.Cards)
}
