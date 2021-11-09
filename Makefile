SHELL=/bin/bash -eo pipefail

PROTO_DIR := ./pkg/grpc/protos
PROTO_GEN_DIR := ./pkg/grpc/go

.PHONY: serve
serve: 
	@go run ./cmd

.PHONY: protoc
protoc:
	protoc -I=$(PROTO_DIR)/auth \
		--go_out=$(PROTO_GEN_DIR)/auth --go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_GEN_DIR)/auth --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/auth/auth.proto

.PHONY: gen
gen:
	go generate -x ./...
