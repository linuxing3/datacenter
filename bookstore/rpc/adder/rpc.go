package main

import (
	"flag"
	"fmt"

	"datacenter/bookstore/rpc/adder/add"
	"datacenter/bookstore/rpc/adder/internal/config"
	"datacenter/bookstore/rpc/adder/internal/server"
	"datacenter/bookstore/rpc/adder/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/add.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewAdderServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		add.RegisterAdderServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
