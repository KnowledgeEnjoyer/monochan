package handlers

import (
	"net/http"
	"text/template"

	"github.com/KnowledgeEnjoyer/monochan/models"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}

func GetBPage(w http.ResponseWriter, r *http.Request) {
	// select * from Threads inner join Posts on Threads.ThreadId = Posts.ThreadId
	threads := models.GetAllThreads()
	tmpl := template.Must(template.ParseFiles("templates/b.html"))

	tmpl.Execute(w, threads)
}
