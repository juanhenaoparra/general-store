package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"../models/assignment"
	"../models/device"
	"../models/ip"
	"../models/transaction"
	"../utils"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
)

func syncTransactions(date *int, transactions *transaction.Repo, ips *ip.Repo, devices *device.Repo, assignments *assignment.AssignmentsRepo, dg *dgo.Dgraph) map[string]string {
	//buyersAssignments map[string]string

	// Read the body and manage errors
	body, err := utils.ExtractDataFrom("transactions", *date)
	if err != nil {
		fmt.Println(err)
	}

	// Parse the transactions response of the server into [][]string
	parsedTransactions := parseTransactionsData(body)
	// Parse the ips and devices
	listIPS, listDevices := extractIpsAndDevices(parsedTransactions)

	// Sync and Save the ips and devices in our AssignmentsRepo structure
	assignments.Ips = SyncIPS(listIPS, ips, dg)
	assignments.Devices = SyncDevices(listDevices, devices, dg)

	// Join all and save in transactions
	utils.GiveMeRepoTransactionStructure(parsedTransactions, date, transactions, assignments)

	return saveTransactions(transactions, dg)

}

func parseTransactionsData(byteText []byte) [][]string {
	text := string(byteText)
	separator := "!"
	// Replace the rare chracter by 'separator', because performance
	replacedRare := strings.ReplaceAll(text, "\u0000", separator)
	// Split transactions by our 'separator' identifier
	separeTransactions := strings.Split(replacedRare, separator+separator)

	var parsed [][]string
	for _, t := range separeTransactions[:len(separeTransactions)-2] {
		parsed = append(parsed, strings.Split(t, separator))
	}
	return parsed
}

func extractIpsAndDevices(parsedTransactions [][]string) ([]string, []string) {
	var listIPS []string
	var listDevices []string

	for _, t := range parsedTransactions {
		if utils.CheckBack(listIPS, t[2]) == false {
			listIPS = append(listIPS, t[2])
		}

		if utils.CheckBack(listDevices, t[3]) == false {
			listDevices = append(listDevices, t[3])
		}
	}

	return listIPS, listDevices
}

func saveTransactions(transactions *transaction.Repo, dg *dgo.Dgraph) map[string]string {
	ctx := context.Background()
	mu := &api.Mutation{
		CommitNow: true,
	}

	fmt.Printf("Length of transactions: %v\n", len(transactions.Transactions))

	pb, err := json.Marshal(transactions.Transactions)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb

	assignments, err := dg.NewTxn().Mutate(ctx, mu)

	if err != nil {
		log.Fatal(err)
	}

	for i := range transactions.Transactions {
		transactions.Transactions[i].UID = assignments.Uids[transactions.Transactions[i].ID]
	}

	return assignments.Uids
}
