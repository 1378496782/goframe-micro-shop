// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2023-11-09 11:41:17
// 生成路径: internal/app/system/dao/sys_notice.go
// 生成人：gfast
// desc:通知公告
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)

// sysNoticeDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type sysNoticeDao struct {
	*internal.SysNoticeDao
}

var (
	// SysNotice is globally public accessible object for table tools_gen_table operations.
	SysNotice = sysNoticeDao{
		internal.NewSysNoticeDao(),
	}
)

// Fill with you ideas below.
