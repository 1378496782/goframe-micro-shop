package goods

import (
	"context"
	"fmt"
	bargain_info "shop-goframe-micro-service-refacotor/app/goods/api/bargain_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"
	utime "shop-goframe-micro-service-refacotor/utility/time"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/goods/v1"
)

func (c *ControllerV1) Bargain_info_Create(ctx context.Context, req *v1.Bargain_info_CreateReq) (res *v1.Bargain_info_CreateRes, err error) {

	//转换参数
	grpcReq := &bargain_info.BargainInfoCreateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		return nil, fmt.Errorf("用户ID类型错误或不存在")
	}

	grpcReq.UserID = int32(userId)

	//调用rpc服务
	grpcRes, err := c.BargainInfoClient.Create(ctx, grpcReq)
	if err != nil {

		return nil, fmt.Errorf("调用bargain_info create微服务错误")
	}

	//转化时间类型与格式

	crtime_string := utime.TimeString(grpcRes.CreatedTime)
	exptime_string := utime.TimeString(grpcRes.ExpiredTime)

	//响应结构体
	res = &v1.Bargain_info_CreateRes{
		Id:          grpcRes.Id,
		User_id:     grpcRes.UserId,
		Goods_id:    grpcRes.GoodsId,
		Counts:      grpcRes.Counts,
		Create_time: crtime_string,
		Expire_time: exptime_string,
	}

	return res, nil
}
