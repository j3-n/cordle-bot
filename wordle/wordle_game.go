package wordle

import (
	"errors"
	"regexp"
)

// The maximum number of guesses allowed in a game of Wordle
const MaxGuesses int = 6
// The length of a guess
const GuessLength int = 5

// Possible states a character in a guess could be
const (
	IncorrectCharacter 	= iota
	IncorrectPosition 	= iota
	CorrectCharacter	= iota
)

// Possible errors that could be encountered when guessing
var (
	ErrOutOfGuesses 	= errors.New("Out of guesses")
	ErrInvalidLength 	= errors.New("Invalid guess length")
	ErrInvalidFormat 	= errors.New("Invalid characters in guess")
)

// Stores information about a game of Wordle (this has no information linking to Discord)
type WordleGame struct {
	Guesses 	[]string
	GoalWord 	string
}

// Guess allows the submission of a guess to a wordle game, this requires a lowercase guess
// An error is returned if the guess string is invalid or the game has no guesses remaining
func (g *WordleGame) Guess(guess string) ([5]int, error){
	// Validate the guess
	err := validateGuess(guess)
	if err != nil{
		return [5]int{}, err
	}
	// Check that the game has remaining guesses
	if g.GuessesRemaining() == 0{
		return [5]int{}, ErrOutOfGuesses
	}

	// Evaluate the guess
	result := evaluateGuess(guess, g.GoalWord)
	// Add it to the guess history
	g.Guesses = append(g.Guesses, guess)

	return result, nil
}

// GuessesRemaining checks how many guesses remain in a given game
func (g *WordleGame) GuessesRemaining() (int){
	return MaxGuesses - len(g.Guesses)
}

// validateGuess returns an appropriate error if the given guess is invalid
func validateGuess(guess string) (error){
	// Check the length of the guess
	if len(guess) != GuessLength{
		return ErrInvalidLength
	}
	// Check that the guess contains only lowercase alphabetical letters
	if !regexp.MustCompile(`^[a-z]+$`).MatchString(guess){
		return ErrInvalidFormat
	}

	return nil
}

// evaluateGuess analyses a wordle guess against a target and returns an array of character statuses
func evaluateGuess(guess string, goal string) ([5]int){
	return [5]int{}
}

// countRunes returns a map with every rune present in the string along with its number of occurrences
func countRunes(s string) (map[rune]int){
	var counts = make(map[rune]int)
	// Count the occurrences of each character in the goal string
	for _, c := range s{
		_, exists := counts[c]
		if exists{
			counts[c]++
		} else {
			counts[c] = 1
		}
	}
	return counts
}