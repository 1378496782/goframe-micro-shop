package goods

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/goods/v1"
	bargain_history "shop-goframe-micro-service-refacotor/app/goods/api/bargain_history/v1"
	utime "shop-goframe-micro-service-refacotor/utility/time"
)

// todo 要不要在h5网关加删除，以及删除的返回参数 待定
func (c *ControllerV1) Bargain_history_Delete(ctx context.Context, req *v1.Bargain_history_DeleteReq) (res *v1.Bargain_history_DeleteRes, err error) {
	grpcRes, err := c.BargainHistoryClient.Delete(ctx, &bargain_history.BargainHistoryDeleteReq{
		BargainId: req.Bargain_id,
		UserId:    req.User_id,
		Id:        req.Id,
	})
	if err != nil {
		return nil, err
	}

	// 转换时间格式
	create_time := utime.TimeString(grpcRes.CreatedTime)
	delete_time := utime.TimeString(grpcRes.DeletedTime)

	// 构建完整的响应对象
	res = &v1.Bargain_history_DeleteRes{
		Id:          grpcRes.Id,
		Bargain_id:  grpcRes.BargainId,
		User_id:     grpcRes.UserId,
		Create_time: create_time,
		Delete_time: delete_time,
	}

	return res, nil
}
