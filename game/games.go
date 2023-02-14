package game

import "cordle/wordle"

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
	Players	[]string
	Games	[]wordle.WordleGame
}