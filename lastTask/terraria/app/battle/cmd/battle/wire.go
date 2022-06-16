//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"terraria/app/battle/internal/biz"
	"terraria/app/battle/internal/conf"
	"terraria/app/battle/internal/data"
	"terraria/app/battle/internal/server"
	"terraria/app/battle/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(data.ProviderSet, server.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
