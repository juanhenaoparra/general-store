package query

import (
	"context"
	"encoding/json"
	"log"

	"../../db"
	"../../models/datemodel"
)

// SearchDateByTimestamp returns a Date Struct given a timestamp
func SearchDateByTimestamp(ctx *context.Context, timestamp string) string {
	dg := db.NewClient()

	variables := map[string]string{"$timestamp": timestamp}
	q := `query Date($timestamp: string) {
					date(func: eq(timestamp, $timestamp), first: 1) {
						uid
					}
				}`

	resp, err := dg.NewTxn().QueryWithVars(*ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	type SingleDate struct {
		Date []datemodel.Date
	}

	var objmap SingleDate

	err = json.Unmarshal(resp.Json, &objmap)
	if err != nil {
		log.Fatal(err)
	}

	if len(objmap.Date) > 0 {
		return objmap.Date[0].UID
	}

	return ""
}
