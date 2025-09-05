SHELL := /bin/bash

SPEC_URL := https://raw.githubusercontent.com/cloud-hypervisor/cloud-hypervisor/master/vmm/src/api/openapi/cloud-hypervisor.yaml
SPEC_DIR := spec
SPEC_FILE := $(SPEC_DIR)/cloud-hypervisor.yaml

.PHONY: fetch-spec generate build clean

fetch-spec:
	@mkdir -p $(SPEC_DIR)
	curl -fsSL -o $(SPEC_FILE) $(SPEC_URL)
	@echo "Fetched spec to $(SPEC_FILE)"

generate:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0 \
		-config oapi-codegen.yaml $(SPEC_FILE)
	@echo "Generated client into chclient/"

build:
	go build ./...

clean:
	rm -f chclient/gen_client.go
	@echo "Cleaned generated files"
