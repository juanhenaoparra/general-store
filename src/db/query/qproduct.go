package query

import (
	"context"
	"log"

	"../../db"
)

// SearchProductPriceByBuyerID returns a product list given an buyer id
func SearchProductPriceByBuyerID(ctx *context.Context, id string) string {
	dg := db.NewClient()

	variables := map[string]string{"$id": id}
	q := `query Product($id: string) {
					products (func: eq(id,"$id"), first: 1) 	{
						uid,
						id,
						dgraph.type,
						~by_buyer{
							have_products {
								{price}
							}
						}
					},
				}`

	resp, err := dg.NewTxn().QueryWithVars(*ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	return string(resp.Json)
}

// SearchProductsByTopLow returns a product list given an top and low
func SearchProductsByTopLow(ctx *context.Context, top string, low string) string {
	dg := db.NewClient()

	variables := map[string]string{"$top": top, "$low": low}
	q := `query Product($top: int, $low: int) {
					products (func: type("Product")) @filter(gt(price,$low) AND lt(price, $top)){
						name,
						price,
					},
				}`

	resp, err := dg.NewTxn().QueryWithVars(*ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	return string(resp.Json)
}
