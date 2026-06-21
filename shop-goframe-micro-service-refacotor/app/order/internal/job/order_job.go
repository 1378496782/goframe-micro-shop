package job

import (
	"context"
	"sync"
	"time"

	order_info "shop-goframe-micro-service-refacotor/app/order/internal/logic/order_info"

	"github.com/gogf/gf/v2/frame/g"
)

// StartOrderCompensateJob 启动订单销量补偿后台任务。
//
// 这个任务和 HTTP 接口不一样：
// - HTTP 接口需要用户/API Fox 主动请求才会执行
// - 后台任务会随着 order-service 启动自动运行，并按固定间隔扫描异常数据
func StartOrderCompensateJob(ctx context.Context) {
	// 用 goroutine 启动后台任务，避免阻塞 order-service 后面的 gRPC 服务启动。
	go func() {
		// ticker 每隔 30 秒向 ticker.C 发送一次信号，用来触发一轮补偿。
		ticker := time.NewTicker(300 * time.Second)
		defer ticker.Stop()

		// running 用来防止同一个服务实例内重复执行：
		// 如果上一轮补偿还没跑完，下一次 ticker 到了就直接跳过。
		var mu sync.Mutex
		running := false

		// run 表示执行一轮完整的补偿任务。
		run := func() {
			// 进入任务前先抢本进程内的执行权。
			mu.Lock()
			if running {
				mu.Unlock()
				g.Log().Warning(ctx, "订单销量补偿任务仍在执行，本轮跳过")
				return
			}
			running = true
			mu.Unlock()

			// 无论本轮成功还是失败，函数退出前都要释放 running。
			defer func() {
				mu.Lock()
				running = false
				mu.Unlock()
			}()

			// 第一步：把卡在 sales_status=3（同步中）太久的订单恢复成同步失败。
			// 这样下一步 CompensateFailedSales 才能重新补偿它们。
			resetCount, err := order_info.ResetStuckSalesSyncing(ctx, 10, 20)
			if err != nil {
				g.Log().Errorf(ctx, "恢复卡住的销量同步任务失败: %v", err)
				return
			}

			// 第二步：扫描 sales_status=2（同步失败）的订单，重新增加商品销量。
			compensateCount, err := order_info.CompensateFailedSales(ctx, 20)
			if err != nil {
				g.Log().Errorf(ctx, "订单销量补偿任务执行失败: %v", err)
				return
			}

			g.Log().Infof(ctx, "订单销量补偿任务执行完成, reset_count:%d, compensate_count:%d", resetCount, compensateCount)
		}

		// 服务刚启动时先立即跑一轮，不用等第一个 30 秒。
		run()

		for {
			select {
			// ctx.Done() 收到信号时，说明服务准备退出，后台任务也应该停止。
			case <-ctx.Done():
				g.Log().Info(ctx, "订单销量补偿任务停止")
				return
			// ticker.C 到点后执行下一轮补偿。
			case <-ticker.C:
				run()
			}
		}
	}()
}
