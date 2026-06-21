package goods_info

import (
	"context"
	"fmt"
	"os"
	"testing"

	v1 "shop-goframe-micro-service-refacotor/app/goods/api/goods_info/v1"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/app/goods/internal/model/entity"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func TestMain(m *testing.M) {
	dbLink := os.Getenv("GOODS_TEST_DB_LINK")
	if dbLink == "" {
		dbLink = "mysql:root:CHANGE_ME_MYSQL_PASSWORD@tcp(127.0.0.1:3306)/goods"
	}

	if err := gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			{
				Type: "mysql",
				Link: dbLink,
			},
		},
	}); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestDeductStockInvalidParams(t *testing.T) {
	ctx := context.Background()
	controller := &Controller{}

	tests := []struct {
		name string
		req  *v1.DeductStockReq
	}{
		{
			name: "empty goods ids",
			req:  &v1.DeductStockReq{},
		},
		{
			name: "goods ids and counts length mismatch",
			req: &v1.DeductStockReq{
				GoodsIds: []uint32{1, 2},
				Counts:   []uint32{1},
			},
		},
		{
			name: "count must be greater than zero",
			req: &v1.DeductStockReq{
				GoodsIds: []uint32{1},
				Counts:   []uint32{0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := controller.DeductStock(ctx, tt.req)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			if gerror.Code(err) != gcode.CodeInvalidParameter {
				t.Fatalf("expected invalid parameter error, got: %v", err)
			}
		})
	}
}

func TestDeductStockSuccessAndMergeRepeatedGoods(t *testing.T) {
	ctx := context.Background()
	controller := &Controller{}

	goodsId1 := insertTestGoods(t, ctx, "deduct stock merge goods 1", 10)
	goodsId2 := insertTestGoods(t, ctx, "deduct stock merge goods 2", 8)

	_, err := controller.DeductStock(ctx, &v1.DeductStockReq{
		GoodsIds: []uint32{goodsId1, goodsId1, goodsId2},
		Counts:   []uint32{2, 3, 1},
	})
	if err != nil {
		t.Fatalf("DeductStock error: %v", err)
	}

	assertGoodsStock(t, ctx, goodsId1, 5)
	assertGoodsStock(t, ctx, goodsId2, 7)
}

func TestDeductStockRollbackWhenStockNotEnough(t *testing.T) {
	ctx := context.Background()
	controller := &Controller{}

	goodsId1 := insertTestGoods(t, ctx, "deduct stock rollback goods 1", 10)
	goodsId2 := insertTestGoods(t, ctx, "deduct stock rollback goods 2", 2)

	_, err := controller.DeductStock(ctx, &v1.DeductStockReq{
		GoodsIds: []uint32{goodsId1, goodsId2},
		Counts:   []uint32{3, 5},
	})
	if err == nil {
		t.Fatal("expected stock not enough error, got nil")
	}
	if gerror.Code(err) != gcode.CodeValidationFailed {
		t.Fatalf("expected validation failed error, got: %v", err)
	}

	assertGoodsStock(t, ctx, goodsId1, 10)
	assertGoodsStock(t, ctx, goodsId2, 2)
}

func insertTestGoods(t *testing.T, ctx context.Context, name string, stock int) uint32 {
	t.Helper()

	now := gtime.Now()
	id, err := dao.GoodsInfo.Ctx(ctx).Data(g.Map{
		"name":               fmt.Sprintf("%s %d", name, now.UnixNano()),
		"pic_url":            "",
		"price":              100,
		"level1_category_id": 1,
		"level2_category_id": 0,
		"level3_category_id": 0,
		"brand":              "test",
		"stock":              stock,
		"sale":               0,
		"tags":               "test",
		"detail_info":        "",
		"sort":               0,
		"enable_bargain":     0,
		"bargain_price":      0,
		"created_at":         now,
		"updated_at":         now,
	}).InsertAndGetId()
	if err != nil {
		t.Fatalf("insert test goods error: %v", err)
	}

	goodsId := uint32(id)
	t.Cleanup(func() {
		if _, err := dao.GoodsInfo.Ctx(ctx).Where("id", goodsId).Delete(); err != nil {
			t.Logf("cleanup test goods %d error: %v", goodsId, err)
		}
	})

	return goodsId
}

func assertGoodsStock(t *testing.T, ctx context.Context, goodsId uint32, expected int) {
	t.Helper()

	var goods *entity.GoodsInfo
	err := dao.GoodsInfo.Ctx(ctx).Where("id", goodsId).Scan(&goods)
	if err != nil {
		t.Fatalf("query goods %d error: %v", goodsId, err)
	}
	if goods == nil {
		t.Fatalf("goods %d not found", goodsId)
	}
	if goods.Stock != expected {
		t.Fatalf("goods %d stock expected %d, got %d", goodsId, expected, goods.Stock)
	}
}
