API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: grpc
grpc:
	protoc -I=. \
		   -I=./third_party \
		   --go_out=paths=source_relative:. \
		   --go-grpc_out=paths=source_relative:. \
		   $(API_PROTO_FILES)
		   