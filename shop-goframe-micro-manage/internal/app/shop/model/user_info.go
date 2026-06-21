// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2025-09-09 15:51:53
// 生成路径: internal/app/shop/model/user_info.go
// 生成人：gfast
// desc:用户
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
	comModel "github.com/tiger1103/gfast/v3/internal/app/common/model"
)

// UserInfoInfoRes is the golang structure for table user_info.
type UserInfoInfoRes struct {
	gmeta.Meta   `orm:"table:user_info"`
	Id           int         `orm:"id,primary" json:"id" dc:""`                     //
	Name         string      `orm:"name" json:"name" dc:"用户名"`                      // 用户名
	Avatar       string      `orm:"avatar" json:"avatar" dc:"头像"`                   // 头像
	Password     string      `orm:"password" json:"password" dc:""`                 //
	UserSalt     string      `orm:"user_salt" json:"userSalt" dc:"加密盐 生成密码用"`       // 加密盐 生成密码用
	Sex          int         `orm:"sex" json:"sex" dc:"1男 2女"`                      // 1男 2女
	Status       int         `orm:"status" json:"status" dc:"1正常 2拉黑冻结"`            // 1正常 2拉黑冻结
	Sign         string      `orm:"sign" json:"sign" dc:"个性签名"`                     // 个性签名
	SecretAnswer string      `orm:"secret_answer" json:"secretAnswer" dc:"密保问题的答案"` // 密保问题的答案
	CreatedAt    *gtime.Time `orm:"created_at" json:"createdAt" dc:""`              //
	UpdatedAt    *gtime.Time `orm:"updated_at" json:"updatedAt" dc:""`              //
	DeletedAt    *gtime.Time `orm:"deleted_at" json:"deletedAt" dc:""`              //
}

type UserInfoListRes struct {
	Id        int         `json:"id" dc:""`
	Name      string      `json:"name" dc:"用户名"`
	Avatar    string      `json:"avatar" dc:"头像"`
	Sex       int         `json:"sex" dc:"1男 2女"`
	Status    int         `json:"status" dc:"1正常 2拉黑冻结"`
	Sign      string      `json:"sign" dc:"个性签名"`
	CreatedAt *gtime.Time `json:"createdAt" dc:""`
}

// UserInfoSearchReq 分页请求参数
type UserInfoSearchReq struct {
	comModel.PageReq
	Id        string `p:"id" dc:""`                                                       //
	Name      string `p:"name" dc:"用户名"`                                                  //用户名
	Avatar    string `p:"avatar" dc:"头像"`                                                 //头像
	Sex       string `p:"sex" v:"sex@integer#1男 2女需为整数" dc:"1男 2女"`                       //1男 2女
	Status    string `p:"status" v:"status@integer#1正常 2拉黑冻结需为整数" dc:"1正常 2拉黑冻结"`         //1正常 2拉黑冻结
	Sign      string `p:"sign" dc:"个性签名"`                                                 //个性签名
	CreatedAt string `p:"createdAt" v:"createdAt@datetime#需为YYYY-MM-DD hh:mm:ss格式" dc:""` //
}

// UserInfoSearchRes 列表返回结果
type UserInfoSearchRes struct {
	comModel.ListRes
	List []*UserInfoListRes `json:"list"`
}

// UserInfoAddReq 添加操作请求参数
type UserInfoAddReq struct {
	Name   string `p:"name" v:"required#用户名不能为空" dc:"用户名"`
	Avatar string `p:"avatar"  dc:"头像"`
	Sex    int    `p:"sex"  dc:"1男 2女"`
	Status int    `p:"status" v:"required#1正常 2拉黑冻结不能为空" dc:"1正常 2拉黑冻结"`
	Sign   string `p:"sign"  dc:"个性签名"`
}

// UserInfoEditReq 修改操作请求参数
type UserInfoEditReq struct {
	Id     int    `p:"id" v:"required#主键ID不能为空" dc:""`
	Name   string `p:"name" v:"required#用户名不能为空" dc:"用户名"`
	Avatar string `p:"avatar"  dc:"头像"`
	Sex    int    `p:"sex"  dc:"1男 2女"`
	Status int    `p:"status" v:"required#1正常 2拉黑冻结不能为空" dc:"1正常 2拉黑冻结"`
	Sign   string `p:"sign"  dc:"个性签名"`
}
