package database

import (
	"fmt"
)

func AddUser(user User) {
	conn := Connect()
	query := fmt.Sprintf(
		`insert into users(id, name, wins, losses, draws, games, elo, level) 
		values(%s);`,
		user.ToSqlAdd())

	insert, err := conn.db.Query(query)

	if err != nil {
		panic(err.Error())
	}
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

	if err != nil {
		panic(err.Error())
	}
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

	if err != nil {
		panic(err.Error())
	}
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

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var user User
	result.Next()
	err = result.StructScan(&user)

	if err != nil {
		panic(err.Error())
	}
	return user
}

func GetUsers() []User {
	conn := Connect()
	result, err := conn.db.Queryx(("select * from users;"))

	if err != nil {
		panic(err.Error())
	}
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

	if err != nil {
		panic(err.Error())
	}

	return name
}

func GetStats(id int) Stats {
	conn := Connect()
	result, err := conn.db.Queryx(fmt.Sprintf(
		"select wins, losses, draws, games, elo, level from users where id=%d;",
		id))

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var stats Stats
	result.Next()
	err = result.StructScan(&stats)

	if err != nil {
		panic(err.Error())
	}
	return stats
}
