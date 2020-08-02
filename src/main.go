package main

import (
	// Core imports
	"fmt"
	"net/http"

	// Import Chi Router
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	// Local imports
	"./handler"
)

func main() {

	// Start chi router in port:3333
	port := "3333"
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Route to get all repos
	r.Get("/sync", handler.SyncAll)

	// Start the server
	fmt.Printf("Server Listening at %v port\n", port)
	http.ListenAndServe(":"+port, r)
}
