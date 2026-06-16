package logic

import (
	"context"
	"errors"
	"fmt"
	cartApi "shop-goframe-micro-service-refacotor/app/goods/api/cart_info/v1"
	goods_info "shop-goframe-micro-service-refacotor/app/goods/api/goods_info/v1"
	v1 "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"
	"shop-goframe-micro-service-refacotor/app/order/api/pbentity"
	"shop-goframe-micro-service-refacotor/app/order/internal/consts"
	"shop-goframe-micro-service-refacotor/app/order/internal/dao"
	"shop-goframe-micro-service-refacotor/app/order/internal/model/entity"
	cart "shop-goframe-micro-service-refacotor/app/order/utility/cart_info"
	goods "shop-goframe-micro-service-refacotor/app/order/utility/goods_info"
	"shop-goframe-micro-service-refacotor/app/order/utility/rabbitmq"
	"shop-goframe-micro-service-refacotor/utility"
	consts_ctrl "shop-goframe-micro-service-refacotor/utility/consts"
	"shop-goframe-micro-service-refacotor/utility/metrics"
	grabbitmq "shop-goframe-micro-service-refacotor/utility/rabbitmq"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/util/gconv"
)

// Create 创建订单（包含完整的事务处理）
func Create(ctx context.Context, req *v1.OrderInfoCreateReq) (int32, string, error) {
	//对订单请求进行校验
	if len(req.OrderGoodsInfo) == 0 {
		return 0, "", errors.New("订单必须至少包含一个商品")
	}
	//订单金额必须等于商品金额之和
	var totalGoodsPrice uint32
	var totalCouponPrice uint32
	var goodsIds []uint32
	for _, item := range req.OrderGoodsInfo {
		totalGoodsPrice += item.Price
		totalCouponPrice += item.CouponPrice
		goodsIds = append(goodsIds, item.GoodsId)
	}
	if req.Price != totalGoodsPrice {
		return 0, "", fmt.Errorf("订单总价[%d]与商品总价[%d]不符", req.Price, totalGoodsPrice)
	}

	if req.ActualPrice != req.Price-req.CouponPrice {
		return 0, "", fmt.Errorf("订单实际支付价格[%d]不等于订单总价[%d]减去优惠券价格[%d]", req.ActualPrice, req.Price, req.CouponPrice)
	}
	if req.CouponPrice < totalCouponPrice {
		return 0, "", fmt.Errorf("订单优惠券价格[%d]小于商品优惠券价格[%d]", req.CouponPrice, totalCouponPrice)
	}
	// 库存校验
	goodsStockMap, err := goods.Client.GetGoodsStock(ctx, &goods_info.GetGoodsStockReq{GoodsIds: goodsIds})
	if err != nil {
		return 0, "", fmt.Errorf("调用 goods 模块失败,err:%v", err)
	}
	fmt.Println("goodsStockMap", goodsStockMap)
	for _, item := range req.OrderGoodsInfo {
		if item.Count > uint32(goodsStockMap.GoodsStock[item.GoodsId]) {
			return 0, "", fmt.Errorf("商品{%d}库存不足", item.GoodsId)
		}
	}

	// 计算OrderGoodsItem中分摊的coupon_price
	var preAssignedCouponPrice uint32
	var orderGoodsList []entity.OrderGoodsInfo
	var itemsToAllocate []*entity.OrderGoodsInfo
	var allocatableItemsTotalPrice uint32

	if err := gconv.Structs(req.OrderGoodsInfo, &orderGoodsList); err != nil {
		return 0, "", fmt.Errorf("订单商品数据转换失败: %v", err)
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

	// 开启事务
	db := g.DB()
	tx, err := db.Begin(ctx)
	if err != nil {
		// 记录订单创建失败的业务指标
		metrics.RecordOrderCreate(ctx, false)
		return 0, "", fmt.Errorf("开启事务失败: %v", err)
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
		// 记录订单创建失败的业务指标
		metrics.RecordOrderCreate(ctx, false)
		return 0, "", fmt.Errorf("订单数据转换失败: %v", err)
	}

	// 设置订单特有字段
	order.Number = utility.GenerateOrderNumber()
	//order.Status = 1 // 6待确认
	if req.CouponId > 0 {
		order.Status = int(consts.OrderStatusPendingConfirm) // 使用优惠券，待确认
	} else {
		order.Status = int(consts.OrderStatusPendingPayment) // 不使用优惠券，待支付
	}
	order.CreatedAt = gtime.Now()
	order.UpdatedAt = gtime.Now()

	// 使用事务插入主订单
	result, err := dao.OrderInfo.Ctx(ctx).TX(tx).InsertAndGetId(order)
	if err != nil {
		// 记录订单创建失败的业务指标
		metrics.RecordOrderCreate(ctx, false)
		return 0, "", fmt.Errorf("插入订单失败: %v", err)
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
			// 记录订单创建失败的业务指标
			metrics.RecordOrderCreate(ctx, false)
			return 0, "", fmt.Errorf("插入订单商品失败: %v", err)
		}
	}

	// 提交事务
	if err = tx.Commit(); err != nil {
		// 记录订单创建失败的业务指标
		metrics.RecordOrderCreate(ctx, false)
		return 0, "", fmt.Errorf("提交事务失败: %v", err)
	}

	success = true

	// 订单创建成功后，发布订单创建事件，用于后续操作（如删除购物车商品）
	var goodsInfos []*grabbitmq.OrderGoodsInfo
	for _, item := range req.OrderGoodsInfo {
		goodsInfos = append(goodsInfos, &grabbitmq.OrderGoodsInfo{
			GoodsId: int(item.GoodsId),
			Count:   int(item.Count),
		})
	}
	event := grabbitmq.OrderCreatedEvent{
		UserId:    req.UserId,
		OrderId:   uint32(orderId),
		GoodsIds:  goodsIds,
		GoodsInfo: goodsInfos,
	}

	go grabbitmq.PublishOrderCreatedEvent(event)

	// 订单创建成功后，如果有优惠券使用，发送订单确认消息给goods服务进行优惠券扣减
	if req.CouponId > 0 {
		go grabbitmq.PublishCouponConfirmEvent(orderId, int32(req.UserId), int32(req.CouponId))
	}

	// 订单创建成功后，异步发送延迟消息
	delay := rabbitmq.GetOrderTimeoutDelay(ctx)
	go grabbitmq.PublishOrderTimeoutEvent(int(orderId), delay)

	// 记录订单创建成功的业务指标
	metrics.RecordOrderCreate(ctx, true)

	// 对于每个商品，更新库存指标
	for _, goodsInfo := range goodsInfos {
		metrics.UpdateInventory(ctx, fmt.Sprintf("%d", goodsInfo.GoodsId), int64(goodsStockMap.GoodsStock[uint32(goodsInfo.GoodsId)]-int32(goodsInfo.Count)))
	}

	return orderId, order.Number, nil
}

// sendOrderTimeoutMessage 发送订单超时消息
//func sendOrderTimeoutMessage(ctx context.Context, orderId int32) {
//	// 获取配置的延迟时间
//	delay := rabbitmq.GetOrderTimeoutDelay(ctx)
//
//	// 使用静态方法发送订单超时消息
//	err := rabbitmq.SendOrderTimeoutMessageStatic(ctx, orderId, delay)
//	if err != nil {
//		g.Log().Errorf(ctx, "发送订单超时消息失败, 订单ID: %d, 错误: %v", orderId, err)
//	}
//}

// GetDetail 获取订单详情
func GetDetail(ctx context.Context, orderId uint32, userId uint32) (*pbentity.OrderInfo, []*pbentity.OrderGoodsInfo, error) {
	// 查询主订单
	var order entity.OrderInfo
	err := dao.OrderInfo.Ctx(ctx).WherePri(orderId).Scan(&order)
	if err != nil {
		return nil, nil, fmt.Errorf("查询订单失败: %v", err)
	}
	// 检查订单是否存在
	if order.Id == 0 {
		return nil, nil, gerror.NewCode(gcode.CodeNotFound, "订单不存在")
	}
	// 校验订单归属
	if order.UserId != int(userId) {
		return nil, nil, gerror.NewCode(gcode.CodeNotAuthorized, "无权访问此订单")
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

// getlist V4版本分步联表查询 修改嵌套内容
func GetList(ctx context.Context, req *v1.OrderInfoGetListReq) ([]*v1.OrderListInfo, int, error) {
	// 1. 查询订单主表
	var orders []*entity.OrderInfo
	err := g.Model("order_info").
		Where("user_id", req.UserId).
		Where("status", req.Status).
		Page(int(req.Page), int(req.Size)).
		Order("id DESC").
		Scan(&orders)
	if err != nil {
		return nil, 0, gerror.Wrap(err, "查询订单失败")
	}

	// 2. 查询总数
	total, err := g.Model("order_info").
		Where("user_id", req.UserId).
		Where("status", req.Status).
		Count()
	if err != nil {
		return nil, 0, gerror.Wrap(err, "查询订单总数失败")
	}

	if len(orders) == 0 {
		return []*v1.OrderListInfo{}, total, nil
	}

	// 3. 构建结果
	result := make([]*v1.OrderListInfo, 0, len(orders))
	for _, order := range orders {
		// 查询商品信息
		var goods []*entity.OrderGoodsInfo
		err := g.Model("order_goods_info").
			Where("order_id", order.Id).
			Scan(&goods)
		if err != nil {
			g.Log().Errorf(ctx, "查询订单商品失败, order_id=%d: %v", order.Id, err)
			continue
		}

		// 转换商品信息
		goodsInfo := make([]*v1.OrderListGoodsInfo, 0, len(goods))
		for _, g := range goods {
			goodsInfo = append(goodsInfo, &v1.OrderListGoodsInfo{
				GoodsId: int32(g.GoodsId),
				Count:   int32(g.Count),
			})
		}

		// 构建订单信息
		result = append(result, &v1.OrderListInfo{
			Id:          int32(order.Id),
			UserId:      int32(order.UserId),
			Number:      order.Number,
			Status:      int32(order.Status),
			Price:       int32(order.Price),
			ActualPrice: int32(order.ActualPrice),
			GoodsInfo:   goodsInfo,
		})
	}

	g.Log().Debugf(ctx, "成功查询到 %d 条订单数据", len(result))
	return result, total, nil
}

// UpdateOrderStatus 更新订单状态
func UpdateOrderStatus(ctx context.Context, orderId, status int) error {
	updateData := g.Map{
		"status":     status,
		"updated_at": gtime.Now(),
	}

	// 只有当订单状态变为已支付(2)时才设置支付时间
	if status == int(consts.OrderStatusPaid) {
		updateData["pay_at"] = gtime.Now()
	}

	_, err := dao.OrderInfo.Ctx(ctx).Where("id", orderId).Update(updateData)
	if err != nil {
		return fmt.Errorf("更新订单状态失败: %v", err)
	}

	g.Log().Infof(ctx, "订单状态更新成功, 订单ID: %d, 新状态: %d", orderId, status)
	return nil
}

type OrderStatusTransitionReq struct {
	UserId     uint32
	OrderId    uint32
	FromStatus int
	ToStatus   int
}

// UpdateOrderStatusIfMatch 只有当前状态匹配时才更新订单状态
func UpdateOrderStatusIfMatch(ctx context.Context, req *OrderStatusTransitionReq) (bool, error) {
	updateData := g.Map{
		"status":     req.ToStatus,
		"updated_at": gtime.Now(),
	}

	// 只有当订单状态变为已支付(2)时才设置支付时间
	if req.ToStatus == int(consts.OrderStatusPaid) {
		updateData["pay_at"] = gtime.Now()
	}

	result, err := dao.OrderInfo.Ctx(ctx).Where(g.Map{
		"id":      req.OrderId,
		"user_id": req.UserId,
		"status":  req.FromStatus,
	}).Update(updateData)
	if err != nil {
		return false, fmt.Errorf("更新订单状态失败: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("获取订单状态更新结果失败: %v", err)
	}
	if rows == 0 {
		return false, nil
	}

	g.Log().Infof(ctx, "订单状态更新成功, 订单ID: %d, 原状态: %d, 新状态: %d", req.OrderId, req.FromStatus, req.ToStatus)
	return true, nil
}

// UpdateOrderStatus 更新订单状态
// true  = 本次真的把订单从待支付改成已支付
// false = 订单已经处理过，或者状态不是待支付
// error = 数据库错误
func UpdateOrderStatusByNumber(ctx context.Context, number, transactionId string, status int) (bool, error) {
	// exists, err := dao.OrderInfo.Ctx(ctx).
	// 	Where("number", number).
	// 	Where("status", consts.OrderStatusPaid).
	// 	Exist()
	// if err != nil {
	// 	return gerror.WrapCode(gcode.CodeDbOperationError, err)
	// }
	// if exists {
	// 	g.Log().Infof(ctx, "{%s}订单的状态已修改，不需要再修改", number)
	// 	return nil
	// }

	updateData := g.Map{
		"status":         status,
		"updated_at":     gtime.Now(),
		"transaction_id": transactionId,
	}

	// 更新订单状态
	result, err := dao.OrderInfo.Ctx(ctx).Where(g.Map{
		"number": number,
		"status": consts.OrderStatusPendingPayment,
	}).Update(updateData)
	if err != nil {
		return false, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return false, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	if rows == 0 {
		g.Log().Infof(ctx, "订单状态未更新，可能已处理或状态不允许更新, 订单编号: %s", number)
		return false, nil
	}

	g.Log().Infof(ctx, "订单状态更新成功, 订单编号: %s, 新状态: %d", number, status)
	return true, nil
}

// HandlePaidOrder 处理支付成功后的业务动作。
// 返回 nil 且未更新订单时，表示该支付回调已经处理过，可以安全忽略。
func HandlePaidOrder(ctx context.Context, orderNumber, transactionId string) error {
	success, err := UpdateOrderStatusByNumber(ctx, orderNumber, transactionId, int(consts.OrderStatusPaid))
	if err != nil {
		return err
	}
	if !success {
		return nil
	}

	if err = IncreaseOrderGoodsSales(ctx, orderNumber); err != nil {
		// 即使 IncreaseOrderGoodsSales 失败，也先记录错误日志并返回 nil。这样支付回调不会一直失败重试；补偿我们后面专门做。
		g.Log().Errorf(ctx, "增加订单销量失败, 订单编号: %s, 错误: %v", orderNumber, err)
		return nil
	}

	return nil
}

// HandleCouponResult 处理优惠券扣减结果
// goods服务通过userid和couponid在user_coupon_info表中定位数据
// 如果找到数据且状态为"待使用"，则修改为"已使用"并返回成功
// 如果未找到数据或状态不是"待使用"，则返回失败
func HandleCouponResult(ctx context.Context, orderId int, success bool, message string) error {

	if success {
		// 优惠券扣减成功，订单状态改为待支付(1)
		err := UpdateOrderStatus(ctx, orderId, int(consts.OrderStatusPendingPayment))
		if err != nil {
			g.Log().Errorf(ctx, "优惠券扣减成功，但更新订单状态失败, 订单ID: %d, 错误: %v", orderId, err)
			return err
		}
		g.Log().Infof(ctx, "优惠券扣减成功，订单状态已更新为待支付, 订单ID: %d", orderId)
	} else {
		// 优惠券扣减失败（未找到数据或状态不是"待使用"），订单状态改为已取消(7)
		err := UpdateOrderStatus(ctx, orderId, int(consts.OrderStatusCancelled))
		if err != nil {
			g.Log().Errorf(ctx, "优惠券扣减失败，但更新订单状态失败, 订单ID: %d, 错误: %v", orderId, err)
			return err
		}
		g.Log().Warningf(ctx, "优惠券扣减失败，订单状态已更新为已取消, 订单ID: %d, 原因: %s", orderId, message)
	}

	return nil
}

// GetCount 获取各类订单数量
func GetCount(ctx context.Context, userId uint32) (*v1.OrderInfoGetCountRes, error) {
	var results []struct {
		Status int    `json:"status"`
		Count  uint32 `json:"count"`
	}
	err := dao.OrderInfo.Ctx(ctx).
		Fields("status, COUNT(*) as count").
		Where("user_id", userId).
		Group("status").
		Scan(&results)
	if err != nil {
		return nil, gerror.Wrap(err, "查询订单数量失败")
	}

	res := &v1.OrderInfoGetCountRes{}
	for _, item := range results {
		switch consts.OrderStatus(item.Status) {
		case consts.OrderStatusPendingPayment:
			res.Pending += item.Count
		case consts.OrderStatusPaid:
			res.Shipping += item.Count
		case consts.OrderStatusShipped:
			res.Delivered += item.Count
		case consts.OrderStatusReceived, consts.OrderStatusCompleted:
			res.Completed += item.Count
			//TODO 不代表售后
			//case consts.OrderStatusCancelled:
			//res.AfterSale += item.Count // Assuming cancelled orders are "afterSale" for now
		}
	}

	return res, nil
}

func HandleOrderTimeoutResult(ctx context.Context, orderId int) error {
	// 更新字段
	updateData := g.Map{
		"status":     consts.OrderStatusCancelled,
		"updated_at": gtime.Now(), // 可选：更新时间戳
	}
	// 更新订单状态
	result, err := dao.OrderInfo.Ctx(ctx).Where("id=? AND status=?", orderId, consts.OrderStatusPendingPayment).Update(updateData)
	if err != nil {
		return gerror.WrapCode(gcode.CodeDbOperationError, err)
	}

	row, _ := result.RowsAffected()
	if row == 0 {
		g.Log().Infof(ctx, "订单已取消，无需再取消, orderId=%d", orderId)
		return nil
	}

	g.Log().Infof(ctx, "订单状态更新成功, 订单编号:{%d}, 新状态: %d", orderId, consts.OrderStatusPendingPayment)
	return nil
}

func GetOrderDetail(ctx context.Context, orderId int) ([]*grabbitmq.OrderGoodsInfo, error) {
	// 查询主订单
	var order entity.OrderInfo
	err := dao.OrderInfo.Ctx(ctx).WherePri(orderId).Scan(&order)
	if err != nil {
		return nil, fmt.Errorf("查询订单失败: %v", err)
	}
	// 检查订单是否存在
	if order.Id == 0 {
		return nil, gerror.NewCode(gcode.CodeNotFound, "订单不存在")
	}
	// 查询订单商品
	var goodsList []*entity.OrderGoodsInfo
	err = dao.OrderGoodsInfo.Ctx(ctx).Where("order_id", orderId).Scan(&goodsList)
	if err != nil {
		return nil, fmt.Errorf("查询订单商品失败: %v", err)
	}

	// 转换订单商品数据
	var orderGoodsInfo []*grabbitmq.OrderGoodsInfo
	for _, goods := range goodsList {
		orderGoodsInfo = append(orderGoodsInfo, &grabbitmq.OrderGoodsInfo{
			GoodsId: goods.GoodsId,
			Count:   goods.Count,
		})
	}

	return orderGoodsInfo, nil
}

func Preview(ctx context.Context, req *v1.OrderInfoPreviewReq) (*v1.OrderInfoPreviewRes, error) {
	if len(req.CartIds) == 0 || req.UserId == 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "参数错误")
	}

	cartRes, err := cart.Client.GetSelectedItems(ctx, &cartApi.CartInfoGetSelectedItemsReq{
		UserId:  req.UserId,
		CartIds: req.CartIds,
	})
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}

	items := []*v1.OrderInfoPreviewItem{}
	totalPrice := uint64(0)
	totalCount := uint32(0)
	for _, item := range cartRes.Items {
		subTotal := item.GoodsPrice * uint64(item.Count)

		items = append(items, &v1.OrderInfoPreviewItem{
			CartId:     item.Id,
			GoodsId:    item.GoodsId,
			GoodsName:  item.GoodsName,
			GoodsPrice: item.GoodsPrice,
			Count:      item.Count,
			SubTotal:   subTotal,
		})

		totalPrice += subTotal
		totalCount += item.Count
	}

	return &v1.OrderInfoPreviewRes{
		Items:      items,
		TotalPrice: totalPrice,
		TotalCount: totalCount,
	}, nil
}

/*
1. 参数校验
2. 查选中的购物车商品
3. 校验必须全部命中
4. 校验商品有效性和库存，同时计算金额
5. 组装订单主表
6. 事务插入订单主表和订单商品表
*/
func CreateFromCart(ctx context.Context, req *v1.OrderInfoCreateFromCartReq) (orderId int32, orderNumber string, err error) {
	// 1. 参数校验
	if len(req.CartIds) == 0 || req.UserId == 0 {
		return 0, "", gerror.NewCode(gcode.CodeInvalidParameter, "参数错误")
	}

	// 2. 查选中的购物车商品
	cartRes, err := cart.Client.GetSelectedItems(ctx, &cartApi.CartInfoGetSelectedItemsReq{
		UserId:  req.UserId,
		CartIds: req.CartIds,
	})
	if err != nil {
		return 0, "", gerror.WrapCode(gcode.CodeDbOperationError, err)
	}

	// 3. 校验必须全部命中
	if len(cartRes.Items) != len(req.CartIds) {
		return 0, "", gerror.NewCode(gcode.CodeInvalidParameter, "购物车商品数量与请求参数不一致")
	}

	// 4. 校验商品有效性和库存，同时计算金额
	totalPrice := uint64(0)
	orderGoodsList := []entity.OrderGoodsInfo{}

	goodsIds := []uint32{}
	counts := []uint32{}
	for _, item := range cartRes.Items {
		goodsIds = append(goodsIds, item.GoodsId)
		counts = append(counts, item.Count)
		if item.GoodsId == 0 || item.GoodsName == "" || item.GoodsPrice <= 0 || item.Count <= 0 {
			return 0, "", gerror.NewCode(gcode.CodeInvalidParameter, "商品名称、价格或数量无效")
		}
		if item.GoodsStock < item.Count {
			return 0, "", gerror.NewCode(gcode.CodeInvalidParameter, "商品库存不足")
		}

		subTotal := item.GoodsPrice * uint64(item.Count)
		totalPrice += subTotal

		orderGoodsList = append(orderGoodsList, entity.OrderGoodsInfo{
			GoodsId:     int(item.GoodsId),
			Count:       int(item.Count),
			Price:       int(subTotal),
			CouponPrice: 0,
			ActualPrice: int(subTotal),
			CreatedAt:   gtime.Now(),
			UpdatedAt:   gtime.Now(),
		})
	}

	// 调 goods-service DeductStock
	if _, err = goods.Client.DeductStock(ctx, &goods_info.DeductStockReq{
		GoodsIds: goodsIds,
		Counts:   counts,
	}); err != nil {
		g.Log().Errorf(ctx, "扣减商品库存失败: %v", err)
		return 0, "", fmt.Errorf("扣减商品库存失败: %v", err)
	}

	// 5. 组装订单主表
	order := entity.OrderInfo{
		Number:           utility.GenerateOrderNumber(),
		UserId:           int(req.UserId),
		Remark:           req.Remark,
		Status:           int(consts.OrderStatusPendingPayment),
		ConsigneeName:    req.ConsigneeName,
		ConsigneePhone:   req.ConsigneePhone,
		ConsigneeAddress: req.ConsigneeAddress,
		Price:            int(totalPrice),
		CouponPrice:      0,
		ActualPrice:      int(totalPrice),
		CreatedAt:        gtime.Now(),
		UpdatedAt:        gtime.Now(),
	}

	// 6. 事务插入
	db := g.DB()
	tx, err := db.Begin(ctx)
	if err != nil {
		return 0, "", fmt.Errorf("开启事务失败: %v", err)
	}

	success := false
	defer func() {
		if !success {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				g.Log().Errorf(ctx, "事务回滚失败: %v", rollbackErr)
			}
		}
	}()

	id, err := dao.OrderInfo.Ctx(ctx).TX(tx).InsertAndGetId(order)
	if err != nil {
		return 0, "", fmt.Errorf("插入订单失败: %v", err)
	}

	orderId = int32(id)

	for i := range orderGoodsList {
		orderGoodsList[i].OrderId = int(orderId)
	}

	if len(orderGoodsList) > 0 {
		if _, err = dao.OrderGoodsInfo.Ctx(ctx).TX(tx).Insert(orderGoodsList); err != nil {
			return 0, "", fmt.Errorf("插入订单商品失败: %v", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return 0, "", fmt.Errorf("提交事务失败: %v", err)
	}

	success = true

	// 7. 删除购物车商品
	if _, err = cart.Client.DeleteSelectedItems(ctx, &cartApi.CartInfoDeleteSelectedItemsReq{
		UserId:  req.UserId,
		CartIds: req.CartIds,
	}); err != nil {
		g.Log().Errorf(ctx, "删除购物车商品失败: %v", err)
	}

	return orderId, order.Number, nil
}

func CancelOrder(ctx context.Context, req *v1.CancelOrderReq) (res *v1.CancelOrderRes, err error) {
	infoError := consts_ctrl.InfoError(consts_ctrl.OrderInfo, consts_ctrl.GetOrderRecord)

	record, err := dao.OrderInfo.Ctx(ctx).Where("id", req.Id).One()

	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}

	if record.IsEmpty() {
		return nil, gerror.NewCode(gcode.CodeNotFound, "订单不存在")
	}

	// 查 order_goods_info 商品快照
	var goodsList []*entity.OrderGoodsInfo
	err = dao.OrderGoodsInfo.Ctx(ctx).Where("order_id", req.Id).Scan(&goodsList)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, "查询订单商品失败")
	}

	// 只有待支付订单才允许取消
	ok, err := UpdateOrderStatusIfMatch(ctx, &OrderStatusTransitionReq{
		UserId:     req.UserId,
		OrderId:    req.Id,
		FromStatus: int(consts.OrderStatusPendingPayment),
		ToStatus:   int(consts.OrderStatusCancelled),
	})
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, "系统错误，取消失败")
	}
	if !ok {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "订单不存在或状态不允许取消")
	}

	// 恢复库存
	goodsIds := []uint32{}
	counts := []uint32{}
	for _, item := range goodsList {
		goodsIds = append(goodsIds, uint32(item.GoodsId))
		counts = append(counts, uint32(item.Count))
	}
	if _, err = goods.Client.RestoreStock(ctx, &goods_info.RestoreStockReq{
		GoodsIds: goodsIds,
		Counts:   counts,
	}); err != nil {
		g.Log().Errorf(ctx, "恢复商品库存失败: %v", err)

		// 恢复库存失败，回滚订单状态
		ok, err = UpdateOrderStatusIfMatch(ctx, &OrderStatusTransitionReq{
			UserId:     req.UserId,
			OrderId:    req.Id,
			FromStatus: int(consts.OrderStatusCancelled),
			ToStatus:   int(consts.OrderStatusPendingPayment),
		})
		if err != nil {
			g.Log().Errorf(ctx, "回滚订单状态失败: %v", err)
			return nil, gerror.WrapCode(gcode.CodeDbOperationError, err)
		}
		if !ok {
			g.Log().Warningf(ctx, "回滚订单状态未生效, 订单ID: %d", req.Id)
		}

		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "恢复商品库存失败")
	}

	return &v1.CancelOrderRes{
		Message: "订单取消成功",
	}, nil
}

func IncreaseOrderGoodsSales(ctx context.Context, orderNumber string) error {
	// 根据订单号查订单ID
	var order entity.OrderInfo
	err := dao.OrderInfo.Ctx(ctx).Where(dao.OrderInfo.Columns().Number, orderNumber).Scan(&order)
	if err != nil {
		g.Log().Errorf(ctx, "查询订单失败: %v", err)
		return err
	}
	if order.Id == 0 {
		return gerror.NewCode(gcode.CodeNotFound, "订单不存在")
	}

	// 查询订单商品
	var orderGoodsList []*entity.OrderGoodsInfo
	err = dao.OrderGoodsInfo.Ctx(ctx).Where(dao.OrderGoodsInfo.Columns().OrderId, order.Id).Scan(&orderGoodsList)
	if err != nil {
		g.Log().Errorf(ctx, "查询订单商品失败: %v", err)
		return err
	}
	if len(orderGoodsList) == 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "订单商品为空")
	}

	goodsIds := make([]uint32, 0, len(orderGoodsList))
	counts := make([]uint32, 0, len(orderGoodsList))
	for _, item := range orderGoodsList {
		goodsIds = append(goodsIds, uint32(item.GoodsId))
		counts = append(counts, uint32(item.Count))
	}

	// 调 goods-service IncreaseSales
	_, err = goods.Client.IncreaseSales(ctx, &goods_info.IncreaseSalesReq{
		GoodsIds: goodsIds,
		Counts:   counts,
	})
	if err != nil {
		g.Log().Errorf(ctx, "增加商品销售量失败: %v", err)
		return err
	}

	return nil
}
