package config

// GameConfig stores game specific settings
type GameConfig struct {
	AnswersPath       string `json:"answersPath"`
	GuessesPath       string `json:"guessesPath"`
	EmojisPath        string `json:"emojisPath"`
	ChallengeDuration int    `json:"challengeDuration"`
	InactivityTimeout int    `json:"inactivityTimeout"`
	InactivityWarning int    `json:"inactivityWarning"`
	PostGameDelay     int    `json:"postGameDelay"`
}
