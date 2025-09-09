// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2025-09-09 15:10:49
// 生成路径: internal/app/shop/model/coupon_info.go
// 生成人：gfast
// desc:优惠券
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
	comModel "github.com/tiger1103/gfast/v3/internal/app/common/model"
)

// CouponInfoInfoRes is the golang structure for table coupon_info.
type CouponInfoInfoRes struct {
	gmeta.Meta `orm:"table:coupon_info"`
	Id         int         `orm:"id,primary" json:"id" dc:"ID"`                 // ID
	GoodsId    int         `orm:"goods_id" json:"goodsId" dc:"关联商品id（0表示全场通用）"` // 关联商品id（0表示全场通用）
	Name       string      `orm:"name" json:"name" dc:"名称"`                     // 名称
	Type       int         `orm:"type" json:"type" dc:"类型"`                     // 类型
	Amount     int         `orm:"amount" json:"amount" dc:"优惠金额（元）"`            // 优惠金额（元）
	Deadline   *gtime.Time `orm:"deadline" json:"deadline" dc:"过期时间"`           // 过期时间
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt" dc:"创建时间"`        // 创建时间
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt" dc:"更新时间"`        // 更新时间
	DeletedAt  *gtime.Time `orm:"deleted_at" json:"deletedAt" dc:"删除时间（软删除）"`   // 删除时间（软删除）
}

type CouponInfoListRes struct {
	Id        int         `json:"id" dc:"ID"`
	Name      string      `json:"name" dc:"名称"`
	Type      int         `json:"type" dc:"类型"`
	Amount    int         `json:"amount" dc:"优惠金额（元）"`
	Deadline  *gtime.Time `json:"deadline" dc:"过期时间"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}

// CouponInfoSearchReq 分页请求参数
type CouponInfoSearchReq struct {
	comModel.PageReq
	Id        string `p:"id" dc:"ID"`                                                             //ID
	Name      string `p:"name" dc:"名称"`                                                           //名称
	Type      string `p:"type" v:"type@integer#类型需为整数" dc:"类型"`                                   //类型
	Amount    string `p:"amount" v:"amount@integer#优惠金额（元）需为整数" dc:"优惠金额（元）"`                     //优惠金额（元）
	Deadline  string `p:"deadline" v:"deadline@datetime#过期时间需为YYYY-MM-DD hh:mm:ss格式" dc:"过期时间"`   //过期时间
	CreatedAt string `p:"createdAt" v:"createdAt@datetime#创建时间需为YYYY-MM-DD hh:mm:ss格式" dc:"创建时间"` //创建时间
}

// CouponInfoSearchRes 列表返回结果
type CouponInfoSearchRes struct {
	comModel.ListRes
	List []*CouponInfoListRes `json:"list"`
}

// CouponInfoAddReq 添加操作请求参数
type CouponInfoAddReq struct {
	Name     string      `p:"name" v:"required#名称不能为空" dc:"名称"`
	Type     int         `p:"type"  dc:"类型"`
	Amount   int         `p:"amount"  dc:"优惠金额（元）"`
	Deadline *gtime.Time `p:"deadline" v:"required#过期时间不能为空" dc:"过期时间"`
}

// CouponInfoEditReq 修改操作请求参数
type CouponInfoEditReq struct {
	Id       int         `p:"id" v:"required#主键ID不能为空" dc:"ID"`
	Name     string      `p:"name" v:"required#名称不能为空" dc:"名称"`
	Type     int         `p:"type"  dc:"类型"`
	Amount   int         `p:"amount"  dc:"优惠金额（元）"`
	Deadline *gtime.Time `p:"deadline" v:"required#过期时间不能为空" dc:"过期时间"`
}
