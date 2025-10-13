package order

import (
	"bytes"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"io"
	"net/http"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
	order_info "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"
)

func (c *ControllerV1) Notify(ctx context.Context, req *v1.NotifyReq) (res *v1.NotifyRes, err error) {
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
	// 恢复 body，防止其他中间件重用失败
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	req.RawBody = string(body)

	// headers（只取验签相关的）
	headers := map[string]string{
		"Wechatpay-Signature": r.Header.Get("Wechatpay-Signature"),
		"Wechatpay-Timestamp": r.Header.Get("Wechatpay-Timestamp"),
		"Wechatpay-Nonce":     r.Header.Get("Wechatpay-Nonce"),
		"Wechatpay-Serial":    r.Header.Get("Wechatpay-Serial"),
		"X-Bypass-Verify":     r.Header.Get("X-Bypass-Verify"), // 测试代码，本地测试用
	}
	req.Headers = headers

	// 调用 gRPC 服务进行验签/解密/业务处理
	_, err = c.OrderInfoClient.Notify(ctx, &order_info.NotifyReq{
		RawBody: string(body),
		Headers: headers,
	})
	if err != nil {
		g.Log().Errorf(ctx, "微信支付回调处理失败,err:%v", err)
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
