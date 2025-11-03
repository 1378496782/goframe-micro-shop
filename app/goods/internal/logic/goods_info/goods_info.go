package goods_info

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"runtime/debug"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/app/goods/internal/model/entity"
	"shop-goframe-micro-service-refacotor/app/goods/utility/goodsRedis"
	"shop-goframe-micro-service-refacotor/utility/rabbitmq"
	"sync"
)

func ReturnStock(ctx context.Context, req *rabbitmq.OrderStockReturnEvent) ([]*rabbitmq.OrderGoodsInfo, error) {
	// 使用带缓冲的通道来收集每个商品的库存返还结果
	resultChan := make(chan *rabbitmq.OrderGoodsInfo, len(req.GoodsInfo))

	var wg sync.WaitGroup
	wg.Add(len(req.GoodsInfo))

	// 并发处理每个商品的库存返还
	for _, stockInfo := range req.GoodsInfo {
		go func(stockInfo *rabbitmq.OrderGoodsInfo) {
			defer wg.Done()

			// 捕获可能的 panic，避免影响其他 goroutine
			defer func() {
				if r := recover(); r != nil {
					resultChan <- &rabbitmq.OrderGoodsInfo{
						GoodsId: stockInfo.GoodsId,
						Count:   stockInfo.Count,
					}
					g.Log().Errorf(ctx, "panic,: %v", r)
				}
			}()

			// 获取当前库存
			var goodsInfo entity.GoodsInfo
			err := dao.GoodsInfo.Ctx(ctx).Where("id=?", stockInfo.GoodsId).Fields("stock").Scan(&goodsInfo)
			if err != nil {
				// 如果查询失败，将错误信息返回
				resultChan <- &rabbitmq.OrderGoodsInfo{
					GoodsId: stockInfo.GoodsId,
					Count:   stockInfo.Count,
				}
				return
			}

			// 更新库存
			newStock := goodsInfo.Stock + stockInfo.Count
			g.Log().Infof(ctx, "商品{%d}新库存:%d", goodsInfo.Stock, newStock)
			_, err = dao.GoodsInfo.Ctx(ctx).Where("id=?", stockInfo.GoodsId).Data(g.Map{"stock": newStock}).Update()
			if err != nil {
				// 如果更新失败，将错误信息返回
				resultChan <- &rabbitmq.OrderGoodsInfo{
					GoodsId: stockInfo.GoodsId,
					Count:   stockInfo.Count,
				}
				return
			}

			// 返回成功结果
			g.Log().Infof(ctx, "库存更新成功: %v", stockInfo.GoodsId)
		}(stockInfo) // 启动并发 goroutine
	}

	// 等待所有 goroutine 完成
	wg.Wait()
	close(resultChan)

	// 收集所有结果
	var resultArr []*rabbitmq.OrderGoodsInfo
	for res := range resultChan {
		resultArr = append(resultArr, res)
	}

	// 返回结果
	return resultArr, nil
}

func ReduceStock(ctx context.Context, req *rabbitmq.OrderCreatedEvent) error {
	var cacheGoods []uint32

	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, goods := range req.GoodsInfo {
			// 查询当前库存
			var goodsInfo entity.GoodsInfo
			if err := dao.GoodsInfo.Ctx(ctx).TX(tx).
				Where("id = ?", goods.GoodsId).
				Fields("stock").
				Scan(&goodsInfo); err != nil {
				return gerror.Wrapf(err, "查询商品{%d}库存失败", goods.GoodsId)
			}

			// 判断库存是否足够
			if goodsInfo.Stock < goods.Count {
				return gerror.Newf("订单{%d}中商品{%d}库存不足(当前:%d, 需要:%d)", req.OrderId, goods.GoodsId, goodsInfo.Stock, goods.Count)
			}

			// 计算剩余库存并直接更新为新值（非递减操作）
			newStock := goodsInfo.Stock - goods.Count
			g.Log().Infof(ctx, "商品{%d}新库存:%d", goodsInfo.Stock, newStock)
			if _, err := dao.GoodsInfo.Ctx(ctx).TX(tx).
				Where("id = ?", goods.GoodsId).
				Data(g.Map{"stock": newStock}).
				Update(); err != nil {
				return gerror.Wrapf(err, "更新商品{%d}库存失败", goods.GoodsId)
			}

			cacheGoods = append(cacheGoods, uint32(goods.GoodsId))
		}

		return nil
	})

	// 异步删除缓存
	go func() {
		defer func() {
			if r := recover(); r != nil {
				g.Log().Errorf(context.Background(), "[Recover] panic: %v\nStack:\n%s", r, debug.Stack())
			}
		}()
		//if len(cacheGoods) == 0 {
		//	return
		//}

		if err := goodsRedis.DeleteKeys(ctx, cacheGoods); err != nil {
			g.Log().Errorf(ctx, "{%d}订单删除缓存失败, 商品key:%v, 错误:%v", req.OrderId, cacheGoods, err)
		}
	}()

	if err != nil {
		return err
	}
	return nil
}
