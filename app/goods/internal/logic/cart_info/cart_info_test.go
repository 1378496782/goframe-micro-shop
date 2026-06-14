package cart_info

import (
	"context"
	"testing"

	v1 "shop-goframe-micro-service-refacotor/app/goods/api/cart_info/v1"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func TestGetSelectedItems(t *testing.T) {
	ctx := context.Background()

	res, err := GetSelectedItems(ctx, &v1.CartInfoGetSelectedItemsReq{
		UserId:  6,
		CartIds: []uint32{4, 5, 6, 7, 8},
	})
	if err != nil {
		t.Fatalf("GetSelectedItems error: %v", err)
	}

	t.Logf("items length: %d", len(res.Items))

	for _, item := range res.Items {
		t.Log(item)
		// t.Logf(
		// 	"cart_id=%d user_id=%d goods_id=%d goods_name=%s count=%d price=%d stock=%d",
		// 	item.Id,
		// 	item.UserId,
		// 	item.GoodsId,
		// 	item.GoodsName,
		// 	item.Count,
		// 	item.GoodsPrice,
		// 	item.GoodsStock,
		// )
	}
}
