package banner

import (
	"context"
	position_info "shop-goframe-micro-service-refacotor/app/banner/api/position_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-admin/api/banner/v1"
)

func (c *ControllerV1) PositionInfoDelete(ctx context.Context, req *v1.PositionInfoDeleteReq) (res *v1.PositionInfoDeleteRes, err error) {
	// 调用gRPC服务
	_, err = c.PositionInfoClient.Delete(ctx, &position_info.PositionInfoDeleteReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &v1.PositionInfoDeleteRes{}, nil
}
