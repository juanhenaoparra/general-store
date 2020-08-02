package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// SyncAll the repos with one date
func SyncAll(w http.ResponseWriter, r *http.Request) {
	// Get date param
	date := r.URL.Query().Get("date")

	if date == "" {
		date = strconv.FormatInt(time.Now().Unix(), 10)
	}

	fmt.Printf("Fecha a sincronizar: %v\n", date)

	buyers := syncBuyers(date)
	products := syncProducts(date)
	transactions := syncTransactions(date)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(buyers.Buyers[:10])
	json.NewEncoder(w).Encode(products.Products[:10])
	json.NewEncoder(w).Encode(transactions.Transactions[:10])
}
