package util

import "log"

// Utility function for checking errors
func CheckError(err error, msg string){
	if err != nil {
		log.Fatal(msg);
	}
}