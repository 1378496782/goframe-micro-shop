// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2025-09-09 15:10:49
// 生成路径: api/v1/shop/coupon_info.go
// 生成人：gfast
// desc:优惠券相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package shop

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
)

// CouponInfoSearchReq 分页请求参数
type CouponInfoSearchReq struct {
	g.Meta `path:"/list" tags:"优惠券" method:"get" summary:"优惠券列表"`
	commonApi.Author
	model.CouponInfoSearchReq
}

// CouponInfoSearchRes 列表返回结果
type CouponInfoSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.CouponInfoSearchRes
}

// CouponInfoExportReq 导出请求
type CouponInfoExportReq struct {
	g.Meta `path:"/export" tags:"优惠券" method:"get" summary:"优惠券导出"`
	commonApi.Author
	model.CouponInfoSearchReq
}

// CouponInfoExportRes 导出响应
type CouponInfoExportRes struct {
	commonApi.EmptyRes
}
type CouponInfoExcelTemplateReq struct {
	g.Meta `path:"/excelTemplate" tags:"优惠券" method:"get" summary:"导出模板文件"`
	commonApi.Author
}
type CouponInfoExcelTemplateRes struct {
	commonApi.EmptyRes
}
type CouponInfoImportReq struct {
	g.Meta `path:"/import" tags:"优惠券" method:"post" summary:"优惠券导入"`
	commonApi.Author
	File *ghttp.UploadFile `p:"file" type:"file" dc:"选择上传文件"  v:"required#上传文件必须"`
}
type CouponInfoImportRes struct {
	commonApi.EmptyRes
}

// CouponInfoAddReq 添加操作请求参数
type CouponInfoAddReq struct {
	g.Meta `path:"/add" tags:"优惠券" method:"post" summary:"优惠券添加"`
	commonApi.Author
	*model.CouponInfoAddReq
}

// CouponInfoAddRes 添加操作返回结果
type CouponInfoAddRes struct {
	commonApi.EmptyRes
}

// CouponInfoEditReq 修改操作请求参数
type CouponInfoEditReq struct {
	g.Meta `path:"/edit" tags:"优惠券" method:"put" summary:"优惠券修改"`
	commonApi.Author
	*model.CouponInfoEditReq
}

// CouponInfoEditRes 修改操作返回结果
type CouponInfoEditRes struct {
	commonApi.EmptyRes
}

// CouponInfoGetReq 获取一条数据请求
type CouponInfoGetReq struct {
	g.Meta `path:"/get" tags:"优惠券" method:"get" summary:"获取优惠券信息"`
	commonApi.Author
	Id int `p:"id" v:"required#主键必须"` //通过主键获取
}

// CouponInfoGetRes 获取一条数据结果
type CouponInfoGetRes struct {
	g.Meta `mime:"application/json"`
	*model.CouponInfoInfoRes
}

// CouponInfoDeleteReq 删除数据请求
type CouponInfoDeleteReq struct {
	g.Meta `path:"/delete" tags:"优惠券" method:"delete" summary:"删除优惠券"`
	commonApi.Author
	Ids []int `p:"ids" v:"required#主键必须"` //通过主键删除
}

// CouponInfoDeleteRes 删除数据返回
type CouponInfoDeleteRes struct {
	commonApi.EmptyRes
}
