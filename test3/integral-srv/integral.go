package main

import (
	"flag"
	"fmt"

	"2102Aweek3/integral-srv/internal/config"
	"2102Aweek3/integral-srv/internal/server"
	"2102Aweek3/integral-srv/internal/svc"
	"2102Aweek3/integral-srv/pb/integral"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "2102Aweek3/inits"
)

var configFile = flag.String("f", "etc/integral.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		integral.RegisterIntegralServer(grpcServer, server.NewIntegralServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
