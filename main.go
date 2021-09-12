package main

import (
	"net/http"

	urls_routes "github.com/arnaugomez/url_shortener_yaml/urls/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/{path}", urls_routes.GetUrl).Methods("GET")
	r.HandleFunc("/", urls_routes.PostUrl).Methods("POST")
	http.ListenAndServe(":8080", r)
}
