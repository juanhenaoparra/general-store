package utils

import (
	"fmt"
	"strconv"
	"strings"

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
			ID:    p[0],
			Name:  p[1],
			Price: conversionPrice,
			DType: myDtype,
		})
	}

	return products
}

// GiveMeRepoTransactionStructure given an string matrix iterate over them and returns Repo Product Type
func GiveMeRepoTransactionStructure(textArray []string, separator string, date *int) transaction.Repo {
	var transactions transaction.Repo
	myDtype := []string{"Transaction"}

	for _, t := range textArray[:len(textArray)-2] {
		transactionListed := strings.Split(t, separator)
		productIDList := GetListOfProducts(transactionListed[4])

		transactions.Add(transaction.Transaction{
			ID:       transactionListed[0],
			BuyerID:  transactionListed[1],
			IP:       transactionListed[2],
			Device:   transactionListed[3],
			Products: productIDList,
			Date:     date,
			DType:    myDtype,
		})
	}

	return transactions
}

func GetListOfProducts(text string) []string {
	text = text[1 : len(text)-1]

	return strings.Split(text, ",")
}
