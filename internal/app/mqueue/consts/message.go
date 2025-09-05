/**
 * @Company: 云南奇讯科技有限公司
 * @Author: yxf
 * @Description:
 * @Date: 2023/7/31 11:12
 */

package consts

import "time"

type SendMsgMethod int

// MsgDelayLevel 参数1-18 对应时间: 1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h
type MsgDelayLevel int

const (
	MsgIdLength         = 32
	SendMsgPublish      = 1
	SendMsgPublishAsync = 2
	SendMsgDelay        = 3
)

const (
	MsgDelay1s MsgDelayLevel = iota + 1
	MsgDelay5s
	MsgDelay10s
	MsgDelay30s
	MsgDelay1m
	MsgDelay2m
	MsgDelay3m
	MsgDelay4m
	MsgDelay5m
	MsgDelay6m
	MsgDelay7m
	MsgDelay8m
	MsgDelay9m
	MsgDelay10m
	MsgDelay20m
	MsgDelay30m
	MsgDelay1h
	MsgDelay2h
)

func (m MsgDelayLevel) Value() time.Duration {
	delay := 1
	switch m {
	case MsgDelay1s:
		delay = 1
	case MsgDelay5s:
		delay = 5
	case MsgDelay10s:
		delay = 10
	case MsgDelay30s:
		delay = 30
	case MsgDelay1m:
		delay = 60
	case MsgDelay2m:
		delay = 60 * 2
	case MsgDelay3m:
		delay = 60 * 3
	case MsgDelay4m:
		delay = 60 * 4
	case MsgDelay5m:
		delay = 60 * 5
	case MsgDelay6m:
		delay = 60 * 6
	case MsgDelay7m:
		delay = 60 * 7
	case MsgDelay8m:
		delay = 60 * 8
	case MsgDelay9m:
		delay = 60 * 9
	case MsgDelay10m:
		delay = 60 * 10
	case MsgDelay20m:
		delay = 60 * 20
	case MsgDelay30m:
		delay = 60 * 30
	case MsgDelay1h:
		delay = 60 * 60 * 1
	case MsgDelay2h:
		delay = 60 * 60 * 2
	}
	return time.Duration(delay) * time.Second
}
