package interaction

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	comment "shop-goframe-micro-service-refacotor/app/interaction/api/comment_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/interaction/v1"
)

func (c *ControllerV1) CommentInfoCreate(ctx context.Context, req *v1.CommentInfoCreateReq) (res *v1.CommentInfoCreateRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &comment.CommentInfoCreateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	// 调用gRPC服务
	grpcRes, err := c.CommentInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.CommentInfoCreateRes{Id: grpcRes.Id}, nil
}
