package server

import (
	v1 "terraria/api/battle/v1"
	"terraria/app/battle/internal/conf"
	"terraria/app/battle/internal/service"
	"terraria/pkg"

	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, t *conf.Trace, monster *service.MonsterService, weapon *service.WeaponService, logger log.Logger) *grpc.Server {
	tp, _ := pkg.TracerProvider(t.Endpoint)
	var opts = []grpc.ServerOption{
		grpc.Middleware(

			recovery.Recovery(),
			tracing.Server(tracing.WithTracerProvider(tp)),
			metrics.Server(
				metrics.WithSeconds(prom.NewHistogram(_metricSeconds)),
				metrics.WithRequests(prom.NewCounter(_metricRequests)),
			),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterMonsterServer(srv, monster)
	v1.RegisterWeaponServer(srv, weapon)
	return srv
}
