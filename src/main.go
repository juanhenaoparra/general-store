package main

import (
	// Core imports

	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	// Chi router import
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	// Local imports
	"./models/buyer"
	"./models/product"
)

// Datasource Type
type Datasource struct {
	Buyers       string
	Products     string
	Transactions string
}

func getDataPaths() Datasource {
	path := "repo.json"
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var urls Datasource

	err = json.Unmarshal(byteValue, &urls)

	if err != nil {
		fmt.Println("Error", err)
	}

	return urls
}

// Get all buyers from endpoint
func syncBuyers(w http.ResponseWriter, r *http.Request) {
	// Search in repo.json the urls
	urls := getDataPaths()

	date, dateErr := strconv.Atoi(r.URL.Query().Get("date"))

	if dateErr != nil {
		fmt.Println(dateErr)
	}

	response, err := http.Get(urls.Buyers + "?date=" + strconv.Itoa(date))

	// Managing errors
	if err != nil {
		fmt.Println(err)
	}

	// Close Body on finish scoped function
	defer response.Body.Close()

	// Read the body and manage errors
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
	}

	// Create a var buyers of type Repo
	var buyers buyer.Repo
	buyers.Date = date
	json.Unmarshal(body, &buyers.Buyers) // Parse the bytes readed from the body in a Repo Property

	// This makes a for loop for every buyer and set the date
	// buyers.SetDateAll()

	// Set headers and return encoded
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(buyers)
}

func syncProducts(w http.ResponseWriter, r *http.Request) {
	// Search in repo.json the urls
	urls := getDataPaths()

	date, dateErr := strconv.Atoi(r.URL.Query().Get("date"))

	if dateErr != nil {
		fmt.Println(dateErr)
	}

	response, err := http.Get(urls.Products + "?date=" + strconv.Itoa(date))

	// Managing errors
	if err != nil {
		fmt.Println(err)
	}

	// Close Body on finish scoped function
	defer response.Body.Close()

	// Read the body and manage errors
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
	}

	// Replace ' char with ,
	replaced := strings.ReplaceAll(string(body), "'", ",")

	// Create a csv reader for the replaced string
	reader := csv.NewReader(strings.NewReader(replaced))
	results, _ := reader.ReadAll()

	// Create a var products of type Repo
	var products product.Repo
	products.Date = date

	for _, p := range results {
		conversionPrice, err := strconv.Atoi(p[2])

		if err != nil {
			fmt.Println(err)
		}

		products.Add(product.Product{
			ID:    p[0],
			Name:  p[1],
			Price: conversionPrice,
		})
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
