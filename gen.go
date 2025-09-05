//go:build tools

// Package tools wires go:generate to produce the Cloud Hypervisor REST client.
// Run:
//   make fetch-spec
//   make generate
package tools

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config oapi-codegen.yaml spec/cloud-hypervisor.yaml
