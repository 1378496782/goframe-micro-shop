// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2025-10-11 23:55:29
// 生成路径: internal/app/shop/logic/order_goods_info.go
// 生成人：gfast
// desc:订单物品表
// company:云南奇讯科技有限公司
// ==========================================================================

package orderGoodsInfo

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/internal/app/shop/dao"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/shop/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/shop/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterOrderGoodsInfo(New())
}

func New() service.IOrderGoodsInfo {
	return &sOrderGoodsInfo{}
}

type sOrderGoodsInfo struct{}

func (s *sOrderGoodsInfo) List(ctx context.Context, req *model.OrderGoodsInfoSearchReq) (listRes *model.OrderGoodsInfoSearchRes, err error) {
	listRes = new(model.OrderGoodsInfoSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.OrderGoodsInfo.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.OrderGoodsInfo.Columns().Id+" = ?", req.Id)
		}
		if req.OrderId != "" {
			m = m.Where(dao.OrderGoodsInfo.Columns().OrderId+" = ?", gconv.Int(req.OrderId))
		}
		if req.GoodsId != "" {
			m = m.Where(dao.OrderGoodsInfo.Columns().GoodsId+" = ?", gconv.Int(req.GoodsId))
		}
		if req.GoodsOptionsId != "" {
			m = m.Where(dao.OrderGoodsInfo.Columns().GoodsOptionsId+" = ?", gconv.Int(req.GoodsOptionsId))
		}
		if req.Count != "" {
			m = m.Where(dao.OrderGoodsInfo.Columns().Count+" = ?", gconv.Int(req.Count))
		}
		if req.Price != "" {
			m = m.Where(dao.OrderGoodsInfo.Columns().Price+" = ?", gconv.Int(req.Price))
		}
		if req.CouponPrice != "" {
			m = m.Where(dao.OrderGoodsInfo.Columns().CouponPrice+" = ?", gconv.Int(req.CouponPrice))
		}
		if req.ActualPrice != "" {
			m = m.Where(dao.OrderGoodsInfo.Columns().ActualPrice+" = ?", gconv.Int(req.ActualPrice))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.OrderGoodsInfo.Columns().CreatedAt+" >=? AND "+dao.OrderGoodsInfo.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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
		var res []*model.OrderGoodsInfoListRes
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.OrderGoodsInfoListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.OrderGoodsInfoListRes{
				Id:             v.Id,
				OrderId:        v.OrderId,
				GoodsId:        v.GoodsId,
				GoodsOptionsId: v.GoodsOptionsId,
				Count:          v.Count,
				Remark:         v.Remark,
				Price:          v.Price,
				CouponPrice:    v.CouponPrice,
				ActualPrice:    v.ActualPrice,
				CreatedAt:      v.CreatedAt,
			}
		}
	})
	return
}

func (s *sOrderGoodsInfo) GetById(ctx context.Context, id int) (res *model.OrderGoodsInfoInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.OrderGoodsInfo.Ctx(ctx).WithAll().Where(dao.OrderGoodsInfo.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sOrderGoodsInfo) Add(ctx context.Context, req *model.OrderGoodsInfoAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.OrderGoodsInfo.Ctx(ctx).Insert(do.OrderGoodsInfo{
			OrderId:        req.OrderId,
			GoodsId:        req.GoodsId,
			GoodsOptionsId: req.GoodsOptionsId,
			Count:          req.Count,
			Remark:         req.Remark,
			Price:          req.Price,
			CouponPrice:    req.CouponPrice,
			ActualPrice:    req.ActualPrice,
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sOrderGoodsInfo) Edit(ctx context.Context, req *model.OrderGoodsInfoEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.OrderGoodsInfo.Ctx(ctx).WherePri(req.Id).Update(do.OrderGoodsInfo{
			OrderId:        req.OrderId,
			GoodsId:        req.GoodsId,
			GoodsOptionsId: req.GoodsOptionsId,
			Count:          req.Count,
			Remark:         req.Remark,
			Price:          req.Price,
			CouponPrice:    req.CouponPrice,
			ActualPrice:    req.ActualPrice,
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sOrderGoodsInfo) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.OrderGoodsInfo.Ctx(ctx).Delete(dao.OrderGoodsInfo.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

// GetByOrderId 根据订单ID获取订单商品列表
func (s *sOrderGoodsInfo) GetByOrderId(ctx context.Context, orderId int) (list []*model.OrderGoodsInfoListRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.OrderGoodsInfo.Ctx(ctx).
			Where(dao.OrderGoodsInfo.Columns().OrderId, orderId).
			Scan(&list)
		liberr.ErrIsNil(ctx, err, "获取订单商品列表失败")
	})
	return
}

// GetOrderGoodsDetail 获取订单商品详情（包含商品信息）
func (s *sOrderGoodsInfo) GetOrderGoodsDetail(ctx context.Context, id int) (res *model.OrderGoodsDetailRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 获取订单商品信息
		var orderGoods *entity.OrderGoodsInfo
		err = dao.OrderGoodsInfo.Ctx(ctx).Where(dao.OrderGoodsInfo.Columns().Id, id).Scan(&orderGoods)
		liberr.ErrIsNil(ctx, err, "获取订单商品信息失败")
		if orderGoods == nil {
			liberr.ErrIsNil(ctx, gerror.New("订单商品不存在"), "订单商品不存在")
		}
		
		// 获取商品信息
		var goods *entity.GoodsInfo
		err = dao.GoodsInfo.Ctx(ctx).Where(dao.GoodsInfo.Columns().Id, orderGoods.GoodsId).Scan(&goods)
		liberr.ErrIsNil(ctx, err, "获取商品信息失败")
		
		// 构建返回结果
		res = &model.OrderGoodsDetailRes{
			Id:             orderGoods.Id,
			OrderId:        orderGoods.OrderId,
			GoodsId:        orderGoods.GoodsId,
			GoodsName:      goods.Name,
			GoodsImage:     goods.PicUrl,
			GoodsOptionsId: orderGoods.GoodsOptionsId,
			Count:          orderGoods.Count,
			Remark:         orderGoods.Remark,
			Price:          orderGoods.Price,
			CouponPrice:    orderGoods.CouponPrice,
			ActualPrice:    orderGoods.ActualPrice,
			CreatedAt:      orderGoods.CreatedAt,
		}
	})
	return
}

// AddOrderGoods 添加订单商品，同时更新订单总金额
func (s *sOrderGoodsInfo) AddOrderGoods(ctx context.Context, req *model.OrderGoodsAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 开启事务
		err = dao.OrderGoodsInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			// 插入订单商品
			_, err = dao.OrderGoodsInfo.Ctx(ctx).TX(tx).Insert(do.OrderGoodsInfo{
				OrderId:        req.OrderId,
				GoodsId:        req.GoodsId,
				GoodsOptionsId: req.GoodsOptionsId,
				Count:          req.Count,
				Remark:         req.Remark,
				Price:          req.Price,
				CouponPrice:    req.CouponPrice,
				ActualPrice:    req.ActualPrice,
			})
			if err != nil {
				return err
			}
			
			// 计算订单总金额
			totalPrice := req.ActualPrice * req.Count
			couponTotalPrice := req.CouponPrice * req.Count
			actualTotalPrice := req.ActualPrice * req.Count
			
			// 更新订单总金额
			_, err = dao.OrderInfo.Ctx(ctx).TX(tx).
				Where(dao.OrderInfo.Columns().Id, req.OrderId).
				Increment(dao.OrderInfo.Columns().Price, totalPrice)
			if err != nil {
				return err
			}
			
			_, err = dao.OrderInfo.Ctx(ctx).TX(tx).
				Where(dao.OrderInfo.Columns().Id, req.OrderId).
				Increment(dao.OrderInfo.Columns().CouponPrice, couponTotalPrice)
			if err != nil {
				return err
			}
			
			_, err = dao.OrderInfo.Ctx(ctx).TX(tx).
				Where(dao.OrderInfo.Columns().Id, req.OrderId).
				Increment(dao.OrderInfo.Columns().ActualPrice, actualTotalPrice)
			if err != nil {
				return err
			}
			
			return nil
		})
		liberr.ErrIsNil(ctx, err, "添加订单商品失败")
	})
	return
}
