// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2023-11-09 17:37:25
// 生成路径: internal/app/system/service/sys_notice_read.go
// 生成人：gfast
// desc:已读记录
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"

	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

type ISysNoticeRead interface {
	List(ctx context.Context, req *model.SysNoticeReadSearchReq) (res *model.SysNoticeReadSearchRes, err error)
	ReadNotice(ctx context.Context, req *model.SysNoticeReadNoticeReq) (err error)
	GetById(ctx context.Context, Id int64) (res *model.SysNoticeReadInfoRes, err error)
	Add(ctx context.Context, req *model.SysNoticeReadAddReq) (err error)
	Edit(ctx context.Context, req *model.SysNoticeReadEditReq) (err error)
	Delete(ctx context.Context, Id []int64) (err error)
}

var localSysNoticeRead ISysNoticeRead

func SysNoticeRead() ISysNoticeRead {
	if localSysNoticeRead == nil {
		panic("implement not found for interface ISysNoticeRead, forgot register?")
	}
	return localSysNoticeRead
}

func RegisterSysNoticeRead(i ISysNoticeRead) {
	localSysNoticeRead = i
}
