package blackjack_players

import (
	"blackjack/cards"
	"fmt"
)

type PlayerTurn interface {
	ShouldHit() bool
	Hit(cards.Card)
	CheckBust() bool
	GetHandValue() uint8
	PrintCards()
}

func RunTurn(p PlayerTurn, c *cards.Deck) uint8 {
	defer func() {
		fmt.Println("Final hand")
		p.PrintCards()
		fmt.Printf("Final hand value %v\n", p.GetHandValue())
	}()
	p.PrintCards()
	for {
		ans := p.ShouldHit()
		if ans {
			p.Hit(c.DealOneFromDeck())
		} else {
			return p.GetHandValue()
		}
		if p.CheckBust() {
			return p.GetHandValue()
		}
	}

}
