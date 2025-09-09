// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2025-09-09 15:10:49
// 生成路径: internal/app/shop/logic/coupon_info.go
// 生成人：gfast
// desc:优惠券
// company:云南奇讯科技有限公司
// ==========================================================================

package couponInfo

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
	service.RegisterCouponInfo(New())
}

func New() service.ICouponInfo {
	return &sCouponInfo{}
}

type sCouponInfo struct{}

func (s *sCouponInfo) List(ctx context.Context, req *model.CouponInfoSearchReq) (listRes *model.CouponInfoSearchRes, err error) {
	listRes = new(model.CouponInfoSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.CouponInfo.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.CouponInfo.Columns().Id+" = ?", req.Id)
		}
		if req.Name != "" {
			m = m.Where(dao.CouponInfo.Columns().Name+" like ?", "%"+req.Name+"%")
		}
		if req.Type != "" {
			m = m.Where(dao.CouponInfo.Columns().Type+" = ?", gconv.Int(req.Type))
		}
		if req.Amount != "" {
			m = m.Where(dao.CouponInfo.Columns().Amount+" = ?", gconv.Int(req.Amount))
		}
		if req.Deadline != "" {
			m = m.Where(dao.CouponInfo.Columns().Deadline+" = ?", gconv.Time(req.Deadline))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.CouponInfo.Columns().CreatedAt+" >=? AND "+dao.CouponInfo.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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
		order := "goods_id asc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		var res []*model.CouponInfoListRes
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.CouponInfoListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.CouponInfoListRes{
				Id:        v.Id,
				Name:      v.Name,
				Type:      v.Type,
				Amount:    v.Amount,
				Deadline:  v.Deadline,
				CreatedAt: v.CreatedAt,
			}
		}
	})
	return
}

func (s *sCouponInfo) GetExportData(ctx context.Context, req *model.CouponInfoSearchReq) (listRes []*model.CouponInfoInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.CouponInfo.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.CouponInfo.Columns().Id+" = ?", req.Id)
		}
		if req.Name != "" {
			m = m.Where(dao.CouponInfo.Columns().Name+" like ?", "%"+req.Name+"%")
		}
		if req.Type != "" {
			m = m.Where(dao.CouponInfo.Columns().Type+" = ?", gconv.Int(req.Type))
		}
		if req.Amount != "" {
			m = m.Where(dao.CouponInfo.Columns().Amount+" = ?", gconv.Int(req.Amount))
		}
		if req.Deadline != "" {
			m = m.Where(dao.CouponInfo.Columns().Deadline+" = ?", gconv.Time(req.Deadline))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.CouponInfo.Columns().CreatedAt+" >=? AND "+dao.CouponInfo.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
		}
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		order := "goods_id asc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&listRes)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

func (s *sCouponInfo) Import(ctx context.Context, file *ghttp.UploadFile) (err error) {
	if file == nil {
		err = errors.New("请上传数据文件")
		return
	}
	var data []do.CouponInfo
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
		data = make([]do.CouponInfo, len(rows)-1)
		for k, v := range rows {
			if k == 0 {
				continue
			}
			for kv, vv := range v {
				d[kv] = vv
			}
			data[k-1] = do.CouponInfo{
				GoodsId:   gconv.Int64(d[0]),
				Name:      d[1],
				Type:      gconv.Int64(d[2]),
				Amount:    gconv.Int64(d[3]),
				Deadline:  gconv.GTime(d[4]),
				CreatedAt: gconv.GTime(d[5]),
				UpdatedAt: gconv.GTime(d[6]),
				DeletedAt: gconv.GTime(d[7]),
			}
		}
		if len(data) > 0 {
			err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				_, err = dao.CouponInfo.Ctx(ctx).Batch(500).Insert(data)
				return err
			})
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

func (s *sCouponInfo) GetById(ctx context.Context, id int) (res *model.CouponInfoInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.CouponInfo.Ctx(ctx).WithAll().Where(dao.CouponInfo.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sCouponInfo) Add(ctx context.Context, req *model.CouponInfoAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.CouponInfo.Ctx(ctx).Insert(do.CouponInfo{
			Name:     req.Name,
			Type:     req.Type,
			Amount:   req.Amount,
			Deadline: req.Deadline,
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sCouponInfo) Edit(ctx context.Context, req *model.CouponInfoEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.CouponInfo.Ctx(ctx).WherePri(req.Id).Update(do.CouponInfo{
			Name:     req.Name,
			Type:     req.Type,
			Amount:   req.Amount,
			Deadline: req.Deadline,
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sCouponInfo) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.CouponInfo.Ctx(ctx).Delete(dao.CouponInfo.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}
