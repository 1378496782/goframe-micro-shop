package search

import (
	"context"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/search/v1"
)

type ISearchV1 interface {
	SearchGoods(ctx context.Context, req *v1.SearchGoodsReq) (res *v1.SearchGoodsRes, err error)
}
