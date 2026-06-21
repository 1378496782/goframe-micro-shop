// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2025-10-11 23:55:29
// 生成路径: internal/app/shop/service/order_goods_info.go
// 生成人：gfast
// desc:订单物品表
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"

	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
)

type IOrderGoodsInfo interface {
	List(ctx context.Context, req *model.OrderGoodsInfoSearchReq) (res *model.OrderGoodsInfoSearchRes, err error)
	GetById(ctx context.Context, Id int) (res *model.OrderGoodsInfoInfoRes, err error)
	Add(ctx context.Context, req *model.OrderGoodsInfoAddReq) (err error)
	Edit(ctx context.Context, req *model.OrderGoodsInfoEditReq) (err error)
	Delete(ctx context.Context, Id []int) (err error)
	GetByOrderId(ctx context.Context, orderId int) (list []*model.OrderGoodsInfoListRes, err error)
	GetOrderGoodsDetail(ctx context.Context, id int) (res *model.OrderGoodsDetailRes, err error)
	AddOrderGoods(ctx context.Context, req *model.OrderGoodsAddReq) (err error)
}

var localOrderGoodsInfo IOrderGoodsInfo

func OrderGoodsInfo() IOrderGoodsInfo {
	if localOrderGoodsInfo == nil {
		panic("implement not found for interface IOrderGoodsInfo, forgot register?")
	}
	return localOrderGoodsInfo
}

func RegisterOrderGoodsInfo(i IOrderGoodsInfo) {
	localOrderGoodsInfo = i
}
