package client

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type GrpcClientOptions []grpc.ClientOption

func NewClientOption() GrpcClientOptions {
	return []grpc.ClientOption{grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(recovery.Recovery())}
}
