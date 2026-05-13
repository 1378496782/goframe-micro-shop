package refund_info

import (
	"context"
	"errors"
	"fmt"
	"shop-goframe-micro-service-refacotor/app/order/internal/consts"
	"shop-goframe-micro-service-refacotor/app/order/internal/dao"
	"shop-goframe-micro-service-refacotor/app/order/internal/model/entity"
	"shop-goframe-micro-service-refacotor/app/order/utility/payment"
	"shop-goframe-micro-service-refacotor/utility"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RefundInfoService 退款服务接口
type RefundInfoService interface {
	// CreateRefund 创建退款申请
	CreateRefund(ctx context.Context, req *CreateRefundReq) (*CreateRefundRes, error)
	// ProcessRefund 处理退款（调用支付平台）
	ProcessRefund(ctx context.Context, refundId int, order *entity.OrderInfo) error
	// UpdateRefundStatus 更新退款状态
	UpdateRefundStatus(ctx context.Context, refundId string, status int) error
	// GetRefundByRefundId 根据第三方退款编号获取退款记录
	GetRefundByRefundId(ctx context.Context, refundId string) (*entity.RefundInfo, error)
	// GetRefundById 根据ID获取退款记录
	GetRefundById(ctx context.Context, id int) (*entity.RefundInfo, error)
	// HandleRefundNotify 处理退款回调
	HandleRefundNotify(ctx context.Context, req interface{}) error
}

// CreateRefundReq 创建退款请求
type CreateRefundReq struct {
	OrderId      int    `json:"orderId"`
	GoodsId      int    `json:"goodsId"`
	Reason       string `json:"reason"`
	RefundAmount int    `json:"refundAmount"`
	UserId       int    `json:"userId"`
}

// CreateRefundRes 创建退款响应
type CreateRefundRes struct {
	Id     int    `json:"id"`
	Number string `json:"number"`
}

// refundInfoService 退款服务实现
type refundInfoService struct{}

// NewRefundInfoService 创建退款服务实例
func NewRefundInfoService() RefundInfoService {
	return &refundInfoService{}
}

// CreateRefund 创建退款申请
func (s *refundInfoService) CreateRefund(ctx context.Context, req *CreateRefundReq) (*CreateRefundRes, error) {
	// 在MVP阶段，简化订单检查
	// 实际环境中需要完整的订单状态验证
	if req.OrderId <= 0 {
		return nil, errors.New("订单ID无效")
	}

	// 查询订单是否已存在退款记录
	exist, _ := dao.RefundInfo.Ctx(ctx).
		Where("order_id", req.OrderId).
		One()
	if !exist.IsEmpty() {
		return nil, errors.New("该订单已存在退款申请，请勿重复操作")
	}

	// 创建退款记录
	refund := &entity.RefundInfo{
		OrderId:      req.OrderId,
		GoodsId:      req.GoodsId,
		Reason:       req.Reason,
		RefundAmount: req.RefundAmount,
		UserId:       req.UserId,
		Number:       utility.GenerateRefundNumber(),
		RefundStatus: int(consts.RefundOrderStatusNone),
	}

	// 由于缺少订单信息，直接设置为待审核状态
	refund.Status = int(consts.RefundStatusPending)

	// 保存退款记录
	id, err := dao.RefundInfo.Ctx(ctx).InsertAndGetId(refund)
	if err != nil {
		return nil, errors.New("创建退款记录失败")
	}

	// 退款创建后由其他流程处理，不立即处理

	return &CreateRefundRes{
		Id:     int(id),
		Number: refund.Number,
	}, nil
}

// ProcessRefund 处理退款
func (s *refundInfoService) ProcessRefund(ctx context.Context, refundId int, order *entity.OrderInfo) error {
	// 获取退款记录
	refund, err := s.GetRefundById(ctx, refundId)
	if err != nil {
		return err
	}

	// 调用支付平台退款接口
	refundReq := &payment.RefundReq{
		TransactionId: order.TransactionId,
		OutRefundNo:   refund.Number,
		Reason:        refund.Reason,
		TotalAmount:   int64(order.ActualPrice),
		RefundAmount:  int64(refund.RefundAmount),
	}

	// 调用微信支付退款
	thirdPartyRefundId, err := payment.Refund(ctx, refundReq)
	if err != nil {
		return err
	}

	// 更新退款状态为处理中
	_, err = dao.RefundInfo.Ctx(ctx).Where("id", refundId).Data(g.Map{
		"refund_status": int(consts.RefundOrderStatusProcessing),
		"refund_id":     thirdPartyRefundId,
		"updated_at":    gtime.Now(),
	}).Update()
	if err != nil {
		g.Log().Errorf(ctx, "更新退款状态失败: %v", err)
		return err
	}

	g.Log().Infof(ctx, "已向微信平台发送退款申请，订单号=%d，退款单号=%s,退款号=%s", order.Id, refund.Number, thirdPartyRefundId)
	return nil
}

// RetryProcessRefund 重试处理退款
func (s *refundInfoService) RetryProcessRefund(ctx context.Context, refundId int, orderInfo map[string]interface{}) {
	maxRetries := 3
	retryInterval := 5 * time.Second

	// 获取退款记录
	refund, err := s.GetRefundById(ctx, refundId)
	if err != nil {
		g.Log().Errorf(ctx, "获取退款记录失败: %v", err)
		return
	}

	// 创建模拟订单信息
	order := &entity.OrderInfo{
		Id:            refund.OrderId,
		TransactionId: fmt.Sprintf("TRANS_%d", refund.OrderId),
		ActualPrice:   refund.RefundAmount,
	}

	for i := 0; i < maxRetries; i++ {
		err := s.ProcessRefund(ctx, refundId, order)
		if err == nil {
			g.Log().Infof(ctx, "退款处理成功，退款ID: %d", refundId)
			return
		}

		g.Log().Errorf(ctx, "退款处理失败，将在%f秒后重试，第%d次尝试: %v", retryInterval.Seconds(), i+1, err)
		time.Sleep(retryInterval)
		// 指数退避
		retryInterval = retryInterval * 2
	}

	g.Log().Errorf(ctx, "退款处理失败，已达到最大重试次数，退款ID: %d", refundId)
}

// UpdateRefundStatus 更新退款状态
func (s *refundInfoService) UpdateRefundStatus(ctx context.Context, refundId string, status int) error {
	// 检查是否已经是成功状态
	exists, err := dao.RefundInfo.Ctx(ctx).
		Where("refund_id", refundId).
		Where("refund_status", consts.RefundOrderStatusSuccess).
		Exist()
	if err != nil {
		return err
	}
	if exists {
		g.Log().Infof(ctx, "{%s}退款记录的状态已修改，不需要再修改", refundId)
		return nil
	}

	// 更新退款状态
	_, err = dao.RefundInfo.Ctx(ctx).Where("refund_id", refundId).Data(g.Map{
		"refund_status": status,
		"updated_at":    gtime.Now(),
	}).Update()
	if err != nil {
		return err
	}

	g.Log().Infof(ctx, "退款状态更新成功, 退款编号:{%s}, 新状态: %d", refundId, status)
	return nil
}

// GetRefundByRefundId 根据第三方退款编号获取退款记录
func (s *refundInfoService) GetRefundByRefundId(ctx context.Context, refundId string) (*entity.RefundInfo, error) {
	var refund entity.RefundInfo
	err := dao.RefundInfo.Ctx(ctx).Where("refund_id", refundId).Scan(&refund)
	if err != nil {
		return nil, err
	}
	if refund.Id == 0 {
		return nil, errors.New("退款记录不存在")
	}
	return &refund, nil
}

// GetRefundById 根据ID获取退款记录
func (s *refundInfoService) GetRefundById(ctx context.Context, id int) (*entity.RefundInfo, error) {
	var refund entity.RefundInfo
	err := dao.RefundInfo.Ctx(ctx).WherePri(id).Scan(&refund)
	if err != nil {
		return nil, err
	}
	if refund.Id == 0 {
		return nil, errors.New("退款记录不存在")
	}
	return &refund, nil
}

// HandleRefundNotify 处理退款回调
func (s *refundInfoService) HandleRefundNotify(ctx context.Context, req interface{}) error {
	// 调用支付模块处理回调
	refundId, err := payment.RefundNotify(ctx, req)
	if err != nil {
		return err
	}

	// 更新退款状态为成功
	return s.UpdateRefundStatus(ctx, refundId, int(consts.RefundOrderStatusSuccess))
}

// 默认的退款服务实例
var defaultService = NewRefundInfoService()

// CreateRefund 全局便捷函数
func CreateRefund(ctx context.Context, req *CreateRefundReq) (*CreateRefundRes, error) {
	return defaultService.CreateRefund(ctx, req)
}

// ProcessRefund 全局便捷函数
func ProcessRefund(ctx context.Context, refundId int, order *entity.OrderInfo) error {
	return defaultService.ProcessRefund(ctx, refundId, order)
}

// UpdateRefundStatus 全局便捷函数
func UpdateRefundStatus(ctx context.Context, refundId string, status int) error {
	return defaultService.UpdateRefundStatus(ctx, refundId, status)
}

// GetRefundByRefundId 全局便捷函数
func GetRefundByRefundId(ctx context.Context, refundId string) (*entity.RefundInfo, error) {
	return defaultService.GetRefundByRefundId(ctx, refundId)
}

// GetRefundById 全局便捷函数
func GetRefundById(ctx context.Context, id int) (*entity.RefundInfo, error) {
	return defaultService.GetRefundById(ctx, id)
}

// HandleRefundNotify 全局便捷函数
func HandleRefundNotify(ctx context.Context, req interface{}) error {
	return defaultService.HandleRefundNotify(ctx, req)
}
