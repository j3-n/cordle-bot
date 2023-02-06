package database

import (
	"fmt"
)

func (i Interface) connect() {
	connection := i.Connection
	fmt.Println(connection)
}