// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2025-10-11 23:55:29
// 生成路径: api/v1/shop/order_goods_info.go
// 生成人：gfast
// desc:订单物品表相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package shop

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
)

// OrderGoodsInfoSearchReq 分页请求参数
type OrderGoodsInfoSearchReq struct {
	g.Meta `path:"/list" tags:"订单物品表" method:"get" summary:"订单物品表列表"`
	commonApi.Author
	model.OrderGoodsInfoSearchReq
}

// OrderGoodsInfoSearchRes 列表返回结果
type OrderGoodsInfoSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.OrderGoodsInfoSearchRes
}

// OrderGoodsInfoAddReq 添加操作请求参数
type OrderGoodsInfoAddReq struct {
	g.Meta `path:"/add" tags:"订单物品表" method:"post" summary:"订单物品表添加"`
	commonApi.Author
	*model.OrderGoodsInfoAddReq
}

// OrderGoodsInfoAddRes 添加操作返回结果
type OrderGoodsInfoAddRes struct {
	commonApi.EmptyRes
}

// OrderGoodsInfoEditReq 修改操作请求参数
type OrderGoodsInfoEditReq struct {
	g.Meta `path:"/edit" tags:"订单物品表" method:"put" summary:"订单物品表修改"`
	commonApi.Author
	*model.OrderGoodsInfoEditReq
}

// OrderGoodsInfoEditRes 修改操作返回结果
type OrderGoodsInfoEditRes struct {
	commonApi.EmptyRes
}

// OrderGoodsInfoGetReq 获取一条数据请求
type OrderGoodsInfoGetReq struct {
	g.Meta `path:"/get" tags:"订单物品表" method:"get" summary:"获取订单物品表信息"`
	commonApi.Author
	Id int `p:"id" v:"required#主键必须"` //通过主键获取
}

// OrderGoodsInfoGetRes 获取一条数据结果
type OrderGoodsInfoGetRes struct {
	g.Meta `mime:"application/json"`
	*model.OrderGoodsInfoInfoRes
}

// OrderGoodsInfoDeleteReq 删除数据请求
type OrderGoodsInfoDeleteReq struct {
	g.Meta `path:"/delete" tags:"订单物品表" method:"delete" summary:"删除订单物品表"`
	commonApi.Author
	Ids []int `p:"ids" v:"required#主键必须"` //通过主键删除
}

// OrderGoodsInfoDeleteRes 删除数据返回
type OrderGoodsInfoDeleteRes struct {
	commonApi.EmptyRes
}

// OrderGoodsInfoGetByOrderIdReq 根据订单ID获取订单商品列表请求
type OrderGoodsInfoGetByOrderIdReq struct {
	g.Meta `path:"/getByOrderId" tags:"订单物品表" method:"get" summary:"根据订单ID获取订单商品列表"`
	commonApi.Author
	OrderId int `p:"orderId" v:"required#订单ID必须"` //订单ID
}

// OrderGoodsInfoGetByOrderIdRes 根据订单ID获取订单商品列表返回
type OrderGoodsInfoGetByOrderIdRes struct {
	g.Meta `mime:"application/json"`
	List []*model.OrderGoodsInfoListRes `json:"list"`
}

// OrderGoodsInfoGetDetailReq 获取订单商品详细信息请求
type OrderGoodsInfoGetDetailReq struct {
	g.Meta `path:"/getDetail" tags:"订单物品表" method:"get" summary:"获取订单商品详细信息"`
	commonApi.Author
	Id int `p:"id" v:"required#主键必须"` //通过主键获取
}

// OrderGoodsInfoGetDetailRes 获取订单商品详细信息返回
type OrderGoodsInfoGetDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.OrderGoodsDetailRes
}

// OrderGoodsInfoAddOrderGoodsReq 添加订单商品请求
type OrderGoodsInfoAddOrderGoodsReq struct {
	g.Meta `path:"/addOrderGoods" tags:"订单物品表" method:"post" summary:"添加订单商品"`
	commonApi.Author
	*model.OrderGoodsAddReq
}

// OrderGoodsInfoAddOrderGoodsRes 添加订单商品返回
type OrderGoodsInfoAddOrderGoodsRes struct {
	commonApi.EmptyRes
}
