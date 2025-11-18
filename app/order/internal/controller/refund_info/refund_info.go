package refund_info

import (
	"context"
	"shop-goframe-micro-service-refacotor/app/order/api/pbentity"
	v1 "shop-goframe-micro-service-refacotor/app/order/api/refund_info/v1"
	"shop-goframe-micro-service-refacotor/app/order/internal/dao"
	"shop-goframe-micro-service-refacotor/app/order/internal/model/entity"
	"shop-goframe-micro-service-refacotor/app/order/utility/payment"
	"shop-goframe-micro-service-refacotor/utility"
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
	total, err := dao.RefundInfo.GetModel().Ctx(ctx).Count()
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "查询退款记录总数失败")
	}
	response.Total = uint32(total)

	// 查询当前页数据
	refundRecords, err := dao.RefundInfo.GetModel().Ctx(ctx).
		Page(int(req.Page), int(req.Size)).
		All()
	if err != nil {
		return &v1.RefundInfoGetListRes{Data: response}, gerror.WrapCode(gcode.CodeInternalError, err, "查询退款记录失败")
	}

	// 数据转换
	for _, record := range refundRecords {
		var refund entity.RefundInfo
		if err := record.Struct(&refund); err != nil {
			continue
		}

		var pbRefund pbentity.RefundInfo
		if err := gconv.Struct(refund, &pbRefund); err != nil {
			continue
		}

		// 单独处理时间字段
		pbRefund.CreatedAt = utility.SafeConvertTime(refund.CreatedAt)
		pbRefund.UpdatedAt = utility.SafeConvertTime(refund.UpdatedAt)

		response.List = append(response.List, &pbRefund)
	}
	return &v1.RefundInfoGetListRes{Data: response}, nil
}

func (*Controller) GetDetail(ctx context.Context, req *v1.RefundInfoGetDetailReq) (res *v1.RefundInfoGetDetailRes, err error) {
	// 查询退款记录
	record, err := dao.RefundInfo.GetModel().Ctx(ctx).Where("id", req.Id).One()
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
	// 直接使用原有逻辑处理退款创建，以确保功能正常
	var refund entity.RefundInfo
	if err := gconv.Struct(req, &refund); err != nil {
		return nil, err
	}

	// 使用延迟初始化的DAO方法
	// 查询订单是否已存在退款记录
	exist, _ := dao.RefundInfo.GetModel().Ctx(ctx).
		Where("order_id", req.OrderId).
		One()
	if !exist.IsEmpty() {
		return nil, gerror.New("该订单已存在退款申请，请勿重复操作")
	}

	// 售后订单号生成函数
	refund.Number = utility.GenerateRefundNumber()
	refund.RefundStatus = 0 // 初始状态
	refund.Status = 1       // 待处理状态

	id, err := dao.RefundInfo.GetModel().Ctx(ctx).InsertAndGetId(refund)
	if err != nil {
		return nil, err
	}

	// 启动goroutine异步处理退款（模拟服务层功能）
	go func() {
		// 异步处理逻辑占位
	}()

	return &v1.RefundInfoCreateRes{Id: uint32(id)}, nil
}

func (*Controller) RefundNotify(ctx context.Context, req *v1.RefundNotifyReq) (res *v1.RefundNotifyRes, err error) {
	// 1) 微信支付回调验证
	refundId, err := payment.RefundNotify(ctx, req)
	if err != nil {
		return nil, err
	}

	// 2) 直接更新退款状态
	_, err = dao.RefundInfo.GetModel().Ctx(ctx).
		Where("number", refundId).
		Data(map[string]interface{}{
			"refund_status": 2, // 退款成功状态
			"refund_id":     refundId,
		}).Update()
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "更新退款状态失败")
	}

	return nil, nil
}
