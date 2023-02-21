package database

import (
	"fmt"
	"strconv"
)

func AddUser(user User) {
	query := fmt.Sprintf(
		`insert into users(id, name, wins, losses, draws, games, elo, level) 
		values(%d, '%s', %d, %d, %d, %d, %d, %d);`,
		user.Id,
		user.Name,
		user.Wins,
		user.Losses,
		user.Draws,
		user.Games,
		user.Elo,
		user.Level)

	insertRecord(query)
}

func AddUsers(users []User) {
	for _, user := range users {
		AddUser(user)
	}
}

func DeleteUser(id int) {
	query := fmt.Sprintf(
		"id=%d",
		id)

	deleteRecord("users", query)
}

func DeleteUsers(ids []int) {
	var idsStr []string
	for index, id := range ids {
		idsStr[index] = strconv.Itoa(id)
	}
	deleteRecords("users", "id=", idsStr)
}

func UpdateUser(user *User) {
	updates := fmt.Sprintf(
		"wins=%d, losses=%d, draws=%d, games=%d, elo=%d, level=%d",
		user.Wins,
		user.Losses,
		user.Draws,
		user.Games,
		user.Elo,
		user.Level)

	query := fmt.Sprintf(
		"id=%d",
		user.Id)

	updateRecord("users", updates, query)
}

func UpdateUsers(users []User) {
	for _, user := range users {
		UpdateUser(&user)
	}
}

func GetUser(id int) User {
	conn := Connect()
	result, err := conn.db.Query(fmt.Sprintf(
		"select * from users where id=%d;",
		id))

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
	conn := Connect()
	result, err := conn.db.Query(fmt.Sprintf(
		"select * from users;"))

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

}

func GetName(id int) string {
	conn := Connect()
	result, err := conn.db.Query(fmt.Sprintf(
		"select name from users where id=%d",
		id))

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var name string
	result.Next()
	err = result.Scan(&name)

	return name
}

func GetStats(id int) Stats {
	conn := Connect()
	result, err := conn.db.Query(fmt.Sprintf(
		"select wins, losses, draws, games, elo, level from users where id=%d;",
		id))

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
