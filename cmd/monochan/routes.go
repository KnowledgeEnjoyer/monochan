package main

import (
	"net/http"

	"github.com/KnowledgeEnjoyer/monochan/handlers"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("assets/"))))
	r.HandleFunc("/", handlers.GetHomePage)

	/* /b/ routes */
	r.HandleFunc("/b", handlers.GetBPage)

	return r
}
