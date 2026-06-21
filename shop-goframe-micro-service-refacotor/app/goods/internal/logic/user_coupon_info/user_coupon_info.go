package user_coupon_info

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/app/goods/internal/model/entity"
	"shop-goframe-micro-service-refacotor/utility/consts"
	"shop-goframe-micro-service-refacotor/utility/rabbitmq"
)

// IssueCouponToUser 发放优惠券给用户
func IssueCouponToUser(ctx context.Context, userID int) error {
	// 错误类型
	infoError := consts.InfoError(consts.UserCouponInfo, consts.CreateFail)

	// 构建优惠券数据
	userCouponData := &entity.UserCouponInfo{
		UserId:   userID,
		CouponId: 1, //新人优惠券写死id为1
		Status:   0, // 0表示未使用
		Amount:   9999,
	}

	// 插入数据库
	_, err := dao.UserCouponInfo.Ctx(ctx).Insert(userCouponData)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return nil
}

// HandleOrderConfirmMessage 处理订单确认消息
// 通过userid和couponid在user_coupon_info表中定位数据
// 如果找到数据且状态为0（待使用），则修改为1（已使用）并返回成功
// 如果未找到数据或状态不是0（待使用），则返回失败
func HandleOrderConfirmMessage(ctx context.Context, orderId int, userId int, couponId int) error {

	// 查询用户优惠券信息
	var userCoupon entity.UserCouponInfo
	err := dao.UserCouponInfo.Ctx(ctx).
		Where("user_id", userId).
		Where("coupon_id", couponId).
		Where("status", 0). // 0表示待使用
		Scan(&userCoupon)

	if err != nil {
		// 未找到数据或查询失败
		g.Log().Warningf(ctx, "查询用户优惠券失败或未找到数据, 用户ID: %d, 优惠券ID: %d, 错误: %v", userId, couponId, err)
		// 发送失败消息
		rabbitmq.PublishCouponConfirmResultEvent(orderId, false, "优惠券不存在或已使用")
		return nil
	}

	// 更新优惠券状态为已使用
	_, updateErr := dao.UserCouponInfo.Ctx(ctx).
		Where("id", userCoupon.Id).
		Update(g.Map{"status": 1}) // 1表示已使用

	if updateErr != nil {
		g.Log().Errorf(ctx, "更新用户优惠券状态失败, ID: %d, 错误: %v", userCoupon.Id, updateErr)
		// 发送失败消息
		rabbitmq.PublishCouponConfirmResultEvent(orderId, false, "更新优惠券状态失败")
		return updateErr
	}

	g.Log().Infof(ctx, "优惠券使用成功, 用户ID: %d, 优惠券ID: %d", userId, couponId)
	// 发送成功消息
	rabbitmq.PublishCouponConfirmResultEvent(orderId, true, "优惠券使用成功")
	return nil
}
