package service

import (
	v1 "shop-goframe-micro-service-refacotor/app/flash-sale/api/flash_sale/v1"
	"shop-goframe-micro-service-refacotor/app/flash-sale/utility"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// FlashSaleController 秒杀控制器
type FlashSaleController struct{}

// GetFlashSaleGoodsList 获取秒杀商品列表
func (c *FlashSaleController) GetFlashSaleGoodsList(r *ghttp.Request) {
	ctx := r.Context()

	var req v1.FlashSaleGoodsListReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code":    400,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}

	// 调用服务
	res, err := utility.GetFlashSaleService().GetFlashSaleGoodsList(ctx, &req)
	if err != nil {
		g.Log().Error(ctx, "获取秒杀商品列表失败:", err)
		r.Response.WriteJsonExit(g.Map{
			"code":    500,
			"message": "系统错误",
			"data":    nil,
		})
		return
	}

	// 类型转换
	goodsList, ok := res.(*v1.FlashSaleGoodsListRes)
	if !ok {
		g.Log().Error(ctx, "响应类型转换失败")
		r.Response.WriteJsonExit(g.Map{
			"code":    500,
			"message": "系统错误",
			"data":    nil,
		})
		return
	}

	r.Response.WriteJsonExit(g.Map{
		"code":    0,
		"message": "success",
		"data":    goodsList,
	})
}

// GetFlashSaleGoodsDetail 获取秒杀商品详情
func (c *FlashSaleController) GetFlashSaleGoodsDetail(r *ghttp.Request) {
	ctx := r.Context()

	var req v1.FlashSaleGoodsDetailReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code":    400,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}

	// 调用服务
	res, err := utility.GetFlashSaleService().GetFlashSaleGoodsDetail(ctx, &req)
	if err != nil {
		g.Log().Error(ctx, "获取秒杀商品详情失败:", err)
		r.Response.WriteJsonExit(g.Map{
			"code":    500,
			"message": "系统错误",
			"data":    nil,
		})
		return
	}

	// 类型转换
	goodsDetail, ok := res.(*v1.FlashSaleGoodsDetailRes)
	if !ok {
		g.Log().Error(ctx, "响应类型转换失败")
		r.Response.WriteJsonExit(g.Map{
			"code":    500,
			"message": "系统错误",
			"data":    nil,
		})
		return
	}

	r.Response.WriteJsonExit(g.Map{
		"code":    0,
		"message": "success",
		"data":    goodsDetail,
	})
}

// CreateFlashSaleOrder 创建秒杀订单
func (c *FlashSaleController) CreateFlashSaleOrder(r *ghttp.Request) {
	ctx := r.Context()

	// 获取用户ID（从token或session）
	userId := gconv.Uint32(r.GetHeader("X-User-Id"))
	if userId == 0 {
		// 模拟用户ID，实际应该从认证信息获取
		userId = 10001
	}

	var req v1.CreateFlashSaleOrderReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code":    400,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}

	// 设置用户ID
	req.UserId = userId

	// 用户限流检查
	if err := utility.UserRateLimit(ctx, userId, utility.GetFlashSaleCache()); err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code":    429,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// IP限流检查
	clientIP := utility.GetClientIP(ctx)
	if err := utility.IPRateLimit(ctx, clientIP, utility.GetFlashSaleCache()); err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code":    429,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用服务
	res, err := utility.GetFlashSaleService().CreateFlashSaleOrder(ctx, &req)
	if err != nil {
		g.Log().Error(ctx, "创建秒杀订单失败:", err)
		r.Response.WriteJsonExit(g.Map{
			"code":    500,
			"message": "系统错误",
			"data":    nil,
		})
		return
	}

	// 类型转换
	orderRes, ok := res.(*v1.CreateFlashSaleOrderRes)
	if !ok {
		g.Log().Error(ctx, "响应类型转换失败")
		r.Response.WriteJsonExit(g.Map{
			"code":    500,
			"message": "系统错误",
			"data":    nil,
		})
		return
	}

	if !orderRes.Success {
		r.Response.WriteJsonExit(g.Map{
			"code":    400,
			"message": orderRes.Message,
			"data":    orderRes,
		})
		return
	}

	r.Response.WriteJsonExit(g.Map{
		"code":    0,
		"message": "success",
		"data":    orderRes,
	})
}

// GetFlashSaleResult 查询秒杀结果
func (c *FlashSaleController) GetFlashSaleResult(r *ghttp.Request) {
	ctx := r.Context()

	// 获取用户ID（从token或session）
	userId := gconv.Uint32(r.GetHeader("X-User-Id"))
	if userId == 0 {
		// 模拟用户ID，实际应该从认证信息获取
		userId = 10001
	}

	var req v1.GetFlashSaleResultReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code":    400,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}

	// 设置用户ID
	req.UserId = userId

	// 调用服务
	res, err := utility.GetFlashSaleService().GetFlashSaleResult(ctx, &req)
	if err != nil {
		g.Log().Error(ctx, "查询秒杀结果失败:", err)
		r.Response.WriteJsonExit(g.Map{
			"code":    500,
			"message": "系统错误",
			"data":    nil,
		})
		return
	}

	// 类型转换
	resultRes, ok := res.(*v1.GetFlashSaleResultRes)
	if !ok {
		g.Log().Error(ctx, "响应类型转换失败")
		r.Response.WriteJsonExit(g.Map{
			"code":    500,
			"message": "系统错误",
			"data":    nil,
		})
		return
	}

	r.Response.WriteJsonExit(g.Map{
		"code":    0,
		"message": "success",
		"data":    resultRes,
	})
}
