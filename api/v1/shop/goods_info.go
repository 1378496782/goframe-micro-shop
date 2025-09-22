// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2025-09-22 16:30:35
// 生成路径: api/v1/shop/goods_info.go
// 生成人：gfast
// desc:商品相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package shop

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
)

// GoodsInfoSearchReq 分页请求参数
type GoodsInfoSearchReq struct {
	g.Meta `path:"/list" tags:"商品" method:"get" summary:"商品列表"`
	commonApi.Author
	model.GoodsInfoSearchReq
}

// GoodsInfoSearchRes 列表返回结果
type GoodsInfoSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.GoodsInfoSearchRes
}

// GoodsInfoAddReq 添加操作请求参数
type GoodsInfoAddReq struct {
	g.Meta `path:"/add" tags:"商品" method:"post" summary:"商品添加"`
	commonApi.Author
	*model.GoodsInfoAddReq
}

// GoodsInfoAddRes 添加操作返回结果
type GoodsInfoAddRes struct {
	commonApi.EmptyRes
}

// GoodsInfoEditReq 修改操作请求参数
type GoodsInfoEditReq struct {
	g.Meta `path:"/edit" tags:"商品" method:"put" summary:"商品修改"`
	commonApi.Author
	*model.GoodsInfoEditReq
}

// GoodsInfoEditRes 修改操作返回结果
type GoodsInfoEditRes struct {
	commonApi.EmptyRes
}

// GoodsInfoGetReq 获取一条数据请求
type GoodsInfoGetReq struct {
	g.Meta `path:"/get" tags:"商品" method:"get" summary:"获取商品信息"`
	commonApi.Author
	Id uint `p:"id" v:"required#主键必须"` //通过主键获取
}

// GoodsInfoGetRes 获取一条数据结果
type GoodsInfoGetRes struct {
	g.Meta `mime:"application/json"`
	*model.GoodsInfoInfoRes
}

// GoodsInfoDeleteReq 删除数据请求
type GoodsInfoDeleteReq struct {
	g.Meta `path:"/delete" tags:"商品" method:"delete" summary:"删除商品"`
	commonApi.Author
	Ids []uint `p:"ids" v:"required#主键必须"` //通过主键删除
}

// GoodsInfoDeleteRes 删除数据返回
type GoodsInfoDeleteRes struct {
	commonApi.EmptyRes
}
