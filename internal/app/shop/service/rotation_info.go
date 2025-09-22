// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2025-09-22 16:50:04
// 生成路径: internal/app/shop/service/rotation_info.go
// 生成人：gfast
// desc:轮播图
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"

	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
)

type IRotationInfo interface {
	List(ctx context.Context, req *model.RotationInfoSearchReq) (res *model.RotationInfoSearchRes, err error)
	GetById(ctx context.Context, Id int) (res *model.RotationInfoInfoRes, err error)
	Add(ctx context.Context, req *model.RotationInfoAddReq) (err error)
	Edit(ctx context.Context, req *model.RotationInfoEditReq) (err error)
	Delete(ctx context.Context, Id []int) (err error)
}

var localRotationInfo IRotationInfo

func RotationInfo() IRotationInfo {
	if localRotationInfo == nil {
		panic("implement not found for interface IRotationInfo, forgot register?")
	}
	return localRotationInfo
}

func RegisterRotationInfo(i IRotationInfo) {
	localRotationInfo = i
}
