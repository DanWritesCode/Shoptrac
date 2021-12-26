package main

import (
	"./endpoints"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"time"
)

func InitRouter(timeout int) *chi.Mux {
	r := chi.NewRouter()

	corsConf := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(
		middleware.Recoverer,
		middleware.Timeout(time.Duration(timeout)*time.Second),
		corsConf.Handler,
	)

	r.Route("/", func(r chi.Router) {
		// Protected routes
		r.Group(func(r chi.Router) {
      r.Get("/install", endpoints.GetInstall)
			// Shopify App Install Endpoint - handles OAuth with Shopify
			r.Post("/install", endpoints.PostInstall)

			// Dashboard GET Endpoint - Shows Summary
			r.Get("/summary", endpoints.GetSummary)

			// Items/Products view page
			r.Get("/items", endpoints.GetItems)

			r.Get("/orders", endpoints.GetOrders)

			r.Get("/customers", endpoints.GetCustomers)

			r.Get("/revenue", endpoints.GetRevenue)

			r.Get("/expenses", endpoints.GetExpenses)

			// Configuration POST - User Edits StonksUp' Configuration
			r.Post("/config", endpoints.PostConfig)

		})
		r.NotFound(endpoints.Error404)
		r.MethodNotAllowed(endpoints.Error405)
	})

	return r
}
