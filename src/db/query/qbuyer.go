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

// GetBuyersPaginated retrieve users paginated
func GetBuyersPaginated(ctx *context.Context, first string, offset string) buyer.Repo {
	dg := db.NewClient()

	variables := map[string]string{"$first": first, "$offset": offset}
	q := `query Buyers($first: int, $offset: int) {
					buyers (func: type("Buyer"), first: $first, offset: $offset) 	{
						id,
						name,
						age,
					}
				}`

	resp, err := dg.NewTxn().QueryWithVars(*ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	var buyersRepo buyer.Repo

	err = json.Unmarshal(resp.Json, &buyersRepo)
	if err != nil {
		log.Fatal(err)
	}

	return buyersRepo
}

// GetBuyerProfile retrieve user profile by id param
func GetBuyerProfile(ctx *context.Context, id string, first string, offset string) string {
	dg := db.NewClient()

	variables := map[string]string{"$id": id, "$first": first, "$offset": offset}
	q := `query Buyers($id: string, $first: int, $offset: int) {
					buyers (func: eq(id, $id), first: 1) 	{
						id,
						name,
						age,
						~by_buyer (first: $first, offset: $offset) {
							id,
							price,
							time {
								timestamp
							}
						}
					}
				}`

	resp, err := dg.NewTxn().QueryWithVars(*ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	return string(resp.Json)
}
