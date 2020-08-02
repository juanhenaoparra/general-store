package handler

import (
	"fmt"
	"strings"

	"../models/transaction"
	"../utils"
)

func syncTransactions(date *int) transaction.Repo {

	// Read the body and manage errors
	body, err := utils.ExtractDataFrom("transactions", *date)
	if err != nil {
		fmt.Println(err)
	}

	separator := "!"
	// Replace the rare chracter by 'separator', because performance
	replacedRare := strings.ReplaceAll(string(body), "\u0000", separator)
	// Split transactions by our 'separator' identifier
	separeTransactions := strings.Split(replacedRare, separator+separator)

	var transactions transaction.Repo = utils.GiveMeRepoTransactionStructure(separeTransactions, separator, date)

	return transactions
}
