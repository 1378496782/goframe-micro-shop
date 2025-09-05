package rotation_info

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"shop-goframe-micro-service-refacotor/app/banner/api/pbentity"
	v1 "shop-goframe-micro-service-refacotor/app/banner/api/rotation_info/v1"
	"shop-goframe-micro-service-refacotor/app/banner/internal/dao"
	"shop-goframe-micro-service-refacotor/app/banner/internal/model/entity"
	"shop-goframe-micro-service-refacotor/utility"
	"shop-goframe-micro-service-refacotor/utility/consts"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedRotationInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterRotationInfoServer(s.Server, &Controller{})
}

func (c *Controller) GetList(ctx context.Context, req *v1.RotationInfoGetListReq) (*v1.RotationInfoGetListRes, error) {
	// 错误类型
	infoError := consts.InfoError(consts.RotationInfo, consts.GetListFail)
	// 初始化响应结构
	response := &v1.RotationInfoListResponse{
		List:  make([]*pbentity.RotationInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}

	// 查询总数
	total, err := dao.RotationInfo.Ctx(ctx).Count()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	// 查询当前页数据
	rotationRecords, err := dao.RotationInfo.Ctx(ctx).
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
	for _, record := range rotationRecords {
		var rotation entity.RotationInfo
		if err := record.Struct(&rotation); err != nil {
			continue
		}

		var pbRotation pbentity.RotationInfo
		if err := gconv.Struct(rotation, &pbRotation); err != nil {
			continue
		}

		// 单独处理时间字段（gconv无法自动转换）
		pbRotation.CreatedAt = utility.SafeConvertTime(rotation.CreatedAt)
		pbRotation.UpdatedAt = utility.SafeConvertTime(rotation.UpdatedAt)
		pbRotation.DeletedAt = utility.SafeConvertTime(rotation.DeletedAt)

		response.List = append(response.List, &pbRotation)
	}
	return &v1.RotationInfoGetListRes{Data: response}, nil
}

// Create 创建
func (c *Controller) Create(ctx context.Context, req *v1.RotationInfoCreateReq) (*v1.RotationInfoCreateRes, error) {
	// 错误类型
	infoError := consts.InfoError(consts.RotationInfo, consts.CreateFail)
	id, err := dao.RotationInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	return &v1.RotationInfoCreateRes{Id: uint32(id)}, nil
}

// Update 更新
func (c *Controller) Update(ctx context.Context, req *v1.RotationInfoUpdateReq) (*v1.RotationInfoUpdateRes, error) {
	// 错误类型
	infoError := consts.InfoError(consts.RotationInfo, consts.UpdateFail)
	_, err := dao.RotationInfo.Ctx(ctx).Where("id", req.Id).Update(req)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	return &v1.RotationInfoUpdateRes{Id: req.Id}, nil
}

// Delete 删除
func (c *Controller) Delete(ctx context.Context, req *v1.RotationInfoDeleteReq) (*v1.RotationInfoDeleteRes, error) {
	// 错误类型
	infoError := consts.InfoError(consts.RotationInfo, consts.DeleteFail)
	// 只需要关注是否出错，不返回被删除的数据
	_, err := dao.RotationInfo.Ctx(ctx).Where("id", req.Id).Delete()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.RotationInfoDeleteRes{}, nil // 返回空结构体
}
