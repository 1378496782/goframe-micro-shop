// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2025-10-10 23:08:01
// 生成路径: internal/app/shop/dao/internal/order_info.go
// 生成人：gfast
// desc:订单表
// company:云南奇讯科技有限公司
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderInfoDao is the manager for logic model data accessing and custom defined data operations functions management.
type OrderInfoDao struct {
	table   string           // Table is the underlying table name of the DAO.
	group   string           // Group is the database configuration group name of current DAO.
	columns OrderInfoColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// OrderInfoColumns defines and stores column names for table order_info.
type OrderInfoColumns struct {
	Id               string //
	Number           string // 订单编号
	UserId           string // 用户id
	PayType          string // 支付方式 1微信 2支付宝 3云闪付
	Remark           string // 备注
	PayAt            string // 支付时间
	Status           string // 订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价 6待确认 7已取消
	ConsigneeName    string // 收货人姓名
	ConsigneePhone   string // 收货人手机号
	ConsigneeAddress string // 收货人详细地址
	Price            string // 订单金额 单位分
	CouponPrice      string // 优惠券金额 单位分
	ActualPrice      string // 实际支付金额 单位分
	CreatedAt        string //
	UpdatedAt        string //
}

var orderInfoColumns = OrderInfoColumns{
	Id:               "id",
	Number:           "number",
	UserId:           "user_id",
	PayType:          "pay_type",
	Remark:           "remark",
	PayAt:            "pay_at",
	Status:           "status",
	ConsigneeName:    "consignee_name",
	ConsigneePhone:   "consignee_phone",
	ConsigneeAddress: "consignee_address",
	Price:            "price",
	CouponPrice:      "coupon_price",
	ActualPrice:      "actual_price",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewOrderInfoDao creates and returns a new DAO object for table data access.
func NewOrderInfoDao() *OrderInfoDao {
	return &OrderInfoDao{
		group:   "order",
		table:   "order_info",
		columns: orderInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrderInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrderInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrderInfoDao) Columns() OrderInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrderInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrderInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrderInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
