package goods_images

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "shop-goframe-micro-service-refacotor/app/goods/api/goods_images/v1"
	"shop-goframe-micro-service-refacotor/app/goods/api/pbentity"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/app/goods/internal/model/entity"
	"shop-goframe-micro-service-refacotor/utility/consts"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedGoodsImagesServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterGoodsImagesServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.GoodsImagesGetListReq) (res *v1.GoodsImagesGetListRes, err error) {
	// 初始化响应结构
	response := &v1.GoodsImagesListResponse{
		List:  make([]*pbentity.GoodsImages, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}
	// 错误类型
	infoError := consts.InfoError(consts.GoodsImages, consts.GetListFail)
	// 查询总数
	total, err := dao.GoodsImages.Ctx(ctx).Count()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	// 查询当前页数据
	goodsRecords, err := dao.GoodsImages.Ctx(ctx).
		Page(int(req.Page), int(req.Size)).
		All()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 数据转换
	// 在循环中替换手动赋值
	for _, record := range goodsRecords {
		var goods entity.GoodsImages
		if err := record.Struct(&goods); err != nil {
			continue
		}

		var pbGoods pbentity.GoodsImages
		if err := gconv.Struct(goods, &pbGoods); err != nil {
			continue
		}

		response.List = append(response.List, &pbGoods)
	}

	return &v1.GoodsImagesGetListRes{Data: response}, nil
}

func (*Controller) Create(ctx context.Context, req *v1.GoodsImagesCreateReq) (res *v1.GoodsImagesCreateRes, err error) {
	// 错误类型
	infoError := consts.InfoError(consts.GoodsImages, consts.CreateFail)
	// 向数据库中插入数据并获取自动生成的ID
	result, err := dao.GoodsImages.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回创建成功响应，包含新创建的ID
	return &v1.GoodsImagesCreateRes{Id: uint32(result)}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.GoodsImagesDeleteReq) (res *v1.GoodsImagesDeleteRes, err error) {
	// 根据ID从数据库中删除对应信息
	_, err = dao.GoodsImages.Ctx(ctx).Where("id", req.Id).Delete()
	infoError := consts.InfoError(consts.GoodsImages, consts.DeleteFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回删除成功的空响应
	return &v1.GoodsImagesDeleteRes{}, nil // 返回空结构体
}
