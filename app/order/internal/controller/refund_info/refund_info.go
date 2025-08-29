package refund_info

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"shop-goframe-micro-service-refacotor/app/order/api/pbentity"
	v1 "shop-goframe-micro-service-refacotor/app/order/api/refund_info/v1"
	"shop-goframe-micro-service-refacotor/app/order/internal/dao"
	"shop-goframe-micro-service-refacotor/app/order/internal/model/entity"
	"shop-goframe-micro-service-refacotor/utility"

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

	// 售后订单号生成函数
	refund.Number = utility.GenerateRefundNumber()
	refund.Status = 1

	id, err := dao.RefundInfo.Ctx(ctx).InsertAndGetId(refund)
	if err != nil {
		return nil, err
	}

	return &v1.RefundInfoCreateRes{Id: uint32(id)}, nil
}
