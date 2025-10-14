package time

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

var errstring = "转化错误 请检查参数"

// TimestampToDateTimeString 将 google.protobuf.Timestamp 类型转换为 yyyy-mm-dd HH:MM:SS 格式的字符串
// 如果输入为 nil，则返回空字符串
func TimeString(ts *timestamppb.Timestamp) string {
	if ts == nil {
		return errstring
	}
	// 将 Timestamp 转换为 Go 的 time.Time 类型
	return ts.AsTime().Format("2006-01-02 15:04:05")
}
