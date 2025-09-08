package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"google.golang.org/grpc"
	"shop-goframe-micro-service-refacotor/app/goods/internal/controller/cart_info"
	"shop-goframe-micro-service-refacotor/app/goods/internal/controller/category_info"
	"shop-goframe-micro-service-refacotor/app/goods/internal/controller/goods_images"
	"shop-goframe-micro-service-refacotor/app/goods/internal/controller/goods_info"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "goods grpc service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
				)}...,
			)
			s := grpcx.Server.New(c)
			goods_info.Register(s)
			goods_images.Register(s)
			category_info.Register(s)
			cart_info.Register(s)
			s.Run()
			return nil
		},
	}
)
