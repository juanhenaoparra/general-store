package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	// Using DGraph
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"

	"../models/buyer"
	"../utils"
)

// Get all buyers from endpoint
func syncBuyers(date int, dg *dgo.Dgraph) (buyer.Repo, map[string]string) {

	// Read the body and manage errors
	body, err := utils.ExtractDataFrom("buyers", date)
	if err != nil {
		fmt.Println(err)
	}

	// Create a var buyers of type Repo and conversionError
	var buyers buyer.Repo

	json.Unmarshal(body, &buyers.Buyers) // Parse the bytes readed from the body in a Repo Property

	buyers, assignments := saveBuyers(buyers, dg)

	return buyers, assignments
}

func saveBuyers(buyers buyer.Repo, dg *dgo.Dgraph) (buyer.Repo, map[string]string) {
	ctx := context.Background()
	mu := &api.Mutation{
		CommitNow: true,
	}

	myDtype := []string{"Buyer"}
	for i := range buyers.Buyers {
		buyers.Buyers[i].UID = "_:" + buyers.Buyers[i].ID
		buyers.Buyers[i].DType = myDtype
	}

	fmt.Printf("Length of buyers: %v\n", len(buyers.Buyers))

	pb, err := json.Marshal(buyers.Buyers)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb

	assignments, err := dg.NewTxn().Mutate(ctx, mu)

	if err != nil {
		log.Fatal(err)
	}

	for i := range buyers.Buyers {
		buyers.Buyers[i].UID = assignments.Uids[buyers.Buyers[i].ID]
	}

	return buyers, assignments.Uids

}
