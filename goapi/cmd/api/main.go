package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/emmajiugo/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting server on :8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}