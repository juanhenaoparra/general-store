package handler

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"

	"../db/query"
	"../models/datemodel"
)

//SyncDate return a list of handled ips
func SyncDate(date *int, dateModel *datemodel.Date, dg *dgo.Dgraph) map[string]string {
	dateAssignment := make(map[string]string)
	ctx := context.Background()

	dateAssignment["uid"] = query.SearchDateByTimestamp(&ctx, strconv.Itoa(*date))

	if dateAssignment["uid"] == "" {
		dateAssignment["uid"] = saveDate(&ctx, dateModel, dg)

		return dateAssignment
	}

	dateAssignment["uid"] = ""

	return dateAssignment
}

func saveDate(ctx *context.Context, dateModel *datemodel.Date, dg *dgo.Dgraph) string {

	mu := &api.Mutation{
		CommitNow: true,
	}

	pb, err := json.Marshal(dateModel)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb

	assignments, err := dg.NewTxn().Mutate(*ctx, mu)

	if err != nil {
		log.Fatal(err)
	}

	return assignments.Uids[strconv.Itoa(*dateModel.Date)]
}
