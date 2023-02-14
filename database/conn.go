package database

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var lock = &sync.Mutex{}

type Conn struct {
	db sql.DB
}

var instance *Conn

func Connect() *Conn {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		instance = &Conn{connDb()}
	}

	return instance
}

func connDb() sql.DB {
	db, err := sql.Open("mysql", connStr())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	return *db
}
