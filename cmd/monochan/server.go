package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := Routes()

	db, err := sql.Open("mysql", "monochan-super-user:123@(127.0.0.1:3306)/monochan?parseTime=true")
	if err != nil {
		log.Fatal(err)

	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server running localhost:8888")
	http.ListenAndServe(":8888", r)
}
