package refund_info

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"shop-goframe-micro-service-refacotor/app/order/internal/consts"
	"shop-goframe-micro-service-refacotor/app/order/internal/dao"
)

// UpdateOrderStatus 更新订单状态
func UpdateRefundStatusByNumber(ctx context.Context, refundId string, status int) error {
	exists, err := dao.RefundInfo.Ctx(ctx).
		Where("refund_id", refundId).
		Where("refund_status", consts.RefundOrderStatusSuccess).
		Exist()
	if err != nil {
		return gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	if exists {
		g.Log().Infof(ctx, "{%s}退款记录的状态已修改，不需要再修改", refundId)
		return nil
	}

	updateData := g.Map{
		"refund_status": status,
		"updated_at":    gtime.Now(),
	}

	// 更新退款状态
	_, err = dao.RefundInfo.Ctx(ctx).Where("refund_id", refundId).Update(updateData)
	if err != nil {
		return gerror.WrapCode(gcode.CodeDbOperationError, err)
	}

	g.Log().Infof(ctx, "订单状态更新成功, 订单编号:{%s}, 新状态: %d", refundId, status)
	return nil
}
