// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2025-09-05 12:04:35
// 生成路径: internal/app/shop/logic/goods_info.go
// 生成人：王中阳
// desc:商品表
// company:云南奇讯科技有限公司
// ==========================================================================

package goodsInfo

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/internal/app/shop/dao"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/shop/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/library/libUtils"
	"github.com/tiger1103/gfast/v3/library/liberr"
	"github.com/xuri/excelize/v2"
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
		if req.Images != "" {
			m = m.Where(dao.GoodsInfo.Columns().Images+" = ?", req.Images)
		}
		if req.Price != "" {
			m = m.Where(dao.GoodsInfo.Columns().Price+" = ?", gconv.Int(req.Price))
		}
		if req.Level1CategoryId != "" {
			m = m.Where(dao.GoodsInfo.Columns().Level1CategoryId+" = ?", gconv.Int(req.Level1CategoryId))
		}
		if req.Level2CategoryId != "" {
			m = m.Where(dao.GoodsInfo.Columns().Level2CategoryId+" = ?", gconv.Int(req.Level2CategoryId))
		}
		if req.Level3CategoryId != "" {
			m = m.Where(dao.GoodsInfo.Columns().Level3CategoryId+" = ?", gconv.Int(req.Level3CategoryId))
		}
		if req.Brand != "" {
			m = m.Where(dao.GoodsInfo.Columns().Brand+" = ?", req.Brand)
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
		if req.DetailInfo != "" {
			m = m.Where(dao.GoodsInfo.Columns().DetailInfo+" = ?", req.DetailInfo)
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
				Id:                     v.Id,
				Name:                   v.Name,
				Images:                 v.Images,
				Price:                  v.Price,
				Level1CategoryId:       v.Level1CategoryId,
				LinkedLevel1CategoryId: v.LinkedLevel1CategoryId,
				Level2CategoryId:       v.Level2CategoryId,
				LinkedLevel2CategoryId: v.LinkedLevel2CategoryId,
				Level3CategoryId:       v.Level3CategoryId,
				LinkedLevel3CategoryId: v.LinkedLevel3CategoryId,
				Brand:                  v.Brand,
				Stock:                  v.Stock,
				Sale:                   v.Sale,
				Tags:                   v.Tags,
				DetailInfo:             v.DetailInfo,
				CreatedAt:              v.CreatedAt,
			}
		}
	})
	return
}

func (s *sGoodsInfo) GetExportData(ctx context.Context, req *model.GoodsInfoSearchReq) (listRes []*model.GoodsInfoInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.GoodsInfo.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.GoodsInfo.Columns().Id+" = ?", req.Id)
		}
		if req.Name != "" {
			m = m.Where(dao.GoodsInfo.Columns().Name+" like ?", "%"+req.Name+"%")
		}
		if req.Images != "" {
			m = m.Where(dao.GoodsInfo.Columns().Images+" = ?", req.Images)
		}
		if req.Price != "" {
			m = m.Where(dao.GoodsInfo.Columns().Price+" = ?", gconv.Int(req.Price))
		}
		if req.Level1CategoryId != "" {
			m = m.Where(dao.GoodsInfo.Columns().Level1CategoryId+" = ?", gconv.Int(req.Level1CategoryId))
		}
		if req.Level2CategoryId != "" {
			m = m.Where(dao.GoodsInfo.Columns().Level2CategoryId+" = ?", gconv.Int(req.Level2CategoryId))
		}
		if req.Level3CategoryId != "" {
			m = m.Where(dao.GoodsInfo.Columns().Level3CategoryId+" = ?", gconv.Int(req.Level3CategoryId))
		}
		if req.Brand != "" {
			m = m.Where(dao.GoodsInfo.Columns().Brand+" = ?", req.Brand)
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
		if req.DetailInfo != "" {
			m = m.Where(dao.GoodsInfo.Columns().DetailInfo+" = ?", req.DetailInfo)
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.GoodsInfo.Columns().CreatedAt+" >=? AND "+dao.GoodsInfo.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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

func (s *sGoodsInfo) Import(ctx context.Context, file *ghttp.UploadFile) (err error) {
	if file == nil {
		err = errors.New("请上传数据文件")
		return
	}
	var data []do.GoodsInfo
	err = g.Try(ctx, func(ctx context.Context) {
		f, err := file.Open()
		liberr.ErrIsNil(ctx, err)
		defer f.Close()
		exFile, err := excelize.OpenReader(f)
		liberr.ErrIsNil(ctx, err)
		defer exFile.Close()
		rows, err := exFile.GetRows("Sheet1")
		liberr.ErrIsNil(ctx, err)
		if len(rows) == 0 {
			liberr.ErrIsNil(ctx, errors.New("表格内容不能为空"))
		}
		d := make([]interface{}, len(rows[0]))
		data = make([]do.GoodsInfo, len(rows)-1)
		for k, v := range rows {
			if k == 0 {
				continue
			}
			for kv, vv := range v {
				d[kv] = vv
			}
			data[k-1] = do.GoodsInfo{
				Name:             d[0],
				Price:            gconv.Int64(d[1]),
				Level1CategoryId: gconv.Int64(d[2]),
				Level2CategoryId: gconv.Int64(d[3]),
				Level3CategoryId: gconv.Int64(d[4]),
				Brand:            d[5],
				Stock:            gconv.Int64(d[6]),
				Sale:             gconv.Int64(d[7]),
				Tags:             d[8],
				DetailInfo:       d[9],
				CreatedAt:        gconv.GTime(d[10]),
				UpdatedAt:        gconv.GTime(d[11]),
				DeletedAt:        gconv.GTime(d[12]),
			}
		}
		if len(data) > 0 {
			err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				_, err = dao.GoodsInfo.Ctx(ctx).Batch(500).Insert(data)
				return err
			})
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

// LinkedDataSearch 相关连表查询数据
func (s *sGoodsInfo) LinkedGoodsInfoDataSearch(ctx context.Context) (res *model.LinkedGoodsInfoDataSearchRes, err error) {
	res = new(model.LinkedGoodsInfoDataSearchRes)
	res.LinkedGoodsInfoCategoryInfo, err = s.ListGoodsInfoCategoryInfo(ctx)
	liberr.ErrIsNil(ctx, err, "获取关联表信息失败")
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
			Name:             req.Name,
			Images:           req.Images,
			Price:            req.Price,
			Level1CategoryId: req.Level1CategoryId,
			Level2CategoryId: req.Level2CategoryId,
			Level3CategoryId: req.Level3CategoryId,
			Brand:            req.Brand,
			Stock:            req.Stock,
			Sale:             req.Sale,
			Tags:             req.Tags,
			DetailInfo:       req.DetailInfo,
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
			Name:             req.Name,
			Images:           req.Images,
			Price:            req.Price,
			Level1CategoryId: req.Level1CategoryId,
			Level2CategoryId: req.Level2CategoryId,
			Level3CategoryId: req.Level3CategoryId,
			Brand:            req.Brand,
			Stock:            req.Stock,
			Sale:             req.Sale,
			Tags:             req.Tags,
			DetailInfo:       req.DetailInfo,
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

func (s *sGoodsInfo) ListGoodsInfoCategoryInfo(ctx context.Context) (linkedGoodsInfoCategoryInfo []*model.LinkedGoodsInfoCategoryInfo, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.CategoryInfo.
			Ctx(ctx).
			Fields(model.LinkedGoodsInfoCategoryInfo{}).Scan(&linkedGoodsInfoCategoryInfo)
		liberr.ErrIsNil(ctx, err)
	})
	return
}
