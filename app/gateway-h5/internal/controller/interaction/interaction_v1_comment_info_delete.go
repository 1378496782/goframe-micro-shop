package interaction

import (
	"context"
	comment "shop-goframe-micro-service-refacotor/app/interaction/api/comment_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/interaction/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) CommentInfoDelete(ctx context.Context, req *v1.CommentInfoDeleteReq) (res *v1.CommentInfoDeleteRes, err error) {
	grpcReq := &comment.CommentInfoDeleteReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "无法获取用户信息，请重新登录")
	}
	grpcReq.UserId = userId
	// 调用gRPC服务
	grpcRes, err := c.CommentInfoClient.Delete(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.CommentInfoDeleteRes{Id: grpcRes.Id}, nil
}
