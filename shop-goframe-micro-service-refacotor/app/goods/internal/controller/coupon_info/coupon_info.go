package coupon_info

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "shop-goframe-micro-service-refacotor/app/goods/api/coupon_info/v1"
	"shop-goframe-micro-service-refacotor/app/goods/api/pbentity"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/app/goods/internal/model/entity"
	"shop-goframe-micro-service-refacotor/utility"
	"shop-goframe-micro-service-refacotor/utility/consts"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedCouponInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterCouponInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.CouponInfoGetListReq) (res *v1.CouponInfoGetListRes, err error) {
	// 错误类型
	infoError := consts.InfoError(consts.CouponInfo, consts.GetListFail)
	// 初始化响应结构
	response := &v1.CouponInfoListResponse{
		List:  make([]*pbentity.CouponInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}

	// 查询总数
	total, err := dao.CouponInfo.Ctx(ctx).Count()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	// 查询当前页数据
	couponRecords, err := dao.CouponInfo.Ctx(ctx).
		Page(int(req.Page), int(req.Size)).
		All()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 数据转换
	// 在循环中替换手动赋值
	for _, record := range couponRecords {
		var coupon entity.CouponInfo
		if err := record.Struct(&coupon); err != nil {
			continue
		}

		var pbCoupon pbentity.CouponInfo
		if err := gconv.Struct(coupon, &pbCoupon); err != nil {
			continue
		}

		// 单独处理时间字段（gconv无法自动转换）
		pbCoupon.Deadline = utility.SafeConvertTime(coupon.Deadline)
		pbCoupon.CreatedAt = utility.SafeConvertTime(coupon.CreatedAt)
		pbCoupon.UpdatedAt = utility.SafeConvertTime(coupon.UpdatedAt)
		pbCoupon.DeletedAt = utility.SafeConvertTime(coupon.DeletedAt)

		response.List = append(response.List, &pbCoupon)
	}
	return &v1.CouponInfoGetListRes{Data: response}, nil
}

func (*Controller) Create(ctx context.Context, req *v1.CouponInfoCreateReq) (res *v1.CouponInfoCreateRes, err error) {
	// 错误类型
	infoError := consts.InfoError(consts.CouponInfo, consts.CreateFail)
	id, err := dao.CouponInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	return &v1.CouponInfoCreateRes{Id: uint32(id)}, nil
}

func (*Controller) Update(ctx context.Context, req *v1.CouponInfoUpdateReq) (res *v1.CouponInfoUpdateRes, err error) {
	// 错误类型
	infoError := consts.InfoError(consts.CouponInfo, consts.UpdateFail)
	_, err = dao.CouponInfo.Ctx(ctx).Where("id", req.Id).Update(req)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	return &v1.CouponInfoUpdateRes{Id: req.Id}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.CouponInfoDeleteReq) (res *v1.CouponInfoDeleteRes, err error) {
	// 错误类型
	infoError := consts.InfoError(consts.CouponInfo, consts.DeleteFail)
	// 只需要关注是否出错，不返回被删除的数据
	_, err = dao.CouponInfo.Ctx(ctx).Where("id", req.Id).Delete()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.CouponInfoDeleteRes{}, nil // 返回空结构体
}
