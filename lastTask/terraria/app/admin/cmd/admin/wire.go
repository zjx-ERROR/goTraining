// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"terraria/app/admin/internal/biz"
	"terraria/app/admin/internal/conf"
	"terraria/app/admin/internal/data"
	"terraria/app/admin/internal/server"
	"terraria/app/admin/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
