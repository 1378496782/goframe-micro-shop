// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2025-09-09 15:39:41
// 生成路径: internal/app/shop/service/user_coupon_info.go
// 生成人：gfast
// desc:用户优惠券
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
)

type IUserCouponInfo interface {
	List(ctx context.Context, req *model.UserCouponInfoSearchReq) (res *model.UserCouponInfoSearchRes, err error)
	GetExportData(ctx context.Context, req *model.UserCouponInfoSearchReq) (listRes []*model.UserCouponInfoInfoRes, err error)
	Import(ctx context.Context, file *ghttp.UploadFile) (err error)
	GetById(ctx context.Context, Id int) (res *model.UserCouponInfoInfoRes, err error)
	Add(ctx context.Context, req *model.UserCouponInfoAddReq) (err error)
	Edit(ctx context.Context, req *model.UserCouponInfoEditReq) (err error)
	Delete(ctx context.Context, Id []int) (err error)
	LinkedUserCouponInfoDataSearch(ctx context.Context) (res *model.LinkedUserCouponInfoDataSearchRes, err error)
}

var localUserCouponInfo IUserCouponInfo

func UserCouponInfo() IUserCouponInfo {
	if localUserCouponInfo == nil {
		panic("implement not found for interface IUserCouponInfo, forgot register?")
	}
	return localUserCouponInfo
}

func RegisterUserCouponInfo(i IUserCouponInfo) {
	localUserCouponInfo = i
}
