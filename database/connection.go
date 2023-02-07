package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Connection struct {
	Hostname string
	Port     int
	Username string
	Password string
	Database string
}

func ConnString() string {
	content, err := ioutil.ReadFile("database/config.json")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}

	connData := Connection{}
	err = json.Unmarshal(content, &connData)
	if err != nil {
		fmt.Println("Error duing unmarshall() :", err)
	}

	return fmt.Sprintf("postgresql://%s:%s@%s/todos?sslmode=disable",
		connData.Username,
		connData.Password,
		connData.Database)
}

func (i Interface) connect() {

}
