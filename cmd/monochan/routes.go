package main

import (
	"net/http"

	"github.com/KnowledgeEnjoyer/monochan/handlers"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("assets/"))))
	r.HandleFunc("/", handlers.HomePage)

	/* /b/ routes */
	r.HandleFunc("/b", handlers.BPage)

	return r
}
