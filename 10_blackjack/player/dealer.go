package blackjack_players

type DealerPlayer struct {
	Player
}

func (p DealerPlayer) ShouldHit() bool {
	if p.HandValue < 17 {
		return true
	}
	return false
}
