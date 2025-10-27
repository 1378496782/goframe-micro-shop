// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BargainInfo is the golang structure of table bargain_info for DAO operations like Where/Data.
type BargainInfo struct {
	g.Meta      `orm:"table:bargain_info, do:true"`
	Id          any         //
	UserId      any         //
	GoodsId     any         //
	Counts      any         //
	CreatedTime *gtime.Time //
	UpdatedTime *gtime.Time //
	DeletedTime *gtime.Time //
	ExpiredTime *gtime.Time //
}
