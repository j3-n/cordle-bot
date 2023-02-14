package game

import (
	"cordle/wordle"
	"errors"
	"sync"

	"github.com/bwmarrin/discordgo"
)

// These are the available modes that Cordle can be played in
type GameMode int
const (
	Duel 			GameMode = iota
	FreeForAll 		GameMode = iota
	TimeAttack		GameMode = iota
	TeamDeathmatch	GameMode = iota
)

// Game encapsulates the information about an individual game of Cordle, regardless of mode
type Game struct {
	Mode 	GameMode
	Players	[]*discordgo.User
	Games	[]*wordle.WordleGame
}

// List of all current games, thread safe
var games struct{
	g 	[]*Game
	mu 	sync.Mutex
}

// Map specialized gamemodes to their functions to create them
var gameModes = map[GameMode]func(p []*discordgo.User) (*Game, error){
	Duel: NewDuelGame,
}

// NewGame creates a new game and stores it
func NewGame(m GameMode, p []*discordgo.User) (error){
	// Attempt to find the function to create a new game of this type with
	createGame, exists := gameModes[m]
	// If the handler does not exist, return an error
	if !exists{
		return errors.New("Unknown GameMode")
	}
	// If the handler was found, create the new game
	g, err := createGame(p)
	if err != nil{
		return err
	}
	// Lock the games array and append to it
	games.mu.Lock()
	games.g = append(games.g, g)
	games.mu.Unlock()

	return nil
}