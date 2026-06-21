package goods_info_test

import (
	"context"
	"fmt"
	"os"
	v1 "shop-goframe-micro-service-refacotor/app/goods/api/goods_info/v1"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/app/goods/internal/logic/goods_info"
	"shop-goframe-micro-service-refacotor/app/goods/internal/model/entity"
	"testing"

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

func TestRestoreStock(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name     string
		setup    func(t *testing.T) (*v1.RestoreStockReq, func(t *testing.T))
		wantErr  bool
		wantCode gcode.Code
	}{
		{
			name: "invalid params empty goods ids",
			setup: func(t *testing.T) (*v1.RestoreStockReq, func(t *testing.T)) {
				return &v1.RestoreStockReq{}, nil
			},
			wantErr:  true,
			wantCode: gcode.CodeInvalidParameter,
		},
		{
			name: "invalid params goods ids and counts length mismatch",
			setup: func(t *testing.T) (*v1.RestoreStockReq, func(t *testing.T)) {
				return &v1.RestoreStockReq{
					GoodsIds: []uint32{1, 2},
					Counts:   []uint32{1},
				}, nil
			},
			wantErr:  true,
			wantCode: gcode.CodeInvalidParameter,
		},
		{
			name: "invalid params count must be greater than zero",
			setup: func(t *testing.T) (*v1.RestoreStockReq, func(t *testing.T)) {
				return &v1.RestoreStockReq{
					GoodsIds: []uint32{1},
					Counts:   []uint32{0},
				}, nil
			},
			wantErr:  true,
			wantCode: gcode.CodeInvalidParameter,
		},
		{
			name: "success and merge repeated goods",
			setup: func(t *testing.T) (*v1.RestoreStockReq, func(t *testing.T)) {
				goodsId1 := insertTestGoods(t, ctx, "restore stock merge goods 1", 10)
				goodsId2 := insertTestGoods(t, ctx, "restore stock merge goods 2", 8)

				return &v1.RestoreStockReq{
						GoodsIds: []uint32{goodsId1, goodsId1, goodsId2},
						Counts:   []uint32{2, 3, 1},
					}, func(t *testing.T) {
						assertGoodsStock(t, ctx, goodsId1, 15)
						assertGoodsStock(t, ctx, goodsId2, 9)
					}
			},
		},
		{
			name: "rollback when one goods does not exist",
			setup: func(t *testing.T) (*v1.RestoreStockReq, func(t *testing.T)) {
				goodsId := insertTestGoods(t, ctx, "restore stock rollback goods", 10)

				return &v1.RestoreStockReq{
						GoodsIds: []uint32{goodsId, 4294967295},
						Counts:   []uint32{3, 5},
					}, func(t *testing.T) {
						assertGoodsStock(t, ctx, goodsId, 10)
					}
			},
			wantErr:  true,
			wantCode: gcode.CodeValidationFailed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, assert := tt.setup(t)
			got, gotErr := goods_info.RestoreStock(ctx, req)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("RestoreStock() failed: %v", gotErr)
					return
				}
				if gerror.Code(gotErr) != tt.wantCode {
					t.Errorf("RestoreStock() error code = %v, want %v", gerror.Code(gotErr), tt.wantCode)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("RestoreStock() succeeded unexpectedly")
			}
			if got == nil {
				t.Fatal("RestoreStock() returned nil response")
			}
			if assert != nil {
				assert(t)
			}
		})
	}
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
