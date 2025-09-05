// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2023-11-09 17:37:25
// 生成路径: api/v1/system/sys_notice_read.go
// 生成人：gfast
// desc:已读记录相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

// SysNoticeReadSearchReq 分页请求参数
type SysNoticeReadSearchReq struct {
	g.Meta `path:"/list" tags:"系统后台/已读记录" method:"get" summary:"已读记录列表"`
	commonApi.Author
	model.SysNoticeReadSearchReq
}

// SysNoticeReadSearchRes 列表返回结果
type SysNoticeReadSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.SysNoticeReadSearchRes
}

// SysNoticeReadAddReq 添加操作请求参数
type SysNoticeReadAddReq struct {
	g.Meta `path:"/add" tags:"系统后台/已读记录" method:"post" summary:"已读记录添加"`
	commonApi.Author
	*model.SysNoticeReadAddReq
}

// SysNoticeReadAddRes 添加操作返回结果
type SysNoticeReadAddRes struct {
	commonApi.EmptyRes
}

// SysNoticeReadAddReq 添加操作请求参数
type SysNoticeReadNoticeAddReq struct {
	g.Meta `path:"/readNotice" tags:"系统后台/已读记录" method:"post" summary:"已读记录"`
	commonApi.Author
	*model.SysNoticeReadNoticeReq
}

// SysNoticeReadAddRes 添加操作返回结果
type SysNoticeReadNoticeAddRes struct {
	commonApi.EmptyRes
}

// SysNoticeReadEditReq 修改操作请求参数
type SysNoticeReadEditReq struct {
	g.Meta `path:"/edit" tags:"系统后台/已读记录" method:"put" summary:"已读记录修改"`
	commonApi.Author
	*model.SysNoticeReadEditReq
}

// SysNoticeReadEditRes 修改操作返回结果
type SysNoticeReadEditRes struct {
	commonApi.EmptyRes
}

// SysNoticeReadGetReq 获取一条数据请求
type SysNoticeReadGetReq struct {
	g.Meta `path:"/get" tags:"系统后台/已读记录" method:"get" summary:"获取已读记录信息"`
	commonApi.Author
	Id int64 `p:"id" v:"required#主键必须"` //通过主键获取
}

// SysNoticeReadGetRes 获取一条数据结果
type SysNoticeReadGetRes struct {
	g.Meta `mime:"application/json"`
	*model.SysNoticeReadInfoRes
}

// SysNoticeReadDeleteReq 删除数据请求
type SysNoticeReadDeleteReq struct {
	g.Meta `path:"/delete" tags:"系统后台/已读记录" method:"delete" summary:"删除已读记录"`
	commonApi.Author
	Ids []int64 `p:"ids" v:"required#主键必须"` //通过主键删除
}

// SysNoticeReadDeleteRes 删除数据返回
type SysNoticeReadDeleteRes struct {
	commonApi.EmptyRes
}
