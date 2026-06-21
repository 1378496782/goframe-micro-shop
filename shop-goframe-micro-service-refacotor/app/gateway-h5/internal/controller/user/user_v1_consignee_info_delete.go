package user

import (
	"context"
	consignee "shop-goframe-micro-service-refacotor/app/user/api/consignee_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/user/v1"
)

func (c *ControllerV1) ConsigneeInfoDelete(ctx context.Context, req *v1.ConsigneeInfoDeleteReq) (res *v1.ConsigneeInfoDeleteRes, err error) {
	// 调用gRPC服务
	_, err = c.ConsigneeInfoClient.Delete(ctx, &consignee.ConsigneeInfoDeleteReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &v1.ConsigneeInfoDeleteRes{}, nil
}
