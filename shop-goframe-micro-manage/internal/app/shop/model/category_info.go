// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2025-09-05 11:52:54
// 生成路径: internal/app/shop/model/category_info.go
// 生成人：王中阳
// desc:商品分类
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
	comModel "github.com/tiger1103/gfast/v3/internal/app/common/model"
)

// CategoryInfoInfoRes is the golang structure for table category_info.
type CategoryInfoInfoRes struct {
	gmeta.Meta     `orm:"table:category_info"`
	Id             int                             `orm:"id,primary" json:"id" dc:"ID"`        // ID
	ParentId       int                             `orm:"parent_id" json:"parentId" dc:"父级id"` // 父级id
	LinkedParentId *LinkedCategoryInfoCategoryInfo `orm:"with:id=parent_id" json:"linkedParentId"`
	Name           string                          `orm:"name" json:"name" dc:"名称"`              // 名称
	PicUrl         string                          `orm:"pic_url" json:"picUrl" dc:"图片"`         // 图片
	Level          int                             `orm:"level" json:"level" dc:"等级"`            // 等级
	Sort           int                             `orm:"sort" json:"sort" dc:"排序"`              // 排序
	CreatedAt      *gtime.Time                     `orm:"created_at" json:"createdAt" dc:"创建时间"` // 创建时间
	UpdatedAt      *gtime.Time                     `orm:"updated_at" json:"updatedAt" dc:""`     //
	DeletedAt      *gtime.Time                     `orm:"deleted_at" json:"deletedAt" dc:""`     //
}
type LinkedCategoryInfoCategoryInfo struct {
	gmeta.Meta `orm:"table:category_info"`
	Id         int    `orm:"id" json:"id" dc:""`     //
	Name       string `orm:"name" json:"name" dc:""` //
}

type CategoryInfoListRes struct {
	Id             int                             `json:"id" dc:"ID"`
	ParentId       int                             `json:"parentId" dc:"父级id"`
	LinkedParentId *LinkedCategoryInfoCategoryInfo `orm:"with:id=parent_id" json:"linkedParentId" dc:"父级id"`
	Name           string                          `json:"name" dc:"名称"`
	PicUrl         string                          `json:"picUrl" dc:"图片"`
	Level          int                             `json:"level" dc:"等级"`
	Sort           int                             `json:"sort" dc:"排序"`
	CreatedAt      *gtime.Time                     `json:"createdAt" dc:"创建时间"`
}

// CategoryInfoSearchReq 分页请求参数
type CategoryInfoSearchReq struct {
	comModel.PageReq
	Id        string `p:"id" dc:"ID"`                                                             //ID
	ParentId  string `p:"parentId" v:"parentId@integer#父级id需为整数" dc:"父级id"`                       //父级id
	Name      string `p:"name" dc:"名称"`                                                           //名称
	PicUrl    string `p:"picUrl" dc:"图片"`                                                         //图片
	Level     string `p:"level" v:"level@integer#等级需为整数" dc:"等级"`                                 //等级
	Sort      string `p:"sort" v:"sort@integer#排序需为整数" dc:"排序"`                                   //排序
	CreatedAt string `p:"createdAt" v:"createdAt@datetime#创建时间需为YYYY-MM-DD hh:mm:ss格式" dc:"创建时间"` //创建时间
}

// CategoryInfoSearchRes 列表返回结果
type CategoryInfoSearchRes struct {
	comModel.ListRes
	List []*CategoryInfoListRes `json:"list"`
}

// 相关连表查询数据
type LinkedCategoryInfoDataSearchRes struct {
	LinkedCategoryInfoCategoryInfo []*LinkedCategoryInfoCategoryInfo `json:"linkedCategoryInfoCategoryInfo"`
}

// CategoryInfoAddReq 添加操作请求参数
type CategoryInfoAddReq struct {
	ParentId int    `p:"parentId"  dc:"父级id"`
	Name     string `p:"name" v:"required#名称不能为空" dc:"名称"`
	PicUrl   string `p:"picUrl"  dc:"图片"`
	Level    int    `p:"level"  dc:"等级"`
	Sort     int    `p:"sort"  dc:"排序"`
}

// CategoryInfoEditReq 修改操作请求参数
type CategoryInfoEditReq struct {
	Id       int    `p:"id" v:"required#主键ID不能为空" dc:"ID"`
	ParentId int    `p:"parentId"  dc:"父级id"`
	Name     string `p:"name" v:"required#名称不能为空" dc:"名称"`
	PicUrl   string `p:"picUrl"  dc:"图片"`
	Level    int    `p:"level"  dc:"等级"`
	Sort     int    `p:"sort"  dc:"排序"`
}
