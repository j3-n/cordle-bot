package game

import (
	"cordle/wordle"
	"errors"

	"github.com/bwmarrin/discordgo"
)

// NewDuelGame creates a specialized Game struct representing a Cordle Duel Game
func newDuelGame(p []*discordgo.User) (*Game, error) {
	if len(p) != 2{
		// Duels may only consist of two players
		return nil, errors.New("Invalid player count")
	}

	// Create the shared game
	g1 := wordle.NewRandomGame()
	// Manually create a second game with the same goal word
	// This is more efficient than doing a deep copy
	g2 := &wordle.WordleGame{
		Guesses: 	[]string{},
		GoalWord: 	g1.GoalWord,
	}
	
	// Create the game struct and return it
	return &Game{
		Mode: Duel,
		Players: p,
		Games: []*wordle.WordleGame{g1, g2},
	}, nil
}