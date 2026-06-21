package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"shop-goframe-micro-service-refacotor/app/order/internal/consts"
	"shop-goframe-micro-service-refacotor/app/order/internal/dao"
	"shop-goframe-micro-service-refacotor/app/order/internal/model/entity"
	grabbitmq "shop-goframe-micro-service-refacotor/utility/rabbitmq"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Outbox 中继默认参数
const (
	// outboxMaxRetry 超过该重试次数仍失败的消息视为死信，不再被中继捞取
	outboxMaxRetry = 5
	// outboxBackoffBase 首次失败后的退避基数
	outboxBackoffBase = 10 * time.Second
	// outboxBackoffMax 退避上限
	outboxBackoffMax = 10 * time.Minute
)

// ScanAndPublishOutbox 最小版 Outbox 扫描发布：
// 取出 limit 条待发送（pending）消息，逐条投递到 RabbitMQ，发送成功后标记为 sent。
// 返回成功发送的条数。失败的消息保持 pending，留待下次扫描重试。
//
// 不做重试退避，适合手动触发 / 冒烟验证。带退避的版本见 RelayOutboxOnce。
func ScanAndPublishOutbox(ctx context.Context, limit int) (sent int, err error) {
	if limit <= 0 {
		limit = 100
	}

	cols := dao.OrderOutboxMessage.Columns()

	var messages []*entity.OrderOutboxMessage
	err = dao.OrderOutboxMessage.Ctx(ctx).
		Where(cols.Status, int(consts.OutboxStatusPending)).
		OrderAsc(cols.Id).
		Limit(limit).
		Scan(&messages)
	if err != nil {
		return 0, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	if len(messages) == 0 {
		return 0, nil
	}

	rb, err := grabbitmq.NewRabbitMQ(ctx)
	if err != nil {
		return 0, err
	}
	defer rb.Close()

	for _, msg := range messages {
		if err = publishOutboxMessage(rb, msg); err != nil {
			g.Log().Errorf(ctx, "Outbox 发送失败, id=%d, eventId=%s, err=%v", msg.Id, msg.EventId, err)
			continue
		}
		if err = markOutboxSent(ctx, msg.Id); err != nil {
			g.Log().Errorf(ctx, "Outbox 标记 sent 失败, id=%d, err=%v", msg.Id, err)
			continue
		}
		sent++
		g.Log().Infof(ctx, "Outbox 发送成功, id=%d, eventId=%s", msg.Id, msg.EventId)
	}
	return sent, nil
}

// RelayOutboxOnce 执行一轮 Outbox 中继：
//   - 捞取 status=pending，或 status=failed 且已到重试时间、且重试未超限的消息；
//   - 用 CAS 把消息抢占为 sending，避免多实例 / 多协程重复发送；
//   - 投递成功标记 sent；失败则累加 retry_count、按指数退避写 next_retry_at，并记录 last_error；
//   - 重试次数达到 outboxMaxRetry 后不再被捞取（死信，需人工介入）。
//
// 返回本轮成功与失败的条数。
func RelayOutboxOnce(ctx context.Context, limit int) (sent, failed int, err error) {
	if limit <= 0 {
		limit = 100
	}

	cols := dao.OrderOutboxMessage.Columns()
	now := gtime.Now()

	// 1. 捞取可投递的消息：未超重试上限，且（pending 或 到期的 failed）
	var messages []*entity.OrderOutboxMessage
	err = dao.OrderOutboxMessage.Ctx(ctx).
		Where(cols.RetryCount+" < ?", outboxMaxRetry).
		Where("status = ? OR (status = ? AND next_retry_at <= ?)",
			int(consts.OutboxStatusPending), int(consts.OutboxStatusFailed), now).
		OrderAsc(cols.Id).
		Limit(limit).
		Scan(&messages)
	if err != nil {
		return 0, 0, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	if len(messages) == 0 {
		return 0, 0, nil
	}

	rb, err := grabbitmq.NewRabbitMQ(ctx)
	if err != nil {
		return 0, 0, err
	}
	defer rb.Close()

	for _, msg := range messages {
		// 2. CAS 抢占：只有状态仍是捞取时的状态才抢占成功，避免重复发送
		claimed, cErr := claimOutboxMessage(ctx, msg)
		if cErr != nil {
			g.Log().Errorf(ctx, "Outbox 抢占失败, id=%d, err=%v", msg.Id, cErr)
			continue
		}
		if !claimed {
			// 已被其他实例/协程抢走，跳过
			continue
		}

		// 3. 投递
		if pErr := publishOutboxMessage(rb, msg); pErr != nil {
			failed++
			markOutboxFailed(ctx, msg, pErr)
			g.Log().Errorf(ctx, "Outbox 发送失败, id=%d, eventId=%s, retry=%d, err=%v",
				msg.Id, msg.EventId, msg.RetryCount+1, pErr)
			continue
		}

		// 4. 成功
		if uErr := markOutboxSent(ctx, msg.Id); uErr != nil {
			// 已投递成功但状态更新失败：下次会重发，靠消费端按 event_id 幂等去重兜底
			g.Log().Errorf(ctx, "Outbox 标记 sent 失败, id=%d, err=%v", msg.Id, uErr)
			continue
		}
		sent++
		g.Log().Infof(ctx, "Outbox 发送成功, id=%d, eventId=%s", msg.Id, msg.EventId)
	}

	return sent, failed, nil
}

// ResetStuckOutboxSending 恢复卡在 sending 状态的僵尸消息。
//
// claimOutboxMessage 会先把消息抢占为 sending 再投递。如果服务在抢占成功之后、
// 标记 sent/failed 之前崩溃，这条消息会永远停在 sending —— 而 RelayOutboxOnce
// 只捞 pending 和到期的 failed，永远不会再扫到它，导致消息永久卡住。
//
// 本函数把 status=sending 且 updated_at 早于 timeoutMinutes 的记录恢复为 failed、
// next_retry_at 置为当前时间（立即可重试），并累加 retry_count（超时一次记一次失败，
// 避免反复崩溃的消息无限重试）。timeoutMinutes 应明显大于一轮正常投递耗时，
// 避免误伤正在处理中的消息。需在每轮投递前调用。
//
// 注意：被恢复的消息可能其实已投递成功（崩溃发生在 publish 之后），重发依赖
// 消费端按 event_id 幂等去重兜底。
func ResetStuckOutboxSending(ctx context.Context, timeoutMinutes int, limit int) (int, error) {
	if timeoutMinutes <= 0 {
		timeoutMinutes = 5
	}
	if limit <= 0 {
		limit = 100
	}

	cols := dao.OrderOutboxMessage.Columns()
	now := gtime.Now()

	result, err := dao.OrderOutboxMessage.Ctx(ctx).
		Where(cols.Status, int(consts.OutboxStatusSending)).
		Where(fmt.Sprintf("%s < DATE_SUB(NOW(), INTERVAL ? MINUTE)", cols.UpdatedAt), timeoutMinutes).
		Data(g.Map{
			cols.Status:      int(consts.OutboxStatusFailed),
			cols.RetryCount:  gdb.Raw(cols.RetryCount + " + 1"),
			cols.NextRetryAt: now,
			cols.LastError:   "reset from stuck sending",
			cols.UpdatedAt:   now,
		}).
		Limit(limit).
		Update()
	if err != nil {
		return 0, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	return int(rows), nil
}

// publishOutboxMessage 把一条 Outbox 消息投递到 RabbitMQ。
// payload 已是 JSON 字符串，用 RawMessage 原样发出，避免被二次编码。
func publishOutboxMessage(rb *grabbitmq.RabbitMQ, msg *entity.OrderOutboxMessage) error {
	if err := rb.DeclareExchange(msg.Exchange, "topic"); err != nil {
		return err
	}
	return rb.Publish(msg.Exchange, msg.RoutingKey, json.RawMessage(msg.Payload))
}

// claimOutboxMessage 用 CAS 把消息从当前状态抢占为 sending。
// 返回 true 表示抢占成功，false 表示已被他人抢走。
func claimOutboxMessage(ctx context.Context, msg *entity.OrderOutboxMessage) (bool, error) {
	cols := dao.OrderOutboxMessage.Columns()
	result, err := dao.OrderOutboxMessage.Ctx(ctx).
		Where(cols.Id, msg.Id).
		Where(cols.Status, msg.Status).
		Data(g.Map{
			cols.Status:    int(consts.OutboxStatusSending),
			cols.UpdatedAt: gtime.Now(),
		}).Update()
	if err != nil {
		return false, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return false, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	return rows > 0, nil
}

// markOutboxSent 标记消息发送成功。
func markOutboxSent(ctx context.Context, id int64) error {
	cols := dao.OrderOutboxMessage.Columns()
	now := gtime.Now()
	_, err := dao.OrderOutboxMessage.Ctx(ctx).
		Where(cols.Id, id).
		Data(g.Map{
			cols.Status:    int(consts.OutboxStatusSent),
			cols.SentAt:    now,
			cols.UpdatedAt: now,
		}).Update()
	if err != nil {
		return gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	return nil
}

// markOutboxFailed 标记消息发送失败：累加重试次数、按指数退避设置下次重试时间、记录失败原因。
func markOutboxFailed(ctx context.Context, msg *entity.OrderOutboxMessage, cause error) {
	cols := dao.OrderOutboxMessage.Columns()
	now := gtime.Now()
	nextRetryAt := now.Add(outboxBackoff(msg.RetryCount))

	lastErr := cause.Error()
	if len(lastErr) > 500 {
		lastErr = lastErr[:500]
	}

	_, err := dao.OrderOutboxMessage.Ctx(ctx).
		Where(cols.Id, msg.Id).
		Data(g.Map{
			cols.Status:      int(consts.OutboxStatusFailed),
			cols.RetryCount:  msg.RetryCount + 1,
			cols.NextRetryAt: nextRetryAt,
			cols.LastError:   lastErr,
			cols.UpdatedAt:   now,
		}).Update()
	if err != nil {
		g.Log().Errorf(ctx, "Outbox 标记 failed 失败, id=%d, err=%v", msg.Id, err)
	}
}

// outboxBackoff 根据已重试次数计算下次重试的退避时长（指数退避并封顶）。
func outboxBackoff(retryCount int) time.Duration {
	if retryCount < 0 {
		retryCount = 0
	}
	// 限制移位次数，避免溢出
	if retryCount > 6 {
		return outboxBackoffMax
	}
	d := outboxBackoffBase << uint(retryCount)
	if d > outboxBackoffMax {
		return outboxBackoffMax
	}
	return d
}
