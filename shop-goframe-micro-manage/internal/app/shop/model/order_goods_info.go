// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2025-10-11 23:55:29
// 生成路径: internal/app/shop/model/order_goods_info.go
// 生成人：gfast
// desc:订单物品表
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
	comModel "github.com/tiger1103/gfast/v3/internal/app/common/model"
)

// OrderGoodsInfoInfoRes is the golang structure for table order_goods_info.
type OrderGoodsInfoInfoRes struct {
	gmeta.Meta     `orm:"table:order_goods_info"`
	Id             int         `orm:"id,primary" json:"id" dc:"商品维度的订单表"`                        // 商品维度的订单表
	OrderId        int         `orm:"order_id" json:"orderId" dc:"关联的主订单表"`                      // 关联的主订单表
	GoodsId        int         `orm:"goods_id" json:"goodsId" dc:"商品id"`                         // 商品id
	GoodsOptionsId int         `orm:"goods_options_id" json:"goodsOptionsId" dc:"商品规格id sku id"` // 商品规格id sku id
	Count          int         `orm:"count" json:"count" dc:"商品数量"`                              // 商品数量
	Remark         string      `orm:"remark" json:"remark" dc:"备注"`                              // 备注
	Price          int         `orm:"price" json:"price" dc:"订单金额 单位分"`                          // 订单金额 单位分
	CouponPrice    int         `orm:"coupon_price" json:"couponPrice" dc:"优惠券金额 单位分"`            // 优惠券金额 单位分
	ActualPrice    int         `orm:"actual_price" json:"actualPrice" dc:"实际支付金额 单位分"`           // 实际支付金额 单位分
	CreatedAt      *gtime.Time `orm:"created_at" json:"createdAt" dc:""`                         //
	UpdatedAt      *gtime.Time `orm:"updated_at" json:"updatedAt" dc:""`                         //
}

type OrderGoodsInfoListRes struct {
	Id             int         `json:"id" dc:"商品维度的订单表"`
	OrderId        int         `json:"orderId" dc:"关联的主订单表"`
	GoodsId        int         `json:"goodsId" dc:"商品id"`
	GoodsName      string      `json:"goods_name" dc:"商品名称"`
	PicUrl         string      `json:"pic_url" dc:"商品图片"`
	GoodsOptionsId int         `json:"goodsOptionsId" dc:"商品规格id sku id"`
	Count          int         `json:"count" dc:"商品数量"`
	Remark         string      `json:"remark" dc:"备注"`
	Price          int         `json:"price" dc:"订单金额 单位分"`
	CouponPrice    int         `json:"couponPrice" dc:"优惠券金额 单位分"`
	ActualPrice    int         `json:"actualPrice" dc:"实际支付金额 单位分"`
	CreatedAt      *gtime.Time `json:"createdAt" dc:""`
}

// OrderGoodsInfoSearchReq 分页请求参数
type OrderGoodsInfoSearchReq struct {
	comModel.PageReq
	Id             string `p:"id" dc:"商品维度的订单表"`                                                               //商品维度的订单表
	OrderId        string `p:"orderId" v:"orderId@integer#关联的主订单表需为整数" dc:"关联的主订单表"`                           //关联的主订单表
	GoodsId        string `p:"goodsId" v:"goodsId@integer#商品id需为整数" dc:"商品id"`                                 //商品id
	GoodsOptionsId string `p:"goodsOptionsId" v:"goodsOptionsId@integer#商品规格id sku id需为整数" dc:"商品规格id sku id"` //商品规格id sku id
	Count          string `p:"count" v:"count@integer#商品数量需为整数" dc:"商品数量"`                                     //商品数量
	Price          string `p:"price" v:"price@integer#订单金额 单位分需为整数" dc:"订单金额 单位分"`                             //订单金额 单位分
	CouponPrice    string `p:"couponPrice" v:"couponPrice@integer#优惠券金额 单位分需为整数" dc:"优惠券金额 单位分"`               //优惠券金额 单位分
	ActualPrice    string `p:"actualPrice" v:"actualPrice@integer#实际支付金额 单位分需为整数" dc:"实际支付金额 单位分"`             //实际支付金额 单位分
	CreatedAt      string `p:"createdAt" v:"createdAt@datetime#需为YYYY-MM-DD hh:mm:ss格式" dc:""`                 //
}

// OrderGoodsInfoSearchRes 列表返回结果
type OrderGoodsInfoSearchRes struct {
	comModel.ListRes
	List []*OrderGoodsInfoListRes `json:"list"`
}

// OrderGoodsInfoAddReq 添加操作请求参数
type OrderGoodsInfoAddReq struct {
	OrderId        int    `p:"orderId"  dc:"关联的主订单表"`
	GoodsId        int    `p:"goodsId"  dc:"商品id"`
	GoodsOptionsId int    `p:"goodsOptionsId"  dc:"商品规格id sku id"`
	Count          int    `p:"count" v:"required#商品数量不能为空" dc:"商品数量"`
	Remark         string `p:"remark"  dc:"备注"`
	Price          int    `p:"price"  dc:"订单金额 单位分"`
	CouponPrice    int    `p:"couponPrice"  dc:"优惠券金额 单位分"`
	ActualPrice    int    `p:"actualPrice"  dc:"实际支付金额 单位分"`
}

// OrderGoodsInfoEditReq 修改操作请求参数
type OrderGoodsInfoEditReq struct {
	Id             int    `p:"id" v:"required#主键ID不能为空" dc:"商品维度的订单表"`
	OrderId        int    `p:"orderId"  dc:"关联的主订单表"`
	GoodsId        int    `p:"goodsId"  dc:"商品id"`
	GoodsOptionsId int    `p:"goodsOptionsId"  dc:"商品规格id sku id"`
	Count          int    `p:"count" v:"required#商品数量不能为空" dc:"商品数量"`
	Remark         string `p:"remark"  dc:"备注"`
	Price          int    `p:"price"  dc:"订单金额 单位分"`
	CouponPrice    int    `p:"couponPrice"  dc:"优惠券金额 单位分"`
	ActualPrice    int    `p:"actualPrice"  dc:"实际支付金额 单位分"`
}

// OrderGoodsDetailRes 订单商品详细信息返回结果
type OrderGoodsDetailRes struct {
	Id             int         `json:"id" dc:"商品维度的订单表"`
	OrderId        int         `json:"orderId" dc:"关联的主订单表"`
	GoodsId        int         `json:"goodsId" dc:"商品id"`
	GoodsName      string      `json:"goodsName" dc:"商品名称"`
	GoodsImage     string      `json:"goodsImage" dc:"商品图片"`
	GoodsOptionsId int         `json:"goodsOptionsId" dc:"商品规格id sku id"`
	Count          int         `json:"count" dc:"商品数量"`
	Remark         string      `json:"remark" dc:"备注"`
	Price          int         `json:"price" dc:"订单金额 单位分"`
	CouponPrice    int         `json:"couponPrice" dc:"优惠券金额 单位分"`
	ActualPrice    int         `json:"actualPrice" dc:"实际支付金额 单位分"`
	CreatedAt      *gtime.Time `json:"createdAt" dc:""`
}

// OrderGoodsAddReq 添加订单商品请求参数
type OrderGoodsAddReq struct {
	OrderId        int    `p:"orderId" v:"required#订单ID不能为空" dc:"关联的主订单表"`
	GoodsId        int    `p:"goodsId" v:"required#商品ID不能为空" dc:"商品id"`
	GoodsOptionsId int    `p:"goodsOptionsId" dc:"商品规格id sku id"`
	Count          int    `p:"count" v:"required#商品数量不能为空" dc:"商品数量"`
	Remark         string `p:"remark" dc:"备注"`
	Price          int    `p:"price" v:"required#订单金额不能为空" dc:"订单金额 单位分"`
	CouponPrice    int    `p:"couponPrice" dc:"优惠券金额 单位分"`
	ActualPrice    int    `p:"actualPrice" v:"required#实际支付金额不能为空" dc:"实际支付金额 单位分"`
}
