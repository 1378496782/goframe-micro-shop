package goods

import (
	"context"
	add_goods_info "shop-goframe-micro-service-refacotor/app/goods/api/add_goods_info/v1"

	"github.com/gogf/gf/v2/util/gconv"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/goods/v1"
)

func (c *ControllerV1) AddGoodsInfoGetList(ctx context.Context, req *v1.AddGoodsInfoGetListReq) (res *v1.AddGoodsInfoGetListRes, err error) {
	grpcReq := &add_goods_info.AddGoodsInfoGetListReq{}
	err = gconv.Struct(req, grpcReq)
	if err != nil {
		return nil, err
	}
	grpcRes, err := c.AddGoodsInfoClient.GetList(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.AddGoodsInfoGetListRes{
		Total: grpcRes.Data.Total,
	}
	err = gconv.Struct(grpcRes.Data.List, &res.List)
	if err != nil {
		return nil, err
	}

	return res, nil
}
