// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2025-09-09 15:51:53
// 生成路径: internal/app/shop/service/user_info.go
// 生成人：gfast
// desc:用户
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"

	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
)

type IUserInfo interface {
	List(ctx context.Context, req *model.UserInfoSearchReq) (res *model.UserInfoSearchRes, err error)
	GetExportData(ctx context.Context, req *model.UserInfoSearchReq) (listRes []*model.UserInfoInfoRes, err error)
	GetById(ctx context.Context, Id int) (res *model.UserInfoInfoRes, err error)
	Add(ctx context.Context, req *model.UserInfoAddReq) (err error)
	Edit(ctx context.Context, req *model.UserInfoEditReq) (err error)
	Delete(ctx context.Context, Id []int) (err error)
}

var localUserInfo IUserInfo

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("implement not found for interface IUserInfo, forgot register?")
	}
	return localUserInfo
}

func RegisterUserInfo(i IUserInfo) {
	localUserInfo = i
}
