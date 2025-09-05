// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2024-10-23 16:10:12
// 生成路径: internal/app/system/service/sys_attachment.go
// 生成人：gfast
// desc:附件管理
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"
	"github.com/tiger1103/gfast/v3/internal/app/common/model"
)

type ISysAttachment interface {
	List(ctx context.Context, req *model.SysAttachmentSearchReq) (res *model.SysAttachmentSearchRes, err error)
	GetById(ctx context.Context, Id int64) (res *model.SysAttachmentInfoRes, err error)
	GetByMd5(ctx context.Context, md5 string) (res *model.SysAttachmentInfoRes, err error)
	AddUpload(ctx context.Context, req *model.UploadResponse, attr *model.SysAttachmentAddAttribute) (err error)
	Add(ctx context.Context, req *model.SysAttachmentAddReq) (err error)
	Edit(ctx context.Context, req *model.SysAttachmentEditReq) (err error)
	Delete(ctx context.Context, Id []int64) (err error)
	// 附件管理状态修改（状态）
	ChangeStatus(ctx context.Context, id int64, status bool) (err error)
}

var localSysAttachment ISysAttachment

func SysAttachment() ISysAttachment {
	if localSysAttachment == nil {
		panic("implement not found for interface ISysAttachment, forgot register?")
	}
	return localSysAttachment
}

func RegisterSysAttachment(i ISysAttachment) {
	localSysAttachment = i
}
