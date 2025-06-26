package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/emmajiugo/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	
	log.Println("Starting server on port 9090...")
	log.Fatal(http.ListenAndServe(":9090", r))
}