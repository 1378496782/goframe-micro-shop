// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2025-09-22 16:50:04
// 生成路径: internal/app/shop/model/rotation_info.go
// 生成人：gfast
// desc:轮播图
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
	comModel "github.com/tiger1103/gfast/v3/internal/app/common/model"
)

// RotationInfoInfoRes is the golang structure for table rotation_info.
type RotationInfoInfoRes struct {
	gmeta.Meta `orm:"table:rotation_info"`
	Id         int         `orm:"id,primary" json:"id" dc:"ID"`          // ID
	PicUrl     string      `orm:"pic_url" json:"picUrl" dc:"图片"`         // 图片
	Link       string      `orm:"link" json:"link" dc:"跳转链接"`            // 跳转链接
	Sort       int         `orm:"sort" json:"sort" dc:"排序字段"`            // 排序字段
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt" dc:"创建时间"` // 创建时间
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt" dc:""`     //
	DeletedAt  *gtime.Time `orm:"deleted_at" json:"deletedAt" dc:""`     //
}

type RotationInfoListRes struct {
	Id        int         `json:"id" dc:"ID"`
	PicUrl    string      `json:"picUrl" dc:"图片"`
	Link      string      `json:"link" dc:"跳转链接"`
	Sort      int         `json:"sort" dc:"排序字段"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}

// RotationInfoSearchReq 分页请求参数
type RotationInfoSearchReq struct {
	comModel.PageReq
	Id        string `p:"id" dc:"ID"`                                                             //ID
	PicUrl    string `p:"picUrl" dc:"图片"`                                                         //图片
	Link      string `p:"link" dc:"跳转链接"`                                                         //跳转链接
	Sort      string `p:"sort" v:"sort@integer#排序字段需为整数" dc:"排序字段"`                               //排序字段
	CreatedAt string `p:"createdAt" v:"createdAt@datetime#创建时间需为YYYY-MM-DD hh:mm:ss格式" dc:"创建时间"` //创建时间
}

// RotationInfoSearchRes 列表返回结果
type RotationInfoSearchRes struct {
	comModel.ListRes
	List []*RotationInfoListRes `json:"list"`
}

// RotationInfoAddReq 添加操作请求参数
type RotationInfoAddReq struct {
	PicUrl string `p:"picUrl"  dc:"图片"`
	Link   string `p:"link"  dc:"跳转链接"`
	Sort   int    `p:"sort"  dc:"排序字段"`
}

// RotationInfoEditReq 修改操作请求参数
type RotationInfoEditReq struct {
	Id     int    `p:"id" v:"required#主键ID不能为空" dc:"ID"`
	PicUrl string `p:"picUrl"  dc:"图片"`
	Link   string `p:"link"  dc:"跳转链接"`
	Sort   int    `p:"sort"  dc:"排序字段"`
}
