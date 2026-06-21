package dao

import (
	"context"
	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/model"
	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/model/do"
	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FlashSaleGoodsDao 秒杀商品数据访问对象
type FlashSaleGoodsDao struct {
	model.BaseDao
}

var (
	flashSaleGoodsDao = &FlashSaleGoodsDao{}
)

// FlashSaleGoods 获取秒杀商品DAO实例
func FlashSaleGoods() *FlashSaleGoodsDao {
	return flashSaleGoodsDao
}

// GetFlashSaleGoodsById 根据ID获取秒杀商品
func (d *FlashSaleGoodsDao) GetFlashSaleGoodsById(ctx context.Context, id uint32) (*entity.FlashSaleGoods, error) {
	var goods entity.FlashSaleGoods
	err := g.DB().Ctx(ctx).Model(d.Table).Where("id = ?", id).Scan(&goods)
	if err != nil {
		return nil, err
	}
	return &goods, nil
}

// GetFlashSaleGoodsByGoodsId 根据商品ID获取秒杀商品
func (d *FlashSaleGoodsDao) GetFlashSaleGoodsByGoodsId(ctx context.Context, goodsId uint32) (*entity.FlashSaleGoods, error) {
	var goods entity.FlashSaleGoods
	err := g.DB().Ctx(ctx).Model(d.Table).Where("goods_id = ?", goodsId).Scan(&goods)
	if err != nil {
		return nil, err
	}
	return &goods, nil
}

// GetFlashSaleGoodsList 获取秒杀商品列表
func (d *FlashSaleGoodsDao) GetFlashSaleGoodsList(ctx context.Context, activityId uint32) ([]*entity.FlashSaleGoods, error) {
	var goodsList []*entity.FlashSaleGoods
	err := g.DB().Ctx(ctx).Model(d.Table).Where("activity_id = ?", activityId).Scan(&goodsList)
	if err != nil {
		return nil, err
	}
	return goodsList, nil
}

// UpdateFlashSaleStock 更新秒杀商品库存
func (d *FlashSaleGoodsDao) UpdateFlashSaleStock(ctx context.Context, goodsId uint32, stock int) error {
	_, err := g.DB().Ctx(ctx).Model(d.Table).
		Where("goods_id = ?", goodsId).
		Data(g.Map{
			"available_stock": gdb.Raw("available_stock - ?"),
			"updated_at":      gtime.Timestamp(),
		}).Update()
	return err
}

// FlashSaleOrderDao 秒杀订单数据访问对象
type FlashSaleOrderDao struct {
	model.BaseDao
}

var (
	flashSaleOrderDao = &FlashSaleOrderDao{}
)

// FlashSaleOrder 获取秒杀订单DAO实例
func FlashSaleOrder() *FlashSaleOrderDao {
	return flashSaleOrderDao
}

// CreateFlashSaleOrder 创建秒杀订单
func (d *FlashSaleOrderDao) CreateFlashSaleOrder(ctx context.Context, order *do.FlashSaleOrder) error {
	_, err := g.DB().Ctx(ctx).Model(d.Table).Data(order).Insert()
	return err
}

// GetFlashSaleOrderByOrderNo 根据订单号获取秒杀订单
func (d *FlashSaleOrderDao) GetFlashSaleOrderByOrderNo(ctx context.Context, orderNo string) (*entity.FlashSaleOrder, error) {
	var order entity.FlashSaleOrder
	err := g.DB().Ctx(ctx).Model(d.Table).Where("order_no = ?", orderNo).Scan(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
