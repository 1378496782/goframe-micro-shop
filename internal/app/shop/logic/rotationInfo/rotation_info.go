// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2025-09-22 16:50:04
// 生成路径: internal/app/shop/logic/rotation_info.go
// 生成人：gfast
// desc:轮播图
// company:云南奇讯科技有限公司
// ==========================================================================

package rotationInfo

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
	service.RegisterRotationInfo(New())
}

func New() service.IRotationInfo {
	return &sRotationInfo{}
}

type sRotationInfo struct{}

func (s *sRotationInfo) List(ctx context.Context, req *model.RotationInfoSearchReq) (listRes *model.RotationInfoSearchRes, err error) {
	listRes = new(model.RotationInfoSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.RotationInfo.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.RotationInfo.Columns().Id+" = ?", req.Id)
		}
		if req.PicUrl != "" {
			m = m.Where(dao.RotationInfo.Columns().PicUrl+" = ?", req.PicUrl)
		}
		if req.Link != "" {
			m = m.Where(dao.RotationInfo.Columns().Link+" = ?", req.Link)
		}
		if req.Sort != "" {
			m = m.Where(dao.RotationInfo.Columns().Sort+" = ?", gconv.Int(req.Sort))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.RotationInfo.Columns().CreatedAt+" >=? AND "+dao.RotationInfo.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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
		var res []*model.RotationInfoListRes
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.RotationInfoListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.RotationInfoListRes{
				Id:        v.Id,
				PicUrl:    v.PicUrl,
				Link:      v.Link,
				Sort:      v.Sort,
				CreatedAt: v.CreatedAt,
			}
		}
	})
	return
}

func (s *sRotationInfo) GetById(ctx context.Context, id int) (res *model.RotationInfoInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.RotationInfo.Ctx(ctx).WithAll().Where(dao.RotationInfo.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sRotationInfo) Add(ctx context.Context, req *model.RotationInfoAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.RotationInfo.Ctx(ctx).Insert(do.RotationInfo{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sRotationInfo) Edit(ctx context.Context, req *model.RotationInfoEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.RotationInfo.Ctx(ctx).WherePri(req.Id).Update(do.RotationInfo{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sRotationInfo) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.RotationInfo.Ctx(ctx).Delete(dao.RotationInfo.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}
