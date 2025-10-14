package goods

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/goods/v1"
	bargain_info "shop-goframe-micro-service-refacotor/app/goods/api/bargain_info/v1"
	utime "shop-goframe-micro-service-refacotor/utility/time"
)

func (c *ControllerV1) Bargain_info_Delete(ctx context.Context, req *v1.Bargain_info_DeleteReq) (res *v1.Bargain_info_DeleteRes, err error) { // 调用gRPC服务
	grpcRes, err := c.BargainInfoClient.Delete(ctx, &bargain_info.BargainInfoDeleteReq{
		UserId:  req.User_id,
		GoodsId: req.Goods_id,
		Id:      req.Id,
	})
	if err != nil {
		return nil, err
	}

	// 转换时间格式
	create_time := utime.TimeString(grpcRes.CreatedTime)
	delete_time := utime.TimeString(grpcRes.DeletedTime)

	// 构建完整的响应对象
	res = &v1.Bargain_info_DeleteRes{
		Id:          grpcRes.Id,
		User_id:     grpcRes.UserId,
		Create_time: create_time,
		Delete_time: delete_time,
	}

	return res, nil
}
