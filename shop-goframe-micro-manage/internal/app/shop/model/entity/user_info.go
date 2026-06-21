// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-09-09 15:51:53
// 生成路径: internal/app/shop/model/entity/user_info.go
// 生成人：gfast
// desc:用户
// company:云南奇讯科技有限公司
// ==========================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// UserInfo is the golang structure for table user_info.
type UserInfo struct {
	gmeta.Meta   `orm:"table:user_info"`
	Id           int         `orm:"id,primary" json:"id"`              //
	Name         string      `orm:"name" json:"name"`                  // 用户名
	Avatar       string      `orm:"avatar" json:"avatar"`              // 头像
	Password     string      `orm:"password" json:"password"`          //
	UserSalt     string      `orm:"user_salt" json:"userSalt"`         // 加密盐 生成密码用
	Sex          int         `orm:"sex" json:"sex"`                    // 1男 2女
	Status       int         `orm:"status" json:"status"`              // 1正常 2拉黑冻结
	Sign         string      `orm:"sign" json:"sign"`                  // 个性签名
	SecretAnswer string      `orm:"secret_answer" json:"secretAnswer"` // 密保问题的答案
	CreatedAt    *gtime.Time `orm:"created_at" json:"createdAt"`       //
	UpdatedAt    *gtime.Time `orm:"updated_at" json:"updatedAt"`       //
	DeletedAt    *gtime.Time `orm:"deleted_at" json:"deletedAt"`       //
}
