package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"../db"
	"../models/assignment"
	"../models/datemodel"
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
	var dateModel datemodel.Date

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

	// Create connection with Dgraph
	dg := db.NewClient()

	dateModel.UID = "_:" + strconv.Itoa(nDate)
	dateModel.Date = &nDate
	dateModel.DType = []string{"Date"}

	dateAssignments := SyncDate(&nDate, &dateModel, dg)

	if len(dateAssignments["uid"]) <= 0 {
		json.NewEncoder(w).Encode("Error: La fecha ya ha sido sincronizada")
	} else {
		fmt.Printf("Fecha a sincronizar: %v\n", nDate)

		buyers, buyersAssignments := syncBuyers(nDate, dg)
		products, productsAssignments := syncProducts(nDate, dg)

		assignments.Date = dateAssignments
		assignments.Buyers = buyersAssignments
		assignments.Products = productsAssignments

		// Once everything are saved pass params and save in transactions
		assignments.Transactions = syncTransactions(&nDate, &transactions, &ips, &devices, &assignments, dg)

		response := make(map[string]int)

		response["date"] = nDate
		response["buyers"] = len(buyers.Buyers)
		response["products"] = len(products.Products)
		response["ips"] = len(ips.IPS)
		response["devices"] = len(devices.Devices)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
