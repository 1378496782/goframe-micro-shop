// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2023-11-09 11:41:17
// 生成路径: internal/app/system/service/sys_notice.go
// 生成人：gfast
// desc:通知公告
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

type ISysNotice interface {
	List(ctx context.Context, req *model.SysNoticeSearchReq) (res *model.SysNoticeSearchRes, err error)
	ListShow(ctx context.Context, req *model.SysNoticeSearchReq) (res *model.SysNoticeSearchRes, err error)

	GetById(ctx context.Context, Id int64) (res *model.SysNoticeInfoRes, err error)
	Add(ctx context.Context, req *model.SysNoticeAddReq) (err error)
	Edit(ctx context.Context, req *model.SysNoticeEditReq) (err error)
	Delete(ctx context.Context, Id []int64) (err error)
	//IndexData(ctx context.Context) (res *model.SysNoticeIndexRes, err error)
	/*首页相关数据接口*/
	UnReadList(ctx context.Context) (res *model.SysNoticeListRes, err error)
	UnReadCount(ctx context.Context, currentUser uint64) (sysNoticeUnreadIds *model.SysNoticeUnreadCount, err error)
	ReadAll(ctx context.Context, nType string) (err error)
	GetUserNameList(ctx context.Context, search string) (res []*model.SysNoticeUserNickname, err error)
	//NoticeReadLengthAdd(ctx context.Context, id int64) (err error)
	NoticeReadAddUserId(ctx context.Context, req *model.SysNoticeReadAddUserReq) (err error)
	//获取有我的消息的所有私信
	//CurrentUseWithIds(ctx context.Context, currentUserId uint64) (ids []int64, err error)
}

var localSysNotice ISysNotice

func SysNotice() ISysNotice {
	if localSysNotice == nil {
		panic("implement not found for interface ISysNotice, forgot register?")
	}
	return localSysNotice
}

func RegisterSysNotice(i ISysNotice) {
	localSysNotice = i
}
