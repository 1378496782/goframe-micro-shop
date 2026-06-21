package cart_info

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/goods/api/cart_info/v1"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/utility"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type CartGoodsInfo struct {
	// 购物车字段
	Id        uint32      `orm:"id"`
	UserId    uint32      `orm:"user_id"`
	Count     uint32      `orm:"count"`
	CreatedAt *gtime.Time `orm:"created_at"`
	UpdatedAt *gtime.Time `orm:"updated_at"`

	// 商品字段
	GoodsId        uint32      `orm:"goods_id"`
	GoodsName      string      `orm:"goods_name"`
	GoodsPicUrl    string      `orm:"goods_pic_url"`
	GoodsPrice     uint64      `orm:"goods_price"`
	GoodsBrand     string      `orm:"goods_brand"`
	GoodsStock     uint32      `orm:"goods_stock"`
	GoodsSale      uint32      `orm:"goods_sale"`
	GoodsTags      string      `orm:"goods_tags"`
	GoodsCreatedAt *gtime.Time `orm:"goods_created_at"`
	GoodsUpdatedAt *gtime.Time `orm:"goods_updated_at"`
	GoodsSort      uint32      `orm:"goods_sort"`
}

type CartSummary struct {
	TotalPrice uint64 `orm:"total_price"`
	TotalCount uint32 `orm:"total_count"`
}

// GetList 获取购物车列表
func GetList(ctx context.Context, req *v1.CartInfoGetListReq) (*v1.CartInfoListResponse, error) {
	response := &v1.CartInfoListResponse{
		List:       make([]*v1.CartItem, 0),
		Page:       req.Page,
		Size:       req.Size,
		Total:      0,
		TotalPrice: 0,
		TotalCount: 0,
	}

	// 查询总数
	total, err := dao.CartInfo.Ctx(ctx).Where("user_id", req.UserId).Count()
	if err != nil {
		return nil, err
	}
	response.Total = uint32(total)

	// 联表查询购物车和商品信息
	var cartGoodsList []*CartGoodsInfo
	err = dao.CartInfo.Ctx(ctx).
		Page(int(req.Page), int(req.Size)).
		Where("cart_info.user_id", req.UserId).
		LeftJoin("goods_info", "cart_info.goods_id = goods_info.id").
		Fields(`
            cart_info.id,
            cart_info.user_id,
            cart_info.count,
            cart_info.created_at,
            cart_info.updated_at,
            goods_info.id as goods_id,
            goods_info.name as goods_name,
            goods_info.pic_url as goods_pic_url,
            goods_info.price as goods_price,
            goods_info.brand as goods_brand,
            goods_info.stock as goods_stock,
            goods_info.sale as goods_sale,
            goods_info.tags as goods_tags,
            goods_info.created_at as goods_created_at,
            goods_info.updated_at as goods_updated_at,
            goods_info.sort as goods_sort
        `).
		Scan(&cartGoodsList)
	if err != nil {
		return nil, err
	}

	// 数据转换
	for _, item := range cartGoodsList {
		cartItem := &v1.CartItem{
			// 购物车字段
			Id:     item.Id,
			UserId: item.UserId,
			Count:  item.Count,

			// 商品字段
			GoodsId:     item.GoodsId,
			GoodsName:   item.GoodsName,
			GoodsPicUrl: item.GoodsPicUrl,
			GoodsPrice:  item.GoodsPrice,
			GoodsBrand:  item.GoodsBrand,
			GoodsStock:  item.GoodsStock,
			GoodsSale:   item.GoodsSale,
			GoodsTags:   item.GoodsTags,
			GoodsSort:   item.GoodsSort,
		}

		// 设置时间字段
		cartItem.GoodsCreatedAt = utility.SafeConvertTime(item.GoodsCreatedAt)
		cartItem.GoodsUpdatedAt = utility.SafeConvertTime(item.GoodsUpdatedAt)

		response.List = append(response.List, cartItem)
	}

	// 全表查询计算总金额和总数量
	var cartSummary CartSummary
	err = dao.CartInfo.Ctx(ctx).
		Where("user_id", req.UserId).
		LeftJoin("goods_info", "goods_info.id = cart_info.goods_id").
		Fields(`
            COALESCE(sum(goods_info.price * cart_info.count), 0) as total_price,
            COALESCE(sum(cart_info.count), 0) as total_count
        `).
		Scan(&cartSummary)
	if err != nil {
		return nil, err
	}
	response.TotalPrice = cartSummary.TotalPrice
	response.TotalCount = cartSummary.TotalCount

	return response, nil
}

func GetSelectedItems(ctx context.Context, req *v1.CartInfoGetSelectedItemsReq) (*v1.CartInfoGetSelectedItemsRes, error) {
	if req.UserId == 0 || len(req.CartIds) == 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "参数错误")
	}

	// 使用 user_id + cart_ids 查询当前用户的购物车项
	var cartItems []*v1.CartItem
	err := dao.CartInfo.Ctx(ctx).
		Where(dao.CartInfo.Columns().UserId, req.UserId).
		WhereIn(dao.CartInfo.Columns().Id, req.CartIds).
		LeftJoin("goods_info", "goods_info.id = cart_info.goods_id").
		Fields(`
            cart_info.id,
            cart_info.user_id,
            cart_info.count,
            cart_info.created_at,
            cart_info.updated_at,
            goods_info.id as goods_id,
            goods_info.name as goods_name,
            goods_info.pic_url as goods_pic_url,
            goods_info.price as goods_price,
            goods_info.brand as goods_brand,
            goods_info.stock as goods_stock,
            goods_info.sale as goods_sale,
            goods_info.tags as goods_tags,
            goods_info.created_at as goods_created_at,
            goods_info.updated_at as goods_updated_at,
            goods_info.sort as goods_sort
        `).
		Scan(&cartItems)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	return &v1.CartInfoGetSelectedItemsRes{
		Items: cartItems,
	}, nil
}
