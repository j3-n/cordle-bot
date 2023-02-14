package game

import (
	"cordle/wordle"
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
	Games	[]wordle.WordleGame
}

// List of all current games, thread safe
var games struct{
	g 	[]Game
	mu 	sync.Mutex
}