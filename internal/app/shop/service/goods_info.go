// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2025-09-22 16:48:52
// 生成路径: internal/app/shop/service/goods_info.go
// 生成人：gfast
// desc:商品
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"

	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
)

type IGoodsInfo interface {
	List(ctx context.Context, req *model.GoodsInfoSearchReq) (res *model.GoodsInfoSearchRes, err error)
	GetById(ctx context.Context, Id uint) (res *model.GoodsInfoInfoRes, err error)
	Add(ctx context.Context, req *model.GoodsInfoAddReq) (err error)
	Edit(ctx context.Context, req *model.GoodsInfoEditReq) (err error)
	Delete(ctx context.Context, Id []uint) (err error)
}

var localGoodsInfo IGoodsInfo

func GoodsInfo() IGoodsInfo {
	if localGoodsInfo == nil {
		panic("implement not found for interface IGoodsInfo, forgot register?")
	}
	return localGoodsInfo
}

func RegisterGoodsInfo(i IGoodsInfo) {
	localGoodsInfo = i
}
