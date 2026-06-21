package goods_info

import (
	"context"
	"runtime/debug"
	v1 "shop-goframe-micro-service-refacotor/app/goods/api/goods_info/v1"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/app/goods/internal/model/entity"
	"shop-goframe-micro-service-refacotor/app/goods/utility/goodsRedis"
	"shop-goframe-micro-service-refacotor/utility/rabbitmq"
	"slices"
	"sync"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
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

func DeductStock(ctx context.Context, req *v1.DeductStockReq) (res *v1.DeductStockRes, err error) {
	// 1. 基础参数校验：GoodsIds 和 Counts 必须长度一致且非空
	if len(req.GoodsIds) == 0 || len(req.GoodsIds) != len(req.Counts) {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "商品ID与数量不匹配")
	}
	for i := range req.Counts {
		if req.Counts[i] <= 0 {
			return nil, gerror.NewCode(gcode.CodeInvalidParameter, "商品数量必须大于0")
		}
	}

	// 2. 先汇总每个 goodsId 需要扣减的总量（同一个 goodsId 可能出现多次）
	needMap := make(map[uint32]uint32, len(req.GoodsIds))
	for i, goodsId := range req.GoodsIds {
		needMap[goodsId] += req.Counts[i]
	}

	goodsIds := make([]uint32, 0, len(needMap))

	for goodsId := range needMap {
		goodsIds = append(goodsIds, goodsId)
	}
	slices.Sort(goodsIds)

	// 3. 原子扣减：UPDATE ... WHERE stock >= 扣量 AND id = ?，靠 RowsAffected 判断是否成功
	//    用事务保证"要么全扣成功，要么全不扣"
	affectedGoodsIds := make(map[uint32]struct{}, len(needMap))
	err = dao.GoodsInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, goodsId := range goodsIds {
			need := needMap[goodsId]
			result, txErr := tx.Model(dao.GoodsInfo.Table()).
				Where("id", goodsId).
				Where("stock >= ?", need).
				Decrement("stock", need)
			if txErr != nil {
				return txErr
			}
			rows, _ := result.RowsAffected()
			if rows == 0 {
				return gerror.NewCodef(gcode.CodeValidationFailed, "商品 %d 库存不足", goodsId)
			}
			affectedGoodsIds[goodsId] = struct{}{}
		}
		return nil
	})
	if err != nil {
		g.Log().Errorf(ctx, "扣减库存失败: %v", err)
		// 已回滚，直接返回原始错误（保留错误码给调用方）
		return nil, err
	}

	// 4. 库存变更后，主动失效相关 Redis 缓存（避免读到旧 stock）
	for goodsId := range affectedGoodsIds {
		if delErr := goodsRedis.DeleteGoodsDetail(ctx, goodsId); delErr != nil {
			g.Log().Warningf(ctx, "失效商品 %d 缓存失败: %v", goodsId, delErr)
		}
	}

	return &v1.DeductStockRes{}, nil
}

func RestoreStock(ctx context.Context, req *v1.RestoreStockReq) (res *v1.RestoreStockRes, err error) {
	// 1. 基础参数校验：GoodsIds 和 Counts 必须长度一致且非空
	if len(req.GoodsIds) == 0 || len(req.GoodsIds) != len(req.Counts) {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "商品ID与数量不匹配")
	}
	for i := range req.Counts {
		if req.Counts[i] <= 0 {
			return nil, gerror.NewCode(gcode.CodeInvalidParameter, "商品数量必须大于0")
		}
	}

	// 2. 先汇总每个 goodsId 需要返还的总量（同一个 goodsId 可能出现多次）
	needMap := make(map[uint32]uint32, len(req.GoodsIds))
	for i, goodsId := range req.GoodsIds {
		needMap[goodsId] += req.Counts[i]
	}

	goodsIds := make([]uint32, 0, len(needMap))

	for goodsId := range needMap {
		goodsIds = append(goodsIds, goodsId)
	}
	slices.Sort(goodsIds)

	// 3. 原子返还：UPDATE ... WHERE stock >= 扣量 AND id = ?，靠 RowsAffected 判断是否成功
	//    用事务保证"要么全返还成功，要么全不返还"
	affectedGoodsIds := make(map[uint32]struct{}, len(needMap))
	err = dao.GoodsInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, goodsId := range goodsIds {
			need := needMap[goodsId]
			result, txErr := tx.Model(dao.GoodsInfo.Table()).
				Where("id", goodsId).
				Increment("stock", need)
			if txErr != nil {
				return txErr
			}
			rows, _ := result.RowsAffected()
			if rows == 0 {
				return gerror.NewCodef(gcode.CodeValidationFailed, "商品 %d 可能不存在", goodsId)
			}
			affectedGoodsIds[goodsId] = struct{}{}
		}
		return nil
	})
	if err != nil {
		g.Log().Errorf(ctx, "返还库存失败: %v", err)
		// 已回滚，直接返回原始错误（保留错误码给调用方）
		return nil, err
	}

	// 4. 库存变更后，主动失效相关 Redis 缓存（避免读到旧 stock）
	for goodsId := range affectedGoodsIds {
		if delErr := goodsRedis.DeleteGoodsDetail(ctx, goodsId); delErr != nil {
			g.Log().Warningf(ctx, "失效商品 %d 缓存失败: %v", goodsId, delErr)
		}
	}

	return &v1.RestoreStockRes{}, nil
}

func IncreaseSales(ctx context.Context, req *v1.IncreaseSalesReq) (res *v1.IncreaseSalesRes, err error) {
	if len(req.GoodsIds) == 0 || len(req.GoodsIds) != len(req.Counts) {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "商品ID与数量不匹配")
	}
	for i := range req.Counts {
		if req.Counts[i] <= 0 {
			return nil, gerror.NewCode(gcode.CodeInvalidParameter, "商品数量必须大于0")
		}
	}

	needMap := make(map[uint32]uint32, len(req.GoodsIds))
	for i, goodsId := range req.GoodsIds {
		needMap[goodsId] += req.Counts[i]
	}

	goodsIds := make([]uint32, 0, len(needMap))
	for goodsId := range needMap {
		goodsIds = append(goodsIds, goodsId)
	}
	slices.Sort(goodsIds)

	affectedGoodsIds := make(map[uint32]struct{}, len(needMap))
	err = dao.GoodsInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, goodsId := range goodsIds {
			result, txErr := tx.Model(dao.GoodsInfo.Table()).
				Where("id", goodsId).
				Increment("sale", needMap[goodsId])
			if txErr != nil {
				return txErr
			}
			rows, _ := result.RowsAffected()
			if rows == 0 {
				return gerror.NewCodef(gcode.CodeValidationFailed, "商品 %d 可能不存在", goodsId)
			}
			affectedGoodsIds[goodsId] = struct{}{}
		}
		return nil
	})
	if err != nil {
		g.Log().Errorf(ctx, "增加商品销量失败: %v", err)
		return nil, err
	}

	for goodsId := range affectedGoodsIds {
		if delErr := goodsRedis.DeleteGoodsDetail(ctx, goodsId); delErr != nil {
			g.Log().Warningf(ctx, "失效商品 %d 缓存失败: %v", goodsId, delErr)
		}
	}

	return &v1.IncreaseSalesRes{}, nil
}
