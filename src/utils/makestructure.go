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

		transactions.Add(transaction.Transaction{
			ID:       t[0],
			BuyerID:  assignments.Buyers[t[1]],
			IP:       assignments.Ips[t[2]],
			Device:   assignments.Devices[t[3]],
			Products: productUIDList,
			Date:     date,
			DType:    myDtype,
		})
	}

}

func getListOfProducts(assignments *assignment.AssignmentsRepo, text string) []string {
	var productsFinded []string
	text = text[1 : len(text)-1]

	for _, v := range strings.Split(text, ",") {
		productsFinded = append(productsFinded, assignments.Products[v])
	}

	return productsFinded
}
