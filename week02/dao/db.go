package dao

import "database/sql"

var db *sql.DB

func init() {
	db, e := sql.Open("mysql", "<user>:<password>@/<databasename>")
	ErrorCheck(e)

	// close database after all work is done
	defer db.Close()
	PingDB(db)
}

func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err)
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
