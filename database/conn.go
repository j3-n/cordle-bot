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

func Disconnect() {
	if instance != nil {
		closeDb(&instance.db)
	}
}

func connDb() sqlx.DB {
	db, err := sqlx.Open("mysql", connStr())
	checkErr(err)
	//defer db.Close()

	return *db
}

func closeDb(db *sqlx.DB) {
	defer db.Close()
}
