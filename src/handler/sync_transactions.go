package handler

import (
	"fmt"
	"strconv"
	"strings"

	"../models/transaction"
	"../utils"
)

func syncTransactions(date string) transaction.Repo {

	// Read the body and manage errors
	body, err := utils.ExtractDataFrom("transactions", date)
	if err != nil {
		fmt.Println(err)
	}

	separator := "!"
	// Replace the rare chracter by 'separator', because performance
	replacedRare := strings.ReplaceAll(string(body), "\u0000", separator)
	// Split transactions by our 'separator' identifier
	separeTransactions := strings.Split(replacedRare, separator+separator)

	var transactions transaction.Repo = utils.GiveMeRepoTransactionStructure(separeTransactions, separator)

	// Add date to products repo
	var conversionError error
	transactions.Date, conversionError = strconv.Atoi(date)

	if conversionError != nil {
		fmt.Println(conversionError)
	}

	return transactions
}
