// +build wireinject

package main

import (
	"context"
	"fwt/internal/biz"
	"fwt/internal/conf"
	"fwt/internal/data"
	"fwt/internal/server"
	"fwt/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func initApp(*conf.Server, *conf.Data, *conf.Config, log.Logger, context.Context) (*kratos.App, func(), error) {
	panic(wire.Build(service.ProviderSet, data.ProviderSet, biz.ProviderSet, server.ProviderSet, newApp))
}
