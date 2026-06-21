/**
 * @Company: 云南奇讯科技有限公司
 * @Author: yxf
 * @Description:
 * @Date: 2023/7/7 15:36
 */

package model

import "time"

type DiskQueueConfig struct {
	DataPath        string        // 磁盘存储路径
	MaxBytesPerFile int64         // 单个单件最大字节数
	SyncEvery       int64         // 多少次读写后刷盘
	SyncTimeout     time.Duration // 刷盘间隔（秒）
	MaxMsgSize      int32         // 最大消息字节数
	MinMsgSize      int32         // 最小消息字节数
}
