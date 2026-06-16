package goods_info

import (
	"context"
	"fmt"
	v1 "shop-goframe-micro-service-refacotor/app/goods/api/goods_info/v1"
	"shop-goframe-micro-service-refacotor/app/goods/api/pbentity"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/app/goods/internal/model/entity"
	"shop-goframe-micro-service-refacotor/app/goods/utility/goodsRedis"
	"shop-goframe-micro-service-refacotor/utility"
	"shop-goframe-micro-service-refacotor/utility/consts"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"shop-goframe-micro-service-refacotor/app/goods/internal/logic/goods_info"
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

	// 构建查询条件
	query := dao.GoodsInfo.Ctx(ctx)

	// 根据IsHot参数添加筛选条件
	if req.IsHot == 1 {
		query = query.Where("sort > ?", 0)
	}

	// 根据关键字添加筛选条件
	keyword := strings.TrimSpace(req.Keyword)
	if keyword != "" {
		query = query.Where("name like ?", "%"+keyword+"%")
	}

	// 根据分类ID添加筛选条件
	if req.CategoryId > 0 {
		query = query.Where("level1_category_id = ? or level2_category_id = ? or level3_category_id = ?", req.CategoryId, req.CategoryId, req.CategoryId)
	}

	// 根据价格范围添加筛选条件
	if req.PriceMin > 0 && req.PriceMax > 0 {
		query = query.Where("price between ? and ?", req.PriceMin, req.PriceMax)
	} else if req.PriceMin > 0 {
		query = query.Where("price >= ?", req.PriceMin)
	} else if req.PriceMax > 0 {
		query = query.Where("price <= ?", req.PriceMax)
	}

	// 根据排序类型添加排序条件
	switch req.SortType {
	case v1.SortType_PRICE_ASC:
		query = query.OrderAsc(dao.GoodsInfo.Columns().Price)
	case v1.SortType_PRICE_DESC:
		query = query.OrderDesc(dao.GoodsInfo.Columns().Price)
	case v1.SortType_SALE_DESC:
		query = query.OrderDesc(dao.GoodsInfo.Columns().Sale)
	default:
		query = query.OrderDesc(dao.GoodsInfo.Columns().Sort)
	}

	// 根据是否只看有库存添加筛选条件
	if req.OnlyInStock == 1 {
		query = query.Where("stock > ?", 0)
	}

	// 查询总数
	total, err := query.Count()
	fmt.Println("total,", total)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	// 查询当前页数据
	goodsRecords, err := query.
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

		response.List = append(response.List, &pbGoods)
	}

	return &v1.GoodsInfoGetListRes{Data: response}, nil
}

func (*Controller) GetDetail(ctx context.Context, req *v1.GoodsInfoGetDetailReq) (res *v1.GoodsInfoGetDetailRes, err error) {
	// 先尝试从Redis获取
	detail, err := goodsRedis.GetGoodsDetail(ctx, req.Id)
	if err != nil {
		g.Log().Infof(ctx, "商品详情缓存读取失败，降级查询MySQL: goods_id=%d, err=%v", req.Id, err)
		// 继续查询数据库，不直接返回错误
	} else if !detail.IsNil() {
		// 检查是否为空缓存标记
		if detail.String() == goodsRedis.EmptyValue {
			g.Log().Infof(ctx, "商品详情命中空缓存: goods_id=%d", req.Id)
			return nil, gerror.NewCode(gcode.CodeNotFound, "商品不存在")
		}
		// 缓存命中，反序列化数据
		var cachedRes v1.GoodsInfoGetDetailRes
		if err := detail.Struct(&cachedRes); err != nil {
			g.Log().Errorf(ctx, "商品详情缓存反序列化失败，降级查询MySQL: goods_id=%d, err=%v", req.Id, err)
			// 继续查询数据库
		} else {
			g.Log().Infof(ctx, "商品详情命中缓存: goods_id=%d", req.Id)
			return &cachedRes, nil
		}
	}
	// 缓存未命中，查询数据库
	g.Log().Infof(ctx, "商品详情缓存未命中，查询MySQL: goods_id=%d", req.Id)
	infoError := consts.InfoError(consts.GoodsInfo, consts.GetDetailFail)
	record, err := dao.GoodsInfo.Ctx(ctx).Where("id", req.Id).One()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	if record.IsEmpty() {
		g.Log().Infof(ctx, "MySQL未查询到商品，写入空缓存: goods_id=%d", req.Id)
		// 设置空缓存防止缓存穿透
		_ = goodsRedis.SetEmptyGoodsDetail(ctx, req.Id)
		return nil, gerror.NewCode(gcode.CodeNotFound, "商品不存在")
	}
	g.Log().Infof(ctx, "MySQL查询到商品详情: goods_id=%d", req.Id)

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

	// 组装响应
	res = &v1.GoodsInfoGetDetailRes{
		Data: &pbGoods,
	}
	// 同步设置缓存（使用较短的超时时间避免阻塞）
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	if err := goodsRedis.SetGoodsDetail(ctxWithTimeout, pbGoods.Id, res); err != nil {
		g.Log().Warningf(ctx, "设置商品详情缓存失败: goods_id=%d, err=%v", pbGoods.Id, err)
		// 不返回错误，因为主业务已成功
	} else {
		g.Log().Infof(ctx, "商品详情已写入缓存: goods_id=%d", pbGoods.Id)
	}
	return res, nil
}

func (*Controller) Create(ctx context.Context, req *v1.GoodsInfoCreateReq) (res *v1.GoodsInfoCreateRes, err error) {
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
	infoError := consts.InfoError(consts.GoodsInfo, consts.UpdateFail)
	// 根据ID更新数据库中的信息
	_, err = dao.GoodsInfo.Ctx(ctx).Where("id", req.Id).Update(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	// 数据库更新成功后，删除缓存
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	if err := goodsRedis.DeleteGoodsDetail(ctxWithTimeout, req.Id); err != nil {
		g.Log().Warningf(ctx, "删除商品详情数据缓存失败: %v", err)
		// 不返回错误，因为主业务已成功
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

func (*Controller) GetGoodsStock(ctx context.Context, req *v1.GetGoodsStockReq) (res *v1.GetGoodsStockRes, err error) {
	goodsToStack := make(map[uint32]int32)
	var goodsId []uint32
	for _, productId := range req.GoodsIds {
		detail, err := goodsRedis.GetGoodsDetail(ctx, productId)
		if err != nil {
			g.Log().Infof(ctx, "Redis查询失败: %v", err)
			goodsId = append(goodsId, productId)
			continue
		}
		if detail.IsNil() {
			g.Log().Infof(ctx, "缓存未命中: %v", err)
			goodsId = append(goodsId, productId)
			continue
		}
		// 检查是否为空缓存标记
		if detail.String() == "__EMPTY__" {
			g.Log().Info(ctx, "空缓存命中，防止缓存穿透")
			return nil, gerror.New("商品不存在")
		}
		// 缓存命中，反序列化数据
		var cachedRes v1.GoodsInfoGetDetailRes
		if err := detail.Struct(&cachedRes); err != nil {
			g.Log().Errorf(ctx, "缓存数据反序列化失败: %v", err)
			goodsId = append(goodsId, productId)
		}
		if cachedRes.Data.Stock <= 0 {
			g.Log().Infof(ctx, "商品 %d 库存不足或为0，准备查数据库确认", productId)
			goodsId = append(goodsId, productId)
			continue
		}
		g.Log().Info(ctx, "goods detail缓存命中")
		goodsToStack[productId] = cachedRes.Data.Stock
	}
	if len(goodsId) > 0 {
		var goodsInfo []*entity.GoodsInfo
		err = dao.GoodsInfo.Ctx(ctx).Fields("id", "stock").WhereIn("id", goodsId).Scan(&goodsInfo)
		infoError := consts.InfoError(consts.GoodsInfo, consts.GetStockFail)
		if err != nil {
			g.Log().Errorf(ctx, "%v %v", infoError, err)
			return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
		}
		for _, re := range goodsInfo {
			goodsToStack[uint32(re.Id)] = int32(re.Stock)
		}
	}
	return &v1.GetGoodsStockRes{GoodsStock: goodsToStack}, nil
}

func (c *Controller) DeductStock(ctx context.Context, req *v1.DeductStockReq) (res *v1.DeductStockRes, err error) {
	return goods_info.DeductStock(ctx, req)
}

func (*Controller) RestoreStock(ctx context.Context, req *v1.RestoreStockReq) (res *v1.RestoreStockRes, err error) {
	return goods_info.RestoreStock(ctx, req)
}

func (*Controller) IncreaseSales(ctx context.Context, req *v1.IncreaseSalesReq) (res *v1.IncreaseSalesRes, err error) {
	return goods_info.IncreaseSales(ctx, req)
}
