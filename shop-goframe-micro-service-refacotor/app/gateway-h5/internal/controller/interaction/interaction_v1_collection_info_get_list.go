package interaction

import (
	"context"
	goods_info "shop-goframe-micro-service-refacotor/app/goods/api/goods_info/v1"
	collection "shop-goframe-micro-service-refacotor/app/interaction/api/collection_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/interaction/v1"
)

func (c *ControllerV1) CollectionInfoGetList(ctx context.Context, req *v1.CollectionInfoGetListReq) (res *v1.CollectionInfoGetListRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &collection.CollectionInfoGetListReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 从token获取用户的id
	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		return nil, gerror.NewCode(gcode.CodeValidationFailed, "用户ID类型错误或不存在")
	}
	grpcReq.UserId = userId

	// 调用gRPC服务,获取收藏列表
	collectionInfoGetListRes, err := c.CollectionInfoClient.GetList(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 先整理一部分数据
	list := make([]*v1.UserCollectionInfoItem, len(collectionInfoGetListRes.Data.List))
	for i, v := range collectionInfoGetListRes.Data.List {
		list[i] = &v1.UserCollectionInfoItem{}
		list[i].Id = uint32(v.Id)
		list[i].Type = uint32(v.Type)
		list[i].ObjectId = uint32(v.ObjectId)
		list[i].UserId = uint32(v.UserId)
		list[i].CreatedAt = v.CreatedAt
		list[i].UpdatedAt = v.UpdatedAt
	}

	// 再整理其他数据。分类 1是商品
	if req.Type == 1 {
		for _, v := range list {
			goodsDetailReq := &goods_info.GoodsInfoGetDetailReq{}
			goodsDetailReq.Id = v.ObjectId
			goodsDetailRes, err := c.GoodsClient.GetDetail(ctx, goodsDetailReq)
			if err != nil {
				g.Log().Errorf(ctx, "获取商品详情失败, 商品ID: %d, 错误: %v", v.ObjectId, err)
				continue
			}
			if goodsDetailRes == nil || goodsDetailRes.Data == nil {
				g.Log().Warningf(ctx, "商品详情返回为空, 商品ID: %d", v.ObjectId)
				continue
			}

			v.PicUrl = goodsDetailRes.Data.PicUrl
			v.Name = goodsDetailRes.Data.Name
			v.Price = uint64(goodsDetailRes.Data.Price)
		}
	}
	// 分类 2是文章
	if req.Type == 2 {

	}

	// 组装响应
	res = &v1.CollectionInfoGetListRes{}
	res.List = list
	res.Page = collectionInfoGetListRes.Data.Page
	res.Size = collectionInfoGetListRes.Data.Size
	res.Total = collectionInfoGetListRes.Data.Total

	return res, nil
}
