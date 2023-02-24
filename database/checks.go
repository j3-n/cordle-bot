package database

import (
	"database/sql"
	"log"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkErrMsg(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func checkRow(err error) (bool, error) {
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}
	return true, nil
}
