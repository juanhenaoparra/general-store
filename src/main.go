package main

import (
	// Core imports

	"fmt"
	"net/http"

	// Import Chi Router
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"

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

	// ROUTES
	// sync
	r.Get("/sync", handler.SyncAll) // Route to get all repos
	// Buyer related
	r.Get("/buyer", handler.GetBuyersByPage)                       // Receive params first & offset
	r.Get("/buyer/{buyerId}", handler.GetBuyerProfile)             // Receive id and params first & offset
	r.Get("/product/{buyerId}", handler.GetProductPricesByBuyerID) // Receive buyer id
	r.Get("/product/similar", handler.GetSimilarPricesByTopLow)    // Receive params media and calculate similar products

	corsHandler := cors.Default().Handler(r)
	// Start the server
	fmt.Printf("Server Listening at %v port\n", port)
	http.ListenAndServe(":"+port, corsHandler)
}
