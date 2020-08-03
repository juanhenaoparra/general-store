package query

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"../../db"
	"../../models/device"
)

// SearchDeviceByName returns a Device Object given a name
func SearchDeviceByName(ctx context.Context, name string) {
	dg := db.NewClient()

	variables := map[string]string{"$name": name}
	q := `query Device($name: string) {
					device (func: eq(name, $name)) {
						uid
						name
					}
				}`

	resp, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	var d device.Device
	err = json.Unmarshal(resp.Json, &d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(d)
}
