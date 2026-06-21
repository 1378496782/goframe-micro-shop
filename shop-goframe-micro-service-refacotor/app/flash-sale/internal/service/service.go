package service

import (
	"github.com/gogf/gf/v2/container/gmap"
)

var (
	localFlashSaleService = gmap.NewStrAnyMap(true)
)

func FlashSale() *FlashSaleService {
	return localFlashSaleService.GetOrSetFuncLock("FlashSale", func() interface{} {
		return NewFlashSaleService()
	}).(*FlashSaleService)
}
