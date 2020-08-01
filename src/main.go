package main

import (
	// Core imports

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	// Chi router import
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	// Local imports
	"./models/buyer"
	"./utils"
)

// Get all buyers from endpoint
func syncBuyers(w http.ResponseWriter, r *http.Request) {
	// Get date string from query params
	date := r.URL.Query().Get("date")

	// Read the body and manage errors
	body, err := utils.ExtractDataFrom("buyers", date)
	if err != nil {
		fmt.Println(err)
	}

	// Create a var buyers of type Repo and conversionError
	var buyers buyer.Repo
	var conversionError error

	buyers.Date, conversionError = strconv.Atoi(date)

	if conversionError != nil {
		fmt.Println(conversionError)
	}

	json.Unmarshal(body, &buyers.Buyers) // Parse the bytes readed from the body in a Repo Property

	// This makes a for loop for every buyer and set the date
	// buyers.SetDateAll()

	// Set headers and return encoded
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(buyers)
}

func syncProducts(w http.ResponseWriter, r *http.Request) {
	// Get date string from query params
	date := r.URL.Query().Get("date")

	// Read the body and manage errors
	body, err := utils.ExtractDataFrom("products", date)
	if err != nil {
		fmt.Println(err)
	}

	normalizedCsv := utils.GetNormalizeCsv(string(body), "'", ",")

	products := utils.GiveMeProductsStructure(normalizedCsv)

	// Add date to products repo
	var conversionError error
	products.Date, conversionError = strconv.Atoi(date)

	if conversionError != nil {
		fmt.Println(conversionError)
	}

	// // Set headers and return encoded
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)

}

func main() {

	// Start chi router in port:3333
	port := "3333"
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Route to get buyers
	r.Get("/sync/buyers", syncBuyers)
	r.Get("/sync/products", syncProducts)

	// Start the server
	http.ListenAndServe(":"+string(port), r)

}
