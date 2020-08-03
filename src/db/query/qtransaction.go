package query

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"../../db"
	"../../models/transaction"
)

// SearchTransactionByID returns a Transaction Object given an id
func SearchTransactionByID(ctx context.Context, id string) {
	dg := db.NewClient()

	variables := map[string]string{"$id": id}
	q := `query Transaction($id: string) {
					transaction (func: eq(id, $id)) {
						uid
						id
						date
						by_buyer {
							uid
							id
						}
						since_ip {
							uid
							address
						}
						since_device {
							uid
							name
						}
						have_products {
							uid
						}
					}
				}`

	resp, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	var t transaction.Transaction
	err = json.Unmarshal(resp.Json, &t)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t)
}
