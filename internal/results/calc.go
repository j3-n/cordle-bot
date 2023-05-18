package results

import (
	"cordle/internal/users"
)

func winCalc(attempts int, player *users.User) {
	if attempts >= 7 {
		attempts = 6
	}

	player.Wins += 1
	player.Elo += 4 * (7 - attempts)
}

func loseCalc(attempts int, player *users.User) {
	if attempts >= 7 {
		attempts = 6
	}

	player.Losses += 1
	player.Elo -= 4 * attempts

	if player.Elo < 0 {
		player.Elo = 0
	}
}

func drawCalc(player *users.User) {
	player.Draws += 1
	player.Elo += 3
}
