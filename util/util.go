package util

import "log"

// CheckError logs a fatal error if err != nil
func CheckError(err error, msg string){
	if err != nil {
		log.Fatal(msg);
	}
}