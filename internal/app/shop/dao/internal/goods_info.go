// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2025-09-08 11:37:29
// 生成路径: internal/app/shop/dao/internal/goods_info.go
// 生成人：王中阳
// desc:商品
// company:云南奇讯科技有限公司
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GoodsInfoDao is the manager for logic model data accessing and custom defined data operations functions management.
type GoodsInfoDao struct {
	table   string           // Table is the underlying table name of the DAO.
	group   string           // Group is the database configuration group name of current DAO.
	columns GoodsInfoColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// GoodsInfoColumns defines and stores column names for table goods_info.
type GoodsInfoColumns struct {
	Id               string // ID
	Name             string // 名称
	Images           string // 多图
	PicUrl           string // 封面图
	Price            string // 价格(分)
	Level1CategoryId string // 一级分类
	Level2CategoryId string // 二级分类
	Level3CategoryId string // 三级分类
	Brand            string // 品牌
	Stock            string // 库存
	Sale             string // 销量
	Tags             string // 标签
	DetailInfo       string // 商品详情
	CreatedAt        string //
	Sort             string // 排序 倒叙
	UpdatedAt        string //
	DeletedAt        string //
}

var goodsInfoColumns = GoodsInfoColumns{
	Id:               "id",
	Name:             "name",
	Images:           "images",
	PicUrl:           "pic_url",
	Price:            "price",
	Level1CategoryId: "level1_category_id",
	Level2CategoryId: "level2_category_id",
	Level3CategoryId: "level3_category_id",
	Brand:            "brand",
	Stock:            "stock",
	Sale:             "sale",
	Tags:             "tags",
	DetailInfo:       "detail_info",
	CreatedAt:        "created_at",
	Sort:             "sort",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
}

// NewGoodsInfoDao creates and returns a new DAO object for table data access.
func NewGoodsInfoDao() *GoodsInfoDao {
	return &GoodsInfoDao{
		group:   "goods",
		table:   "goods_info",
		columns: goodsInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GoodsInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GoodsInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GoodsInfoDao) Columns() GoodsInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GoodsInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GoodsInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GoodsInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
