package handlers

import (
	"net/http"
	"text/template"

	"github.com/KnowledgeEnjoyer/monochan/models"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}

func BPage(w http.ResponseWriter, r *http.Request) {
	threads := models.GetAllThreads()
	tmpl := template.Must(template.ParseFiles("templates/b.html"))

	tmpl.Execute(w, threads)
}
