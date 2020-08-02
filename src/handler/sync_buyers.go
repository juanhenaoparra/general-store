package handler

import (
	"encoding/json"
	"fmt"
	"strconv"

	"../models/buyer"
	"../utils"
)

// Get all buyers from endpoint
func syncBuyers(date string) buyer.Repo {

	// Read the body and manage errors
	body, err := utils.ExtractDataFrom("buyers", date)
	if err != nil {
		fmt.Println(err)
	}

	// Create a var buyers of type Repo and conversionError
	var buyers buyer.Repo
	var conversionError error

	buyers.Date, conversionError = strconv.Atoi(date)

	if conversionError != nil {
		fmt.Println(conversionError)
	}

	json.Unmarshal(body, &buyers.Buyers) // Parse the bytes readed from the body in a Repo Property

	// This makes a for loop for every buyer and set the date
	// buyers.SetDateAll()

	return buyers
}
