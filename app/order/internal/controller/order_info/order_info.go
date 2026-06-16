package order_info

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"
	orderStatus "shop-goframe-micro-service-refacotor/app/order/internal/consts"
	order_info "shop-goframe-micro-service-refacotor/app/order/internal/logic/order_info"
	"shop-goframe-micro-service-refacotor/app/order/utility/payment"
	"shop-goframe-micro-service-refacotor/utility/consts"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct {
	v1.UnimplementedOrderInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterOrderInfoServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.OrderInfoCreateReq) (res *v1.OrderInfoCreateRes, err error) {
	infoError := consts.InfoError(consts.OrderInfo, consts.CreateFail)
	// 调用login层创建订单
	orderId, orderNumber, err := order_info.Create(ctx, req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.OrderInfoCreateRes{Id: uint32(orderId), Number: orderNumber}, nil
}

func (*Controller) GetDetail(ctx context.Context, req *v1.OrderInfoGetDetailReq) (res *v1.OrderInfoGetDetailRes, err error) {
	infoError := consts.InfoError(consts.OrderInfo, consts.GetDetailFail)

	// 调用Service层获取订单详情
	pbOrder, pbGoodsList, err := order_info.GetDetail(ctx, req.Id, req.UserId)
	if err != nil {
		// Check for specific errors from logic layer
		if gerror.Code(err) == gcode.CodeNotAuthorized {
			g.Log().Warningf(ctx, "User %d attempted to access order %d without permission", req.UserId, req.Id)
			return nil, err // Forward the permission denied error
		}
		if gerror.Code(err) == gcode.CodeNotFound {
			g.Log().Debugf(ctx, "Order %d not found", req.Id)
			return nil, err // Forward the not found error
		}
		// General error handling, same as Create
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	return &v1.OrderInfoGetDetailRes{
		OrderInfo:       pbOrder,
		OrderGoodsInfos: pbGoodsList,
	}, nil
}

// getlist方法 v2
func (c *Controller) GetList(ctx context.Context, req *v1.OrderInfoGetListReq) (res *v1.OrderInfoGetListRes, err error) {
	// 初始化响应结构
	response := &v1.OrderInfoListResponse{
		List:  make([]*v1.OrderListInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}

	infoError := consts.InfoError(consts.OrderInfo, consts.GetListFail)

	// 初始化分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 || req.Size > 50 {
		req.Size = 10
	}

	// 调用Service层获取数据
	pbOrders, total, err := order_info.GetList(ctx, req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, "查询订单列表失败")
	}

	// 设置响应数据
	response.List = pbOrders
	response.Total = uint32(total)
	response.Page = req.Page
	response.Size = req.Size

	return &v1.OrderInfoGetListRes{Data: response}, nil
}

func (*Controller) Payment(ctx context.Context, req *v1.PaymentReq) (res *v1.PaymentRes, err error) {
	return payment.WeChatPayment(ctx, req)
}

func (*Controller) Notify(ctx context.Context, req *v1.NotifyReq) (res *v1.NotifyRes, err error) {
	// 1) 微信支付回调验证
	orderNumber, transactionId, err := payment.Notify(ctx, req)
	if err != nil {
		return nil, err
	}

	// 2) 修改订单状态
	success, err := order_info.UpdateOrderStatusByNumber(ctx, orderNumber, transactionId, int(orderStatus.OrderStatusPaid))
	if err != nil {
		return nil, err
	}
	if !success {
		return &v1.NotifyRes{}, nil
	}

	return &v1.NotifyRes{}, nil
}

func (*Controller) GetCount(ctx context.Context, req *v1.OrderInfoGetCountReq) (res *v1.OrderInfoGetCountRes, err error) {
	infoError := consts.InfoError(consts.OrderInfo, consts.GetCountFail)
	res, err = order_info.GetCount(ctx, req.UserId)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return res, nil
}

func (*Controller) CancelOrder(ctx context.Context, req *v1.CancelOrderReq) (res *v1.CancelOrderRes, err error) {
	return order_info.CancelOrder(ctx, req)
}

func (*Controller) Preview(ctx context.Context, req *v1.OrderInfoPreviewReq) (res *v1.OrderInfoPreviewRes, err error) {
	res, err = order_info.Preview(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (*Controller) CreateFromCart(ctx context.Context, req *v1.OrderInfoCreateFromCartReq) (res *v1.OrderInfoCreateRes, err error) {
	// 调用login层创建订单
	orderId, orderNumber, err := order_info.CreateFromCart(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.OrderInfoCreateRes{
		Id:     uint32(orderId),
		Number: orderNumber,
	}, nil
}
