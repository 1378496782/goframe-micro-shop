package utility

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
	"time"
)

func SafeConvertTime(t *gtime.Time) *timestamppb.Timestamp {
	if t == nil || t.IsZero() {
		return nil
	}
	return timestamppb.New(t.Time)
}

// GenerateOrderNumber 生成订单编号
func GenerateOrderNumber() string {
	return fmt.Sprintf("ORD%s%04d", time.Now().Format("20060102150405"), rand.Intn(9999))
}

// GenerateRefundNumber 生成售后订单编号
func GenerateRefundNumber() string {
	return fmt.Sprintf("REF%s%04d", time.Now().Format("20060102150405"), rand.Intn(9999))
}

// GetOrderBy 排序方式判断函数
func GetOrderBy(sort uint32) string {
	if sort == 2 {
		return "sort desc" // 传2：倒序，sort值越大越靠前
	}
	return "sort asc" // 默认或传1：升序，sort值越小越靠前
}
