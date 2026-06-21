package job

import (
	"context"
	"sync"
	"time"

	comment_info "shop-goframe-micro-service-refacotor/app/interaction/internal/logic/comment_info"
	"shop-goframe-micro-service-refacotor/app/interaction/utility/interactionRedis"

	"github.com/gogf/gf/v2/frame/g"
)

func StartCommentLikeCountCalibrateJob(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()

		var mu sync.Mutex
		running := false

		run := func() {
			mu.Lock()
			if running {
				mu.Unlock()
				g.Log().Warning(ctx, "评论点赞计数校准任务仍在执行，本轮跳过")
				return
			}
			running = true
			mu.Unlock()

			defer func() {
				mu.Lock()
				running = false
				mu.Unlock()
			}()

			lockValue, acquired, err := interactionRedis.TryAcquireCommentLikeCountCalibrateLock(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "获取评论点赞计数校准分布式锁失败: %v", err)
				return
			}
			if !acquired {
				g.Log().Info(ctx, "其他实例正在执行评论点赞计数校准，本轮跳过")
				return
			}
			defer func() {
				if err := interactionRedis.ReleaseCommentLikeCountCalibrateLock(ctx, lockValue); err != nil {
					g.Log().Warningf(ctx, "释放评论点赞计数校准分布式锁失败: %v", err)
				}
			}()

			result, err := comment_info.CalibrateCommentLikeCount(ctx, 500)
			if err != nil {
				g.Log().Errorf(ctx, "评论点赞计数校准失败: %v", err)
				return
			}
			if result.Fixed > 0 {
				g.Log().Infof(ctx, "评论点赞计数校准完成, scanned:%d, fixed:%d", result.Scanned, result.Fixed)
			}
		}

		run()

		for {
			select {
			case <-ctx.Done():
				g.Log().Info(ctx, "评论点赞计数校准任务停止")
				return
			case <-ticker.C:
				run()
			}
		}
	}()
}
