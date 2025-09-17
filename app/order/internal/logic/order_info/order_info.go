package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	v1 "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"
	"shop-goframe-micro-service-refacotor/app/order/api/pbentity"
	"shop-goframe-micro-service-refacotor/app/order/internal/dao"
	"shop-goframe-micro-service-refacotor/app/order/internal/model/entity"
	"shop-goframe-micro-service-refacotor/app/order/utility/rabbitmq"
	"shop-goframe-micro-service-refacotor/utility"

	"github.com/gogf/gf/v2/util/gconv"
)

// Create 创建订单（包含完整的事务处理）
func Create(ctx context.Context, req *v1.OrderInfoCreateReq) (int32, error) {
	//对订单请求进行校验
	if len(req.OrderGoodsInfo) == 0 {
		return 0, errors.New("订单必须至少包含一个商品")
	}
	//订单金额必须等于商品金额之和
	var totalGoodsPrice uint32
	var totalCouponPrice uint32
	for _, item := range req.OrderGoodsInfo {
		totalGoodsPrice += item.Price
		totalCouponPrice += item.CouponPrice
	}
	if req.Price != totalGoodsPrice {
		return 0, fmt.Errorf("订单总价[%d]与商品总价[%d]不符", req.Price, totalGoodsPrice)
	}

	if req.ActualPrice != req.Price-req.CouponPrice {
		return 0, fmt.Errorf("订单实际支付价格[%d]不等于订单总价[%d]减去优惠券价格[%d]", req.ActualPrice, req.Price, req.CouponPrice)
	}
	if req.CouponPrice < totalCouponPrice {
		return 0, fmt.Errorf("订单优惠券价格[%d]小于商品优惠券价格[%d]", req.CouponPrice, totalCouponPrice)
	}

	// 计算OrderGoodsItem中分摊的coupon_price
	var preAssignedCouponPrice uint32
	var orderGoodsList []entity.OrderGoodsInfo
	var itemsToAllocate []*entity.OrderGoodsInfo
	var allocatableItemsTotalPrice uint32

	if err := gconv.Structs(req.OrderGoodsInfo, &orderGoodsList); err != nil {
		return 0, fmt.Errorf("订单商品数据转换失败: %v", err)
	}

	for i := 0; i < len(orderGoodsList); i++ {
		item := &orderGoodsList[i]
		if item.CouponPrice > 0 {
			preAssignedCouponPrice += uint32(item.CouponPrice)
		} else {
			itemsToAllocate = append(itemsToAllocate, item)
			allocatableItemsTotalPrice += uint32(item.Price)
		}
	}

	couponPriceToAllocate := req.CouponPrice - preAssignedCouponPrice

	if couponPriceToAllocate > 0 && len(itemsToAllocate) > 0 {
		if allocatableItemsTotalPrice > 0 {
			var allocatedSoFar int = 0
			for i, item := range itemsToAllocate {
				if i == len(itemsToAllocate)-1 {
					item.CouponPrice = int(couponPriceToAllocate) - allocatedSoFar
					item.ActualPrice = item.Price - item.CouponPrice
				} else {
					// 使用uint64进行计算以防止溢出
					share := (uint64(item.Price) * uint64(couponPriceToAllocate)) / uint64(allocatableItemsTotalPrice)
					item.CouponPrice = int(share)
					item.ActualPrice = item.Price - item.CouponPrice
					allocatedSoFar += item.CouponPrice
				}
			}
		}
	}
	glog.Info(ctx, itemsToAllocate)
	glog.Info(ctx, orderGoodsList)
	// 开启事务
	db := g.DB()
	tx, err := db.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("开启事务失败: %v", err)
	}

	// 确保事务回滚
	var success bool
	defer func() {
		if !success {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				g.Log().Errorf(ctx, "事务回滚失败: %v", rollbackErr)
			}
		}
	}()

	// 使用 gconv.Struct 转换主订单
	var order entity.OrderInfo
	if err := gconv.Struct(req, &order); err != nil {
		return 0, fmt.Errorf("订单数据转换失败: %v", err)
	}

	// 设置订单特有字段
	order.Number = utility.GenerateOrderNumber()
	order.Status = 1
	order.CreatedAt = gtime.Now()
	order.UpdatedAt = gtime.Now()

	// 使用事务插入主订单
	result, err := dao.OrderInfo.Ctx(ctx).TX(tx).InsertAndGetId(order)
	if err != nil {
		return 0, fmt.Errorf("插入订单失败: %v", err)
	}
	orderId := int32(result)

	// 设置订单商品公共字段
	for i := range orderGoodsList {
		orderGoodsList[i].OrderId = int(orderId)
		orderGoodsList[i].CreatedAt = gtime.Now()
		orderGoodsList[i].UpdatedAt = gtime.Now()
	}

	// 订单商品列表不为空时，执行批量插入操作
	if len(orderGoodsList) > 0 {
		_, err = dao.OrderGoodsInfo.Ctx(ctx).TX(tx).Insert(orderGoodsList)
		if err != nil {
			return 0, fmt.Errorf("插入订单商品失败: %v", err)
		}
	}

	// 提交事务
	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("提交事务失败: %v", err)
	}

	success = true

	// 订单创建成功后，异步发送延迟消息
	go sendOrderTimeoutMessage(ctx, orderId)

	return orderId, nil
}

// sendOrderTimeoutMessage 发送订单超时消息
func sendOrderTimeoutMessage(ctx context.Context, orderId int32) {
	// 获取配置的延迟时间
	delay := rabbitmq.GetOrderTimeoutDelay(ctx)

	// 使用静态方法发送订单超时消息
	err := rabbitmq.SendOrderTimeoutMessageStatic(ctx, orderId, delay)
	if err != nil {
		g.Log().Errorf(ctx, "发送订单超时消息失败, 订单ID: %d, 错误: %v", orderId, err)
	}
}

// GetDetail 获取订单详情
func GetDetail(ctx context.Context, orderId uint32) (*pbentity.OrderInfo, []*pbentity.OrderGoodsInfo, error) {
	// 查询主订单
	var order entity.OrderInfo
	err := dao.OrderInfo.Ctx(ctx).WherePri(orderId).Scan(&order)
	if err != nil {
		return nil, nil, fmt.Errorf("查询订单失败: %v", err)
	}

	// 查询订单商品
	var goodsList []*entity.OrderGoodsInfo
	err = dao.OrderGoodsInfo.Ctx(ctx).Where("order_id", orderId).Scan(&goodsList)
	if err != nil {
		return nil, nil, fmt.Errorf("查询订单商品失败: %v", err)
	}

	// 转换订单数据
	var pbOrder pbentity.OrderInfo
	if err := gconv.Struct(order, &pbOrder); err != nil {
		return nil, nil, fmt.Errorf("转换订单数据失败: %v", err)
	}
	pbOrder.CreatedAt = utility.SafeConvertTime(order.CreatedAt)
	pbOrder.UpdatedAt = utility.SafeConvertTime(order.UpdatedAt)

	// 转换订单商品数据
	var pbGoodsList []*pbentity.OrderGoodsInfo
	for _, goods := range goodsList {
		var pbGoods pbentity.OrderGoodsInfo
		if err := gconv.Struct(goods, &pbGoods); err != nil {
			continue
		}
		pbGoods.CreatedAt = utility.SafeConvertTime(goods.CreatedAt)
		pbGoods.UpdatedAt = utility.SafeConvertTime(goods.UpdatedAt)
		pbGoodsList = append(pbGoodsList, &pbGoods)
	}

	return &pbOrder, pbGoodsList, nil
}

// GetList 获取订单列表
func GetList(ctx context.Context, req *v1.OrderInfoGetListReq) ([]*pbentity.OrderInfo, int, error) {
	// 构建查询条件
	model := dao.OrderInfo.Ctx(ctx)

	// 按订单编号查询
	if req.Number != "" {
		model = model.Where("number", req.Number)
	}

	// 按用户ID查询
	if req.UserId != 0 {
		model = model.Where("user_id", req.UserId)
	}

	// 按支付方式查询：1微信 2支付宝 3云闪付
	if req.PayType != 0 {
		model = model.Where("pay_type", req.PayType)
	}

	// 按订单状态查询：1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价
	if req.Status != 0 {
		model = model.Where("status", req.Status)
	}

	// 按收货人手机号查询
	if req.ConsigneePhone != "" {
		model = model.Where("consignee_phone", req.ConsigneePhone)
	}

	// 查询订单金额大于等于指定值（单位：分）
	if req.PriceGte != 0 {
		model = model.Where("price >= ?", req.PriceGte)
	}

	// 查询订单金额小于等于指定值（单位：分）
	if req.PriceLte != 0 {
		model = model.Where("price <= ?", req.PriceLte)
	}

	// 查询支付时间大于等于指定时间
	if req.PayAtGte != nil {
		model = model.Where("pay_at >= ?", req.PayAtGte.AsTime())
	}

	// 查询支付时间小于等于指定时间
	if req.PayAtLte != nil {
		model = model.Where("pay_at <= ?", req.PayAtLte.AsTime())
	}

	// 查询创建时间大于等于指定时间
	if req.DateGte != nil {
		model = model.Where("created_at >= ?", req.DateGte.AsTime())
	}

	// 查询创建时间小于等于指定时间
	if req.DateLte != nil {
		model = model.Where("created_at <= ?", req.DateLte.AsTime())
	}

	// 查询总数
	total, err := model.Count()
	if err != nil {
		return nil, 0, err
	}

	// 查询当前页数据
	orderRecords, err := model.Page(int(req.Page), int(req.Size)).All()
	if err != nil {
		return nil, 0, err
	}

	// 数据转换
	var pbOrders []*pbentity.OrderInfo
	for _, record := range orderRecords {
		var order entity.OrderInfo
		if err := record.Struct(&order); err != nil {
			continue
		}

		var pbOrder pbentity.OrderInfo
		if err := gconv.Struct(order, &pbOrder); err != nil {
			continue
		}

		// 单独处理时间字段
		pbOrder.CreatedAt = utility.SafeConvertTime(order.CreatedAt)
		pbOrder.UpdatedAt = utility.SafeConvertTime(order.UpdatedAt)

		pbOrders = append(pbOrders, &pbOrder)
	}

	return pbOrders, total, nil
}
