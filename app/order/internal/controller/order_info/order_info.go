package order_info

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"
	"shop-goframe-micro-service-refacotor/app/order/api/pbentity"
	order_info "shop-goframe-micro-service-refacotor/app/order/internal/logic/order_info"
	"shop-goframe-micro-service-refacotor/utility/consts"
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
	orderId, err := order_info.Create(ctx, req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.OrderInfoCreateRes{Id: uint32(orderId)}, nil
}

func (*Controller) GetDetail(ctx context.Context, req *v1.OrderInfoGetDetailReq) (res *v1.OrderInfoGetDetailRes, err error) {
	infoError := consts.InfoError(consts.OrderInfo, consts.GetDetailFile)

	// 调用Service层获取订单详情
	pbOrder, pbGoodsList, err := order_info.GetDetail(ctx, req.Id)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	return &v1.OrderInfoGetDetailRes{
		OrderInfo:       pbOrder,
		OrderGoodsInfos: pbGoodsList,
	}, nil
}

func (c *Controller) GetList(ctx context.Context, req *v1.OrderInfoGetListReq) (res *v1.OrderInfoGetListRes, err error) {
	// 初始化响应结构
	response := &v1.OrderInfoListResponse{
		List:  make([]*pbentity.OrderInfo, 0),
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
