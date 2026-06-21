// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-09-09 15:10:49
// 生成路径: internal/app/shop/model/entity/coupon_info.go
// 生成人：gfast
// desc:优惠券
// company:云南奇讯科技有限公司
// ==========================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// CouponInfo is the golang structure for table coupon_info.
type CouponInfo struct {
	gmeta.Meta `orm:"table:coupon_info"`
	Id         int         `orm:"id,primary" json:"id"`        // ID
	GoodsId    int         `orm:"goods_id" json:"goodsId"`     // 关联商品id（0表示全场通用）
	Name       string      `orm:"name" json:"name"`            // 名称
	Type       int         `orm:"type" json:"type"`            // 类型
	Amount     int         `orm:"amount" json:"amount"`        // 优惠金额（元）
	Deadline   *gtime.Time `orm:"deadline" json:"deadline"`    // 过期时间
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt"` // 创建时间
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt"` // 更新时间
	DeletedAt  *gtime.Time `orm:"deleted_at" json:"deletedAt"` // 删除时间（软删除）
}
