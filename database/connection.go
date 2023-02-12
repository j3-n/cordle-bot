package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	Username string
	Password string
	Address  string
	Port     string
	Database string
}

func connStr() string {
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
	db, err := sql.Open("mysql", connStr())
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    fmt.Println("Success!")
}
