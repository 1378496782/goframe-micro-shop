// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2025-10-10 23:08:01
// 生成路径: internal/app/shop/logic/order_info.go
// 生成人：gfast
// desc:订单表
// company:云南奇讯科技有限公司
// ==========================================================================

package orderInfo

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
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
	service.RegisterOrderInfo(New())
}

func New() service.IOrderInfo {
	return &sOrderInfo{}
}

type sOrderInfo struct{}

func (s *sOrderInfo) List(ctx context.Context, req *model.OrderInfoSearchReq) (listRes *model.OrderInfoSearchRes, err error) {
	listRes = new(model.OrderInfoSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.OrderInfo.Ctx(ctx)
		if req.Id != "" {
			m = m.Where(dao.OrderInfo.Columns().Id+" = ?", req.Id)
		}
		if req.Number != "" {
			m = m.Where(dao.OrderInfo.Columns().Number+" = ?", req.Number)
		}
		if req.UserId != "" {
			m = m.Where(dao.OrderInfo.Columns().UserId+" = ?", gconv.Int(req.UserId))
		}
		if req.PayType != "" {
			m = m.Where(dao.OrderInfo.Columns().PayType+" = ?", gconv.Int(req.PayType))
		}
		if req.PayAt != "" {
			m = m.Where(dao.OrderInfo.Columns().PayAt+" = ?", gconv.Time(req.PayAt))
		}
		if req.Status != "" {
			m = m.Where(dao.OrderInfo.Columns().Status+" = ?", gconv.Int(req.Status))
		}
		if req.ConsigneeName != "" {
			m = m.Where(dao.OrderInfo.Columns().ConsigneeName+" like ?", "%"+req.ConsigneeName+"%")
		}
		if req.ConsigneePhone != "" {
			m = m.Where(dao.OrderInfo.Columns().ConsigneePhone+" = ?", req.ConsigneePhone)
		}
		if req.ConsigneeAddress != "" {
			m = m.Where(dao.OrderInfo.Columns().ConsigneeAddress+" = ?", req.ConsigneeAddress)
		}
		if req.Price != "" {
			m = m.Where(dao.OrderInfo.Columns().Price+" = ?", gconv.Int(req.Price))
		}
		if req.CouponPrice != "" {
			m = m.Where(dao.OrderInfo.Columns().CouponPrice+" = ?", gconv.Int(req.CouponPrice))
		}
		if req.ActualPrice != "" {
			m = m.Where(dao.OrderInfo.Columns().ActualPrice+" = ?", gconv.Int(req.ActualPrice))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.OrderInfo.Columns().CreatedAt+" >=? AND "+dao.OrderInfo.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
		}
		
		// 获取总数
		listRes.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取总行数失败")
		
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		listRes.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		order := dao.OrderInfo.Table() + ".id asc"
		if req.OrderBy != "" {
			// 如果用户提供的排序字段没有表名前缀，则添加order_info表前缀
			if !strings.Contains(req.OrderBy, ".") {
				order = dao.OrderInfo.Table() + "." + req.OrderBy
			} else {
				order = req.OrderBy
			}
		}
		
		// 先查询订单信息（不包含商品信息，确保分页准确）
		var orderList []*model.OrderInfoListRes
		err = m.Page(req.PageNum, req.PageSize).
			Order(order).
			Scan(&orderList)
		liberr.ErrIsNil(ctx, err, "获取订单数据失败")
		
		// 如果没有订单，直接返回
		if len(orderList) == 0 {
			listRes.List = make([]*model.OrderInfoListRes, 0)
			return
		}
		
		// 收集所有订单ID
		var orderIds []int
		for _, order := range orderList {
			orderIds = append(orderIds, order.Id)
		}
		
		// 批量查询订单商品信息
		var orderGoodsList []*model.OrderGoodsInfoListRes
		err = dao.OrderGoodsInfo.Ctx(ctx).
			WhereIn(dao.OrderGoodsInfo.Columns().OrderId, orderIds).
			Scan(&orderGoodsList)
		liberr.ErrIsNil(ctx, err, "获取订单商品信息失败")
		
		// 按订单ID分组商品信息
		orderGoodsMap := make(map[int][]*model.OrderGoodsInfoListRes)
		var allGoodsIds []int
		for _, goods := range orderGoodsList {
			orderGoodsMap[goods.OrderId] = append(orderGoodsMap[goods.OrderId], goods)
			allGoodsIds = append(allGoodsIds, goods.GoodsId)
		}
		
		// 批量查询商品详情信息（跨库查询）
		var goodsInfoMap = make(map[int]*model.GoodsInfoInfoRes) // 商品ID -> 商品信息
		if len(allGoodsIds) > 0 {
			var goodsInfoList []*model.GoodsInfoInfoRes
			err = dao.GoodsInfo.Ctx(ctx).
				WhereIn(dao.GoodsInfo.Columns().Id, allGoodsIds).
				Scan(&goodsInfoList)
			if err != nil {
				g.Log().Warningf(ctx, "批量获取商品信息失败，错误: %v", err)
			} else {
				// 将商品信息转换为map便于查找
				for _, goods := range goodsInfoList {
					goodsInfoMap[int(goods.Id)] = goods
				}
			}
		}
		
		// 组装最终结果
		listRes.List = make([]*model.OrderInfoListRes, 0, len(orderList))
		for _, order := range orderList {
			// 获取订单商品信息
			goodsList := orderGoodsMap[order.Id]
			
			// 为每个商品设置详细信息
			for _, goods := range goodsList {
				if goodsInfo, exists := goodsInfoMap[goods.GoodsId]; exists {
					goods.GoodsName = goodsInfo.Name
					goods.PicUrl = goodsInfo.PicUrl
				} else {
					// 如果获取商品信息失败，记录日志但继续处理
					g.Log().Warningf(ctx, "获取商品信息失败，商品ID: %d", goods.GoodsId)
					// 设置默认值
					goods.GoodsName = "商品信息获取失败"
					goods.PicUrl = ""
				}
			}
			
			// 设置订单商品信息
			order.GoodsInfo = goodsList
			listRes.List = append(listRes.List, order)
		}
	})
	return
}

func (s *sOrderInfo) GetById(ctx context.Context, id int) (res *model.OrderInfoInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.OrderInfo.Ctx(ctx).WithAll().Where(dao.OrderInfo.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sOrderInfo) Add(ctx context.Context, req *model.OrderInfoAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.OrderInfo.Ctx(ctx).Insert(do.OrderInfo{
			Number:           req.Number,
			UserId:           req.UserId,
			PayType:          req.PayType,
			Remark:           req.Remark,
			PayAt:            req.PayAt,
			Status:           req.Status,
			ConsigneeName:    req.ConsigneeName,
			ConsigneePhone:   req.ConsigneePhone,
			ConsigneeAddress: req.ConsigneeAddress,
			Price:            req.Price,
			CouponPrice:      req.CouponPrice,
			ActualPrice:      req.ActualPrice,
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sOrderInfo) Edit(ctx context.Context, req *model.OrderInfoEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.OrderInfo.Ctx(ctx).WherePri(req.Id).Update(do.OrderInfo{
			Number:           req.Number,
			UserId:           req.UserId,
			PayType:          req.PayType,
			Remark:           req.Remark,
			PayAt:            req.PayAt,
			Status:           req.Status,
			ConsigneeName:    req.ConsigneeName,
			ConsigneePhone:   req.ConsigneePhone,
			ConsigneeAddress: req.ConsigneeAddress,
			Price:            req.Price,
			CouponPrice:      req.CouponPrice,
			ActualPrice:      req.ActualPrice,
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sOrderInfo) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.OrderInfo.Ctx(ctx).Delete(dao.OrderInfo.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

// Ship 订单发货
func (s *sOrderInfo) Ship(ctx context.Context, id int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 检查订单是否存在
		count, err := dao.OrderInfo.Ctx(ctx).Where(dao.OrderInfo.Columns().Id, id).Count()
		liberr.ErrIsNil(ctx, err, "查询订单失败")
		if count == 0 {
			liberr.ErrIsNil(ctx, gerror.New("订单不存在"), "订单不存在")
		}
		
		// 更新订单状态为已发货 (假设状态码3表示已发货)
		_, err = dao.OrderInfo.Ctx(ctx).WherePri(id).Update(do.OrderInfo{
			Status: 3, // 3表示已发货
		})
		liberr.ErrIsNil(ctx, err, "发货失败")
	})
	return
}

// Refund 订单退款
func (s *sOrderInfo) Refund(ctx context.Context, id int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 检查订单是否存在
		count, err := dao.OrderInfo.Ctx(ctx).Where(dao.OrderInfo.Columns().Id, id).Count()
		liberr.ErrIsNil(ctx, err, "查询订单失败")
		if count == 0 {
			liberr.ErrIsNil(ctx, gerror.New("订单不存在"), "订单不存在")
		}
		
		// 更新订单状态为已退款 (假设状态码8表示已退款)
		_, err = dao.OrderInfo.Ctx(ctx).WherePri(id).Update(do.OrderInfo{
			Status: 8, // 8表示已退款
		})
		liberr.ErrIsNil(ctx, err, "退款失败")
	})
	return
}

// GetOrderProducts 获取订单商品列表
func (s *sOrderInfo) GetOrderProducts(ctx context.Context, orderId int) (list []*model.OrderGoodsInfoListRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 检查订单是否存在
		count, err := dao.OrderInfo.Ctx(ctx).Where(dao.OrderInfo.Columns().Id, orderId).Count()
		liberr.ErrIsNil(ctx, err, "查询订单失败")
		if count == 0 {
			liberr.ErrIsNil(ctx, gerror.New("订单不存在"), "订单不存在")
		}
		
		// 获取订单商品列表
		err = dao.OrderGoodsInfo.Ctx(ctx).Where(dao.OrderGoodsInfo.Columns().OrderId, orderId).Scan(&list)
		liberr.ErrIsNil(ctx, err, "获取订单商品列表失败")
	})
	return
}
