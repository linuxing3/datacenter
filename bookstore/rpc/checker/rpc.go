package main

import (
	"flag"
	"fmt"

	"datacenter/bookstore/rpc/checker/check"
	"datacenter/bookstore/rpc/checker/internal/config"
	"datacenter/bookstore/rpc/checker/internal/server"
	"datacenter/bookstore/rpc/checker/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/check.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewCheckerServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		check.RegisterCheckerServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
