package database

import "fmt"

type Interface struct {
	Connection string
}

func getRecord(selection string, table string, query string) User {
	record, err := db.Query(fmt.Sprintf("select %s from %s where %s;", selection, table, query))
	if err != nil {
		panic(err.Error())
	}
	defer record.Close()
	return nil
}

func getRecords(table string, query string) {
	record, err := db.Query(fmt.Sprintf("select * from %s where %s;", table, query))
	if err != nil {
		panic(err.Error())
	}
	defer record.Close()
	return nil
}

func getTable(table string) {
	record, err := db.Query(fmt.Sprintf("select * from %s;", table))
	if err != nil {
		panic(err.Error())
	}
	defer record.Close()
	return nil
}

func updateRecord() {
	
}

func updateRecords() {
	
}

func deleteRecord() {
	
}

func deleteRecords() {
	
}

func deleteTable() {
	
}

func deleteTables() {
	
}