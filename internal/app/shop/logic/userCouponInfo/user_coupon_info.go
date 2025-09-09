// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2025-09-09 15:39:41
// 生成路径: internal/app/shop/logic/user_coupon_info.go
// 生成人：gfast
// desc:用户优惠券
// company:云南奇讯科技有限公司
// ==========================================================================

package userCouponInfo

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
	service.RegisterUserCouponInfo(New())
}

func New() service.IUserCouponInfo {
	return &sUserCouponInfo{}
}

type sUserCouponInfo struct{}

func (s *sUserCouponInfo) List(ctx context.Context, req *model.UserCouponInfoSearchReq) (listRes *model.UserCouponInfoSearchRes, err error) {
	listRes = new(model.UserCouponInfoSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.UserCouponInfo.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.UserCouponInfo.Columns().Id+" = ?", req.Id)
		}
		if req.UserId != "" {
			m = m.Where(dao.UserCouponInfo.Columns().UserId+" = ?", gconv.Int(req.UserId))
		}
		if req.CouponId != "" {
			m = m.Where(dao.UserCouponInfo.Columns().CouponId+" = ?", gconv.Int(req.CouponId))
		}
		if req.Status != "" {
			m = m.Where(dao.UserCouponInfo.Columns().Status+" = ?", gconv.Int(req.Status))
		}
		if req.Amount != "" {
			m = m.Where(dao.UserCouponInfo.Columns().Amount+" = ?", gconv.Int(req.Amount))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.UserCouponInfo.Columns().CreatedAt+" >=? AND "+dao.UserCouponInfo.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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
		var res []*model.UserCouponInfoListRes
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.UserCouponInfoListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.UserCouponInfoListRes{
				Id:             v.Id,
				UserId:         v.UserId,
				CouponId:       v.CouponId,
				LinkedCouponId: v.LinkedCouponId,
				Status:         v.Status,
				Amount:         v.Amount,
				CreatedAt:      v.CreatedAt,
			}
		}
	})
	return
}

func (s *sUserCouponInfo) GetExportData(ctx context.Context, req *model.UserCouponInfoSearchReq) (listRes []*model.UserCouponInfoInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.UserCouponInfo.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.UserCouponInfo.Columns().Id+" = ?", req.Id)
		}
		if req.UserId != "" {
			m = m.Where(dao.UserCouponInfo.Columns().UserId+" = ?", gconv.Int(req.UserId))
		}
		if req.CouponId != "" {
			m = m.Where(dao.UserCouponInfo.Columns().CouponId+" = ?", gconv.Int(req.CouponId))
		}
		if req.Status != "" {
			m = m.Where(dao.UserCouponInfo.Columns().Status+" = ?", gconv.Int(req.Status))
		}
		if req.Amount != "" {
			m = m.Where(dao.UserCouponInfo.Columns().Amount+" = ?", gconv.Int(req.Amount))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.UserCouponInfo.Columns().CreatedAt+" >=? AND "+dao.UserCouponInfo.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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

func (s *sUserCouponInfo) Import(ctx context.Context, file *ghttp.UploadFile) (err error) {
	if file == nil {
		err = errors.New("请上传数据文件")
		return
	}
	var data []do.UserCouponInfo
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
		data = make([]do.UserCouponInfo, len(rows)-1)
		for k, v := range rows {
			if k == 0 {
				continue
			}
			for kv, vv := range v {
				d[kv] = vv
			}
			data[k-1] = do.UserCouponInfo{
				UserId:    gconv.Int64(d[0]),
				CouponId:  gconv.Int64(d[1]),
				Status:    gconv.Int64(d[2]),
				Amount:    gconv.Int64(d[3]),
				CreatedAt: gconv.GTime(d[4]),
				UpdatedAt: gconv.GTime(d[5]),
				DeletedAt: gconv.GTime(d[6]),
			}
		}
		if len(data) > 0 {
			err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				_, err = dao.UserCouponInfo.Ctx(ctx).Batch(500).Insert(data)
				return err
			})
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

// LinkedDataSearch 相关连表查询数据
func (s *sUserCouponInfo) LinkedUserCouponInfoDataSearch(ctx context.Context) (res *model.LinkedUserCouponInfoDataSearchRes, err error) {
	res = new(model.LinkedUserCouponInfoDataSearchRes)
	res.LinkedUserCouponInfoCouponInfo, err = s.ListUserCouponInfoCouponInfo(ctx)
	liberr.ErrIsNil(ctx, err, "获取关联表信息失败")
	return
}

func (s *sUserCouponInfo) GetById(ctx context.Context, id int) (res *model.UserCouponInfoInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.UserCouponInfo.Ctx(ctx).WithAll().Where(dao.UserCouponInfo.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sUserCouponInfo) Add(ctx context.Context, req *model.UserCouponInfoAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.UserCouponInfo.Ctx(ctx).Insert(do.UserCouponInfo{
			UserId:   req.UserId,
			CouponId: req.CouponId,
			Status:   req.Status,
			Amount:   req.Amount,
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sUserCouponInfo) Edit(ctx context.Context, req *model.UserCouponInfoEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.UserCouponInfo.Ctx(ctx).WherePri(req.Id).Update(do.UserCouponInfo{
			UserId:   req.UserId,
			CouponId: req.CouponId,
			Status:   req.Status,
			Amount:   req.Amount,
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sUserCouponInfo) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.UserCouponInfo.Ctx(ctx).Delete(dao.UserCouponInfo.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

func (s *sUserCouponInfo) ListUserCouponInfoCouponInfo(ctx context.Context) (linkedUserCouponInfoCouponInfo []*model.LinkedUserCouponInfoCouponInfo, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.CouponInfo.
			Ctx(ctx).
			Fields(model.LinkedUserCouponInfoCouponInfo{}).Scan(&linkedUserCouponInfoCouponInfo)
		liberr.ErrIsNil(ctx, err)
	})
	return
}
