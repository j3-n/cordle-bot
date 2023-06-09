package database

import (
	"cordle/internal/pkg/sql"
	"cordle/internal/pkg/util"
	"sync"
)

type Adder interface {
	AddUser(user User)
	AddUsers(users []User)
}

type Updater interface {
	UpdateUser(user *User)
	UpdateUsers(users *[]User)
}

type Reader interface {
	ReadUser(id string) User
	ReadUsers() []User
	ReadTop() []User
	ReadStats(id string) Stats
}

type Checker interface {
	CheckUser(id string) bool
}

type Deleter interface {
	DeleteUser(id string)
	DeleteUsers(ids []string)
}

type Manager interface {
	Adder
	Updater
	Reader
	Checker
	Deleter
}

type Db struct {
	ClientMu sync.Mutex
	Client   sql.Clientx
}

func NewDb(connStr string) *Db {
	return &Db{
		Client: *sql.NewClientx(connStr),
	}
}

func (d *Db) Close() {
	err := d.Client.Close()
	util.CheckErr(err)
}
