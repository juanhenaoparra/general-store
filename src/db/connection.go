package db

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

// NewClient returns a dgp.Graph pointer
func NewClient() *dgo.Dgraph {
	port := "9080"
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}

// Setup creates the main environment for types and indexes
func Setup(c *dgo.Dgraph) {
	// Install a schema into dgraph.
	err := c.Alter(context.Background(), &api.Operation{
		Schema: `
			id: string @index(exact) @upsert .
			name: string @index(exact) .
			device_type: string @index(exact) @upsert .
			age: int @index(int) .
			address: string @index(exact) @upsert .
			price: int @index(int) .
			timestamp: int @index(int) @upsert .
			time: uid @reverse .
			by_buyer: uid @reverse .
			since_ip: uid @reverse .
			since_device: uid @reverse .
			have_products: [uid] @count @reverse .

			type Date {
				timestamp
			}

			type Buyer {
				id
				name
				age
			}

			type Product {
				id
				name
				price
			}

			type Ip {
				address
			}

			type Device {
				device_type
			}

			type Transaction {
				id
				time
				by_buyer
				since_ip
				since_device
				have_products
			}
		`,
	})

	// Remaining fields for Transaction Type
	// by_buyer
	// since_ip
	// since_device
	// have_products

	if err != nil {
		log.Fatal(err)
	}
}
