package util

import (
	"database/sql"
	"log"
	"strings"
)

// CheckError logs a fatal error if err != nil
func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func CheckErrMsg(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func CheckRow(err error) (bool, error) {
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CheckTable(err error) (bool, error) {
	if err != nil {
		if strings.HasSuffix(err.Error(), "doesn't exist") {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
