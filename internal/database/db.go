package database

import (
	"cordle/internal/config"
	"cordle/internal/pkg/sql"
	"cordle/internal/pkg/util"
	"cordle/internal/users"
	"sync"
)

type Adder interface {
	AddUserDefault(id string) error
	AddUser(user users.User) error
}

type Updater interface {
	UpdateUser(user users.User) error
}

type Reader interface {
	ReadUser(id string) (users.User, error)
	ListUsers() ([]users.User, error)
	ReadTop() ([]users.User, error)
}

type Checker interface {
	CheckUser(id string) bool
}

type Deleter interface {
	DeleteUser(id string) error
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
