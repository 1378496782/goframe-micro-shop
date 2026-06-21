// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2025-10-10 23:08:01
// 生成路径: internal/app/shop/service/order_info.go
// 生成人：gfast
// desc:订单表
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"

	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
)

type IOrderInfo interface {
	List(ctx context.Context, req *model.OrderInfoSearchReq) (res *model.OrderInfoSearchRes, err error)
	GetById(ctx context.Context, Id int) (res *model.OrderInfoInfoRes, err error)
	Add(ctx context.Context, req *model.OrderInfoAddReq) (err error)
	Edit(ctx context.Context, req *model.OrderInfoEditReq) (err error)
	Delete(ctx context.Context, Id []int) (err error)
	Ship(ctx context.Context, Id int) (err error)  // 发货
	Refund(ctx context.Context, Id int) (err error) // 退款
	GetOrderProducts(ctx context.Context, orderId int) (list []*model.OrderGoodsInfoListRes, err error) // 获取订单商品列表
}

var localOrderInfo IOrderInfo

func OrderInfo() IOrderInfo {
	if localOrderInfo == nil {
		panic("implement not found for interface IOrderInfo, forgot register?")
	}
	return localOrderInfo
}

func RegisterOrderInfo(i IOrderInfo) {
	localOrderInfo = i
}
