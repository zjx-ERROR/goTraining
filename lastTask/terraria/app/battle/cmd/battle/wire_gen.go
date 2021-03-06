// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"terraria/app/battle/internal/biz"
	"terraria/app/battle/internal/conf"
	"terraria/app/battle/internal/data"
	"terraria/app/battle/internal/server"
	"terraria/app/battle/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	monsterRepo := data.NewMonsterRepo(dataData, logger)
	monsterUsecase := biz.NewMonsterUsecase(monsterRepo, logger)
	monsterService := service.NewMonsterService(monsterUsecase, logger)
	weaponRepo := data.NewWeaponRepo(dataData, logger)
	weaponUsecase := biz.NewWeaponUsecase(weaponRepo, logger)
	weaponService := service.NewWeaponService(weaponUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, monsterService, weaponService, logger)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
