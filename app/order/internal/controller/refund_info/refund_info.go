package refund_info

import (
	"context"
	"shop-goframe-micro-service-refacotor/app/order/api/pbentity"
	v1 "shop-goframe-micro-service-refacotor/app/order/api/refund_info/v1"
	"shop-goframe-micro-service-refacotor/app/order/internal/consts"
	"shop-goframe-micro-service-refacotor/app/order/internal/dao"
	"shop-goframe-micro-service-refacotor/app/order/internal/logic/refund_info"
	"shop-goframe-micro-service-refacotor/app/order/internal/model/entity"
	"shop-goframe-micro-service-refacotor/app/order/utility/payment"
	"shop-goframe-micro-service-refacotor/utility"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedRefundInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterRefundInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.RefundInfoGetListReq) (res *v1.RefundInfoGetListRes, err error) {
	// 初始化响应结构
	response := &v1.RefundInfoListResponse{
		List:  make([]*pbentity.RefundInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}

	// 查询总数
	total, err := dao.RefundInfo.Ctx(ctx).Count()
	if err != nil {
		return &v1.RefundInfoGetListRes{Data: response}, nil
	}
	response.Total = uint32(total)

	// 查询当前页数据
	refundRecords, err := dao.RefundInfo.Ctx(ctx).
		Page(int(req.Page), int(req.Size)).
		All()
	if err != nil {
		return &v1.RefundInfoGetListRes{Data: response}, nil
	}

	// 数据转换
	// 在循环中替换手动赋值
	for _, record := range refundRecords {
		var refund entity.RefundInfo
		if err := record.Struct(&refund); err != nil {
			continue
		}

		var pbRefund pbentity.RefundInfo
		if err := gconv.Struct(refund, &pbRefund); err != nil {
			continue
		}

		// 单独处理时间字段（gconv无法自动转换）
		pbRefund.CreatedAt = utility.SafeConvertTime(refund.CreatedAt)
		pbRefund.UpdatedAt = utility.SafeConvertTime(refund.UpdatedAt)

		response.List = append(response.List, &pbRefund)
	}
	return &v1.RefundInfoGetListRes{Data: response}, nil
}

func (*Controller) GetDetail(ctx context.Context, req *v1.RefundInfoGetDetailReq) (res *v1.RefundInfoGetDetailRes, err error) {
	// 查询退款记录
	record, err := dao.RefundInfo.Ctx(ctx).Where("id", req.Id).One()
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "查询退款记录失败")
	}
	if record.IsEmpty() {
		return nil, gerror.NewCode(gcode.CodeNotFound, "退款记录不存在")
	}

	// 转换为实体结构
	var refund entity.RefundInfo
	if err := record.Struct(&refund); err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "数据转换失败")
	}

	// 转换为protobuf结构
	var pbRefund pbentity.RefundInfo
	if err := gconv.Struct(refund, &pbRefund); err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "数据转换失败")
	}

	// 处理时间字段
	pbRefund.CreatedAt = utility.SafeConvertTime(refund.CreatedAt)
	pbRefund.UpdatedAt = utility.SafeConvertTime(refund.UpdatedAt)

	return &v1.RefundInfoGetDetailRes{
		Data: &pbRefund,
	}, nil
}

func (*Controller) Create(ctx context.Context, req *v1.RefundInfoCreateReq) (res *v1.RefundInfoCreateRes, err error) {
	var refund *entity.RefundInfo
	if err := gconv.Struct(req, &refund); err != nil {
		return nil, err
	}

	// 查询订单是否存在以及订单状态是否是已付款
	order := &entity.OrderInfo{}
	err = dao.OrderInfo.Ctx(ctx).WherePri(req.OrderId).Scan(order)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "查询订单记录失败")
	}
	if order == nil || consts.OrderStatus(order.Status) == consts.OrderStatusPendingPayment ||
		consts.OrderStatus(order.Status) == consts.OrderStatusCancelled {
		return nil, gerror.New("订单不存在或状态不支持退款")
	}

	// 查询订单是否已存在退款记录
	exist, _ := dao.RefundInfo.Ctx(ctx).
		Where("order_id", req.OrderId).
		One()
	if !exist.IsEmpty() {
		return nil, gerror.New("该订单已存在退款申请，请勿重复操作")
	}

	// 售后订单号生成函数
	refund.Number = utility.GenerateRefundNumber()
	refund.RefundStatus = int(consts.RefundOrderStatusNone)
	switch consts.OrderStatus(order.Status) {
	// 已支付未发货的情况
	case consts.OrderStatusPaid:
		refund.Status = int(consts.RefundStatusApproved)
	default:
		refund.Status = int(consts.RefundStatusPending)
	}

	id, err := dao.RefundInfo.Ctx(ctx).InsertAndGetId(refund)
	if err != nil {
		return nil, err
	}

	// todo 优化：微信支付接口失败，需要有一个重试机制
	if refund.Status == int(consts.RefundStatusApproved) {
		refundReq := &payment.RefundReq{
			TransactionId: order.TransactionId,
			OutRefundNo:   order.Number,
			Reason:        req.Reason,
			TotalAmount:   int64(order.ActualPrice),
			RefundAmount:  int64(order.ActualPrice),
		}
		refundId, err := payment.Refund(ctx, refundReq)
		if err != nil {
			return nil, err
		}
		g.Log().Infof(ctx, "已向微信平台发送退款申请，订单号=%d，退款单号=%s,退款号=%s", order.Id, refund.Number, refundId)
		_, err = dao.RefundInfo.Ctx(ctx).Where("id", id).Data(g.Map{
			"refund_status": int(consts.RefundOrderStatusProcessing),
			"refund_id":     refundId,
		}).Update()
		if err != nil {
			g.Log().Errorf(ctx, "%v,%v", err, "更新退款状态失败")
		}
	}
	return &v1.RefundInfoCreateRes{Id: uint32(id)}, nil
}

func (*Controller) RefundNotify(ctx context.Context, req *v1.RefundNotifyReq) (res *v1.RefundNotifyRes, err error) {
	// 1) 微信支付回调验证
	refundId, err := payment.RefundNotify(ctx, req)
	if err != nil {
		return nil, err
	}

	// 2) 修改订单状态
	if err = refund_info.UpdateRefundStatusByNumber(ctx, refundId, int(consts.RefundOrderStatusSuccess)); err != nil {
		return nil, err
	}

	return nil, nil
}
