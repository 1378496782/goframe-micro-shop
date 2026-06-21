// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BargainHistory is the golang structure of table bargain_history for DAO operations like Where/Data.
type BargainHistory struct {
	g.Meta      `orm:"table:bargain_history, do:true"`
	Id          any         //
	BargainId   any         //
	Amount      any         //
	UserId      any         //
	CreatedTime *gtime.Time //
	DeletedTime *gtime.Time //
}
