package job

import (
	"context"
	"sync"
	"time"

	order_info "shop-goframe-micro-service-refacotor/app/order/internal/logic/order_info"

	"github.com/gogf/gf/v2/frame/g"
)

// StartOutboxRelayJob 启动 Outbox 中继后台任务。
//
// 它和订单销量补偿任务一样，随 order-service 启动自动运行：
// 按固定间隔轮询 order_outbox_message 里待发送 / 到期可重试的消息，
// 投递到 RabbitMQ，成功标记 sent，失败按指数退避重试。
func StartOutboxRelayJob(ctx context.Context) {
	go func() {
		// 每 5 秒扫一轮。Outbox 消息要尽快发出去，间隔比销量补偿任务短。
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		// running 防止同一进程内上一轮还没跑完就重入。
		var mu sync.Mutex
		running := false

		run := func() {
			mu.Lock()
			if running {
				mu.Unlock()
				g.Log().Warning(ctx, "Outbox 中继任务仍在执行，本轮跳过")
				return
			}
			running = true
			mu.Unlock()

			defer func() {
				mu.Lock()
				running = false
				mu.Unlock()
			}()

			// 先恢复卡在 sending 的僵尸消息（进程崩溃残留），再投递。
			// 超时设 5 分钟，远大于一轮正常投递耗时，避免误伤处理中的消息。
			resetCount, err := order_info.ResetStuckOutboxSending(ctx, 5, 100)
			if err != nil {
				g.Log().Errorf(ctx, "Outbox 恢复 sending 僵尸消息失败: %v", err)
				// 恢复失败不影响本轮正常投递，继续往下走
			} else if resetCount > 0 {
				g.Log().Infof(ctx, "Outbox 恢复 sending 僵尸消息, reset:%d", resetCount)
			}

			sent, failed, err := order_info.RelayOutboxOnce(ctx, 100)
			if err != nil {
				g.Log().Errorf(ctx, "Outbox 中继任务执行失败: %v", err)
				return
			}
			// 没有消息时静默，避免日志刷屏。
			if sent > 0 || failed > 0 {
				g.Log().Infof(ctx, "Outbox 中继任务执行完成, sent:%d, failed:%d", sent, failed)
			}
		}

		// 启动后先立即跑一轮。
		run()

		for {
			select {
			case <-ctx.Done():
				g.Log().Info(ctx, "Outbox 中继任务停止")
				return
			case <-ticker.C:
				run()
			}
		}
	}()
}
