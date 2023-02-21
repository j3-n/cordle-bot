package main

import "cordle/database"

func main() {
	newUser := database.User{
		Id:     567223,
		Name:   "felix",
		Wins:   56,
		Losses: 18,
		Draws:  50,
		Games:  124,
		Elo:    701,
		Level:  51}

	database.AddUser(newUser)
}
