package database

import (
	"sync"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var lock = &sync.Mutex{}

type Conn struct {
	db sqlx.DB
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

func connDb() sqlx.DB {
	db, err := sqlx.Open("mysql", connStr())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.MustExec(schema)

	return *db
}
