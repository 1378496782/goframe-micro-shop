// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2025-10-10 23:08:00
// 生成路径: api/v1/shop/order_info.go
// 生成人：gfast
// desc:订单表相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package shop

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
)

// OrderInfoSearchReq 分页请求参数
type OrderInfoSearchReq struct {
	g.Meta `path:"/list" tags:"订单表" method:"get" summary:"订单表列表"`
	commonApi.Author
	model.OrderInfoSearchReq
}

// OrderInfoSearchRes 列表返回结果
type OrderInfoSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.OrderInfoSearchRes
}

// OrderInfoAddReq 添加操作请求参数
type OrderInfoAddReq struct {
	g.Meta `path:"/add" tags:"订单表" method:"post" summary:"订单表添加"`
	commonApi.Author
	*model.OrderInfoAddReq
}

// OrderInfoAddRes 添加操作返回结果
type OrderInfoAddRes struct {
	commonApi.EmptyRes
}

// OrderInfoEditReq 修改操作请求参数
type OrderInfoEditReq struct {
	g.Meta `path:"/edit" tags:"订单表" method:"put" summary:"订单表修改"`
	commonApi.Author
	*model.OrderInfoEditReq
}

// OrderInfoEditRes 修改操作返回结果
type OrderInfoEditRes struct {
	commonApi.EmptyRes
}

// OrderInfoGetReq 获取一条数据请求
type OrderInfoGetReq struct {
	g.Meta `path:"/get" tags:"订单表" method:"get" summary:"获取订单表信息"`
	commonApi.Author
	Id int `p:"id" v:"required#主键必须"` //通过主键获取
}

// OrderInfoGetRes 获取一条数据结果
type OrderInfoGetRes struct {
	g.Meta `mime:"application/json"`
	*model.OrderInfoInfoRes
}

// OrderInfoDeleteReq 删除数据请求
type OrderInfoDeleteReq struct {
	g.Meta `path:"/delete" tags:"订单表" method:"delete" summary:"删除订单表"`
	commonApi.Author
	Ids []int `p:"ids" v:"required#主键必须"` //通过主键删除
}

// OrderInfoDeleteRes 删除数据返回
type OrderInfoDeleteRes struct {
	commonApi.EmptyRes
}

// OrderInfoShipReq 发货请求参数
type OrderInfoShipReq struct {
	g.Meta `path:"/ship" tags:"订单表" method:"put" summary:"订单发货"`
	commonApi.Author
	Id int `p:"id" v:"required#订单ID必须"` // 订单ID
}

// OrderInfoShipRes 发货操作返回结果
type OrderInfoShipRes struct {
	commonApi.EmptyRes
}

// OrderInfoRefundReq 退款请求参数
type OrderInfoRefundReq struct {
	g.Meta `path:"/refund" tags:"订单表" method:"put" summary:"订单退款"`
	commonApi.Author
	Id int `p:"id" v:"required#订单ID必须"` // 订单ID
}

// OrderInfoRefundRes 退款操作返回结果
type OrderInfoRefundRes struct {
	commonApi.EmptyRes
}

// OrderInfoGetProductsReq 获取订单商品列表请求参数
type OrderInfoGetProductsReq struct {
	g.Meta `path:"/getProducts" tags:"订单表" method:"get" summary:"获取订单商品列表"`
	commonApi.Author
	OrderId int `p:"orderId" v:"required#订单ID必须"` // 订单ID
}

// OrderInfoGetProductsRes 获取订单商品列表返回结果
type OrderInfoGetProductsRes struct {
	g.Meta `mime:"application/json"`
	List []*model.OrderGoodsInfoListRes `json:"list"`
}
