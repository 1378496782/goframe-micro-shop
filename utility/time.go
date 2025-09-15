package utility

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
	"time"
)

const (
	KC_RAND_KIND_NUM = 0 // 纯数字
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

// Krand 随机字符串
func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
