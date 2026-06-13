package flash_sale

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/flash_sale/v1"
)

func (c *ControllerV1) GetFlashSaleResult(ctx context.Context, req *v1.GetFlashSaleResultReq) (res *v1.GetFlashSaleResultRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
