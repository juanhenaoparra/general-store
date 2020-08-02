package handler

import (
	"encoding/json"
	"fmt"

	"../models/buyer"
	"../utils"
)

// Get all buyers from endpoint
func syncBuyers(date int) buyer.Repo {

	// Read the body and manage errors
	body, err := utils.ExtractDataFrom("buyers", date)
	if err != nil {
		fmt.Println(err)
	}

	// Create a var buyers of type Repo and conversionError
	var buyers buyer.Repo
	myDtype := []string{"Buyer"}

	json.Unmarshal(body, &buyers.Buyers) // Parse the bytes readed from the body in a Repo Property

	for i := range buyers.Buyers {
		buyers.Buyers[i].DType = myDtype
	}

	return buyers
}
