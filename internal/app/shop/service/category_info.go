// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2025-09-05 11:52:54
// 生成路径: internal/app/shop/service/category_info.go
// 生成人：王中阳
// desc:商品分类
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
)

type ICategoryInfo interface {
	List(ctx context.Context, req *model.CategoryInfoSearchReq) (res *model.CategoryInfoSearchRes, err error)
	GetExportData(ctx context.Context, req *model.CategoryInfoSearchReq) (listRes []*model.CategoryInfoInfoRes, err error)
	Import(ctx context.Context, file *ghttp.UploadFile) (err error)
	GetById(ctx context.Context, Id int) (res *model.CategoryInfoInfoRes, err error)
	Add(ctx context.Context, req *model.CategoryInfoAddReq) (err error)
	Edit(ctx context.Context, req *model.CategoryInfoEditReq) (err error)
	Delete(ctx context.Context, Id []int) (err error)
	LinkedCategoryInfoDataSearch(ctx context.Context) (res *model.LinkedCategoryInfoDataSearchRes, err error)
}

var localCategoryInfo ICategoryInfo

func CategoryInfo() ICategoryInfo {
	if localCategoryInfo == nil {
		panic("implement not found for interface ICategoryInfo, forgot register?")
	}
	return localCategoryInfo
}

func RegisterCategoryInfo(i ICategoryInfo) {
	localCategoryInfo = i
}
