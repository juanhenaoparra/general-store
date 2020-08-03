package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"../db"
	"../models/assignment"
	"../models/device"
	"../models/ip"
	"../models/transaction"
)

// SyncAll the repos with one date
func SyncAll(w http.ResponseWriter, r *http.Request) {

	// Create Main Variables
	var nDate int
	var err error
	var assignments assignment.AssignmentsRepo
	var transactions transaction.Repo
	var ips ip.Repo
	var devices device.Repo

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

	fmt.Printf("Fecha a sincronizar: %v\n", nDate)

	// Create connection with Dgraph
	dg := db.NewClient()

	_, buyersAssignments := syncBuyers(nDate, dg)
	_, productsAssignments := syncProducts(nDate, dg)

	assignments.Buyers = buyersAssignments
	assignments.Products = productsAssignments

	// Once everything are saved pass params and save in transactions
	assignments.Transactions = syncTransactions(&nDate, &transactions, &ips, &devices, &assignments, dg)

	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(buyers.Buyers[:10])
	// json.NewEncoder(w).Encode(products.Products[:10])
	// fmt.Printf("Transactions: %v, IPS: %v\n", len(transactions.GetAll()), len(ips.GetAll()))
	json.NewEncoder(w).Encode(transactions.Transactions[:100])
	// json.NewEncoder(w).Encode(ips.IPS[:10])
}
