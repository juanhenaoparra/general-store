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

	// Route to get all repos
	r.Get("/sync", handler.SyncAll)

	// Start the server
	fmt.Printf("Server Listening at %v port\n", port)
	http.ListenAndServe(":"+port, r)
}
