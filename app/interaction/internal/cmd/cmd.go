package cmd

import (
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/controller/collection_info"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/controller/comment_info"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/controller/praise_info"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/job"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "interaction grpc service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
				)}...,
			)
			s := grpcx.Server.New(c)
			praise_info.Register(s)
			comment_info.Register(s)
			collection_info.Register(s)
			job.StartCommentLikeCountCalibrateJob(ctx)
			s.Run()
			return nil
		},
	}
)
