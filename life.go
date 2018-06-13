package lorm

import "database/sql"

var db *sql.DB


//OpenDB is a method to initialize a DB
func OpenDB(driver, address string) error {
	var err error
	db, err = sql.Open(driver, address)
	if(err!=nil){
		panic(err)
	}
	return err
}

//CloseDB is a method to close a database
func CloseDB() error {
	var err error
	err = db.Close()
	return err
}

//ConnectDB is a method to get a DB
func ConnectDB() *sql.DB {
	return db
}
