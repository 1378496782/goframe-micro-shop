// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfo is the golang structure of table user_info for DAO operations like Where/Data.
type UserInfo struct {
	g.Meta       `orm:"table:user_info, do:true"`
	Id           any         //
	Name         any         // 用户名
	Avatar       any         // 头像
	Password     any         //
	UserSalt     any         // 加密盐 生成密码用
	OpenId       any         // 微信openid
	Phone        any         // 手机号
	Sex          any         // 1男 2女
	Status       any         // 1正常 2拉黑冻结
	Sign         any         // 个性签名
	SecretAnswer any         // 密保问题的答案
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
	DeletedAt    *gtime.Time //
}
