package user_coupon_info

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/app/goods/internal/model/entity"
	"shop-goframe-micro-service-refacotor/utility/consts"
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
