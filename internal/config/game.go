package config

// GameConfig stores game specific settings
type GameConfig struct {
	ChallengeDuration int `json:"challengeDuration"`
	InactivityTimeout int `json:"inactivityTimeout"`
	InactivityWarning int `json:"inactivityWarning"`
}
