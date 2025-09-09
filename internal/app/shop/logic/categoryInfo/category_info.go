// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2025-09-05 11:52:54
// 生成路径: internal/app/shop/logic/category_info.go
// 生成人：王中阳
// desc:商品分类
// company:云南奇讯科技有限公司
// ==========================================================================

package categoryInfo

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
	"github.com/tiger1103/gfast/v3/library/liberr"
	"github.com/xuri/excelize/v2"
)

func init() {
	service.RegisterCategoryInfo(New())
}

func New() service.ICategoryInfo {
	return &sCategoryInfo{}
}

type sCategoryInfo struct{}

func (s *sCategoryInfo) List(ctx context.Context, req *model.CategoryInfoSearchReq) (listRes *model.CategoryInfoSearchRes, err error) {
	listRes = new(model.CategoryInfoSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.CategoryInfo.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.CategoryInfo.Columns().Id+" = ?", req.Id)
		}
		if req.ParentId != "" {
			m = m.Where(dao.CategoryInfo.Columns().ParentId+" = ?", gconv.Int(req.ParentId))
		}
		if req.Name != "" {
			m = m.Where(dao.CategoryInfo.Columns().Name+" like ?", "%"+req.Name+"%")
		}
		if req.PicUrl != "" {
			m = m.Where(dao.CategoryInfo.Columns().PicUrl+" = ?", req.PicUrl)
		}
		if req.Level != "" {
			m = m.Where(dao.CategoryInfo.Columns().Level+" = ?", gconv.Int(req.Level))
		}
		if req.Sort != "" {
			m = m.Where(dao.CategoryInfo.Columns().Sort+" = ?", gconv.Int(req.Sort))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.CategoryInfo.Columns().CreatedAt+" >=? AND "+dao.CategoryInfo.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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
		var res []*model.CategoryInfoListRes
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.CategoryInfoListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.CategoryInfoListRes{
				Id:             v.Id,
				ParentId:       v.ParentId,
				LinkedParentId: v.LinkedParentId,
				Name:           v.Name,
				PicUrl:         v.PicUrl,
				Level:          v.Level,
				Sort:           v.Sort,
				CreatedAt:      v.CreatedAt,
			}
		}
	})
	return
}

func (s *sCategoryInfo) GetExportData(ctx context.Context, req *model.CategoryInfoSearchReq) (listRes []*model.CategoryInfoInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.CategoryInfo.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.CategoryInfo.Columns().Id+" = ?", req.Id)
		}
		if req.ParentId != "" {
			m = m.Where(dao.CategoryInfo.Columns().ParentId+" = ?", gconv.Int(req.ParentId))
		}
		if req.Name != "" {
			m = m.Where(dao.CategoryInfo.Columns().Name+" like ?", "%"+req.Name+"%")
		}
		if req.PicUrl != "" {
			m = m.Where(dao.CategoryInfo.Columns().PicUrl+" = ?", req.PicUrl)
		}
		if req.Level != "" {
			m = m.Where(dao.CategoryInfo.Columns().Level+" = ?", gconv.Int(req.Level))
		}
		if req.Sort != "" {
			m = m.Where(dao.CategoryInfo.Columns().Sort+" = ?", gconv.Int(req.Sort))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.CategoryInfo.Columns().CreatedAt+" >=? AND "+dao.CategoryInfo.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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

func (s *sCategoryInfo) Import(ctx context.Context, file *ghttp.UploadFile) (err error) {
	if file == nil {
		err = errors.New("请上传数据文件")
		return
	}
	var data []do.CategoryInfo
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
		data = make([]do.CategoryInfo, len(rows)-1)
		for k, v := range rows {
			if k == 0 {
				continue
			}
			for kv, vv := range v {
				d[kv] = vv
			}
			data[k-1] = do.CategoryInfo{
				ParentId:  gconv.Int64(d[0]),
				Name:      d[1],
				Level:     gconv.Int64(d[2]),
				Sort:      gconv.Int64(d[3]),
				CreatedAt: gconv.GTime(d[4]),
				UpdatedAt: gconv.GTime(d[5]),
				DeletedAt: gconv.GTime(d[6]),
			}
		}
		if len(data) > 0 {
			err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				_, err = dao.CategoryInfo.Ctx(ctx).Batch(500).Insert(data)
				return err
			})
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

// LinkedDataSearch 相关连表查询数据
func (s *sCategoryInfo) LinkedCategoryInfoDataSearch(ctx context.Context) (res *model.LinkedCategoryInfoDataSearchRes, err error) {
	res = new(model.LinkedCategoryInfoDataSearchRes)
	res.LinkedCategoryInfoCategoryInfo, err = s.ListCategoryInfoCategoryInfo(ctx)
	liberr.ErrIsNil(ctx, err, "获取关联表信息失败")
	return
}

func (s *sCategoryInfo) GetById(ctx context.Context, id int) (res *model.CategoryInfoInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.CategoryInfo.Ctx(ctx).WithAll().Where(dao.CategoryInfo.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sCategoryInfo) Add(ctx context.Context, req *model.CategoryInfoAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.CategoryInfo.Ctx(ctx).Insert(do.CategoryInfo{
			ParentId: req.ParentId,
			Name:     req.Name,
			PicUrl:   req.PicUrl,
			Level:    req.Level,
			Sort:     req.Sort,
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sCategoryInfo) Edit(ctx context.Context, req *model.CategoryInfoEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.CategoryInfo.Ctx(ctx).WherePri(req.Id).Update(do.CategoryInfo{
			ParentId: req.ParentId,
			Name:     req.Name,
			PicUrl:   req.PicUrl,
			Level:    req.Level,
			Sort:     req.Sort,
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sCategoryInfo) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.CategoryInfo.Ctx(ctx).Delete(dao.CategoryInfo.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

func (s *sCategoryInfo) ListCategoryInfoCategoryInfo(ctx context.Context) (linkedCategoryInfoCategoryInfo []*model.LinkedCategoryInfoCategoryInfo, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.CategoryInfo.
			Ctx(ctx).
			Fields(model.LinkedCategoryInfoCategoryInfo{}).Scan(&linkedCategoryInfoCategoryInfo)
		liberr.ErrIsNil(ctx, err)
	})
	return
}
