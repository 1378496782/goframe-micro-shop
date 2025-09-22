// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2025-09-22 16:48:52
// 生成路径: internal/app/shop/logic/goods_info.go
// 生成人：gfast
// desc:商品
// company:云南奇讯科技有限公司
// ==========================================================================

package goodsInfo

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/internal/app/shop/dao"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/shop/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/library/libUtils"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterGoodsInfo(New())
}

func New() service.IGoodsInfo {
	return &sGoodsInfo{}
}

type sGoodsInfo struct{}

func (s *sGoodsInfo) List(ctx context.Context, req *model.GoodsInfoSearchReq) (listRes *model.GoodsInfoSearchRes, err error) {
	listRes = new(model.GoodsInfoSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.GoodsInfo.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.GoodsInfo.Columns().Id+" = ?", req.Id)
		}
		if req.Name != "" {
			m = m.Where(dao.GoodsInfo.Columns().Name+" like ?", "%"+req.Name+"%")
		}
		if req.PicUrl != "" {
			m = m.Where(dao.GoodsInfo.Columns().PicUrl+" = ?", req.PicUrl)
		}
		if req.Images != "" {
			m = m.Where(dao.GoodsInfo.Columns().Images+" = ?", req.Images)
		}
		if req.Price != "" {
			m = m.Where(dao.GoodsInfo.Columns().Price+" = ?", gconv.Int(req.Price))
		}
		if req.Stock != "" {
			m = m.Where(dao.GoodsInfo.Columns().Stock+" = ?", gconv.Int(req.Stock))
		}
		if req.Sale != "" {
			m = m.Where(dao.GoodsInfo.Columns().Sale+" = ?", gconv.Int(req.Sale))
		}
		if req.Tags != "" {
			m = m.Where(dao.GoodsInfo.Columns().Tags+" = ?", req.Tags)
		}
		if req.Sort != "" {
			m = m.Where(dao.GoodsInfo.Columns().Sort+" = ?", gconv.Int(req.Sort))
		}
		if req.EnableBargain != "" {
			m = m.Where(dao.GoodsInfo.Columns().EnableBargain+" = ?", gconv.Int(req.EnableBargain))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.GoodsInfo.Columns().CreatedAt+" >=? AND "+dao.GoodsInfo.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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
		var res []*model.GoodsInfoListRes
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.GoodsInfoListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.GoodsInfoListRes{
				Id:            v.Id,
				Name:          v.Name,
				PicUrl:        v.PicUrl,
				Images:        v.Images,
				Price:         v.Price,
				Stock:         v.Stock,
				Sale:          v.Sale,
				Tags:          v.Tags,
				Sort:          v.Sort,
				EnableBargain: v.EnableBargain,
				CreatedAt:     v.CreatedAt,
			}
		}
	})
	return
}

func (s *sGoodsInfo) GetById(ctx context.Context, id uint) (res *model.GoodsInfoInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.GoodsInfo.Ctx(ctx).WithAll().Where(dao.GoodsInfo.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sGoodsInfo) Add(ctx context.Context, req *model.GoodsInfoAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		for _, obj := range req.Images {
			obj.Url, err = libUtils.GetFilesPath(ctx, obj.Url)
			liberr.ErrIsNil(ctx, err)
		}
		_, err = dao.GoodsInfo.Ctx(ctx).Insert(do.GoodsInfo{
			Name:          req.Name,
			PicUrl:        req.PicUrl,
			Images:        req.Images,
			Price:         req.Price,
			Stock:         req.Stock,
			Sale:          req.Sale,
			Tags:          req.Tags,
			Sort:          req.Sort,
			EnableBargain: req.EnableBargain,
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sGoodsInfo) Edit(ctx context.Context, req *model.GoodsInfoEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		for _, obj := range req.Images {
			obj.Url, err = libUtils.GetFilesPath(ctx, obj.Url)
			liberr.ErrIsNil(ctx, err)
		}
		_, err = dao.GoodsInfo.Ctx(ctx).WherePri(req.Id).Update(do.GoodsInfo{
			Name:          req.Name,
			PicUrl:        req.PicUrl,
			Images:        req.Images,
			Price:         req.Price,
			Stock:         req.Stock,
			Sale:          req.Sale,
			Tags:          req.Tags,
			Sort:          req.Sort,
			EnableBargain: req.EnableBargain,
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sGoodsInfo) Delete(ctx context.Context, ids []uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.GoodsInfo.Ctx(ctx).Delete(dao.GoodsInfo.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}
