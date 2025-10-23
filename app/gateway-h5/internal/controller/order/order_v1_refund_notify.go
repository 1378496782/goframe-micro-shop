package order

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"io"
	"net/http"
	refund_info "shop-goframe-micro-service-refacotor/app/order/api/refund_info/v1"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
)

func (c *ControllerV1) RefundNotify(ctx context.Context, req *v1.RefundNotifyReq) (res *v1.RefundNotifyRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		r.Response.WriteHeader(http.StatusBadRequest)
		r.Response.WriteJson(g.Map{
			"code":    "FAIL",
			"message": "invalid request",
		})
		return
	}

	// 读取完整 body
	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		r.Response.WriteHeader(http.StatusBadRequest)
		r.Response.WriteJson(g.Map{
			"code":    "FAIL",
			"message": "invalid request",
		})
		return
	}

	// headers（只取验签相关的）
	headers := map[string]string{
		"Wechatpay-Signature": r.Header.Get("Wechatpay-Signature"),
		"Wechatpay-Timestamp": r.Header.Get("Wechatpay-Timestamp"),
		"Wechatpay-Nonce":     r.Header.Get("Wechatpay-Nonce"),
		"Wechatpay-Serial":    r.Header.Get("Wechatpay-Serial"),
		"X-Bypass-Verify":     r.Header.Get("X-Bypass-Verify"), // 测试代码，本地测试用
	}

	// 调用 gRPC 服务进行验签/解密/业务处理
	_, err = c.RefundInfoClient.RefundNotify(ctx, &refund_info.RefundNotifyReq{
		RawBody: string(body),
		Headers: headers,
	})
	if err != nil {
		g.Log().Errorf(ctx, "微信退款回调处理失败,err:%v", err)
		r.Response.WriteHeader(http.StatusInternalServerError)
		r.Response.WriteJson(g.Map{
			"code":    "FAIL",
			"message": "网络繁忙",
		})
		return
	}
	r.Response.WriteHeader(http.StatusOK)
	return nil, nil
}
