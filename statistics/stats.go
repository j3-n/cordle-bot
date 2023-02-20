package statistics

import (
	"cordle/database"
	"fmt"
)

func GetStats(id int) string {
	stats := database.GetStats(id)
	name := database.GetName(id)

	return fmt.Sprintf(
		"``%s's stats:\nWins: %d\nLosses: %d\nDraws: %d\nGames: %d\nElo: %d\nLevel: %d``",
		name,
		stats.Wins,
		stats.Losses,
		stats.Draws,
		stats.Games,
		stats.Elo,
		stats.Level)
}