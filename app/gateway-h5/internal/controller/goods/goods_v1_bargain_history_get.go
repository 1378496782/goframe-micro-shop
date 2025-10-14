package goods

import (
	"context"
	"fmt"
	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/goods/v1"
	bargain_history "shop-goframe-micro-service-refacotor/app/goods/api/bargain_history/v1"
	utime "shop-goframe-micro-service-refacotor/utility/time"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) Bargain_history_Get(ctx context.Context, req *v1.Bargain_history_GetReq) (res *v1.Bargain_history_GetRes, err error) {
	//gconv转换结构体
	grpcReq := &bargain_history.BargainHistoryGetListReq{}

	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	//调用grpc服务
	grpcRes, err := c.BargainHistoryClient.GetList(ctx, grpcReq)

	if err != nil {

		return nil, fmt.Errorf("调用bargain_history getlist微服务错误")
	}
	//转化时间类型与格式
	crtime_string := utime.TimeString(grpcRes.CreatedTime)

	res = &v1.Bargain_history_GetRes{
		Id:          grpcRes.Id,
		Bargain_id:  grpcRes.BargainId,
		Amount:      grpcRes.Amount,
		User_id:     grpcRes.UserId,
		Create_time: crtime_string,
	}

	return res, nil

}
