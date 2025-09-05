package position_info

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"shop-goframe-micro-service-refacotor/app/banner/api/pbentity"
	v1 "shop-goframe-micro-service-refacotor/app/banner/api/position_info/v1"
	"shop-goframe-micro-service-refacotor/app/banner/internal/dao"
	"shop-goframe-micro-service-refacotor/app/banner/internal/model/entity"
	"shop-goframe-micro-service-refacotor/utility"
	"shop-goframe-micro-service-refacotor/utility/consts"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedPositionInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterPositionInfoServer(s.Server, &Controller{})
}

func (c *Controller) GetList(ctx context.Context, req *v1.PositionInfoGetListReq) (*v1.PositionInfoGetListRes, error) {
	// 错误类型
	infoError := consts.InfoError(consts.PositionInfo, consts.GetListFail)
	// 初始化响应结构
	response := &v1.PositionInfoListResponse{
		List:  make([]*pbentity.PositionInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}

	// 查询总数
	total, err := dao.PositionInfo.Ctx(ctx).Count()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	// 查询当前页数据
	// 查询当前页数据，sort值越小越靠前
	positionRecords, err := dao.PositionInfo.Ctx(ctx).
		Order(utility.GetOrderBy(req.Sort)). // sort=2倒序 默认升序
		Page(int(req.Page), int(req.Size)).
		All()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 数据转换
	// 在循环中替换手动赋值
	for _, record := range positionRecords {
		var position entity.PositionInfo
		if err := record.Struct(&position); err != nil {
			continue
		}

		var pbPosition pbentity.PositionInfo
		if err := gconv.Struct(position, &pbPosition); err != nil {
			continue
		}

		// 单独处理时间字段（gconv无法自动转换）
		pbPosition.CreatedAt = utility.SafeConvertTime(position.CreatedAt)
		pbPosition.UpdatedAt = utility.SafeConvertTime(position.UpdatedAt)
		pbPosition.DeletedAt = utility.SafeConvertTime(position.DeletedAt)

		response.List = append(response.List, &pbPosition)
	}
	return &v1.PositionInfoGetListRes{Data: response}, nil
}

// Create 创建
func (c *Controller) Create(ctx context.Context, req *v1.PositionInfoCreateReq) (*v1.PositionInfoCreateRes, error) {
	// 错误类型
	infoError := consts.InfoError(consts.PositionInfo, consts.CreateFail)
	id, err := dao.PositionInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	return &v1.PositionInfoCreateRes{Id: uint32(id)}, nil
}

// Update 更新
func (c *Controller) Update(ctx context.Context, req *v1.PositionInfoUpdateReq) (*v1.PositionInfoUpdateRes, error) {
	// 错误类型
	infoError := consts.InfoError(consts.PositionInfo, consts.UpdateFail)
	_, err := dao.PositionInfo.Ctx(ctx).Where("id", req.Id).Update(req)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	return &v1.PositionInfoUpdateRes{Id: req.Id}, nil
}

// Delete 删除
func (c *Controller) Delete(ctx context.Context, req *v1.PositionInfoDeleteReq) (*v1.PositionInfoDeleteRes, error) {
	// 错误类型
	infoError := consts.InfoError(consts.PositionInfo, consts.DeleteFail)
	// 只需要关注是否出错，不返回被删除的数据
	_, err := dao.PositionInfo.Ctx(ctx).Where("id", req.Id).Delete()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.PositionInfoDeleteRes{}, nil // 返回空结构体
}
