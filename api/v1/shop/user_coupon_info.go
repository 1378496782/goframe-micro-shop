// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2025-09-09 15:39:40
// 生成路径: api/v1/shop/user_coupon_info.go
// 生成人：gfast
// desc:用户优惠券相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package shop

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
)

// UserCouponInfoSearchReq 分页请求参数
type UserCouponInfoSearchReq struct {
	g.Meta `path:"/list" tags:"用户优惠券" method:"get" summary:"用户优惠券管理列表"`
	commonApi.Author
	model.UserCouponInfoSearchReq
}

// UserCouponInfoSearchRes 列表返回结果
type UserCouponInfoSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.UserCouponInfoSearchRes
}

// UserCouponInfoExportReq 导出请求
type UserCouponInfoExportReq struct {
	g.Meta `path:"/export" tags:"用户优惠券" method:"get" summary:"用户优惠券管理导出"`
	commonApi.Author
	model.UserCouponInfoSearchReq
}

// UserCouponInfoExportRes 导出响应
type UserCouponInfoExportRes struct {
	commonApi.EmptyRes
}
type UserCouponInfoExcelTemplateReq struct {
	g.Meta `path:"/excelTemplate" tags:"用户优惠券" method:"get" summary:"导出模板文件"`
	commonApi.Author
}
type UserCouponInfoExcelTemplateRes struct {
	commonApi.EmptyRes
}
type UserCouponInfoImportReq struct {
	g.Meta `path:"/import" tags:"用户优惠券" method:"post" summary:"用户优惠券管理导入"`
	commonApi.Author
	File *ghttp.UploadFile `p:"file" type:"file" dc:"选择上传文件"  v:"required#上传文件必须"`
}
type UserCouponInfoImportRes struct {
	commonApi.EmptyRes
}

// 相关连表查询数据
type LinkedUserCouponInfoDataSearchReq struct {
	g.Meta `path:"/linkedData" tags:"用户优惠券" method:"get" summary:"用户优惠券管理关联表数据"`
	commonApi.Author
}

// 相关连表查询数据
type LinkedUserCouponInfoDataSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.LinkedUserCouponInfoDataSearchRes
}

// UserCouponInfoAddReq 添加操作请求参数
type UserCouponInfoAddReq struct {
	g.Meta `path:"/add" tags:"用户优惠券" method:"post" summary:"用户优惠券管理添加"`
	commonApi.Author
	*model.UserCouponInfoAddReq
}

// UserCouponInfoAddRes 添加操作返回结果
type UserCouponInfoAddRes struct {
	commonApi.EmptyRes
}

// UserCouponInfoEditReq 修改操作请求参数
type UserCouponInfoEditReq struct {
	g.Meta `path:"/edit" tags:"用户优惠券" method:"put" summary:"用户优惠券管理修改"`
	commonApi.Author
	*model.UserCouponInfoEditReq
}

// UserCouponInfoEditRes 修改操作返回结果
type UserCouponInfoEditRes struct {
	commonApi.EmptyRes
}

// UserCouponInfoGetReq 获取一条数据请求
type UserCouponInfoGetReq struct {
	g.Meta `path:"/get" tags:"用户优惠券" method:"get" summary:"获取用户优惠券管理信息"`
	commonApi.Author
	Id int `p:"id" v:"required#主键必须"` //通过主键获取
}

// UserCouponInfoGetRes 获取一条数据结果
type UserCouponInfoGetRes struct {
	g.Meta `mime:"application/json"`
	*model.UserCouponInfoInfoRes
}

// UserCouponInfoDeleteReq 删除数据请求
type UserCouponInfoDeleteReq struct {
	g.Meta `path:"/delete" tags:"用户优惠券" method:"delete" summary:"删除用户优惠券管理"`
	commonApi.Author
	Ids []int `p:"ids" v:"required#主键必须"` //通过主键删除
}

// UserCouponInfoDeleteRes 删除数据返回
type UserCouponInfoDeleteRes struct {
	commonApi.EmptyRes
}
