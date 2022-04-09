package main

import (
	"context"
	"fwt/internal/conf"
	"os"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var (
	version = "1.0"

	Bc       conf.Bootstrap
	flagconf string = "../../configs/config.yaml"
)

func newApp(logger log.Logger, gs *grpc.Server, config *conf.Config, ctx context.Context) *kratos.App {
	// nacos注册器
	r := NacosNamingClient(config.NacosIp, config.NacosPort)
	return kratos.New(
		kratos.ID(strconv.Itoa(int(time.Now().Unix()))),
		kratos.Name("shop"),
		kratos.Version(version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		kratos.Registrar(r),
		kratos.Context(ctx),
	)
}

func NacosNamingClient(ip string, port uint64) *nacos.Registry {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(ip, port),
	}
	cc := constant.ClientConfig{
		LogLevel: "error",
	}
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ServerConfigs: sc,
			ClientConfig:  &cc,
		},
	)
	if err != nil {
		panic(err)
	}
	return nacos.New(client)
}

func main() {
	c := config.New(
		config.WithSource(file.NewSource(flagconf)),
	)
	defer c.Close()
	if err := c.Load(); err != nil {
		panic(err)
	}
	if err := c.Scan(&Bc); err != nil {
		panic(err)
	}
	ctx := context.Background()
	logger := log.NewStdLogger(os.Stdout)
	app, cleanup, err := initApp(Bc.Server, Bc.Data, Bc.Config, logger, ctx)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
