package query

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"../../db"
	"../../models/ip"
)

// SearchIPByAddress returns a IP Object given an address
func SearchIPByAddress(ctx context.Context, address string) {
	dg := db.NewClient()

	variables := map[string]string{"$address": address}
	q := `query Ip($id: string) {
					ip (func: eq(address, $address)) {
						uid
						address
					}
				}`

	resp, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	var b ip.IP
	err = json.Unmarshal(resp.Json, &b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(b)
}
