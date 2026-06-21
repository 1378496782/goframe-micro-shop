package order

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
	goods_info "shop-goframe-micro-service-refacotor/app/goods/api/goods_info/v1"
	order_info "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) OrderInfoGetList(ctx context.Context, req *v1.OrderInfoGetListReq) (res *v1.OrderInfoGetListRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &order_info.OrderInfoGetListReq{}

	grpcGoodsReq := &goods_info.GoodsInfoGetDetailReq{}
	//调用getdetails 并返回goods id与goods name picurl
	//想办法调用
	g.Dump(grpcGoodsReq)

	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 通过 token 获取用户的id
	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		// 处理类型不匹配的情况
		panic("用户ID类型错误或不存在")
	}
	grpcReq.UserId = userId

	// 调用gRPC服务
	grpcRes, err := c.OrderInfoClient.GetList(ctx, grpcReq)
	for i, info := range grpcRes.Data.List {
		g.Dump(i, info)
	}

	if err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.OrderInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}

	// 批量转换列表项
	if err := gconv.Structs(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}

	//读取res中的goodsid并存入数组
	var goodsIDs []uint32
	for _, order := range grpcRes.Data.List {
		for _, goodsItem := range order.GoodsInfo {
			goodsIDs = append(goodsIDs, uint32(goodsItem.GoodsId))
		}
	}

	// 打印收集到的商品ID数组
	g.Dump("收集到的商品ID列表:", goodsIDs)

	// 批量获取商品详情
	goodsMap := make(map[uint32]*goods_info.GoodsInfoGetDetailRes)
	for _, goodsID := range goodsIDs {
		// 调用goods微服务的GetDetail方法
		goodsDetailReq := &goods_info.GoodsInfoGetDetailReq{
			Id: goodsID,
		}
		goodsDetailRes, err := c.GoodsClient.GetDetail(ctx, goodsDetailReq)
		if err != nil {
			g.Log().Error(ctx, "获取商品详情失败:", err)
			continue
		}
		goodsMap[goodsID] = goodsDetailRes
	}

	// // 批量转换列表项
	// if err := gconv.Structs(grpcRes.Data.List, &res.List); err != nil {
	// 	return nil, err
	// }

	// 将商品详情信息填充到订单列表中
	for i, order := range res.List {
		for j, goodsItem := range order.GoodsInfo {
			goodsID := uint32(goodsItem.GoodsId)
			if goodsDetail, ok := goodsMap[goodsID]; ok && goodsDetail.Data != nil {
				res.List[i].GoodsInfo[j].GoodsName = goodsDetail.Data.Name
				res.List[i].GoodsInfo[j].GoodsPrice = goodsDetail.Data.Price
				res.List[i].GoodsInfo[j].PicUrl = goodsDetail.Data.PicUrl
			}
		}
	}
	return res, nil
}
