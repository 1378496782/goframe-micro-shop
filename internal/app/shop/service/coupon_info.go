// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2025-09-09 15:10:49
// 生成路径: internal/app/shop/service/coupon_info.go
// 生成人：gfast
// desc:优惠券
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
)

type ICouponInfo interface {
	List(ctx context.Context, req *model.CouponInfoSearchReq) (res *model.CouponInfoSearchRes, err error)
	GetExportData(ctx context.Context, req *model.CouponInfoSearchReq) (listRes []*model.CouponInfoInfoRes, err error)
	Import(ctx context.Context, file *ghttp.UploadFile) (err error)
	GetById(ctx context.Context, Id int) (res *model.CouponInfoInfoRes, err error)
	Add(ctx context.Context, req *model.CouponInfoAddReq) (err error)
	Edit(ctx context.Context, req *model.CouponInfoEditReq) (err error)
	Delete(ctx context.Context, Id []int) (err error)
}

var localCouponInfo ICouponInfo

func CouponInfo() ICouponInfo {
	if localCouponInfo == nil {
		panic("implement not found for interface ICouponInfo, forgot register?")
	}
	return localCouponInfo
}

func RegisterCouponInfo(i ICouponInfo) {
	localCouponInfo = i
}
