package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func HelloBicha(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/HelloWorld.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", HelloBicha)

	fmt.Println("Server running localhost:8888")
	http.ListenAndServe(":8888", nil)
}
