package handlers

import (
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/emmajiugo/goapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddleware.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		// Middleware for account routes
		router.Use(middleware.Authorization)

		router.Get("/coin", GetCoinBalance)
	})
}