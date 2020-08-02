package handler

import (
	"fmt"

	"../models/product"
	"../utils"
)

func syncProducts(date int) product.Repo {

	// Read the body and manage errors
	body, err := utils.ExtractDataFrom("products", date)
	if err != nil {
		fmt.Println(err)
	}

	normalizedCsv := utils.GetNormalizeCsv(string(body), "'", ",")

	products := utils.GiveMeRepoProductStructure(normalizedCsv)

	return products

}
