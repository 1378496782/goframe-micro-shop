package collection_info

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "shop-goframe-micro-service-refacotor/app/interaction/api/collection_info/v1"
	"shop-goframe-micro-service-refacotor/app/interaction/api/pbentity"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/dao"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/model/entity"
	"shop-goframe-micro-service-refacotor/utility"
	"shop-goframe-micro-service-refacotor/utility/consts"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedCollectionInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterCollectionInfoServer(s.Server, &Controller{})
}

// GetList 列表
func (*Controller) GetList(ctx context.Context, req *v1.CollectionInfoGetListReq) (res *v1.CollectionInfoGetListRes, err error) {
	// 初始化响应结构
	response := &v1.CollectionInfoListResponse{
		List:  make([]*pbentity.CollectionInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}
	// 错误类型
	infoError := consts.InfoError(consts.CollectionInfo, consts.GetListFail)
	// 查询总数
	total, err := dao.CollectionInfo.Ctx(ctx).Count()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	// 查询当前页数据
	collectionRecords, err := dao.CollectionInfo.Ctx(ctx).
		Page(int(req.Page), int(req.Size)).
		All()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 数据转换
	// 在循环中替换手动赋值
	for _, record := range collectionRecords {
		var collection entity.CollectionInfo
		if err := record.Struct(&collection); err != nil {
			continue
		}

		var pbCollection pbentity.CollectionInfo
		if err := gconv.Struct(collection, &pbCollection); err != nil {
			continue
		}

		// 单独处理时间字段（gconv无法自动转换）
		pbCollection.CreatedAt = utility.SafeConvertTime(collection.CreatedAt)
		pbCollection.UpdatedAt = utility.SafeConvertTime(collection.UpdatedAt)

		response.List = append(response.List, &pbCollection)
	}

	return &v1.CollectionInfoGetListRes{Data: response}, nil
}

// Create 创建
func (*Controller) Create(ctx context.Context, req *v1.CollectionInfoCreateReq) (res *v1.CollectionInfoCreateRes, err error) {
	// 错误类型
	infoError := consts.InfoError(consts.CollectionInfo, consts.CreateFail)
	// 向数据库中插入数据并获取自动生成的ID
	result, err := dao.CollectionInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回创建成功响应，包含新创建的ID
	return &v1.CollectionInfoCreateRes{Id: uint32(result)}, nil
}

// Delete 删除
func (*Controller) Delete(ctx context.Context, req *v1.CollectionInfoDeleteReq) (res *v1.CollectionInfoDeleteRes, err error) {
	// 根据ID从数据库中删除对应信息
	_, err = dao.CollectionInfo.Ctx(ctx).Where("id", req.Id).Delete()
	infoError := consts.InfoError(consts.CollectionInfo, consts.DeleteFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回删除成功的空响应
	return &v1.CollectionInfoDeleteRes{}, nil // 返回空结构体
}
