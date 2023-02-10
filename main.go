package main

import (
	"cordle/database"
	"database/sql"
	"fmt"
	
	_ "github.com/go-sql-driver/mysql"
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

	db, err := sql.Open("mysql", connStr)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    fmt.Println("Success!")

	insert, err := db.Query("insert into users(id, name, wins, losses, draws, games, elo, level) values(456123, 'mother teresa', 69, 0, 0, 69, 69420, 456)")
    if err !=nil {
        panic(err.Error())
    }
    defer insert.Close()
    fmt.Println("Yay, values added!")
}
