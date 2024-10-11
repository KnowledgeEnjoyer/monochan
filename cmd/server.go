package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}

func bPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/b.html"))
	tmpl.Execute(w, nil)

}

func main() {
	r := mux.NewRouter()
	db, err := sql.Open("mysql", "monochan-super-user:123@(127.0.0.1:3306)/monochan?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("assets/"))))

	r.HandleFunc("/", homePage)
	r.HandleFunc("/b", bPage)

	fmt.Println("Server running localhost:8888")
	http.ListenAndServe(":8888", r)
}
