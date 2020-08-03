package main

import (
	// Core imports

	"fmt"
	"net/http"

	// Import Chi Router
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	// Local imports
	"./db"
	"./handler"
)

func main() {
	// Create main ports
	port := "3333"

	// Create Chi Router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// ---------- UNCOMMENT AND RUN THIS TO SETUP THE TYPES AND FIELDS ----------
	dg := db.NewClient()
	db.Setup(dg)
	// ---------- UNCOMMENT AND RUN THIS TO SETUP THE TYPES AND FIELDS ----------

	// Routes
	r.Get("/sync", handler.SyncAll) // Route to get all repos

	// r.Route("/buyer", func(r chi.Router) {
	// 	r.Use(handler.BuyerCtx)
	// })
	r.Get("/buyer", handler.GetBuyersByPage)           // Receive params first & offset
	r.Get("/buyer/{buyerId}", handler.GetBuyerProfile) // Receive id and params first & offset

	// Start the server
	fmt.Printf("Server Listening at %v port\n", port)
	http.ListenAndServe(":"+port, r)
}
