package cart_info

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/goods/api/cart_info/v1"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/utility"

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
		response.TotalPrice += item.GoodsPrice * uint64(item.Count)
		response.TotalCount += item.Count
	}

	return response, nil
}
