// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2025-10-10 23:08:01
// 生成路径: internal/app/shop/model/order_info.go
// 生成人：gfast
// desc:订单表
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
	comModel "github.com/tiger1103/gfast/v3/internal/app/common/model"
)

// OrderInfoInfoRes is the golang structure for table order_info.
type OrderInfoInfoRes struct {
	gmeta.Meta       `orm:"table:order_info"`
	Id               int         `orm:"id,primary" json:"id" dc:""`                                               //
	Number           string      `orm:"number" json:"number" dc:"订单编号"`                                           // 订单编号
	UserId           int         `orm:"user_id" json:"userId" dc:"用户id"`                                          // 用户id
	PayType          int         `orm:"pay_type" json:"payType" dc:"支付方式 1微信 2支付宝 3云闪付"`                          // 支付方式 1微信 2支付宝 3云闪付
	Remark           string      `orm:"remark" json:"remark" dc:"备注"`                                             // 备注
	PayAt            *gtime.Time `orm:"pay_at" json:"payAt" dc:"支付时间"`                                            // 支付时间
	Status           int         `orm:"status" json:"status" dc:"订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价 6待确认 7已取消"` // 订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价 6待确认 7已取消
	ConsigneeName    string      `orm:"consignee_name" json:"consigneeName" dc:"收货人姓名"`                           // 收货人姓名
	ConsigneePhone   string      `orm:"consignee_phone" json:"consigneePhone" dc:"收货人手机号"`                        // 收货人手机号
	ConsigneeAddress string      `orm:"consignee_address" json:"consigneeAddress" dc:"收货人详细地址"`                   // 收货人详细地址
	Price            int         `orm:"price" json:"price" dc:"订单金额 单位分"`                                         // 订单金额 单位分
	CouponPrice      int         `orm:"coupon_price" json:"couponPrice" dc:"优惠券金额 单位分"`                           // 优惠券金额 单位分
	ActualPrice      int         `orm:"actual_price" json:"actualPrice" dc:"实际支付金额 单位分"`                          // 实际支付金额 单位分
	CreatedAt        *gtime.Time `orm:"created_at" json:"createdAt" dc:""`                                        //
	UpdatedAt        *gtime.Time `orm:"updated_at" json:"updatedAt" dc:""`                                        //
}

type OrderInfoListRes struct {
	Id               int         `json:"id" dc:"订单id"`
	Number           string      `json:"number" dc:"订单编号"`
	UserId           int         `json:"userId" dc:"用户id"`
	PayType          int         `json:"payType" dc:"支付方式 1微信 2支付宝 3云闪付"`
	Remark           string      `json:"remark" dc:"备注"`
	PayAt            *gtime.Time `json:"payAt" dc:"支付时间"`
	Status           int         `json:"status" dc:"订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价 6待确认 7已取消"`
	ConsigneeName    string      `json:"consigneeName" dc:"收货人姓名"`
	ConsigneePhone   string      `json:"consigneePhone" dc:"收货人手机号"`
	ConsigneeAddress string      `json:"consigneeAddress" dc:"收货人详细地址"`
	Price            int         `json:"price" dc:"订单金额 单位分"`
	CouponPrice      int         `json:"couponPrice" dc:"优惠券金额 单位分"`
	ActualPrice      int         `json:"actualPrice" dc:"实际支付金额 单位分"`
	CreatedAt        *gtime.Time `json:"createdAt" dc:""`
	GoodsInfo        []*OrderGoodsInfoListRes `json:"goods_info" dc:"订单商品信息"`
}

// OrderInfoSearchReq 分页请求参数
type OrderInfoSearchReq struct {
	comModel.PageReq
	Id               string `p:"id" dc:""`                                                                                                                         //
	Number           string `p:"number" dc:"订单编号"`                                                                                                                 //订单编号
	UserId           string `p:"userId" v:"userId@integer#用户id需为整数" dc:"用户id"`                                                                                     //用户id
	PayType          string `p:"payType" v:"payType@integer#支付方式 1微信 2支付宝 3云闪付需为整数" dc:"支付方式 1微信 2支付宝 3云闪付"`                                                       //支付方式 1微信 2支付宝 3云闪付
	PayAt            string `p:"payAt" v:"payAt@datetime#支付时间需为YYYY-MM-DD hh:mm:ss格式" dc:"支付时间"`                                                                   //支付时间
	Status           string `p:"status" v:"status@integer#订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价 6待确认 7已取消需为整数" dc:"订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价 6待确认 7已取消"` //订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价 6待确认 7已取消
	ConsigneeName    string `p:"consigneeName" dc:"收货人姓名"`                                                                                                         //收货人姓名
	ConsigneePhone   string `p:"consigneePhone" dc:"收货人手机号"`                                                                                                       //收货人手机号
	ConsigneeAddress string `p:"consigneeAddress" dc:"收货人详细地址"`                                                                                                    //收货人详细地址
	Price            string `p:"price" v:"price@integer#订单金额 单位分需为整数" dc:"订单金额 单位分"`                                                                               //订单金额 单位分
	CouponPrice      string `p:"couponPrice" v:"couponPrice@integer#优惠券金额 单位分需为整数" dc:"优惠券金额 单位分"`                                                                 //优惠券金额 单位分
	ActualPrice      string `p:"actualPrice" v:"actualPrice@integer#实际支付金额 单位分需为整数" dc:"实际支付金额 单位分"`                                                               //实际支付金额 单位分
	CreatedAt        string `p:"createdAt" v:"createdAt@datetime#需为YYYY-MM-DD hh:mm:ss格式" dc:""`                                                                   //
}

// OrderInfoSearchRes 列表返回结果
type OrderInfoSearchRes struct {
	comModel.ListRes
	List []*OrderInfoListRes `json:"list"`
}

// OrderInfoAddReq 添加操作请求参数
type OrderInfoAddReq struct {
	Number           string      `p:"number" v:"required#订单编号不能为空" dc:"订单编号"`
	UserId           int         `p:"userId"  dc:"用户id"`
	PayType          int         `p:"payType"  dc:"支付方式 1微信 2支付宝 3云闪付"`
	Remark           string      `p:"remark"  dc:"备注"`
	PayAt            *gtime.Time `p:"payAt"  dc:"支付时间"`
	Status           int         `p:"status" v:"required#订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价 6待确认 7已取消不能为空" dc:"订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价 6待确认 7已取消"`
	ConsigneeName    string      `p:"consigneeName" v:"required#收货人姓名不能为空" dc:"收货人姓名"`
	ConsigneePhone   string      `p:"consigneePhone"  dc:"收货人手机号"`
	ConsigneeAddress string      `p:"consigneeAddress"  dc:"收货人详细地址"`
	Price            int         `p:"price"  dc:"订单金额 单位分"`
	CouponPrice      int         `p:"couponPrice"  dc:"优惠券金额 单位分"`
	ActualPrice      int         `p:"actualPrice"  dc:"实际支付金额 单位分"`
}

// OrderInfoEditReq 修改操作请求参数
type OrderInfoEditReq struct {
	Id               int         `p:"id" v:"required#主键ID不能为空" dc:""`
	Number           string      `p:"number" v:"required#订单编号不能为空" dc:"订单编号"`
	UserId           int         `p:"userId"  dc:"用户id"`
	PayType          int         `p:"payType"  dc:"支付方式 1微信 2支付宝 3云闪付"`
	Remark           string      `p:"remark"  dc:"备注"`
	PayAt            *gtime.Time `p:"payAt"  dc:"支付时间"`
	Status           int         `p:"status" v:"required#订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价 6待确认 7已取消不能为空" dc:"订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价 6待确认 7已取消"`
	ConsigneeName    string      `p:"consigneeName" v:"required#收货人姓名不能为空" dc:"收货人姓名"`
	ConsigneePhone   string      `p:"consigneePhone"  dc:"收货人手机号"`
	ConsigneeAddress string      `p:"consigneeAddress"  dc:"收货人详细地址"`
	Price            int         `p:"price"  dc:"订单金额 单位分"`
	CouponPrice      int         `p:"couponPrice"  dc:"优惠券金额 单位分"`
	ActualPrice      int         `p:"actualPrice"  dc:"实际支付金额 单位分"`
}
