//go:generate stringer -type=Suit,Rank

package cards

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // this is a special case
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	Ace   = Rank(1)
	Two   = Rank(2)
	Three = Rank(3)
	Four  = Rank(4)
	Five  = Rank(5)
	Six   = Rank(6)
	Seven = Rank(7)
	Eight = Rank(8)
	Nine  = Rank(9)
	Ten   = Rank(10)
	Jack  = Rank(10)
	Queen = Rank(10)
	King  = Rank(10)
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

type Deck []Card

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New(opts ...func(Deck) Deck) Deck {
	var cards Deck
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	// basically runs every option function defined as an arg in new
	// for example could add sorting as a function
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func DefaultSort(cards Deck) Deck {
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(less func(cards Deck) func(i, j int) bool) func(Deck) Deck {
	return func(cards Deck) Deck {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Less(cards Deck) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

func Shuffle(cards Deck) Deck {
	ret := make(Deck, len(cards))
	perm := shuffleRand.Perm(len(cards))
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return ret
}

func Jokers(n int) func(Deck) Deck {
	return func(cards Deck) Deck {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Rank: Rank(i),
				Suit: Joker,
			})
		}
		return cards
	}
}

func Filter(f func(card Card) bool) func(Deck) Deck {
	return func(cards Deck) Deck {
		var ret Deck
		for _, c := range cards {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}

func Deal(n int) func(Deck) Deck {
	return func(cards Deck) Deck {
		var ret Deck
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}

// seems like strange snytax https://stackoverflow.com/questions/18566499/how-to-remove-an-item-from-a-slice-by-calling-a-method-on-the-slice
func (dck *Deck) DealFromDeck(n int) []Card {
	d := *dck  // set var d to underlying value of Deck dck=pointer
	c := d[:n] // slice d up
	d = d[n:]  // slice d up
	*dck = d   // set underlying value of dck = d
	// strange
	return c
}

func (dck *Deck) DealOneFromDeck() Card {
	return dck.DealFromDeck(1)[0]
}
