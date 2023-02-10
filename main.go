package main

import (
	"cordle/database"
	"database/sql"
	"fmt"
	//"os"
	//_ "github.com/lib/pq"
	//"github.com/gofiber/fiber/v2"
)

func main() {
	person := database.User{
		Id:     1,
		Name:   "dave",
		Wins:   1,
		Losses: 2,
		Draws:  3,
		Games:  4,
		Elo:    5,
		Level:  6,
	}
	fmt.Println(person.ToString())

	connStr := database.ConnString()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connected")
	}

	fmt.Println(db)

	rows, err := db.Query("SELECT * FROM USERS;")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows)
}
