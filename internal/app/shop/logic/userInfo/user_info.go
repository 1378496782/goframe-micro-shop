// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2025-09-09 15:51:53
// 生成路径: internal/app/shop/logic/user_info.go
// 生成人：gfast
// desc:用户
// company:云南奇讯科技有限公司
// ==========================================================================

package userInfo

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/internal/app/shop/dao"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/shop/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterUserInfo(New())
}

func New() service.IUserInfo {
	return &sUserInfo{}
}

type sUserInfo struct{}

func (s *sUserInfo) List(ctx context.Context, req *model.UserInfoSearchReq) (listRes *model.UserInfoSearchRes, err error) {
	listRes = new(model.UserInfoSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.UserInfo.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.UserInfo.Columns().Id+" = ?", req.Id)
		}
		if req.Name != "" {
			m = m.Where(dao.UserInfo.Columns().Name+" like ?", "%"+req.Name+"%")
		}
		if req.Avatar != "" {
			m = m.Where(dao.UserInfo.Columns().Avatar+" = ?", req.Avatar)
		}
		if req.Sex != "" {
			m = m.Where(dao.UserInfo.Columns().Sex+" = ?", gconv.Int(req.Sex))
		}
		if req.Status != "" {
			m = m.Where(dao.UserInfo.Columns().Status+" = ?", gconv.Int(req.Status))
		}
		if req.Sign != "" {
			m = m.Where(dao.UserInfo.Columns().Sign+" = ?", req.Sign)
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.UserInfo.Columns().CreatedAt+" >=? AND "+dao.UserInfo.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
		}
		listRes.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取总行数失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		listRes.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		order := "id asc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		var res []*model.UserInfoListRes
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.UserInfoListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.UserInfoListRes{
				Id:        v.Id,
				Name:      v.Name,
				Avatar:    v.Avatar,
				Sex:       v.Sex,
				Status:    v.Status,
				Sign:      v.Sign,
				CreatedAt: v.CreatedAt,
			}
		}
	})
	return
}

func (s *sUserInfo) GetExportData(ctx context.Context, req *model.UserInfoSearchReq) (listRes []*model.UserInfoInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.UserInfo.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.UserInfo.Columns().Id+" = ?", req.Id)
		}
		if req.Name != "" {
			m = m.Where(dao.UserInfo.Columns().Name+" like ?", "%"+req.Name+"%")
		}
		if req.Avatar != "" {
			m = m.Where(dao.UserInfo.Columns().Avatar+" = ?", req.Avatar)
		}
		if req.Sex != "" {
			m = m.Where(dao.UserInfo.Columns().Sex+" = ?", gconv.Int(req.Sex))
		}
		if req.Status != "" {
			m = m.Where(dao.UserInfo.Columns().Status+" = ?", gconv.Int(req.Status))
		}
		if req.Sign != "" {
			m = m.Where(dao.UserInfo.Columns().Sign+" = ?", req.Sign)
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.UserInfo.Columns().CreatedAt+" >=? AND "+dao.UserInfo.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
		}
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		order := "id asc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&listRes)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

func (s *sUserInfo) GetById(ctx context.Context, id int) (res *model.UserInfoInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.UserInfo.Ctx(ctx).WithAll().Where(dao.UserInfo.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sUserInfo) Add(ctx context.Context, req *model.UserInfoAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.UserInfo.Ctx(ctx).Insert(do.UserInfo{
			Name:   req.Name,
			Avatar: req.Avatar,
			Sex:    req.Sex,
			Status: req.Status,
			Sign:   req.Sign,
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sUserInfo) Edit(ctx context.Context, req *model.UserInfoEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.UserInfo.Ctx(ctx).WherePri(req.Id).Update(do.UserInfo{
			Name:   req.Name,
			Avatar: req.Avatar,
			Sex:    req.Sex,
			Status: req.Status,
			Sign:   req.Sign,
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sUserInfo) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.UserInfo.Ctx(ctx).Delete(dao.UserInfo.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}
