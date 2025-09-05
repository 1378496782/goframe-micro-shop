// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2025-09-05 12:04:34
// 生成路径: internal/app/shop/service/goods_info.go
// 生成人：王中阳
// desc:商品表
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
)

type IGoodsInfo interface {
	List(ctx context.Context, req *model.GoodsInfoSearchReq) (res *model.GoodsInfoSearchRes, err error)
	GetExportData(ctx context.Context, req *model.GoodsInfoSearchReq) (listRes []*model.GoodsInfoInfoRes, err error)
	Import(ctx context.Context, file *ghttp.UploadFile) (err error)
	GetById(ctx context.Context, Id uint) (res *model.GoodsInfoInfoRes, err error)
	Add(ctx context.Context, req *model.GoodsInfoAddReq) (err error)
	Edit(ctx context.Context, req *model.GoodsInfoEditReq) (err error)
	Delete(ctx context.Context, Id []uint) (err error)
	LinkedGoodsInfoDataSearch(ctx context.Context) (res *model.LinkedGoodsInfoDataSearchRes, err error)
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
