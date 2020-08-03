package query

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"../../db"
	"../../models/product"
)

// SearchProductByID returns a Product Object given an id
func SearchProductByID(ctx context.Context, id string) {
	dg := db.NewClient()

	variables := map[string]string{"$id": id}
	q := `query Product($id: string) {
					product (func: eq(id, $id)) {
						uid
						id
						name
						price
					}
				}`

	resp, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	var p product.Product
	err = json.Unmarshal(resp.Json, &p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
}
