// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2023-11-09 17:37:25
// 生成路径: internal/app/system/logic/sys_notice_read.go
// 生成人：gfast
// desc:已读记录
// company:云南奇讯科技有限公司
// ==========================================================================

package logic

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterSysNoticeRead(New())
}

func New() service.ISysNoticeRead {
	return &sSysNoticeRead{}
}

type sSysNoticeRead struct{}

func (s *sSysNoticeRead) List(ctx context.Context, req *model.SysNoticeReadSearchReq) (listRes *model.SysNoticeReadSearchRes, err error) {
	listRes = new(model.SysNoticeReadSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysNoticeRead.Ctx(ctx).WithAll()
		listRes.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取总行数失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		listRes.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		order := "id desc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		var res []*model.SysNoticeReadListRes
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.SysNoticeReadListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.SysNoticeReadListRes{
				Id:        v.Id,
				NoticeId:  v.NoticeId,
				UserId:    v.UserId,
				CreatedAt: v.CreatedAt,
			}
		}
	})
	return
}

func (s *sSysNoticeRead) GetById(ctx context.Context, id int64) (res *model.SysNoticeReadInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysNoticeRead.Ctx(ctx).WithAll().Where(dao.SysNoticeRead.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sSysNoticeRead) Add(ctx context.Context, req *model.SysNoticeReadAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysNoticeRead.Ctx(ctx).Insert(do.SysNoticeRead{})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sSysNoticeRead) ReadNotice(ctx context.Context, req *model.SysNoticeReadNoticeReq) (err error) {
	currentUser := service.Context().Get(ctx).User.Id

	err = g.Try(ctx, func(ctx context.Context) {
		//判断是否有数据
		count, err := dao.SysNoticeRead.Ctx(ctx).Where(dao.SysNoticeRead.Columns().NoticeId, req.NoticeId).
			Where(dao.SysNoticeRead.Columns().UserId, currentUser).Count()
		if err != nil {
			return
		}
		if count > 0 {
			_, err := g.DB().Exec(ctx, " update sys_notice_read set clicks=(clicks+1) where user_id =? and notice_id=?", currentUser, req.NoticeId)
			if err != nil {
				return
			}
		} else {
			_, err = dao.SysNoticeRead.Ctx(ctx).Insert(do.SysNoticeRead{
				NoticeId:  req.NoticeId,
				UserId:    currentUser,
				CreatedAt: gtime.Now(),
				Clicks:    0,
			})
			if err != nil {
				return
			}
		}
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}
func (s *sSysNoticeRead) Edit(ctx context.Context, req *model.SysNoticeReadEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysNoticeRead.Ctx(ctx).WherePri(req.Id).Update(do.SysNoticeRead{})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sSysNoticeRead) Delete(ctx context.Context, ids []int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysNoticeRead.Ctx(ctx).Delete(dao.SysNoticeRead.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}
