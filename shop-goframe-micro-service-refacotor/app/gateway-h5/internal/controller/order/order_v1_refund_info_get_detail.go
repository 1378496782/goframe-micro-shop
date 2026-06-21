package order

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	refund_info "shop-goframe-micro-service-refacotor/app/order/api/refund_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
)

func (c *ControllerV1) RefundInfoGetDetail(ctx context.Context, req *v1.RefundInfoGetDetailReq) (res *v1.RefundInfoGetDetailRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &refund_info.RefundInfoGetDetailReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.RefundInfoClient.GetDetail(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 检查数据是否存在
	if grpcRes.Data == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "退款记录不存在")
	}

	// 批量转换列表项
	res = &v1.RefundInfoGetDetailRes{}
	if err := gconv.Struct(grpcRes.Data, res); err != nil {
		return nil, err
	}

	return res, nil
}
