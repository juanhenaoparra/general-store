package query

import (
	"context"
	"encoding/json"
	"log"

	"../../db"
	"../../models/buyer"
)

// SearchBuyerByID returns a Buyer Object given an id
func SearchBuyerByID(ctx context.Context, id string) string {
	dg := db.NewClient()

	variables := map[string]string{"$id": id}
	q := `query Buyer($id: string) {
					buyer(func: eq(id, $id), first: 1) {
						uid
					}
				}`

	resp, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	type SingleBuyer struct {
		Buyer []buyer.Buyer
	}

	var objmap SingleBuyer

	err = json.Unmarshal(resp.Json, &objmap)
	if err != nil {
		log.Fatal(err)
	}

	if len(objmap.Buyer) > 0 {
		return objmap.Buyer[0].UID
	}

	return ""
}
