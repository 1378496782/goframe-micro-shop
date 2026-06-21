// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2025-09-09 15:51:53
// 生成路径: api/v1/shop/user_info.go
// 生成人：gfast
// desc:用户相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package shop

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
)

// UserInfoSearchReq 分页请求参数
type UserInfoSearchReq struct {
	g.Meta `path:"/list" tags:"用户" method:"get" summary:"用户列表"`
	commonApi.Author
	model.UserInfoSearchReq
}

// UserInfoSearchRes 列表返回结果
type UserInfoSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.UserInfoSearchRes
}

// UserInfoExportReq 导出请求
type UserInfoExportReq struct {
	g.Meta `path:"/export" tags:"用户" method:"get" summary:"用户导出"`
	commonApi.Author
	model.UserInfoSearchReq
}

// UserInfoExportRes 导出响应
type UserInfoExportRes struct {
	commonApi.EmptyRes
}

// UserInfoAddReq 添加操作请求参数
type UserInfoAddReq struct {
	g.Meta `path:"/add" tags:"用户" method:"post" summary:"用户添加"`
	commonApi.Author
	*model.UserInfoAddReq
}

// UserInfoAddRes 添加操作返回结果
type UserInfoAddRes struct {
	commonApi.EmptyRes
}

// UserInfoEditReq 修改操作请求参数
type UserInfoEditReq struct {
	g.Meta `path:"/edit" tags:"用户" method:"put" summary:"用户修改"`
	commonApi.Author
	*model.UserInfoEditReq
}

// UserInfoEditRes 修改操作返回结果
type UserInfoEditRes struct {
	commonApi.EmptyRes
}

// UserInfoGetReq 获取一条数据请求
type UserInfoGetReq struct {
	g.Meta `path:"/get" tags:"用户" method:"get" summary:"获取用户信息"`
	commonApi.Author
	Id int `p:"id" v:"required#主键必须"` //通过主键获取
}

// UserInfoGetRes 获取一条数据结果
type UserInfoGetRes struct {
	g.Meta `mime:"application/json"`
	*model.UserInfoInfoRes
}

// UserInfoDeleteReq 删除数据请求
type UserInfoDeleteReq struct {
	g.Meta `path:"/delete" tags:"用户" method:"delete" summary:"删除用户"`
	commonApi.Author
	Ids []int `p:"ids" v:"required#主键必须"` //通过主键删除
}

// UserInfoDeleteRes 删除数据返回
type UserInfoDeleteRes struct {
	commonApi.EmptyRes
}
