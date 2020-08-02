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
	var nDate int
	var err error

	// Get date param
	date := r.URL.Query().Get("date")

	if date == "" {
		nDate = int(time.Now().Unix())
	} else {
		nDate, err = strconv.Atoi(date)

		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Printf("Fecha a sincronizar: %v\n", date)

	buyers := syncBuyers(nDate)
	products := syncProducts(nDate)
	transactions := syncTransactions(&nDate)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(buyers.Buyers[:10])
	json.NewEncoder(w).Encode(products.Products[:10])
	json.NewEncoder(w).Encode(transactions.Transactions[:10])
}
