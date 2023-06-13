package config

// GameConfig stores game specific settings
type GameConfig struct {
	AnswersPath       string `json:"answersPath"`
	GuessesPath       string `json:"guessesPath"`
	ChallengeDuration int    `json:"challengeDuration"`
	InactivityTimeout int    `json:"inactivityTimeout"`
	InactivityWarning int    `json:"inactivityWarning"`
}
