package database

import (
	"fmt"
)

func AddUser() {

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

func GetUser(username string) User {
	data := getRecord("*", "users", "username=")
	fmt.Println(data)
	return User{}
}

func GetUsers() {

}

func GetUserStats() {

}
