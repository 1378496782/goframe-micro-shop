// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2025-09-22 16:50:03
// 生成路径: api/v1/shop/rotation_info.go
// 生成人：gfast
// desc:轮播图相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package shop

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
)

// RotationInfoSearchReq 分页请求参数
type RotationInfoSearchReq struct {
	g.Meta `path:"/list" tags:"轮播图" method:"get" summary:"轮播图列表"`
	commonApi.Author
	model.RotationInfoSearchReq
}

// RotationInfoSearchRes 列表返回结果
type RotationInfoSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.RotationInfoSearchRes
}

// RotationInfoAddReq 添加操作请求参数
type RotationInfoAddReq struct {
	g.Meta `path:"/add" tags:"轮播图" method:"post" summary:"轮播图添加"`
	commonApi.Author
	*model.RotationInfoAddReq
}

// RotationInfoAddRes 添加操作返回结果
type RotationInfoAddRes struct {
	commonApi.EmptyRes
}

// RotationInfoEditReq 修改操作请求参数
type RotationInfoEditReq struct {
	g.Meta `path:"/edit" tags:"轮播图" method:"put" summary:"轮播图修改"`
	commonApi.Author
	*model.RotationInfoEditReq
}

// RotationInfoEditRes 修改操作返回结果
type RotationInfoEditRes struct {
	commonApi.EmptyRes
}

// RotationInfoGetReq 获取一条数据请求
type RotationInfoGetReq struct {
	g.Meta `path:"/get" tags:"轮播图" method:"get" summary:"获取轮播图信息"`
	commonApi.Author
	Id int `p:"id" v:"required#主键必须"` //通过主键获取
}

// RotationInfoGetRes 获取一条数据结果
type RotationInfoGetRes struct {
	g.Meta `mime:"application/json"`
	*model.RotationInfoInfoRes
}

// RotationInfoDeleteReq 删除数据请求
type RotationInfoDeleteReq struct {
	g.Meta `path:"/delete" tags:"轮播图" method:"delete" summary:"删除轮播图"`
	commonApi.Author
	Ids []int `p:"ids" v:"required#主键必须"` //通过主键删除
}

// RotationInfoDeleteRes 删除数据返回
type RotationInfoDeleteRes struct {
	commonApi.EmptyRes
}
