package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type AiinBEE struct {
	ThreadTitle string
	ThreadBody  string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}

func bPage(w http.ResponseWriter, r *http.Request) {
	data := AiinBEE{
		ThreadTitle: "AEWHOOOO",
		ThreadBody:  "CHANZINHO NOVO, BÊ!",
	}
	tmpl := template.Must(template.ParseFiles("templates/b.html"))
	tmpl.Execute(w, data)

}

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/b", bPage)

	fmt.Println("Server running localhost:8888")
	http.ListenAndServe(":8888", nil)
}
