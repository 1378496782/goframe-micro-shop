// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderInfo is the golang structure of table order_info for DAO operations like Where/Data.
type OrderInfo struct {
	g.Meta           `orm:"table:order_info, do:true"`
	Id               any         //
	Number           any         // 订单编号
	UserId           any         // 用户id
	PayType          any         // 支付方式 1微信 2支付宝 3云闪付
	Remark           any         // 备注
	PayAt            *gtime.Time // 支付时间
	Status           any         // 订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价 6待确认 7已取消 8退款
	TransactionId    any         // 第三方支付交易号
	ConsigneeName    any         // 收货人姓名
	ConsigneePhone   any         // 收货人手机号
	ConsigneeAddress any         // 收货人详细地址
	Price            any         // 订单金额 单位分
	CouponPrice      any         // 优惠券金额 单位分
	ActualPrice      any         // 实际支付金额 单位分
	SalesStatus      any         // 销量同步状态：0未同步 1已同步 2同步失败
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
}
