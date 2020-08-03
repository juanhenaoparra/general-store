package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"../models/ip"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
)

//SyncIPS return a list of handled ips
func SyncIPS(listIPS []string, ips *ip.Repo, dg *dgo.Dgraph) map[string]string {

	for _, v := range listIPS {
		ips.Add(ip.IP{
			UID:     "_:" + v,
			Address: v,
		})
	}

	ipsAssignments := saveIps(ips, dg)

	return ipsAssignments

}

func saveIps(ips *ip.Repo, dg *dgo.Dgraph) map[string]string {
	ctx := context.Background()
	mu := &api.Mutation{
		CommitNow: true,
	}

	fmt.Printf("Length of ips: %v\n", len(ips.IPS))

	pb, err := json.Marshal(ips.IPS)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb

	assignments, err := dg.NewTxn().Mutate(ctx, mu)

	if err != nil {
		log.Fatal(err)
	}

	for i := range ips.IPS {
		ips.IPS[i].UID = assignments.Uids[ips.IPS[i].Address]
	}

	return assignments.Uids

}
