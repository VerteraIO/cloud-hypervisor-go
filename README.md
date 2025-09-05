# cloud-hypervisor-go (experimental)

A standalone Go client for the Cloud Hypervisor REST API, generated from the upstream OpenAPI spec.

- Upstream spec: https://github.com/cloud-hypervisor/cloud-hypervisor/tree/master/vmm/src/api/openapi
- Direct raw: https://raw.githubusercontent.com/cloud-hypervisor/cloud-hypervisor/master/vmm/src/api/openapi/cloud-hypervisor.yaml
- License: Apache-2.0 (mirrors upstream)

## Layout

```
clients/cloud-hypervisor-go/
├─ go.mod
├─ Makefile
├─ README.md
├─ oapi-codegen.yaml           # codegen config
├─ gen.go                      # go:generate wiring
├─ spec/
│  └─ cloud-hypervisor.yaml    # vendored OpenAPI spec (fetched)
└─ chclient/                   # generated code (client + types)
   └─ gen_client.go
```

## Prereqs

- Go 1.21+ (tested on 1.24)

## Generate client

```
make fetch-spec   # downloads the latest upstream spec into spec/cloud-hypervisor.yaml
make generate     # runs oapi-codegen v2 against the vendored spec
make build        # builds a tiny smoke test (compiles the package)
```

Alternatively, to run without make:

```
# Fetch spec
curl -fsSL -o spec/cloud-hypervisor.yaml \
  https://raw.githubusercontent.com/cloud-hypervisor/cloud-hypervisor/master/vmm/src/api/openapi/cloud-hypervisor.yaml

# Generate code
go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0 \
  -config oapi-codegen.yaml spec/cloud-hypervisor.yaml
```

## Usage (example)

```go
package main

import (
    "context"
    "log"
    ch "github.com/VerteraIO/cloud-hypervisor-go/chclient"
)

func main() {
    // CH API typically listens on a local unix socket with a HTTP bridge, but
    // here we assume it’s reachable over HTTP at http://localhost/api/v1
    cfg := ch.WithServer("http://localhost/api/v1")

    // Create a client with default http.Client
    c, err := ch.NewClient(cfg)
    if err != nil { log.Fatal(err) }

    // Ping the VMM
    resp, err := c.VmmPing(context.Background())
    if err != nil { log.Fatal(err) }
    log.Printf("CH version: %s", *resp.Version)
}
```

## Notes

- We vendor the spec to ensure reproducible codegen.
- You can pin the upstream revision by replacing the fetch URL with a raw URL including a commit SHA.
- If the upstream spec changes in a breaking way, update oapi-codegen.yaml or regenerate accordingly.
