package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	user_coupon "shop-goframe-micro-service-refacotor/app/goods/api/user_coupon_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/goods/v1"
)

func (c *ControllerV1) UserCouponInfoGetList(ctx context.Context, req *v1.UserCouponInfoGetListReq) (res *v1.UserCouponInfoGetListRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &user_coupon.UserCouponInfoGetListReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		// 处理类型不匹配的情况
		panic("用户ID类型错误或不存在")
	}
	grpcReq.UserId = userId
	// 调用gRPC服务
	grpcRes, err := c.UserCouponInfoClient.GetList(ctx, grpcReq)

	if err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.UserCouponInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}

	// 批量转换列表项
	if err := gconv.Structs(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}

	return res, nil
}
