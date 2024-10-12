package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetDBInstance() *sql.DB {
	db, err := sql.Open("mysql", "monochan-super-user:123@(127.0.0.1:3306)/monochan?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
