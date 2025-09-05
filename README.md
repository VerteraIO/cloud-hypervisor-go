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

## Usage

### HTTP endpoint

```go
package main

import (
    "context"
    "log"
    ch "github.com/VerteraIO/cloud-hypervisor-go/chclient"
)

func main() {
    // CH API reachable over HTTP
    client, err := ch.NewClient(ch.WithServer("http://localhost/api/v1"))
    if err != nil { log.Fatal(err) }

    // Ping the VMM
    resp, err := client.VmmPing(context.Background())
    if err != nil { log.Fatal(err) }
    log.Printf("CH version: %s", *resp.Version)
}
```

### Unix socket (typical local setup)

```go
package main

import (
    "context"
    "log"
    ch "github.com/VerteraIO/cloud-hypervisor-go/chclient"
    "github.com/VerteraIO/cloud-hypervisor-go/unixhttp"
)

func main() {
    // Create HTTP client for Unix socket communication
    httpc := unixhttp.NewClient("/run/cloud-hypervisor.sock")
    
    client, err := ch.NewClient(
        ch.WithServer("http://unix/api/v1"),
        ch.WithHTTPClient(httpc),
    )
    if err != nil { log.Fatal(err) }

    // Ping the VMM
    resp, err := client.VmmPing(context.Background())
    if err != nil { log.Fatal(err) }
    log.Printf("CH version: %s", *resp.Version)
}
```

## Notes

- We vendor the spec to ensure reproducible codegen.
- You can pin the upstream revision by replacing the fetch URL with a raw URL including a commit SHA.
- If the upstream spec changes in a breaking way, update oapi-codegen.yaml or regenerate accordingly.
