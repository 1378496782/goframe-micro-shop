package collection_info

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "shop-goframe-micro-service-refacotor/app/interaction/api/collection_info/v1"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/logic/collection_info"
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

	// 错误类型
	infoError := consts.InfoError(consts.CollectionInfo, consts.GetListFail)

	// 调用logic层
	list, total, err := collection_info.GetList(ctx, req)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回创建成功响应
	res = &v1.CollectionInfoGetListRes{}
	res.Data = &v1.CollectionInfoListResponse{}
	res.Data.List = list
	res.Data.Total = uint32(total)
	res.Data.Page = req.Page
	res.Data.Size = req.Size
	return res, nil
}

// Create 创建
func (*Controller) Create(ctx context.Context, req *v1.CollectionInfoCreateReq) (res *v1.CollectionInfoCreateRes, err error) {
	// 错误类型
	infoError := consts.InfoError(consts.CollectionInfo, consts.CreateFail)

	// 调用logic层
	id, err := collection_info.Create(ctx, req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回创建成功响应，包含新创建的ID
	res = &v1.CollectionInfoCreateRes{}
	res.Id = uint32(id)
	return res, nil
}

// Delete 删除
func (*Controller) Delete(ctx context.Context, req *v1.CollectionInfoDeleteReq) (res *v1.CollectionInfoDeleteRes, err error) {

	// 系统运行的错误
	infoError := consts.InfoError(consts.CollectionInfo, consts.DeleteFail)

	// 调用logic层
	id, err := collection_info.Delete(ctx, req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回删除成功的空响应
	res = &v1.CollectionInfoDeleteRes{}
	res.Id = uint32(id)
	return res, nil // 返回空结构体
}
