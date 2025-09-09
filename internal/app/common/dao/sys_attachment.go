// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2024-10-23 16:10:12
// 生成路径: internal/app/system/dao/sys_attachment.go
// 生成人：gfast
// desc:附件管理
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/common/dao/internal"
)

// sysAttachmentDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type sysAttachmentDao struct {
	*internal.SysAttachmentDao
}

var (
	// SysAttachment is globally public accessible object for table tools_gen_table operations.
	SysAttachment = sysAttachmentDao{
		internal.NewSysAttachmentDao(),
	}
)

// Fill with you ideas below.
