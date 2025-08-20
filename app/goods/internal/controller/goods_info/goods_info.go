package goods_info

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "shop-goframe-micro-service-refacotor/app/goods/api/goods_info/v1"
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
	v1.UnimplementedGoodsInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterGoodsInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.GoodsInfoGetListReq) (res *v1.GoodsInfoGetListRes, err error) {
	// 初始化响应结构
	response := &v1.GoodsInfoListResponse{
		List:  make([]*pbentity.GoodsInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}
	// 错误类型
	infoError := consts.InfoError(consts.GoodsInfo, consts.GetListFail)
	// 查询总数
	total, err := dao.GoodsInfo.Ctx(ctx).Count()
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	// 查询当前页数据
	goodsRecords, err := dao.GoodsInfo.Ctx(ctx).
		Page(int(req.Page), int(req.Size)).
		All()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 数据转换
	// 在循环中替换手动赋值
	for _, record := range goodsRecords {
		var goods entity.GoodsInfo
		if err := record.Struct(&goods); err != nil {
			continue
		}

		var pbGoods pbentity.GoodsInfo
		if err := gconv.Struct(goods, &pbGoods); err != nil {
			continue
		}

		// 单独处理时间字段（gconv无法自动转换）
		pbGoods.CreatedAt = utility.SafeConvertTime(goods.CreatedAt)
		pbGoods.UpdatedAt = utility.SafeConvertTime(goods.UpdatedAt)
		pbGoods.UpdatedAt = utility.SafeConvertTime(goods.UpdatedAt)

		response.List = append(response.List, &pbGoods)
	}

	return &v1.GoodsInfoGetListRes{Data: response}, nil
}

func (*Controller) GetDetail(ctx context.Context, req *v1.GoodsInfoGetDetailReq) (res *v1.GoodsInfoGetDetailRes, err error) {
	// 错误类型
	infoError := consts.InfoError(consts.GoodsInfo, consts.GetDetailFile)
	record, err := dao.GoodsInfo.Ctx(ctx).Where("id", req.Id).One()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	if record.IsEmpty() {
		g.Log().Errorf(ctx, "%v %v", infoError+"查询商品不存在", err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError+"查询商品不存在")
	}

	// 转换为实体结构
	var goods entity.GoodsInfo
	if err := record.Struct(&goods); err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "数据转换失败")
	}

	// 转换为protobuf结构
	var pbGoods pbentity.GoodsInfo
	if err := gconv.Struct(goods, &pbGoods); err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "数据转换失败")
	}

	/// 单独处理时间字段（gconv无法自动转换）
	pbGoods.CreatedAt = utility.SafeConvertTime(goods.CreatedAt)
	pbGoods.UpdatedAt = utility.SafeConvertTime(goods.UpdatedAt)

	return &v1.GoodsInfoGetDetailRes{
		Data: &pbGoods,
	}, nil
}

func (*Controller) Create(ctx context.Context, req *v1.GoodsInfoCreateReq) (res *v1.GoodsInfoCreateRes, err error) {
	// 定义一个实体对象，用于接收转换后的请求数据
	var goodsInfo *entity.GoodsInfo
	// 将请求参数req转换为实体对象goodsInfo
	if err := gconv.Struct(req, &goodsInfo); err != nil {
		return nil, err
	}
	// 错误类型
	infoError := consts.InfoError(consts.GoodsInfo, consts.CreateFail)
	// 向数据库中插入数据并获取自动生成的ID
	result, err := dao.GoodsInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回创建成功响应，包含新创建的ID
	return &v1.GoodsInfoCreateRes{Id: uint32(result)}, nil
}

func (*Controller) Update(ctx context.Context, req *v1.GoodsInfoUpdateReq) (res *v1.GoodsInfoUpdateRes, err error) {
	// 定义一个实体对象，用于接收转换后的请求数据
	var goodsInfo *entity.GoodsInfo
	// 将请求参数req转换为实体对象goodsInfo
	if err := gconv.Struct(req, &goodsInfo); err != nil {
		return nil, err
	}
	infoError := consts.InfoError(consts.GoodsInfo, consts.UpdateFail)
	// 根据ID更新数据库中的信息
	_, err = dao.GoodsInfo.Ctx(ctx).Where("id", req.Id).Update(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回更新成功响应，包含被更新ID
	return &v1.GoodsInfoUpdateRes{Id: req.Id}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.GoodsInfoDeleteReq) (res *v1.GoodsInfoDeleteRes, err error) {
	// 根据ID从数据库中删除对应信息
	_, err = dao.GoodsInfo.Ctx(ctx).Where("id", req.Id).Delete()
	infoError := consts.InfoError(consts.GoodsInfo, consts.DeleteFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回删除成功的空响应
	return &v1.GoodsInfoDeleteRes{}, nil // 返回空结构体
}
