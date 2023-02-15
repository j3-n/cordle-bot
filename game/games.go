package game

import (
	"cordle/wordle"
	"sync"

	"github.com/bwmarrin/discordgo"
)

// These are the available modes that Cordle can be played in
type GameMode int

const (
	Duel           GameMode = iota
	FreeForAll     GameMode = iota
	TimeAttack     GameMode = iota
	TeamDeathmatch GameMode = iota
)

// Game encapsulates the information about an individual game of Cordle, regardless of mode
type Game struct {
	Mode    GameMode
	Players []*discordgo.User
	Games   []*wordle.WordleGame
}

// Thread safe map of channel IDs to the games in them
var games struct {
	g  map[string]*Game
	mu sync.Mutex
}

// Initialise the games map
func init() {
	games.g = make(map[string]*Game)
}
