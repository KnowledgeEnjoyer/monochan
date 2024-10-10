package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}

func main() {
	fs := http.FileServer(http.Dir("assets/"))

	http.HandleFunc("/", homePage)
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server running localhost:8888")
	http.ListenAndServe(":8888", nil)
}
