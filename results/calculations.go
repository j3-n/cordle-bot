package results

import(
	"cordle/database"
)

func winCalc(attempts int, player database.User) database.User {
	player.Wins += 1
	player.Games += 1
	player.Elo += 4 * (7 - attempts)
	player.Level += 5 * (7 - attempts)

	return player
}

func loseCalc(attempts int, player database.User) database.User {
	player.Losses += 1
	player.Games += 1
	player.Elo -= 4 * attempts
	player.Level += 3

	return player
}

func drawCalc(player database.User) database.User {
	player.Draws += 1
	player.Games += 1
	player.Elo += 7
	player.Level += 10

	return player
}