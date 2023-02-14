package database

import (
	"fmt"
)

func AddUser(user User) {
	query := fmt.Sprintf("insert into users(id, name, wins, losses, draws, games, elo, level) values(%d, '%s', %d, %d, %d, %d, %d, %d);",
		user.Id,
		user.Name,
		user.Wins,
		user.Losses,
		user.Draws,
		user.Games,
		user.Elo,
		user.Level)
		
	insertRecord(query);
}

func AddUsers() {

}

func DeleteUser() {

}

func DeleteUsers() {

}

func UpdateUser() {

}

func UpdateUsers() {

}

func GetUser(id int) User {
	conn := Connect()
	result, err := conn.db.Query(fmt.Sprintf("select * from users where id=%d;", id))
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var user User
	result.Next()
	err = result.Scan(
		&user.Id,
		&user.Name,
		&user.Wins,
		&user.Losses,
		&user.Draws,
		&user.Games,
		&user.Elo,
		&user.Level)

	if err != nil {
		panic(err.Error())
	}
	return user
}

func GetUsers() {

}

func GetStats(id int) Stats {
	conn := Connect()
	result, err := conn.db.Query(fmt.Sprintf("select wins, losses, draws, games, elo, level from users where id=%d;", id))
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var stats Stats
	result.Next()
	err = result.Scan(
		&stats.Wins,
		&stats.Losses,
		&stats.Draws,
		&stats.Games,
		&stats.Elo,
		&stats.Level)

	if err != nil {
		panic(err.Error())
	}
	return stats
}
