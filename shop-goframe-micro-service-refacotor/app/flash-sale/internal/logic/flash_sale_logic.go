package logic

import (
	"context"
	"fmt"
	"time"

	v1 "shop-goframe-micro-service-refacotor/app/flash-sale/api/flash_sale/v1"
	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/consts"
	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/model"
	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/mq"
	"shop-goframe-micro-service-refacotor/app/flash-sale/utility"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// FlashSale 秒杀服务实现
type FlashSale struct{}

// NewFlashSale 创建秒杀服务
func NewFlashSale() *FlashSale {
	return &FlashSale{}
}

// GetFlashSaleGoodsList 获取秒杀商品列表
func (s *FlashSale) GetFlashSaleGoodsList(ctx context.Context, req *v1.FlashSaleGoodsListReq) (*v1.FlashSaleGoodsListRes, error) {
	// 模拟数据 - MVP阶段简化实现
	goodsList := []*v1.FlashSaleGoodsInfo{
		{
			GoodsId:        1001,
			ActivityId:     req.ActivityId,
			Title:          "iPhone 15 Pro Max",
			Description:    "Apple iPhone 15 Pro Max 256GB 钛金属",
			OriginalPrice:  999900, // 9999元
			SalePrice:      899900, // 8999元
			TotalStock:     100,
			AvailableStock: 95,
			StartTime:      gconv.Uint32(time.Now().Unix() - 3600),   // 1小时前开始
			EndTime:        gconv.Uint32(time.Now().Unix() + 3600*2), // 2小时后结束
			Status:         1,                                        // 进行中
			ImageUrl:       "https://example.com/iphone15.jpg",
		},
		{
			GoodsId:        1002,
			ActivityId:     req.ActivityId,
			Title:          "MacBook Pro M3",
			Description:    "Apple MacBook Pro 14英寸 M3芯片",
			OriginalPrice:  1499900, // 14999元
			SalePrice:      1299900, // 12999元
			TotalStock:     50,
			AvailableStock: 48,
			StartTime:      gconv.Uint32(time.Now().Unix() - 1800),   // 30分钟前开始
			EndTime:        gconv.Uint32(time.Now().Unix() + 3600*3), // 3小时后结束
			Status:         1,                                        // 进行中
			ImageUrl:       "https://example.com/macbook.jpg",
		},
	}

	return &v1.FlashSaleGoodsListRes{
		Total: gconv.Int32(len(goodsList)),
		List:  goodsList,
	}, nil
}

// GetFlashSaleGoodsDetail 获取秒杀商品详情
func (s *FlashSale) GetFlashSaleGoodsDetail(ctx context.Context, req *v1.FlashSaleGoodsDetailReq) (*v1.FlashSaleGoodsDetailRes, error) {
	// 模拟数据 - MVP阶段简化实现
	goodsInfo := &v1.FlashSaleGoodsInfo{
		GoodsId:        req.GoodsId,
		ActivityId:     req.ActivityId,
		Title:          "iPhone 15 Pro Max",
		Description:    "Apple iPhone 15 Pro Max 256GB 钛金属",
		OriginalPrice:  999900,
		SalePrice:      899900,
		TotalStock:     100,
		AvailableStock: 95,
		StartTime:      gconv.Uint32(time.Now().Unix() - 3600),
		EndTime:        gconv.Uint32(time.Now().Unix() + 3600*2),
		Status:         1,
		ImageUrl:       "https://example.com/iphone15.jpg",
	}

	// 计算剩余时间
	now := time.Now().Unix()
	var remainSeconds int64
	if now < int64(goodsInfo.StartTime) {
		remainSeconds = int64(goodsInfo.StartTime) - now
	} else if now < int64(goodsInfo.EndTime) {
		remainSeconds = int64(goodsInfo.EndTime) - now
	}

	return &v1.FlashSaleGoodsDetailRes{
		GoodsInfo:     goodsInfo,
		RemainSeconds: remainSeconds,
		CanBuy:        goodsInfo.Status == 1 && goodsInfo.AvailableStock > 0,
	}, nil
}

// CreateFlashSaleOrder 创建秒杀订单
func (s *FlashSale) CreateFlashSaleOrder(ctx context.Context, req *v1.CreateFlashSaleOrderReq) (*v1.CreateFlashSaleOrderRes, error) {
	// 参数校验
	if req.GoodsId == 0 || req.UserId == 0 || req.Count <= 0 {
		return &v1.CreateFlashSaleOrderRes{
			Success: false,
			Message: "参数错误",
			Status:  consts.FlashSaleStatusFailed,
		}, nil
	}

	// 获取客户端IP
	ip := g.RequestFromCtx(ctx).GetClientIp()

	// 获取缓存实例
	cache := utility.GetFlashSaleCache()
	if cache == nil {
		return &v1.CreateFlashSaleOrderRes{
			Success: false,
			Message: "系统错误",
			Status:  consts.FlashSaleStatusFailed,
		}, nil
	}

	// 创建限流器和防刷检查器
	rateLimiter := utility.NewRateLimiter(cache)
	antiBrush := utility.NewAntiBrushChecker(cache)

	// 防刷检查
	if err := antiBrush.CheckUserBehavior(ctx, req.UserId, ip); err != nil {
		g.Log().Warningf(ctx, "防刷检查失败，用户ID:%d，IP:%s，错误:%v", req.UserId, ip, err)
		return &v1.CreateFlashSaleOrderRes{
			Success: false,
			Message: err.Error(),
			Status:  consts.FlashSaleStatusFailed,
		}, nil
	}

	// 全局限流检查
	if err := rateLimiter.CheckGlobalLimit(ctx); err != nil {
		g.Log().Warningf(ctx, "全局限流检查失败，错误:%v", err)
		return &v1.CreateFlashSaleOrderRes{
			Success: false,
			Message: err.Error(),
			Status:  consts.FlashSaleStatusFailed,
		}, nil
	}

	// 用户限流检查
	if err := rateLimiter.CheckUserLimit(ctx, req.UserId); err != nil {
		g.Log().Warningf(ctx, "用户限流检查失败，用户ID:%d，错误:%v", req.UserId, err)
		return &v1.CreateFlashSaleOrderRes{
			Success: false,
			Message: err.Error(),
			Status:  consts.FlashSaleStatusFailed,
		}, nil
	}

	// IP限流检查
	if err := rateLimiter.CheckIPLimit(ctx, ip); err != nil {
		g.Log().Warningf(ctx, "IP限流检查失败，IP:%s，错误:%v", ip, err)
		return &v1.CreateFlashSaleOrderRes{
			Success: false,
			Message: err.Error(),
			Status:  consts.FlashSaleStatusFailed,
		}, nil
	}

	// 购买限制检查
	if err := rateLimiter.CheckPurchaseLimit(ctx, req.UserId, req.GoodsId); err != nil {
		g.Log().Warningf(ctx, "购买限制检查失败，用户ID:%d，商品ID:%d，错误:%v", req.UserId, req.GoodsId, err)
		return &v1.CreateFlashSaleOrderRes{
			Success: false,
			Message: err.Error(),
			Status:  consts.FlashSaleStatusFailed,
		}, nil
	}

	// 生成结果查询ID
	resultId := fmt.Sprintf("%d_%d_%s", req.UserId, req.GoodsId, gconv.String(grand.N(100000, 999999)))

	// 获取库存管理器
	stockManager := utility.GetFlashSaleStockManager(cache)
	if stockManager == nil {
		return &v1.CreateFlashSaleOrderRes{
			Success: false,
			Message: "系统错误",
			Status:  consts.FlashSaleStatusFailed,
		}, nil
	}

	// 扣减库存（核心逻辑）
	err := stockManager.ReduceStock(ctx, req.GoodsId, int(req.Count))
	success := err == nil
	if err != nil {
		g.Log().Warningf(ctx, "扣减秒杀库存失败，用户ID:%d，商品ID:%d，错误:%v", req.UserId, req.GoodsId, err)
		return &v1.CreateFlashSaleOrderRes{
			Success:  false,
			Message:  err.Error(),
			ResultId: resultId,
			Status:   consts.FlashSaleStatusFailed,
		}, nil
	}

	if !success {
		return &v1.CreateFlashSaleOrderRes{
			Success:  false,
			Message:  "库存不足",
			ResultId: resultId,
			Status:   consts.FlashSaleStatusFailed,
		}, nil
	}

	// 记录用户购买（用于限购检查）
	if err := rateLimiter.RecordPurchase(ctx, req.UserId, req.GoodsId); err != nil {
		g.Log().Warningf(ctx, "记录购买失败，用户ID:%d，商品ID:%d，错误:%v", req.UserId, req.GoodsId, err)
		// 记录失败但不影响主流程，继续处理
	}

	// 生成订单号
	orderNo := fmt.Sprintf("FS%s%d", time.Now().Format("20060102150405"), grand.N(1000, 9999))

	// 记录秒杀结果到缓存
	if err := s.recordFlashSaleResult(ctx, resultId, req.UserId, req.GoodsId, orderNo, consts.FlashSaleStatusSuccess); err != nil {
		g.Log().Warning(ctx, "记录秒杀结果失败:", err)
	}

	// 发送消息到队列进行异步处理
	orderMsg := &model.FlashSaleOrderMessage{
		OrderId:    orderNo,
		UserId:     req.UserId,
		GoodsId:    req.GoodsId,
		Count:      uint32(req.Count),
		Price:      899900, // 模拟价格
		CreateTime: time.Now(),
		Status:     1, // 待处理
		RetryCount: 0,
		MaxRetry:   3,
	}

	if err := mq.PublishFlashSaleOrder(ctx, orderMsg); err != nil {
		g.Log().Error(ctx, "发送秒杀订单消息失败:", err)
		// 消息发送失败，但库存已经扣除，需要记录日志并人工处理
	}

	return &v1.CreateFlashSaleOrderRes{
		Success:  true,
		OrderNo:  orderNo,
		Message:  "秒杀成功，订单处理中",
		ResultId: resultId,
		Status:   consts.FlashSaleStatusSuccess,
	}, nil
}

// GetFlashSaleResult 查询秒杀结果
func (s *FlashSale) GetFlashSaleResult(ctx context.Context, req *v1.GetFlashSaleResultReq) (*v1.GetFlashSaleResultRes, error) {
	if req.ResultId == "" || req.UserId == 0 {
		return nil, gerror.New("参数错误")
	}

	// 从缓存获取结果
	cache := utility.GetFlashSaleCache()
	if cache == nil {
		return &v1.GetFlashSaleResultRes{
			Status:  consts.FlashSaleStatusFailed,
			Message: "系统错误",
		}, nil
	}

	resultKey := fmt.Sprintf(consts.FlashSaleResultCacheKey, req.ResultId)
	resultData, err := cache.Get(ctx, resultKey)
	if err != nil {
		g.Log().Warning(ctx, "获取秒杀结果失败:", err)
		return &v1.GetFlashSaleResultRes{
			Status:  consts.FlashSaleStatusPending,
			Message: "处理中",
		}, nil
	}

	if resultData == nil {
		return &v1.GetFlashSaleResultRes{
			Status:  consts.FlashSaleStatusPending,
			Message: "处理中",
		}, nil
	}

	// 解析结果数据
	resultMap := gconv.Map(resultData)
	return &v1.GetFlashSaleResultRes{
		Status:    gconv.Int32(resultMap["status"]),
		Message:   gconv.String(resultMap["message"]),
		OrderNo:   gconv.String(resultMap["order_no"]),
		GoodsId:   gconv.Uint32(resultMap["goods_id"]),
		PayAmount: gconv.Int64(resultMap["pay_amount"]),
	}, nil
}

// InitFlashSaleStock 初始化秒杀商品库存
func (s *FlashSale) InitFlashSaleStock(ctx context.Context, goodsId uint32, stock int) error {
	cache := utility.GetFlashSaleCache()
	if cache == nil {
		return gerror.New("缓存未初始化")
	}

	stockManager := utility.GetFlashSaleStockManager(cache)
	if stockManager == nil {
		return gerror.New("库存管理器未初始化")
	}

	// 使用ReduceStock方法来初始化库存（先设置一个较大的值，然后减少到目标值）
	currentStock, err := stockManager.GetStock(ctx, goodsId)
	if err != nil {
		// 如果库存不存在，我们需要先设置一个基础值
		// 这里我们暂时使用一个简单的缓存设置
		return cache.Set(ctx, fmt.Sprintf("flash_sale:stock:%d", goodsId), stock, 0)
	}

	_ = currentStock // 避免未使用变量的错误
	return cache.Set(ctx, fmt.Sprintf("flash_sale:stock:%d", goodsId), stock, 0)
}

// ProcessFlashSaleOrder 处理秒杀订单（异步）
func (s *FlashSale) ProcessFlashSaleOrder(ctx context.Context, orderMsg *model.FlashSaleOrderMessage) error {
	g.Log().Infof(ctx, "开始处理秒杀订单消息: OrderId=%s, UserId=%d, GoodsId=%d",
		orderMsg.OrderId, orderMsg.UserId, orderMsg.GoodsId)

	// 模拟订单处理逻辑
	time.Sleep(100 * time.Millisecond) // 模拟处理时间

	// 这里可以添加实际的订单处理逻辑，比如：
	// 1. 创建订单记录
	// 2. 更新订单状态
	// 3. 发送通知等

	g.Log().Infof(ctx, "秒杀订单处理完成: %s", orderMsg.OrderId)
	return nil
}

// checkUserRateLimit 检查用户限流
func (s *FlashSale) checkUserRateLimit(ctx context.Context, userId uint32) error {
	cache := utility.GetFlashSaleCache()
	if cache == nil {
		return nil // 缓存未初始化时不限流
	}

	limitKey := fmt.Sprintf(consts.FlashSaleUserRateLimitKey, userId)

	// 获取当前计数
	count, err := cache.Get(ctx, limitKey)
	if err != nil {
		return err
	}

	currentCount := gconv.Int(count)
	if currentCount >= consts.FlashSaleRateLimitPerSecond {
		return gerror.New("请求过于频繁，请稍后再试")
	}

	// 增加计数
	newCount := currentCount + 1
	if currentCount == 0 {
		// 第一次设置，1秒过期
		err = cache.Set(ctx, limitKey, newCount, time.Second)
	} else {
		// 更新计数，保持原有过期时间
		err = cache.Set(ctx, limitKey, newCount, 0) // 0表示保持原有TTL
	}

	return err
}

// recordFlashSaleResult 记录秒杀结果
func (s *FlashSale) recordFlashSaleResult(ctx context.Context, resultId string, userId, goodsId uint32, orderNo string, status uint32) error {
	cache := utility.GetFlashSaleCache()
	if cache == nil {
		return gerror.New("缓存未初始化")
	}

	resultKey := fmt.Sprintf(consts.FlashSaleResultCacheKey, resultId)
	resultData := map[string]interface{}{
		"user_id":    userId,
		"goods_id":   goodsId,
		"order_no":   orderNo,
		"status":     status,
		"message":    "秒杀成功",
		"pay_amount": 899900, // 模拟金额
		"timestamp":  time.Now().Unix(),
	}

	// 缓存5分钟
	return cache.Set(ctx, resultKey, resultData, 5*time.Minute)
}

// asyncProcessOrder 异步处理订单
func (s *FlashSale) asyncProcessOrder(ctx context.Context, orderNo string, req *v1.CreateFlashSaleOrderReq) {
	// 这里可以发送消息到队列，进行后续处理
	g.Log().Infof(ctx, "异步处理订单: %s, 用户ID: %d, 商品ID: %d", orderNo, req.UserId, req.GoodsId)
}
