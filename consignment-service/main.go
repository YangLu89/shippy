// shippy-service-consignment/main.go
package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/YangLu89/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/YangLu89/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	port        = ":50051"
	defaultHost = "datastore:27017"
)

func main() {

	// Database host from the environment variables
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	// Mgo creates a 'master' session, we need to end that session
	// before the main function closes.
	defer session.Close()

	if err != nil {

		// We're wrapping the error returned from our CreateSession
		// here to add some context to the error.
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		micro.Name("consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("vessel", srv.Client())

	// Init will parse the command line flags.
	srv.Init()

	// Register handlers
	h := &service{session, vesselClient}
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
