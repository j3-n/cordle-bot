package config

type DatabaseConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Port     string `json:"port"`
	Database string `json:"database"`
}
