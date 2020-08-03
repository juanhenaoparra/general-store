package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"../models/product"
	"../utils"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
)

func syncProducts(date int, dg *dgo.Dgraph) (product.Repo, map[string]string) {

	// Read the body and manage errors
	body, err := utils.ExtractDataFrom("products", date)
	if err != nil {
		fmt.Println(err)
	}

	normalizedCsv := utils.GetNormalizeCsv(string(body), "'", ",")

	products := utils.GiveMeRepoProductStructure(normalizedCsv)
	productsAssignments := saveProducts(products, dg)

	return products, productsAssignments

}

func saveProducts(products product.Repo, dg *dgo.Dgraph) map[string]string {
	ctx := context.Background()
	mu := &api.Mutation{
		CommitNow: true,
	}

	fmt.Printf("Length of products: %v\n", len(products.Products))

	pb, err := json.Marshal(products.Products)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb

	assignments, err := dg.NewTxn().Mutate(ctx, mu)

	if err != nil {
		log.Fatal(err)
	}

	for i := range products.Products {
		products.Products[i].UID = assignments.Uids[products.Products[i].ID]
	}

	return assignments.Uids

}
