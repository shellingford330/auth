SHELL=/bin/bash -eo pipefail

.PHONY: protoc
protoc:
	protoc --go_out=./presentation/grpc/go/auth --go_opt=paths=source_relative \
		--go-grpc_out=./presentation/grpc/go/auth --go-grpc_opt=paths=source_relative \
		presentation/grpc/protos/auth.proto