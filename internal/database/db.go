package database

import (
	"cordle/internal/config"
	"cordle/internal/pkg/sql"
	"cordle/internal/pkg/util"
	"sync"
)

type Adder interface {
	AddUser(user User) error
	AddUsers(users []User) error
}

type Updater interface {
	UpdateUser(user *User) error
	UpdateUsers(users *[]User) error
}

type Reader interface {
	ReadUser(id string) (User, error)
	ReadUsers() ([]User, error)
	ReadTop() ([]User, error)
	ReadStats(id string) (Stats, error)
}

type Checker interface {
	CheckUser(id string) bool
}

type Deleter interface {
	DeleteUser(id string) error
	DeleteUsers(ids []string) error
}

type Pinger interface {
	Ping() (bool, error)
}

type Manager interface {
	Adder
	Updater
	Reader
	Checker
	Deleter
	Pinger
}

type Db struct {
	clientMu sync.Mutex
	client   sql.Clientx
}

func NewDb(c config.DatabaseConfig) *Db {
	return &Db{
		client: *sql.NewClientx(c),
	}
}

func (d *Db) Close() {
	err := d.client.Close()
	util.CheckErr(err)
}
