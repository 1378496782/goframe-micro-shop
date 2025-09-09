// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2025-09-09 15:39:40
// 生成路径: internal/app/shop/model/user_coupon_info.go
// 生成人：gfast
// desc:用户优惠券
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
	comModel "github.com/tiger1103/gfast/v3/internal/app/common/model"
)

// UserCouponInfoInfoRes is the golang structure for table user_coupon_info.
type UserCouponInfoInfoRes struct {
	gmeta.Meta     `orm:"table:user_coupon_info"`
	Id             int                             `orm:"id,primary" json:"id" dc:""`           //
	UserId         int                             `orm:"user_id" json:"userId" dc:"用户id"`      // 用户id
	CouponId       int                             `orm:"coupon_id" json:"couponId" dc:"优惠券id"` // 优惠券id
	LinkedCouponId *LinkedUserCouponInfoCouponInfo `orm:"with:id=coupon_id" json:"linkedCouponId"`
	Status         int                             `orm:"status" json:"status" dc:"状态"`               // 状态
	Amount         int                             `orm:"amount" json:"amount" dc:"优惠金额（元）"`          // 优惠金额（元）
	CreatedAt      *gtime.Time                     `orm:"created_at" json:"createdAt" dc:"创建时间"`      // 创建时间
	UpdatedAt      *gtime.Time                     `orm:"updated_at" json:"updatedAt" dc:"更新时间"`      // 更新时间
	DeletedAt      *gtime.Time                     `orm:"deleted_at" json:"deletedAt" dc:"删除时间（软删除）"` // 删除时间（软删除）
}
type LinkedUserCouponInfoCouponInfo struct {
	gmeta.Meta `orm:"table:coupon_info"`
	Id         int    `orm:"id" json:"id" dc:""`          //
	Name       string `orm:"name" json:"name" dc:"优惠券名称"` // 优惠券名称
}

type UserCouponInfoListRes struct {
	Id             int                             `json:"id" dc:""`
	UserId         int                             `json:"userId" dc:"用户id"`
	CouponId       int                             `json:"couponId" dc:"优惠券id"`
	LinkedCouponId *LinkedUserCouponInfoCouponInfo `orm:"with:id=coupon_id" json:"linkedCouponId" dc:"优惠券id"`
	Status         int                             `json:"status" dc:"状态"`
	Amount         int                             `json:"amount" dc:"优惠金额（元）"`
	CreatedAt      *gtime.Time                     `json:"createdAt" dc:"创建时间"`
}

// UserCouponInfoSearchReq 分页请求参数
type UserCouponInfoSearchReq struct {
	comModel.PageReq
	Id        string `p:"id" dc:""`                                                               //
	UserId    string `p:"userId" v:"userId@integer#用户id需为整数" dc:"用户id"`                           //用户id
	CouponId  string `p:"couponId" v:"couponId@integer#优惠券id需为整数" dc:"优惠券id"`                     //优惠券id
	Status    string `p:"status" v:"status@integer#状态需为整数" dc:"状态"`                               //状态
	Amount    string `p:"amount" v:"amount@integer#优惠金额（元）需为整数" dc:"优惠金额（元）"`                     //优惠金额（元）
	CreatedAt string `p:"createdAt" v:"createdAt@datetime#创建时间需为YYYY-MM-DD hh:mm:ss格式" dc:"创建时间"` //创建时间
}

// UserCouponInfoSearchRes 列表返回结果
type UserCouponInfoSearchRes struct {
	comModel.ListRes
	List []*UserCouponInfoListRes `json:"list"`
}

// 相关连表查询数据
type LinkedUserCouponInfoDataSearchRes struct {
	LinkedUserCouponInfoCouponInfo []*LinkedUserCouponInfoCouponInfo `json:"linkedUserCouponInfoCouponInfo"`
}

// UserCouponInfoAddReq 添加操作请求参数
type UserCouponInfoAddReq struct {
	UserId   int `p:"userId"  dc:"用户id"`
	CouponId int `p:"couponId" v:"required#优惠券id不能为空" dc:"优惠券id"`
	Status   int `p:"status" v:"required#状态不能为空" dc:"状态"`
	Amount   int `p:"amount"  dc:"优惠金额（元）"`
}

// UserCouponInfoEditReq 修改操作请求参数
type UserCouponInfoEditReq struct {
	Id       int `p:"id" v:"required#主键ID不能为空" dc:""`
	UserId   int `p:"userId"  dc:"用户id"`
	CouponId int `p:"couponId" v:"required#优惠券id不能为空" dc:"优惠券id"`
	Status   int `p:"status" v:"required#状态不能为空" dc:"状态"`
	Amount   int `p:"amount"  dc:"优惠金额（元）"`
}
