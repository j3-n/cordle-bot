package game

import (
	"cordle/wordle"

	"github.com/bwmarrin/discordgo"
)

// NewDuelGame creates a specialized Game struct representing a Cordle Duel Game
func NewDuelGame(th string, p []*discordgo.User) {
	// Create the shared game
	g1 := wordle.NewRandomGame()
	// Manually create a second game with the same goal word
	// This is more efficient than doing a deep copy
	g2 := &wordle.WordleGame{
		Guesses:  []string{},
		GoalWord: g1.GoalWord,
	}

	// Create the game struct and store it
	games.mu.Lock()
	games.g[th] = &Game{
		Mode:    Duel,
		Players: p,
		Games:   []*wordle.WordleGame{g1, g2},
	}
	games.mu.Unlock()
}
