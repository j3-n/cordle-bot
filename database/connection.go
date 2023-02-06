package database

import (
	"fmt"
)

func (i Interface) connect() {
	connection := i.connection

	fmt.Println(connection)
}