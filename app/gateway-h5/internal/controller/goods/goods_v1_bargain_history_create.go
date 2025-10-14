package goods

import (
	"context"
	"fmt"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/goods/v1"
	bargain_history "shop-goframe-micro-service-refacotor/app/goods/api/bargain_history/v1"
	utime "shop-goframe-micro-service-refacotor/utility/time"
)

func (c *ControllerV1) Bargain_history_Create(ctx context.Context, req *v1.Bargain_history_CreateReq) (res *v1.Bargain_history_CreateRes, err error) {
	//转换参数
	grpcReq := &bargain_history.BargainHistoryCreateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		return nil, fmt.Errorf("用户ID类型错误或不存在")
	}

	grpcReq.UserId = int32(userId)

	//调用rpc服务
	grpcRes, err := c.BargainHistoryClient.Create(ctx, grpcReq)
	if err != nil {

		return nil, fmt.Errorf("调用bargain_history create微服务错误")
	}

	//转化时间类型与格式

	crtime_string := utime.TimeString(grpcRes.CreatedTime)

	//响应结构体
	res = &v1.Bargain_history_CreateRes{
		Id:          grpcRes.Id,
		Bargain_id:  grpcRes.BargainId,
		Amount:      grpcRes.Amount,
		User_id:     grpcRes.UserId,
		Create_time: crtime_string,
	}

	return res, nil
}
