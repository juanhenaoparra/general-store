package utils

import (
	"fmt"
	"strconv"
	"strings"

	"../models/assignment"
	"../models/product"
	"../models/transaction"
)

// GiveMeRepoProductStructure given an string matrix iterate over them and returns Repo Product Type
func GiveMeRepoProductStructure(textArray [][]string) product.Repo {
	var products product.Repo
	myDtype := []string{"Product"}

	for _, p := range textArray {
		conversionPrice, err := strconv.Atoi(p[2])

		if err != nil {
			fmt.Println(err)
		}

		products.Add(product.Product{
			UID:   "_:" + p[0],
			ID:    p[0],
			Name:  p[1],
			Price: conversionPrice,
			DType: myDtype,
		})
	}

	return products
}

// GiveMeRepoTransactionStructure given an string matrix iterate over them and returns Repo Product Type
func GiveMeRepoTransactionStructure(rawTransactions [][]string, date *int, transactions *transaction.Repo, assignments *assignment.AssignmentsRepo) {

	myDtype := []string{"Transaction"}

	// Iterate textArray until before throws error (Because the splittler leaves a blank position at final)
	for _, t := range rawTransactions {
		productUIDList := getListOfProducts(assignments, t[4])
		newTransaction := transaction.Transaction{
			ID:    t[0],
			DType: myDtype,
		}

		newTransaction.Date = make(map[string]string)
		newTransaction.BuyerID = make(map[string]string)
		newTransaction.IP = make(map[string]string)
		newTransaction.Device = make(map[string]string)

		newTransaction.Date = assignments.Date
		newTransaction.BuyerID["uid"] = assignments.Buyers[t[1]]
		newTransaction.IP["uid"] = assignments.Ips[t[2]]
		newTransaction.Device["uid"] = assignments.Devices[t[3]]
		newTransaction.Products = productUIDList

		transactions.Add(newTransaction)

	}

}

func getListOfProducts(assignments *assignment.AssignmentsRepo, text string) []map[string]string {
	var productsFinded []map[string]string
	text = text[1 : len(text)-1]

	for _, v := range strings.Split(text, ",") {
		tempProduct := make(map[string]string)
		tempProduct["uid"] = assignments.Products[v]
		productsFinded = append(productsFinded, tempProduct)
	}

	return productsFinded
}
