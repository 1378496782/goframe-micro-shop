package cmd

import (
	"context"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/internal/controller/banner"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/internal/controller/goods"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/internal/controller/interaction"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/internal/controller/order"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/internal/controller/user"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http gateway-h5 server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 创建控制器实例一次，重复使用

			userController := user.NewV1()
			goodsController := goods.NewV1()
			bannerController := banner.NewV1()
			interactionController := interaction.NewV1()
			orderController := order.NewV1()

			s.Group("/frontend", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				// 无需认证的路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Bind(
						userController.UserInfoRegister,
						userController.UserInfoLogin,
						userController.WxMiniLogin,
						orderController.Notify,
						goodsController.CategoryInfoGetAll,
						goodsController.CategoryInfoGetList,
						goodsController.GoodsInfoGetDetail,
						goodsController.GoodsInfoGetList,
						bannerController,
						goodsController.RecommendGoodsInfoGetList,
					)
				})
				// 需要JWT验证的路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTAuth)
					group.Bind(
						userController.ConsigneeInfoCreate,
						userController.ConsigneeInfoDelete,
						userController.ConsigneeInfoGetList,
						userController.ConsigneeInfoUpdate,
						userController.UserInfo,
						userController.UserInfoUpdatePassword,
						userController.UserInfoUpdate,
						userController.FillPhone,
						goodsController.CartInfoGetList,
						goodsController.CartInfoCreate,
						goodsController.CartInfoDelete,
						goodsController.UserCouponInfoGetList,
						interactionController,
						orderController.Payment,
						orderController.OrderInfoCreate,
						orderController.OrderInfoGetList,
						orderController.OrderInfoGetCount,
						orderController.OrderInfoGetDetail,
						orderController.RefundInfoGetDetail,
						orderController.RefundInfoGetList,
						orderController.RefundInfoCreate,
						goodsController.Bargain_info_Create,
						goodsController.Bargain_info_Get,
						goodsController.Bargain_info_Delete,
						goodsController.Bargain_history_Create,
						goodsController.Bargain_history_Get,
						goodsController.Bargain_history_Delete,
						orderController.CancelOrder,
					)
				})
			})

			// 本地测试微信支付用
			//s.EnableHTTPS("D:/goland/codes/exercise/paymentDemo/cert/shop.dayu.club.pem", "D:/goland/codes/exercise/paymentDemo/cert/shop.dayu.club.key")
			s.Run()
			return nil
		},
	}
)
