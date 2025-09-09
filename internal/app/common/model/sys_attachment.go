// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2024-10-23 16:10:12
// 生成路径: internal/app/system/model/sys_attachment.go
// 生成人：gfast
// desc:附件管理
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SysAttachmentInfoRes is the golang structure for table sys_attachment.
type SysAttachmentInfoRes struct {
	gmeta.Meta  `orm:"table:sys_attachment"`
	Id          int64       `orm:"id,primary" json:"id" dc:"文件ID"`         // 文件ID
	AppId       string      `orm:"app_id" json:"appId" dc:"应用ID"`          // 应用ID
	Drive       uint        `orm:"drive" json:"drive" dc:"上传驱动"`           // 上传驱动
	Name        string      `orm:"name" json:"name" dc:"文件原始名"`            // 文件原始名
	Kind        string      `orm:"kind" json:"kind" dc:"上传类型"`             // 上传类型
	MimeType    string      `orm:"mime_type" json:"mimeType" dc:"扩展类型"`    // 扩展类型
	Path        string      `orm:"path" json:"path" dc:"本地路径"`             // 本地路径
	Size        int64       `orm:"size" json:"size" dc:"文件大小"`             // 文件大小
	Ext         string      `orm:"ext" json:"ext" dc:"扩展名"`                // 扩展名
	Md5         string      `orm:"md5" json:"md5" dc:"md5校验码"`             // md5校验码
	CreatedBy   int64       `orm:"created_by" json:"createdBy" dc:"上传人ID"` // 上传人ID
	CreatedUser LinkUserRes `orm:"with:id=created_by" json:"createdUser"`
	Status      bool        `orm:"status" json:"status" dc:"状态"`          // 状态
	CreatedAt   *gtime.Time `orm:"created_at" json:"createdAt" dc:"创建时间"` // 创建时间
	UpdatedAt   *gtime.Time `orm:"updated_at" json:"updatedAt" dc:"修改时间"` // 修改时间
}

type LinkUserRes struct {
	gmeta.Meta   `orm:"table:sys_user"`
	Id           uint64 `orm:"id"       json:"id"`
	UserNickname string `orm:"user_nickname"    json:"userNickname"`
}

type SysAttachmentListRes struct {
	Id        int64       `json:"id" dc:"文件ID"`
	AppId     string      `json:"appId" dc:"应用ID"`
	Drive     uint        `json:"drive" dc:"上传驱动"`
	Name      string      `json:"name" dc:"文件原始名"`
	Kind      string      `json:"kind" dc:"上传类型"`
	Path      string      `json:"path" dc:"本地路径"`
	Size      int64       `json:"size" dc:"文件大小"`
	Ext       string      `json:"ext" dc:"扩展名"`
	Status    bool        `json:"status" dc:"状态"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`
}

// SysAttachmentSearchReq 分页请求参数
type SysAttachmentSearchReq struct {
	PageReq
	AppId     string   `p:"appId" dc:"应用ID"`                                                              //应用ID
	Drive     string   `p:"drive" dc:"上传驱动"`                                                              //上传驱动
	Name      string   `p:"name" dc:"文件原始名"`                                                              //文件原始名
	Kind      string   `p:"kind" dc:"上传类型"`                                                               //上传类型
	MimeType  string   `p:"mimeType" dc:"扩展类型"`                                                           //扩展类型
	Status    string   `p:"status" v:"status@boolean#状态需为true/false" dc:"状态"`                             //状态
	CreatedAt []string `p:"createdAt" v:"createdAt@datetime-array#创建时间需为YYYY-MM-DD hh:mm:ss格式" dc:"创建时间"` //创建时间
}

// SysAttachmentSearchRes 列表返回结果
type SysAttachmentSearchRes struct {
	ListRes
	List []*SysAttachmentListRes `json:"list"`
}

// SysAttachmentAddReq 添加操作请求参数
type SysAttachmentAddReq struct {
	AppId     string `p:"appId" v:"required#应用ID不能为空" dc:"应用ID"`
	Drive     uint   `p:"drive"  dc:"上传驱动"`
	Name      string `p:"name" v:"required#文件原始名不能为空" dc:"文件原始名"`
	Kind      string `p:"kind"  dc:"上传类型"`
	MimeType  string `p:"mimeType"  dc:"扩展类型"`
	Path      string `p:"path"  dc:"本地路径"`
	Size      int64  `p:"size"  dc:"文件大小"`
	Ext       string `p:"ext"  dc:"扩展名"`
	Md5       string `p:"md5"  dc:"md5校验码"`
	Status    bool   `p:"status" v:"required#状态不能为空" dc:"状态"`
	CreatedBy uint64
}

type SysAttachmentAddAttribute struct {
	Md5    string
	Driver uint
	UserId uint64
	AppId  string
}

// SysAttachmentEditReq 修改操作请求参数
type SysAttachmentEditReq struct {
	Id       int64  `p:"id" v:"required#主键ID不能为空" dc:"文件ID"`
	AppId    uint   `p:"appId" v:"required#应用ID不能为空" dc:"应用ID"`
	Drive    string `p:"drive"  dc:"上传驱动"`
	Name     string `p:"name" v:"required#文件原始名不能为空" dc:"文件原始名"`
	Kind     string `p:"kind"  dc:"上传类型"`
	MimeType string `p:"mimeType"  dc:"扩展类型"`
	Path     string `p:"path"  dc:"本地路径"`
	Size     int64  `p:"size"  dc:"文件大小"`
	Ext      string `p:"ext"  dc:"扩展名"`
	Md5      string `p:"md5"  dc:"md5校验码"`
	Status   bool   `p:"status" v:"required#状态不能为空" dc:"状态"`
}
