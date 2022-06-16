// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"terraria/app/job/internal/biz"
	"terraria/app/job/internal/conf"
	"terraria/app/job/internal/data"
	"terraria/app/job/internal/server"
	"terraria/app/job/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
