package game

import (
	"cordle/internal/database"

	"github.com/bwmarrin/discordgo"
)

// getNewScores is called after a ranked duel ends and is used to update the scores of both players
// Returns the newly updated scores of the users after elo calculations are complete
func getNewScores(w *discordgo.User, l *discordgo.User) (int, int) {
	// Retrieve the users from the database
}

// findOrCreateUser attempts to retrieve a user from the database. If the user does not exist, they are created
func findOrCreateUser(id string) *database.User {
}
