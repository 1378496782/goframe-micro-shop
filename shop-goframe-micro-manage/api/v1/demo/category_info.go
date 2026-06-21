// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2025-09-05 11:32:02
// 生成路径: api/v1/demo/category_info.go
// 生成人：王中阳
// desc:商品分类相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package demo

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/demo/model"
)

// CategoryInfoSearchReq 分页请求参数
type CategoryInfoSearchReq struct {
	g.Meta `path:"/list" tags:"商品分类" method:"get" summary:"商品分类列表"`
	commonApi.Author
	model.CategoryInfoSearchReq
}

// CategoryInfoSearchRes 列表返回结果
type CategoryInfoSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.CategoryInfoSearchRes
}

// CategoryInfoAddReq 添加操作请求参数
type CategoryInfoAddReq struct {
	g.Meta `path:"/add" tags:"商品分类" method:"post" summary:"商品分类添加"`
	commonApi.Author
	*model.CategoryInfoAddReq
}

// CategoryInfoAddRes 添加操作返回结果
type CategoryInfoAddRes struct {
	commonApi.EmptyRes
}

// CategoryInfoEditReq 修改操作请求参数
type CategoryInfoEditReq struct {
	g.Meta `path:"/edit" tags:"商品分类" method:"put" summary:"商品分类修改"`
	commonApi.Author
	*model.CategoryInfoEditReq
}

// CategoryInfoEditRes 修改操作返回结果
type CategoryInfoEditRes struct {
	commonApi.EmptyRes
}

// CategoryInfoGetReq 获取一条数据请求
type CategoryInfoGetReq struct {
	g.Meta `path:"/get" tags:"商品分类" method:"get" summary:"获取商品分类信息"`
	commonApi.Author
	Id int `p:"id" v:"required#主键必须"` //通过主键获取
}

// CategoryInfoGetRes 获取一条数据结果
type CategoryInfoGetRes struct {
	g.Meta `mime:"application/json"`
	*model.CategoryInfoInfoRes
}

// CategoryInfoDeleteReq 删除数据请求
type CategoryInfoDeleteReq struct {
	g.Meta `path:"/delete" tags:"商品分类" method:"delete" summary:"删除商品分类"`
	commonApi.Author
	Ids []int `p:"ids" v:"required#主键必须"` //通过主键删除
}

// CategoryInfoDeleteRes 删除数据返回
type CategoryInfoDeleteRes struct {
	commonApi.EmptyRes
}
