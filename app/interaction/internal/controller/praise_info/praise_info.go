package praise_info

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"shop-goframe-micro-service-refacotor/app/interaction/api/pbentity"
	v1 "shop-goframe-micro-service-refacotor/app/interaction/api/praise_info/v1"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/consts"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/dao"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/model/entity"
	"shop-goframe-micro-service-refacotor/utility"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedPraiseInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterPraiseInfoServer(s.Server, &Controller{})
}

// GetList 列表
func (*Controller) GetList(ctx context.Context, req *v1.PraiseInfoGetListReq) (res *v1.PraiseInfoGetListRes, err error) {
	// 初始化响应结构
	response := &v1.PraiseInfoListResponse{
		List:  make([]*pbentity.PraiseInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}
	// 错误类型
	infoError := consts.InfoError(consts.PraiseInfo, consts.GetListFail)
	// 查询总数
	total, err := dao.PraiseInfo.Ctx(ctx).Count()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	// 查询当前页数据
	praiseRecords, err := dao.PraiseInfo.Ctx(ctx).
		Page(int(req.Page), int(req.Size)).
		All()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 数据转换
	// 在循环中替换手动赋值
	for _, record := range praiseRecords {
		var praise entity.PraiseInfo
		if err := record.Struct(&praise); err != nil {
			continue
		}

		var pbPraise pbentity.PraiseInfo
		if err := gconv.Struct(praise, &pbPraise); err != nil {
			continue
		}

		// 单独处理时间字段（gconv无法自动转换）
		pbPraise.CreatedAt = utility.SafeConvertTime(praise.CreatedAt)
		pbPraise.UpdatedAt = utility.SafeConvertTime(praise.UpdatedAt)

		response.List = append(response.List, &pbPraise)
	}

	return &v1.PraiseInfoGetListRes{Data: response}, nil
}

// Create 创建
func (*Controller) Create(ctx context.Context, req *v1.PraiseInfoCreateReq) (res *v1.PraiseInfoCreateRes, err error) {
	// 定义一个实体对象，用于接收转换后的请求数据
	var praiseInfo *entity.PraiseInfo
	// 将请求参数req转换为实体对象praiseInfo
	if err := gconv.Struct(req, &praiseInfo); err != nil {
		return nil, err
	}
	// 错误类型
	infoError := consts.InfoError(consts.PraiseInfo, consts.CreateFail)
	// 向数据库中插入数据并获取自动生成的ID
	result, err := dao.PraiseInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回创建成功响应，包含新创建的ID
	return &v1.PraiseInfoCreateRes{Id: uint32(result)}, nil
}

// Delete 删除
func (*Controller) Delete(ctx context.Context, req *v1.PraiseInfoDeleteReq) (res *v1.PraiseInfoDeleteRes, err error) {
	// 根据ID从数据库中删除对应信息
	_, err = dao.PraiseInfo.Ctx(ctx).Where("id", req.Id).Delete()
	infoError := consts.InfoError(consts.PraiseInfo, consts.DeleteFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回删除成功的空响应
	return &v1.PraiseInfoDeleteRes{}, nil // 返回空结构体
}
