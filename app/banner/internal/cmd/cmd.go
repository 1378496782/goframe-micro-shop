package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"google.golang.org/grpc"
	"shop-goframe-micro-service-refacotor/app/banner/internal/controller/position_info"
	"shop-goframe-micro-service-refacotor/app/banner/internal/controller/rotation_info"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "banner grpc service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
				)}...,
			)
			s := grpcx.Server.New(c)
			position_info.Register(s)
			rotation_info.Register(s)
			s.Run()
			return nil
		},
	}
)
