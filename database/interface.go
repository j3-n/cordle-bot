package database

type Server interface {
	connect()
	getRecord()
	getRecords()
	updateRecord()
	updateRecords()
	deleteRecord()
	deleteRecords()
	deleteTable()
	deleteTables()
}

type Interface struct {
	connection string
}

func (i Interface) getRecord() {

}

func (i Interface) getRecords() {
	
}

func (i Interface) updateRecord() {
	
}

func (i Interface) updateRecords() {
	
}

func (i Interface) deleteRecord() {
	
}

func (i Interface) deleteRecords() {
	
}

func (i Interface) deleteTable() {
	
}

func (i Interface) deleteTables() {
	
}