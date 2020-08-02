package handler

import (
	"fmt"
	"strconv"

	"../models/product"
	"../utils"
)

func syncProducts(date string) product.Repo {

	// Read the body and manage errors
	body, err := utils.ExtractDataFrom("products", date)
	if err != nil {
		fmt.Println(err)
	}

	normalizedCsv := utils.GetNormalizeCsv(string(body), "'", ",")

	products := utils.GiveMeRepoProductStructure(normalizedCsv)

	// Add date to products repo
	var conversionError error
	products.Date, conversionError = strconv.Atoi(date)

	if conversionError != nil {
		fmt.Println(conversionError)
	}

	return products

}
