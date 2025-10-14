// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package goods

import (
	"context"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/goods/v1"
)

type IGoodsV1 interface {
	Bargain_history_Create(ctx context.Context, req *v1.Bargain_history_CreateReq) (res *v1.Bargain_history_CreateRes, err error)
	Bargain_history_Get(ctx context.Context, req *v1.Bargain_history_GetReq) (res *v1.Bargain_history_GetRes, err error)
	Bargain_history_Delete(ctx context.Context, req *v1.Bargain_history_DeleteReq) (res *v1.Bargain_history_DeleteRes, err error)
	Bargain_info_Create(ctx context.Context, req *v1.Bargain_info_CreateReq) (res *v1.Bargain_info_CreateRes, err error)
	Bargain_info_Get(ctx context.Context, req *v1.Bargain_info_GetReq) (res *v1.Bargain_info_GetRes, err error)
	Bargain_info_Delete(ctx context.Context, req *v1.Bargain_info_DeleteReq) (res *v1.Bargain_info_DeleteRes, err error)
	CartInfoGetList(ctx context.Context, req *v1.CartInfoGetListReq) (res *v1.CartInfoGetListRes, err error)
	CartInfoCreate(ctx context.Context, req *v1.CartInfoCreateReq) (res *v1.CartInfoCreateRes, err error)
	CartInfoDelete(ctx context.Context, req *v1.CartInfoDeleteReq) (res *v1.CartInfoDeleteRes, err error)
	CategoryInfoGetList(ctx context.Context, req *v1.CategoryInfoGetListReq) (res *v1.CategoryInfoGetListRes, err error)
	CategoryInfoGetAll(ctx context.Context, req *v1.CategoryInfoGetAllReq) (res *v1.CategoryInfoGetAllRes, err error)
	GoodsImagesGetList(ctx context.Context, req *v1.GoodsImagesGetListReq) (res *v1.GoodsImagesGetListRes, err error)
	GoodsInfoGetDetail(ctx context.Context, req *v1.GoodsInfoGetDetailReq) (res *v1.GoodsInfoGetDetailRes, err error)
	GoodsInfoGetList(ctx context.Context, req *v1.GoodsInfoGetListReq) (res *v1.GoodsInfoGetListRes, err error)
	RecommendGoodsInfoGetList(ctx context.Context, req *v1.RecommendGoodsInfoGetListReq) (res *v1.RecommendGoodsInfoGetListRes, err error)
	UserCouponInfoGetList(ctx context.Context, req *v1.UserCouponInfoGetListReq) (res *v1.UserCouponInfoGetListRes, err error)
}
