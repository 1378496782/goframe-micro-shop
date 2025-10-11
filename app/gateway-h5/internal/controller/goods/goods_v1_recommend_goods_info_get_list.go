package goods

import (
	"context"

	recommend_goods_info "shop-goframe-micro-service-refacotor/app/goods/api/recommend_goods_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/goods/v1"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) RecommendGoodsInfoGetList(ctx context.Context, req *v1.RecommendGoodsInfoGetListReq) (res *v1.RecommendGoodsInfoGetListRes, err error) {

	grpcReq := &recommend_goods_info.RecommendGoodsInfoGetListReq{}
	err = gconv.Struct(req, grpcReq)
	if err != nil {
		return nil, err
	}
	grpcRes, err := c.RecommendGoodsInfoClient.GetList(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.RecommendGoodsInfoGetListRes{
		Total: grpcRes.Data.Total,
	}
	err = gconv.Struct(grpcRes.Data.List, &res.List)
	if err != nil {
		return nil, err
	}

	return res, nil
}
