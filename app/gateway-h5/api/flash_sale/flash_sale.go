package flash_sale

import (
	"context"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/flash_sale/v1"
)

type IFlashSaleV1 interface {
	FlashSaleGoodsList(ctx context.Context, req *v1.FlashSaleGoodsListReq) (res *v1.FlashSaleGoodsListRes, err error)
	FlashSaleGoodsDetail(ctx context.Context, req *v1.FlashSaleGoodsDetailReq) (res *v1.FlashSaleGoodsDetailRes, err error)
	CreateFlashSaleOrder(ctx context.Context, req *v1.CreateFlashSaleOrderReq) (res *v1.CreateFlashSaleOrderRes, err error)
	GetFlashSaleResult(ctx context.Context, req *v1.GetFlashSaleResultReq) (res *v1.GetFlashSaleResultRes, err error)
}