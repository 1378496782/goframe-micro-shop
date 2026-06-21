package user_coupon_info

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"shop-goframe-micro-service-refacotor/app/goods/api/pbentity"
	v1 "shop-goframe-micro-service-refacotor/app/goods/api/user_coupon_info/v1"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/app/goods/internal/model/entity"
	"shop-goframe-micro-service-refacotor/utility"
	"shop-goframe-micro-service-refacotor/utility/consts"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedUserCouponInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserCouponInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.UserCouponInfoGetListReq) (res *v1.UserCouponInfoGetListRes, err error) {
	// 错误类型
	infoError := consts.InfoError(consts.UserCouponInfo, consts.GetListFail)
	// 初始化响应结构
	response := &v1.UserCouponInfoListResponse{
		List:  make([]*pbentity.UserCouponInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}

	// 查询总数
	total, err := dao.UserCouponInfo.Ctx(ctx).Count()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	// 查询当前页数据
	couponRecords, err := dao.UserCouponInfo.Ctx(ctx).
		Page(int(req.Page), int(req.Size)).
		Where("user_id", req.UserId).
		All()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 数据转换
	// 在循环中替换手动赋值
	for _, record := range couponRecords {
		var coupon entity.UserCouponInfo
		if err := record.Struct(&coupon); err != nil {
			continue
		}

		var pbUserCoupon pbentity.UserCouponInfo
		if err := gconv.Struct(coupon, &pbUserCoupon); err != nil {
			continue
		}

		// 单独处理时间字段（gconv无法自动转换）
		pbUserCoupon.CreatedAt = utility.SafeConvertTime(coupon.CreatedAt)
		pbUserCoupon.UpdatedAt = utility.SafeConvertTime(coupon.UpdatedAt)
		pbUserCoupon.DeletedAt = utility.SafeConvertTime(coupon.DeletedAt)

		response.List = append(response.List, &pbUserCoupon)
	}
	return &v1.UserCouponInfoGetListRes{Data: response}, nil
}

func (*Controller) Create(ctx context.Context, req *v1.UserCouponInfoCreateReq) (res *v1.UserCouponInfoCreateRes, err error) {
	// 错误类型
	infoError := consts.InfoError(consts.UserCouponInfo, consts.CreateFail)
	id, err := dao.UserCouponInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	return &v1.UserCouponInfoCreateRes{Id: uint32(id)}, nil
}

func (*Controller) Update(ctx context.Context, req *v1.UserCouponInfoUpdateReq) (res *v1.UserCouponInfoUpdateRes, err error) {
	// 错误类型
	infoError := consts.InfoError(consts.UserCouponInfo, consts.UpdateFail)
	_, err = dao.UserCouponInfo.Ctx(ctx).Where("id", req.Id).Update(req)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	return &v1.UserCouponInfoUpdateRes{Id: req.Id}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.UserCouponInfoDeleteReq) (res *v1.UserCouponInfoDeleteRes, err error) {
	// 错误类型
	infoError := consts.InfoError(consts.UserCouponInfo, consts.DeleteFail)
	// 只需要关注是否出错，不返回被删除的数据
	_, err = dao.UserCouponInfo.Ctx(ctx).Where("id", req.Id).Delete()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.UserCouponInfoDeleteRes{}, nil // 返回空结构体
}
