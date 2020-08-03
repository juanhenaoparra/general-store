package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"

	"../models/device"
)

//SyncDevices return a list of handled ips
func SyncDevices(listDevices []string, devices *device.Repo, dg *dgo.Dgraph) map[string]string {

	for _, v := range listDevices {
		devices.Add(device.Device{
			UID:  "_:" + v,
			Name: v,
		})
	}

	devicesAssignments := saveDevices(devices, dg)

	return devicesAssignments
}

func saveDevices(devices *device.Repo, dg *dgo.Dgraph) map[string]string {
	ctx := context.Background()
	mu := &api.Mutation{
		CommitNow: true,
	}

	fmt.Printf("Length of devices: %v\n", len(devices.Devices))

	pb, err := json.Marshal(devices.Devices)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb

	assignments, err := dg.NewTxn().Mutate(ctx, mu)

	if err != nil {
		log.Fatal(err)
	}

	for i := range devices.Devices {
		devices.Devices[i].UID = assignments.Uids[devices.Devices[i].Name]
	}

	return assignments.Uids
}
