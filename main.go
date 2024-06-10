package main

import (
	din "github.com/DIN-center/din-sc/apps/din-go/lib/din"
	"log"
)

func main() {
	c, err := din.NewDinClient()
	if err != nil {
		log.Fatalf("Could not instantiate a DIN client")
	}

	var defaultEndpointCollection = "ethereum://mainnet"

	err = c.PrintListAllMethodsByEndpoint(defaultEndpointCollection)
	if err != nil {
		log.Fatalf("Error printing all methods by endpoint: %v", err)
	}

}
