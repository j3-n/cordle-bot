package database

import (
	"encoding/json"
	"fmt"
	"os"
)

type connData struct {
	Username string
	Password string
	Address  string
	Port     string
	Database string
}

func connStr() string {
	content, err := os.ReadFile("database/config.json")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}

	cd := connData{}
	err = json.Unmarshal(content, &cd)
	if err != nil {
		fmt.Println("Error duing unmarshall() :", err)
	}

	return fmt.Sprintf(
		"%s:%s@(%s:%s)/%s",
		cd.Username,
		cd.Password,
		cd.Address,
		cd.Port,
		cd.Database)
}
