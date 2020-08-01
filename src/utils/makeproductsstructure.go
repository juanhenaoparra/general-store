package utils

import (
	"fmt"
	"strconv"

	"../models/product"
)

// GiveMeProductsStructure given an string matrix iterate over them and returns Repo Product Type
func GiveMeProductsStructure(textArray [][]string) product.Repo {
	var products product.Repo

	for _, p := range textArray {
		conversionPrice, err := strconv.Atoi(p[2])

		if err != nil {
			fmt.Println(err)
		}

		products.Add(product.Product{
			ID:    p[0],
			Name:  p[1],
			Price: conversionPrice,
		})
	}

	return products
}
