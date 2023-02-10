package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Connection struct {
	Username string
	Password string
	Address  string
	Port     string
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

	return fmt.Sprintf("%s:%s@(%s:%s)/%s",
		connData.Username,
		connData.Password,
		connData.Address,
		connData.Port,
		connData.Database)
}

func (i Interface) connect() {

}
