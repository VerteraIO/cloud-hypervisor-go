package unixhttp_test

import (
	"context"
	"fmt"
	"log"

	"github.com/VerteraIO/cloud-hypervisor-go/chclient"
	"github.com/VerteraIO/cloud-hypervisor-go/unixhttp"
)

func ExampleNewClient() {
	// Create an HTTP client that communicates over a Unix socket
	httpc := unixhttp.NewClient("/run/cloud-hypervisor.sock")

	// Create the Cloud Hypervisor client
	client, err := chclient.NewClient(
		chclient.WithServer("http://unix/api/v1"),
		chclient.WithHTTPClient(httpc),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the VMM
	resp, err := client.VmmPing(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	if resp.Version != nil {
		fmt.Printf("Cloud Hypervisor version: %s\n", *resp.Version)
	}
}
