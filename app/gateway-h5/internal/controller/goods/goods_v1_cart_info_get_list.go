package goods

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"google.golang.org/protobuf/types/known/timestamppb"
	cart "shop-goframe-micro-service-refacotor/app/goods/api/cart_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/goods/v1"
)

func (c *ControllerV1) CartInfoGetList(ctx context.Context, req *v1.CartInfoGetListReq) (res *v1.CartInfoGetListRes, err error) {
	// 创建 gRPC 请求
	grpcReq := &cart.CartInfoGetListReq{
		Page: req.Page,
		Size: req.Size,
	}

	// 从上下文中获取用户ID
	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		// 处理类型不匹配的情况
		return nil, gerror.New("用户ID类型错误或不存在")
	}
	grpcReq.UserId = userId

	// 调用gRPC服务
	grpcRes, err := c.CartInfoClient.GetList(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.CartInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
		List:  make([]*v1.CartInfoItem, 0, len(grpcRes.Data.List)),
	}

	// 手动转换列表项，因为字段名不完全匹配
	for _, item := range grpcRes.Data.List {
		cartItem := &v1.CartInfoItem{
			// 购物车字段
			Id:     item.Id,
			UserId: item.UserId,
			Count:  item.Count,

			// 商品字段
			GoodsId:     item.GoodsId,
			GoodsName:   item.GoodsName,
			GoodsPicUrl: item.GoodsPicUrl,
			GoodsPrice:  item.GoodsPrice,
			GoodsBrand:  item.GoodsBrand,
			GoodsStock:  item.GoodsStock,
			GoodsSale:   item.GoodsSale,
			GoodsTags:   item.GoodsTags,
			GoodsSort:   item.GoodsSort,
		}

		// 转换时间字段
		if item.GoodsCreatedAt != nil {
			cartItem.GoodsCreatedAt = timestamppb.New(item.GoodsCreatedAt.AsTime())
		}
		if item.GoodsUpdatedAt != nil {
			cartItem.GoodsUpdatedAt = timestamppb.New(item.GoodsUpdatedAt.AsTime())
		}

		res.List = append(res.List, cartItem)
	}

	return res, nil
}
