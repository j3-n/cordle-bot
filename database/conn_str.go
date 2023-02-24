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
	checkErrMsg(err, "Error opening file: ")

	cd := connData{}
	err = json.Unmarshal(content, &cd)
	checkErrMsg(err, "Error duing unmarshall(): ")

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		cd.Username,
		cd.Password,
		cd.Address,
		cd.Port,
		cd.Database)
}
