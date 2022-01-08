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
			// Shopify App Install Endpoints - handle OAuth with Shopify
			r.Get("/install", endpoints.GetInstall)
			r.Post("/install", endpoints.PostInstall)

			// Sync endpoints
			r.Get("/sync", endpoints.GetSync)
			r.Post("/sync", endpoints.PostSync)

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

			// Config - Start oAuth process for various integrated apps
			r.Get("/config/facebook", endpoints.GetFacebook)

		})
		r.NotFound(endpoints.Error404)
		r.MethodNotAllowed(endpoints.Error405)
	})

	return r
}
