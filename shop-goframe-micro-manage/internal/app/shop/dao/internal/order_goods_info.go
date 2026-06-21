// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2025-10-11 23:55:29
// 生成路径: internal/app/shop/dao/internal/order_goods_info.go
// 生成人：gfast
// desc:订单物品表
// company:云南奇讯科技有限公司
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderGoodsInfoDao is the manager for logic model data accessing and custom defined data operations functions management.
type OrderGoodsInfoDao struct {
	table   string                // Table is the underlying table name of the DAO.
	group   string                // Group is the database configuration group name of current DAO.
	columns OrderGoodsInfoColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// OrderGoodsInfoColumns defines and stores column names for table order_goods_info.
type OrderGoodsInfoColumns struct {
	Id             string // 商品维度的订单表
	OrderId        string // 关联的主订单表
	GoodsId        string // 商品id
	GoodsOptionsId string // 商品规格id sku id
	Count          string // 商品数量
	Remark         string // 备注
	Price          string // 订单金额 单位分
	CouponPrice    string // 优惠券金额 单位分
	ActualPrice    string // 实际支付金额 单位分
	CreatedAt      string //
	UpdatedAt      string //
}

var orderGoodsInfoColumns = OrderGoodsInfoColumns{
	Id:             "id",
	OrderId:        "order_id",
	GoodsId:        "goods_id",
	GoodsOptionsId: "goods_options_id",
	Count:          "count",
	Remark:         "remark",
	Price:          "price",
	CouponPrice:    "coupon_price",
	ActualPrice:    "actual_price",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewOrderGoodsInfoDao creates and returns a new DAO object for table data access.
func NewOrderGoodsInfoDao() *OrderGoodsInfoDao {
	return &OrderGoodsInfoDao{
		group:   "order",
		table:   "order_goods_info",
		columns: orderGoodsInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrderGoodsInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrderGoodsInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrderGoodsInfoDao) Columns() OrderGoodsInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrderGoodsInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrderGoodsInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrderGoodsInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
