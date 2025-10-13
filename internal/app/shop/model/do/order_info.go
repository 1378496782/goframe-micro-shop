// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-10-10 23:08:01
// 生成路径: internal/app/shop/model/entity/order_info.go
// 生成人：gfast
// desc:订单表
// company:云南奇讯科技有限公司
// ==========================================================================

package do

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// OrderInfo is the golang structure for table order_info.
type OrderInfo struct {
	gmeta.Meta       `orm:"table:order_info, do:true"`
	Id               interface{} `orm:"id,primary" json:"id"`                      //
	Number           interface{} `orm:"number" json:"number"`                      // 订单编号
	UserId           interface{} `orm:"user_id" json:"userId"`                     // 用户id
	PayType          interface{} `orm:"pay_type" json:"payType"`                   // 支付方式 1微信 2支付宝 3云闪付
	Remark           interface{} `orm:"remark" json:"remark"`                      // 备注
	PayAt            *gtime.Time `orm:"pay_at" json:"payAt"`                       // 支付时间
	Status           interface{} `orm:"status" json:"status"`                      // 订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价 6待确认 7已取消
	ConsigneeName    interface{} `orm:"consignee_name" json:"consigneeName"`       // 收货人姓名
	ConsigneePhone   interface{} `orm:"consignee_phone" json:"consigneePhone"`     // 收货人手机号
	ConsigneeAddress interface{} `orm:"consignee_address" json:"consigneeAddress"` // 收货人详细地址
	Price            interface{} `orm:"price" json:"price"`                        // 订单金额 单位分
	CouponPrice      interface{} `orm:"coupon_price" json:"couponPrice"`           // 优惠券金额 单位分
	ActualPrice      interface{} `orm:"actual_price" json:"actualPrice"`           // 实际支付金额 单位分
	CreatedAt        *gtime.Time `orm:"created_at" json:"createdAt"`               //
	UpdatedAt        *gtime.Time `orm:"updated_at" json:"updatedAt"`               //
}
