CONFIG_PROTO_FILES=$(shell find internal/conf -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: grpc
grpc:
	protoc -I=. \
		   -I=./third_party \
		   --go_out=paths=source_relative:. \
		   --go-grpc_out=paths=source_relative:. \
		   $(API_PROTO_FILES)

.PHONY: config
config:
	protoc -I=. \
		   -I=./third_party \
		   --go_out=paths=source_relative:. \
		   $(CONFIG_PROTO_FILES)

.PHONY: wire
wire:
	cd cmd/shop && wire