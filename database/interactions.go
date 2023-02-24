package database

import (
	"fmt"
)

func AddUser(user User) {
	conn := Connect()
	insert, err := conn.db.Query(fmt.Sprintf(
		`insert into users(id, name, wins, losses, draws, games, elo, level) 
		values(%s);`,
		user.ToSqlAdd()))

	checkErr(err)
	defer insert.Close()
}

func AddUsers(users []User) {
	for _, user := range users {
		AddUser(user)
	}
}

func DeleteUser(id int) {
	conn := Connect()
	query := fmt.Sprintf(
		"id='%d'",
		id)

	delete, err := conn.db.Query(fmt.Sprintf(
		"delete from users where %s;",
		query))

	checkErr(err)
	defer delete.Close()
}

func DeleteUsers(ids []int) {
	for _, id := range ids {
		DeleteUser(id)
	}
}

func UpdateUser(user *User) {
	conn := Connect()
	updates := user.ToSqlUpdate()

	query := fmt.Sprintf(
		"id='%d'",
		user.Id)

	update, err := conn.db.Query(fmt.Sprintf(
		`update users
		set %s
		where %s;`,
		updates,
		query))

	checkErr(err)
	defer update.Close()
}

func UpdateUsers(users []User) {
	for _, user := range users {
		UpdateUser(&user)
	}
}

func GetUser(id int) User {
	conn := Connect()
	result, err := conn.db.Queryx(fmt.Sprintf(
		"select * from users where id=%d;",
		id))

	checkErr(err)
	defer result.Close()

	var user User
	result.Next()
	err = result.StructScan(&user)

	checkErr(err)
	return user
}

func GetUsers() []User {
	conn := Connect()
	result, err := conn.db.Queryx("select * from users;")
	checkErr(err)
	defer result.Close()

	var users []User
	for i := 0; result.Next(); i++ {
		err := result.StructScan(&users[i])
		if err != nil {
			panic(err.Error())
		}
	}

	return users
}

func GetTop() []User {
	conn := Connect()
	results, err := conn.db.Queryx("select * from users order by elo, name asc limit 0,10;")
	checkErr(err)
	defer results.Close()

	topTen := make([]User, 0)
	for i := 0; results.Next(); i++ {
		var user User
		err := results.StructScan(&user)
		topTen = append(topTen, user)
		checkErr(err)
	}

	return topTen
}

func GetName(id int) string {
	conn := Connect()
	result, err := conn.db.Query(fmt.Sprintf(
		"select name from users where id=%d",
		id))

	checkErr(err)
	defer result.Close()

	var name string
	result.Next()
	err = result.Scan(&name)
	checkErr(err)

	return name
}

func GetStats(id int) Stats {
	conn := Connect()
	result, err := conn.db.Queryx(fmt.Sprintf(
		"select wins, losses, draws, games, elo, level from users where id=%d;",
		id))

	checkErr(err)
	defer result.Close()

	var stats Stats
	result.Next()
	err = result.StructScan(&stats)
	checkErr(err)

	return stats
}

func CheckUser(id int) bool {
	conn := Connect()
	err := conn.db.QueryRow(fmt.Sprintf(
		"select id from users where id=%d",
		id)).Scan(&id)

	exists, err := checkRow(err)
	checkErr(err)

	return exists
}
