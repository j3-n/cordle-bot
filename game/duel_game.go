package game

import (
	"cordle/wordle"

	"github.com/bwmarrin/discordgo"
)

// DuelGame holds the information about a DuelGame
type DuelGame struct {
	// games stores a map of user IDs to their game
	games map[string]*wordle.WordleGame
}

// NewDuelGame creates a specialized Game struct representing a Cordle Duel Game
func NewDuelGame(th string, p []*discordgo.User) {
	// Create the shared game
	g0 := wordle.NewRandomGame()
	// Manually create a second game with the same goal word
	// This is more efficient than doing a deep copy
	g1 := &wordle.WordleGame{
		Guesses:  []string{},
		GoalWord: g0.GoalWord,
	}

	// Create the game struct and store it
	games.mu.Lock()
	games.g[th] = &DuelGame{
		games: map[string]*wordle.WordleGame{
			p[0].ID: g0,
			p[1].ID: g1,
		},
	}
	games.mu.Unlock()
}

// PlayerInGame returns true if the given player is a part of the game
func (g *DuelGame) PlayerInGame(p *discordgo.User) bool {
	_, exists := g.games[p.ID]
	return exists
}

// PlayerHasGuesses returns true if the player has guesses remaining in the game
func (g *DuelGame) PlayerHasGuesses(p *discordgo.User) bool {
	return g.games[p.ID].GuessesRemaining() > 0
}

// SubmitGuess allows a guess to be submitted to the game of a given player
// Returns the result as an array of wordle.GuessState
func (g *DuelGame) SubmitGuess(guess string, p *discordgo.User) ([5]wordle.GuessState, error) {
	pg := g.games[p.ID]
	return pg.Guess(guess)
}

// GoalWord returns the goal word for this game
func (g *DuelGame) GoalWord(p *discordgo.User) string {
	return g.games[p.ID].GoalWord
}

// GameWon returns true if the game has been won, as well as the ID of the winner
func (g *DuelGame) GameWon() (bool, string) {
	for id, g := range g.games {
		if g.Won {
			return true, id
		}
	}
	return false, ""
}

// ShouldEndInDraw returns true if the current game has reached a stalemate and should end in a draw
func (g *DuelGame) ShouldEndInDraw() bool {
	for _, g := range g.games {
		if g.GuessesRemaining() > 0 {
			return false
		}
	}
	return true
}
