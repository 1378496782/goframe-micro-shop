package goods

import (
	"context"
	"fmt"

	utime "shop-goframe-micro-service-refacotor/utility/time"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/goods/v1"
	bargain_info "shop-goframe-micro-service-refacotor/app/goods/api/bargain_info/v1"
)

func (c *ControllerV1) Bargain_info_Get(ctx context.Context, req *v1.Bargain_info_GetReq) (res *v1.Bargain_info_GetRes, err error) {
	//gconv转换结构体
	grpcReq := &bargain_info.BargainInfoGetListReq{}

	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	//调用grpc服务
	grpcRes, err := c.BargainInfoClient.GetList(ctx, grpcReq)

	if err != nil {

		return nil, fmt.Errorf("调用bargain_info getlist微服务错误")
	}

	//转化时间类型与格式
	crtime_string := utime.TimeString(grpcRes.CreatedTime)
	uptime_string := utime.TimeString(grpcRes.UpdatedTime)
	expiretime_string := utime.TimeString(grpcRes.ExpiredTime)

	//响应结构体

	res = &v1.Bargain_info_GetRes{
		Id:          grpcRes.Id,
		User_id:     grpcRes.UserId,
		Goods_id:    grpcRes.GoodsId,
		Counts:      grpcRes.Counts,
		Create_time: crtime_string,
		Update_time: uptime_string,
		Expire_time: expiretime_string,
	}

	return res, nil
}
