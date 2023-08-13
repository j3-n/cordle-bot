package game

import (
	"cordle/internal/config"
	"cordle/internal/database"
	"cordle/internal/users"
	"math"

	"github.com/bwmarrin/discordgo"
)

// Database instance to be used for ranking information
var db *database.Db

// OpenDb should be called when the program is ready to connect to the database
func OpenDb() {
	db = database.NewDb(config.Config.Database)
}

// CloseDb should be called before the program quits in order to close the database connection
func CloseDb() {
	db.Close()
}

// updateScores is called after a ranked duel ends and is used to update the scores of both players
// Returns the newly updated scores of the users after elo calculations are complete
// If s is 1, the first user won, if s is 0.5 the game was a draw
func updateScores(w *discordgo.User, l *discordgo.User, s float64) (int, int) {
	// Retrieve the users from the database
	wu := findOrCreateUser(w.ID)
	lu := findOrCreateUser(l.ID)
	// Calculate the new rating scores
	ws := calculateElo(wu.Elo, lu.Elo, s)
	ls := calculateElo(lu.Elo, wu.Elo, 1-s)
	// Update the score in the database
	wu.Elo = ws
	lu.Elo = ls
	db.UpdateUser(wu)
	db.UpdateUser(lu)
	// Return the new scores
	return ws, ls
}

// calculateElo takes a player A and B's score, and returns A's adjusted score based on the outcome score given
func calculateElo(ra int, rb int, score float64) int {
	return ra + int(32*(score-calculateExpectedScore(ra, rb)))
}

// calculateExpectedScore calculates the expected score of player a given Ra and Rb as in the Elo formulae
func calculateExpectedScore(ra int, rb int) float64 {
	var diff float64 = float64(rb-ra) / float64(400)
	return float64(1) / (float64(1) + math.Pow(float64(10), diff))
}

// findOrCreateUser attempts to retrieve a user from the database. If the user does not exist, they are created
func findOrCreateUser(id string) users.User {
	exists, _ := db.CheckUser(id)
	// Create the user if they don't exist
	if !exists {
		db.AddUserDefault(id)
	}
	// Read and return the user
	user, _ := db.ReadUser(id)
	return user
}
