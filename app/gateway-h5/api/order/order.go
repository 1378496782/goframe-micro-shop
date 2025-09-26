// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package order

import (
	"context"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
)

type IOrderV1 interface {
	OrderInfoGetList(ctx context.Context, req *v1.OrderInfoGetListReq) (res *v1.OrderInfoGetListRes, err error)
	OrderInfoCreate(ctx context.Context, req *v1.OrderInfoCreateReq) (res *v1.OrderInfoCreateRes, err error)
	Payment(ctx context.Context, req *v1.PaymentReq) (res *v1.PaymentRes, err error)
	Notify(ctx context.Context, req *v1.NotifyReq) (res *v1.NotifyRes, err error)
	RefundInfoGetList(ctx context.Context, req *v1.RefundInfoGetListReq) (res *v1.RefundInfoGetListRes, err error)
	RefundInfoGetDetail(ctx context.Context, req *v1.RefundInfoGetDetailReq) (res *v1.RefundInfoGetDetailRes, err error)
	RefundInfoCreate(ctx context.Context, req *v1.RefundInfoCreateReq) (res *v1.RefundInfoCreateRes, err error)
}
