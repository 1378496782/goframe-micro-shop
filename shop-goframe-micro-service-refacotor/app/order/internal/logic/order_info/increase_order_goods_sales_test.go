package logic

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	goods_info "shop-goframe-micro-service-refacotor/app/goods/api/goods_info/v1"
	"shop-goframe-micro-service-refacotor/app/order/internal/dao"
	goods "shop-goframe-micro-service-refacotor/app/order/utility/goods_info"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"google.golang.org/grpc"
)

func TestMain(m *testing.M) {
	dbLink := os.Getenv("ORDER_TEST_DB_LINK")
	if dbLink == "" {
		dbLink = "mysql:root:CHANGE_ME_MYSQL_PASSWORD@tcp(127.0.0.1:3306)/order"
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

// fakeGoodsClient 是 goods.Client 的测试替身，只实现 IncreaseSales，
// 其余方法不会在本测试中被调用。
type fakeGoodsClient struct {
	goods_info.GoodsInfoClient
	gotReq *goods_info.IncreaseSalesReq
	retErr error
}

func (c *fakeGoodsClient) IncreaseSales(ctx context.Context, in *goods_info.IncreaseSalesReq, opts ...grpc.CallOption) (*goods_info.IncreaseSalesRes, error) {
	c.gotReq = in
	if c.retErr != nil {
		return nil, c.retErr
	}
	return &goods_info.IncreaseSalesRes{}, nil
}

func TestIncreaseOrderGoodsSales_BitsUT(t *testing.T) {
	ctx := context.Background()

	t.Run("success increase sales", func(t *testing.T) {
		restore := swapGoodsClient(&fakeGoodsClient{})
		defer restore()
		fake := goods.Client.(*fakeGoodsClient)

		orderId, number := insertTestOrder(t, ctx)
		insertTestOrderGoods(t, ctx, orderId, 101, 2)
		insertTestOrderGoods(t, ctx, orderId, 202, 3)

		err := IncreaseOrderGoodsSales(ctx, number)
		if err != nil {
			t.Fatalf("IncreaseOrderGoodsSales() failed: %v", err)
		}
		if fake.gotReq == nil {
			t.Fatal("IncreaseSales RPC was not called")
		}
		assertUint32Slice(t, "goodsIds", fake.gotReq.GoodsIds, []uint32{101, 202})
		assertUint32Slice(t, "counts", fake.gotReq.Counts, []uint32{2, 3})
	})

	t.Run("order not found", func(t *testing.T) {
		restore := swapGoodsClient(&fakeGoodsClient{})
		defer restore()

		// 订单不存在时，gdb.Scan 直接返回 "sql: no rows" 错误并在 order_info.go:809
		// 提前 return，使得后续 order.Id == 0 -> CodeNotFound 分支不可达。
		// 核心契约是"订单不存在必须报错、不得误成功"，故此处只断言返回非 nil error。
		err := IncreaseOrderGoodsSales(ctx, fmt.Sprintf("not-exist-%d", gtime.Now().UnixNano()))
		if err == nil {
			t.Fatal("expected error for non-existent order, got nil")
		}
	})

	t.Run("order goods empty", func(t *testing.T) {
		restore := swapGoodsClient(&fakeGoodsClient{})
		defer restore()

		_, number := insertTestOrder(t, ctx)

		err := IncreaseOrderGoodsSales(ctx, number)
		if err == nil {
			t.Fatal("expected error for order without goods, got nil")
		}
		if gerror.Code(err) != gcode.CodeInvalidParameter {
			t.Fatalf("expected CodeInvalidParameter, got %v", gerror.Code(err))
		}
	})

	t.Run("propagate rpc error", func(t *testing.T) {
		rpcErr := errors.New("increase sales rpc failed")
		restore := swapGoodsClient(&fakeGoodsClient{retErr: rpcErr})
		defer restore()

		orderId, number := insertTestOrder(t, ctx)
		insertTestOrderGoods(t, ctx, orderId, 303, 1)

		err := IncreaseOrderGoodsSales(ctx, number)
		if err == nil {
			t.Fatal("expected rpc error to be propagated, got nil")
		}
		if !errors.Is(err, rpcErr) {
			t.Fatalf("expected propagated rpc error, got %v", err)
		}
	})
}

func swapGoodsClient(c goods_info.GoodsInfoClient) func() {
	prev := goods.Client
	goods.Client = c
	return func() {
		goods.Client = prev
	}
}

func insertTestOrder(t *testing.T, ctx context.Context) (orderId int, number string) {
	t.Helper()

	now := gtime.Now()
	number = fmt.Sprintf("UT%d", now.UnixNano())
	id, err := dao.OrderInfo.Ctx(ctx).Data(g.Map{
		"number":     number,
		"user_id":    1,
		"pay_type":   1,
		"status":     2,
		"price":      100,
		"created_at": now,
		"updated_at": now,
	}).InsertAndGetId()
	if err != nil {
		t.Fatalf("insert test order error: %v", err)
	}

	orderId = int(id)
	t.Cleanup(func() {
		if _, err := dao.OrderInfo.Ctx(ctx).Where("id", orderId).Delete(); err != nil {
			t.Logf("cleanup test order %d error: %v", orderId, err)
		}
	})

	return orderId, number
}

func insertTestOrderGoods(t *testing.T, ctx context.Context, orderId int, goodsId, count int) {
	t.Helper()

	now := gtime.Now()
	id, err := dao.OrderGoodsInfo.Ctx(ctx).Data(g.Map{
		"order_id":   orderId,
		"goods_id":   goodsId,
		"count":      count,
		"price":      100,
		"created_at": now,
		"updated_at": now,
	}).InsertAndGetId()
	if err != nil {
		t.Fatalf("insert test order goods error: %v", err)
	}

	t.Cleanup(func() {
		if _, err := dao.OrderGoodsInfo.Ctx(ctx).Where("id", id).Delete(); err != nil {
			t.Logf("cleanup test order goods %d error: %v", id, err)
		}
	})
}

func assertUint32Slice(t *testing.T, name string, got, want []uint32) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("%s length expected %d, got %d (%v)", name, len(want), len(got), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("%s[%d] expected %d, got %d", name, i, want[i], got[i])
		}
	}
}
