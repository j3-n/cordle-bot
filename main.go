package main

import (
	"cordle/database"
	"fmt"
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

	user := database.GetUser("david")
	fmt.Println(user)
}
