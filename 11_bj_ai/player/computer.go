package blackjack_players

import "math/rand"

type ComputerPlayer struct {
	Player
}

// random for now
func (p ComputerPlayer) ShouldHit() bool {
	if rand.Intn(2) == 1 {
		return true
	}
	return false
}
