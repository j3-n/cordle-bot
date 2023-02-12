package database

type Interactions interface {
	AddUser()
	AddUsers()
	DeleteUser()
	DeleteUsers()
	UpdateUser()
	UpdateUsers()
	GetUser()
	GetUsers()
	GetUserStats()
}

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

func GetUser(username string) {
	return getRecord("*", "users", username)
}


func GetUsers() {

}

func GetUserStats() {

}