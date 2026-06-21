package category_info

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "shop-goframe-micro-service-refacotor/app/goods/api/category_info/v1"
	"shop-goframe-micro-service-refacotor/app/goods/api/pbentity"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/app/goods/internal/model/entity"
	"shop-goframe-micro-service-refacotor/app/goods/utility/goodsRedis"
	"shop-goframe-micro-service-refacotor/utility"
	"shop-goframe-micro-service-refacotor/utility/consts"
	"time"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedCategoryInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterCategoryInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.CategoryInfoGetListReq) (res *v1.CategoryInfoGetListRes, err error) {
	response := &v1.CategoryInfoListResponse{
		List:  make([]*pbentity.CategoryInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}

	// 查询总数 - 同样需要判断sort条件
	countQuery := dao.CategoryInfo.Ctx(ctx)
	if req.Sort != 0 { // 假设0表示未设置sort
		countQuery = countQuery.Where("sort", req.Sort)
	}
	total, err := countQuery.Count()
	if err != nil {
		return &v1.CategoryInfoGetListRes{Data: response}, nil
	}
	response.Total = uint32(total)

	// 构建查询
	query := dao.CategoryInfo.Ctx(ctx).Page(int(req.Page), int(req.Size))

	// 按需添加sort条件
	if req.Sort != 0 {
		query = query.Where("sort", req.Sort)
	}

	// 执行查询
	categoryRecords, err := query.All()
	if err != nil {
		return &v1.CategoryInfoGetListRes{Data: response}, nil
	}

	// 数据转换
	for _, record := range categoryRecords {
		var category entity.CategoryInfo
		if err := record.Struct(&category); err != nil {
			continue
		}

		var pbCategory pbentity.CategoryInfo
		if err := gconv.Struct(category, &pbCategory); err != nil {
			continue
		}

		pbCategory.CreatedAt = utility.SafeConvertTime(category.CreatedAt)
		pbCategory.UpdatedAt = utility.SafeConvertTime(category.UpdatedAt)
		pbCategory.DeletedAt = utility.SafeConvertTime(category.DeletedAt)

		response.List = append(response.List, &pbCategory)
	}

	return &v1.CategoryInfoGetListRes{Data: response}, nil
}

func (*Controller) GetAll(ctx context.Context, req *v1.CategoryInfoGetAllReq) (res *v1.CategoryInfoGetAllRes, err error) {
	// 先尝试从Redis获取
	cachedData, err := goodsRedis.GetCategoryAll(ctx)
	if err != nil {
		g.Log().Infof(ctx, "Redis查询失败: %v", err)
		// 继续查询数据库，不直接返回错误
	} else if !cachedData.IsNil() && !cachedData.IsEmpty() {
		// 缓存命中，反序列化数据
		var cachedRes v1.CategoryInfoGetAllRes
		if err := cachedData.Struct(&cachedRes); err != nil {
			g.Log().Errorf(ctx, "缓存数据反序列化失败: %v", err)
			// 继续查询数据库
		} else {
			g.Log().Info(ctx, "category all data缓存命中")
			return &cachedRes, nil
		}
	}

	// 初始化响应结构
	response := &v1.CategoryInfoGetAllRes{
		List:  make([]*pbentity.CategoryInfo, 0),
		Total: 0,
	}

	// 查询总数
	total, err := dao.CategoryInfo.Ctx(ctx).Count()
	if err != nil {
		return response, err
	}
	response.Total = uint32(total)

	// 查询所有数据
	categoryRecords, err := dao.CategoryInfo.Ctx(ctx).All()
	if err != nil {
		return response, err
	}

	// 预分配切片容量
	if total > 0 {
		response.List = make([]*pbentity.CategoryInfo, 0, total)
	}

	// 数据转换
	for _, record := range categoryRecords {
		var category entity.CategoryInfo
		if err := record.Struct(&category); err != nil {
			continue
		}

		var pbCategory pbentity.CategoryInfo
		if err := gconv.Struct(category, &pbCategory); err != nil {
			continue
		}

		// 处理时间字段
		pbCategory.CreatedAt = utility.SafeConvertTime(category.CreatedAt)
		pbCategory.UpdatedAt = utility.SafeConvertTime(category.UpdatedAt)
		pbCategory.DeletedAt = utility.SafeConvertTime(category.DeletedAt)

		response.List = append(response.List, &pbCategory)
	}
	// 为缓存设置操作设置100毫秒的超时时间 避免Redis操作耗时过长影响主业务流程
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()
	// 设置缓存（使用一周的缓存时间）
	if err := goodsRedis.SetCategoryAll(ctxWithTimeout, response); err != nil {
		g.Log().Warningf(ctx, "设置分类全量数据缓存失败: %v", err)
		// 不返回错误，因为主业务已成功
	}

	return response, nil
}

func (*Controller) Create(ctx context.Context, req *v1.CategoryInfoCreateReq) (res *v1.CategoryInfoCreateRes, err error) {
	// 错误类型
	infoError := consts.InfoError(consts.CategoryInfo, consts.CreateFail)
	// 向数据库中插入数据并获取自动生成的ID
	result, err := dao.CategoryInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回创建成功响应，包含新创建的ID
	return &v1.CategoryInfoCreateRes{Id: uint32(result)}, nil
}

func (*Controller) Update(ctx context.Context, req *v1.CategoryInfoUpdateReq) (res *v1.CategoryInfoUpdateRes, err error) {
	infoError := consts.InfoError(consts.CategoryInfo, consts.UpdateFail)
	// 根据ID更新数据库中的信息
	_, err = dao.CategoryInfo.Ctx(ctx).Where("id", req.Id).Update(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	// 数据库更新成功后，删除缓存
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	if err := goodsRedis.DeleteCategoryAll(ctxWithTimeout); err != nil {
		g.Log().Warningf(ctx, "删除分类全量数据缓存失败: %v", err)
		// 不返回错误，因为主业务已成功
	}

	// 返回更新成功响应，包含被更新ID
	return &v1.CategoryInfoUpdateRes{Id: req.Id}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.CategoryInfoDeleteReq) (res *v1.CategoryInfoDeleteRes, err error) {
	// 根据ID从数据库中删除对应信息
	_, err = dao.CategoryInfo.Ctx(ctx).Where("id", req.Id).Delete()
	infoError := consts.InfoError(consts.CategoryInfo, consts.DeleteFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回删除成功的空响应
	return &v1.CategoryInfoDeleteRes{}, nil // 返回空结构体
}
