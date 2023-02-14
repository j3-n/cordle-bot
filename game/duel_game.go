package game

import (
	"cordle/wordle"
	"errors"

	"github.com/bwmarrin/discordgo"
)

// NewDuelGame creates a specialized Game struct representing a Cordle Duel Game
func NewDuelGame(p []*discordgo.User) (*Game, error) {
	if len(p) != 2{
		// Duels may only consist of two players
		return nil, errors.New("Invalid player count")
	}

	// Create the array of games
	var g []*wordle.WordleGame
	for range p {
		g = append(g, wordle.NewRandomGame())
	}
	
	// Create the game struct and return it
	return &Game{
		Mode: Duel,
		Players: p,
		Games: g,
	}, nil
}