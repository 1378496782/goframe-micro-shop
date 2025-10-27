// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CartInfo is the golang structure of table cart_info for DAO operations like Where/Data.
type CartInfo struct {
	g.Meta    `orm:"table:cart_info, do:true"`
	Id        any         // 购物车表
	UserId    any         //
	GoodsId   any         // 商品id
	Count     any         // 商品数量
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
