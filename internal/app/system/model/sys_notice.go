// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2023-11-09 11:41:17
// 生成路径: internal/app/system/model/sys_notice.go
// 生成人：gfast
// desc:通知公告
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
	comModel "github.com/tiger1103/gfast/v3/internal/app/common/model"
)

// SysNoticeInfoRes is the golang structure for table sys_notice.
type SysNoticeInfoRes struct {
	gmeta.Meta   `orm:"table:sys_notice"`
	Id           int64          `orm:"id,primary" json:"id" dc:"ID"`         // ID
	Title        string         `orm:"title" json:"title" dc:"标题"`           // 标题
	Type         int64          `orm:"type" json:"type" dc:"类型"`             // 类型
	Tag          int            `orm:"tag" json:"tag" dc:"标签"`               // 标签
	Content      string         `orm:"content" json:"content" dc:"内容"`       // 内容
	Remark       string         `orm:"remark" json:"remark" dc:"备注"`         // 备注
	Sort         int            `orm:"sort" json:"sort" dc:"排序"`             // 排序
	Status       int            `orm:"status" json:"status" dc:"状态"`         // 状态
	CreatedBy    int64          `orm:"created_by" json:"createdBy" dc:"发送人"` // 发送人
	CreatedUser  *LinkUserRes   `orm:"with:id=created_by" json:"createdUser"`
	UpdatedBy    int64          `orm:"updated_by" json:"updatedBy" dc:"修改人"` // 修改人
	UpdatedUser  *LinkUserRes   `orm:"with:id=updated_by" json:"updatedUser"`
	CreatedAt    *gtime.Time    `orm:"created_at" json:"createdAt" dc:"创建时间"` // 创建时间
	UpdatedAt    *gtime.Time    `orm:"updated_at" json:"updatedAt" dc:"更新时间"` // 更新时间
	DeletedAt    *gtime.Time    `orm:"deleted_at" json:"deletedAt" dc:"删除时间"` // 删除时间
	Receiver     []uint64       `json:"receiver" orm:"receiver"`
	ReceiverUser []*LinkUserRes `orm:"with:id=receiver" json:"receiverUser"`
}

type SysNoticeListRes struct {
	Id          int64        `json:"id" dc:"ID"`
	Title       string       `json:"title" dc:"标题"`
	Type        int64        `json:"type" dc:"类型"`
	Tag         int          `json:"tag" dc:"标签"`
	Content     string       `json:"content" dc:"内容"`
	Remark      string       `json:"remark" dc:"备注"`
	Sort        int          `json:"sort" dc:"排序"`
	Status      int          `json:"status" dc:"状态"`
	IsRead      bool         `json:"isRead" dc:"当前用户是否已读"`
	ClickNumber int          `json:"clickNumber"`
	CreatedUser *LinkUserRes `orm:"with:id=created_by" json:"createdUser"`
	CreatedBy   int64        `json:"createdBy" dc:"发送人"`
	CreatedAt   *gtime.Time  `json:"createdAt" dc:"创建时间"`
}

// SysNoticeSearchReq 分页请求参数
type SysNoticeSearchReq struct {
	comModel.PageReq
	Id        string `p:"id" dc:"ID"`                                                             //ID
	Title     string `p:"title" dc:"标题"`                                                          //标题
	Type      string `p:"type" v:"type@integer#类型需为整数" dc:"类型"`                                   //类型
	Tag       string `p:"tag" v:"tag@integer#标签需为整数" dc:"标签"`                                     //标签
	Status    string `p:"status" v:"status@integer#状态需为整数" dc:"状态"`                               //状态
	CreatedAt string `p:"createdAt" v:"createdAt@datetime#创建时间需为YYYY-MM-DD hh:mm:ss格式" dc:"创建时间"` //创建时间
	IsTrim    bool   `p:"isTrim"`
}

// SysNoticeSearchRes 列表返回结果
type SysNoticeSearchRes struct {
	comModel.ListRes
	List []*SysNoticeListRes `json:"list"`
}

type SysNoticeUserNickname struct {
	Id           int64  `json:"id"`
	UserNickName string `json:"userNickname"`
}

// SysNoticeAddReq 添加操作请求参数
type SysNoticeAddReq struct {
	Title     string   `p:"title" v:"required#标题不能为空" dc:"标题"`
	Type      int64    `p:"type" v:"required#类型不能为空" dc:"类型"`
	Tag       int      `p:"tag"  dc:"标签"`
	Content   string   `p:"content" v:"required#内容不能为空" dc:"内容"`
	Remark    string   `p:"remark"  dc:"备注"`
	Sort      int      `p:"sort"  dc:"排序"`
	Status    int      `p:"status" v:"required#状态不能为空" dc:"状态"`
	Receiver  []uint64 `p:"receiver"`
	CreatedBy uint64
}

// SysNoticeEditReq 修改操作请求参数
type SysNoticeEditReq struct {
	Id        int64    `p:"id" v:"required#主键ID不能为空" dc:"ID"`
	Title     string   `p:"title" v:"required#标题不能为空" dc:"标题"`
	Type      int64    `p:"type" v:"required#类型不能为空" dc:"类型"`
	Tag       int      `p:"tag"  dc:"标签"`
	Content   string   `p:"content" v:"required#内容不能为空" dc:"内容"`
	Remark    string   `p:"remark"  dc:"备注"`
	Sort      int      `p:"sort"  dc:"排序"`
	Status    int      `p:"status" v:"required#状态不能为空" dc:"状态"`
	Receiver  []uint64 `p:"receiver"`
	UpdatedBy uint64
}

type SysNoticeIndexRes struct {
	Type1Total int               `json:"type1Total"`
	Type2Total int               `json:"type2Total"`
	Type1List  *SysNoticeListRes `json:"type1List"`
	Type2List  *SysNoticeListRes `json:"type2List"`
}

type SysNoticeUnreadCount struct {
	NotifyCount int `json:"notifyCount"`
	NoticeCount int `json:"noticeCount"`
}

type SysNoticeReadAddUserReq struct {
	UserId   int64
	NoticeId int64 `p:"noticeId"`
}
