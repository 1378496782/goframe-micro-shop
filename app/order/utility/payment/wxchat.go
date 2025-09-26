package payment

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"net/http"
	v1 "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"
	"strconv"
	"sync"
	"time"
)

var (
	wechatClient *core.Client
	once         sync.Once
)

// IdempotentCheckFunc 幂等校验函数类型
type IdempotentCheckFunc func(context.Context, string) (bool, error)

type weChatConfig struct {
	mchID      string
	serialNo   string
	apiV3Key   string
	privateKey string
	appID      string
	notifyUrl  string
}

func loadConfigParam() weChatConfig {
	return weChatConfig{
		mchID:      g.Cfg().MustGet(nil, "payment.wechat.mchId").String(),
		serialNo:   g.Cfg().MustGet(nil, "payment.wechat.serialNo").String(),
		apiV3Key:   g.Cfg().MustGet(nil, "payment.wechat.apiV3Key").String(),
		privateKey: g.Cfg().MustGet(nil, "payment.wechat.privateKey").String(),
		appID:      g.Cfg().MustGet(nil, "payment.wechat.appId").String(),
		notifyUrl:  g.Cfg().MustGet(nil, "payment.wechat.notifyUrl").String(),
	}
}

// InitWechatClient 初始化单例微信客户端
func InitWechatClient() error {
	keyData := loadConfigParam()
	privateKey, err := utils.LoadPrivateKey(keyData.privateKey)
	if err != nil {
		return gerror.WrapCode(gcode.CodeOperationFailed, err, "加载私钥失败")
	}

	// 初始化微信客户端
	client, err := core.NewClient(context.Background(), option.WithWechatPayAutoAuthCipher(
		keyData.mchID, keyData.serialNo, privateKey, keyData.apiV3Key))
	if err != nil {
		return gerror.WrapCode(gcode.CodeOperationFailed, err, "初始化客户端失败")
	}
	wechatClient = client
	g.Log().Info(context.Background(), "微信客户端初始化成功")
	return nil
}

// 微信支付
func WeChatPayment(ctx context.Context, req *v1.PaymentReq) (*v1.PaymentRes, error) {
	if wechatClient == nil {
		return nil, gerror.WrapCode(gcode.CodeOperationFailed, errors.New("客户端未初始化"))
	}
	wxConf := loadConfigParam()
	svc := jsapi.JsapiApiService{Client: wechatClient}
	prepayReq := jsapi.PrepayRequest{
		Appid:       core.String(wxConf.appID),
		Mchid:       core.String(wxConf.mchID),
		Description: core.String("小程序商品"),
		OutTradeNo:  core.String(req.Number),
		NotifyUrl:   core.String(wxConf.notifyUrl),
		Amount: &jsapi.Amount{
			Total:    core.Int64(req.Amount),
			Currency: core.String("CNY"),
		},
		Payer: &jsapi.Payer{
			Openid: core.String(req.OpenId),
		},
	}

	prepayResp, _, err := svc.Prepay(ctx, prepayReq)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeOperationFailed, err, "向微信发送请求失败")
	}
	if prepayResp == nil || prepayResp.PrepayId == nil {
		return nil, gerror.WrapCode(gcode.CodeOperationFailed, errors.New("prepay_id 为空"))
	}
	nonceStr, err := genNonceStr(32)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeOperationFailed, err, "生成随机数失败")
	}
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	packageStr := "prepay_id=" + *prepayResp.PrepayId
	toSign := fmt.Sprintf("%s\n%s\n%s\n%s\n", wxConf.appID, timeStamp, nonceStr, packageStr)
	sigRes, err := wechatClient.Sign(ctx, toSign)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeOperationFailed, err, "sign failed")
	}

	return &v1.PaymentRes{
		TimeStamp:  timeStamp,
		NonceStr:   nonceStr,
		Package:    packageStr,
		SignType:   "RSA",
		PaySign:    sigRes.Signature,
		OutTradeNo: req.Number,
	}, nil
}

func Notify(ctx context.Context, req *v1.NotifyReq, checkIdempotent IdempotentCheckFunc) (bool, string, error) {
	// 1) 构造 http.Request 给 wechatpay SDK 使用
	httpReq, err := http.NewRequest("POST", "", bytes.NewBuffer([]byte(req.RawBody)))
	if err != nil {
		return false, "", gerror.WrapCode(gcode.CodeOperationFailed, err, "构造 http 请求失败")
	}
	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}

	//// 测试代码(本地测试用)
	//if req.Headers["X-Bypass-Verify"] == "1" {
	//	res := new(payments.Transaction)
	//	if err := json.Unmarshal([]byte(req.RawBody), res); err != nil {
	//		return false, "", gerror.WrapCode(gcode.CodeOperationFailed, err, "测试模式：解析 transaction 失败")
	//	}
	//	if err != nil {
	//		return false, "", gerror.WrapCode(gcode.CodeOperationFailed, err, "outTradeNo to int error")
	//	}
	//
	//	// 幂等校验
	//	alreadyPaid, err := checkIdempotent(ctx, *res.OutTradeNo)
	//	if err != nil {
	//		return false, "", gerror.WrapCode(gcode.CodeOperationFailed, err, "checkIdempotent 幂等性校验失败")
	//	}
	//	if alreadyPaid {
	//		return true, *res.OutTradeNo, nil
	//	}
	//	return false, *res.OutTradeNo, nil
	//}

	wxConf := loadConfigParam()

	// 2) SDK 的证书访问器 + notify handler（自动使用平台证书）
	certVisitor := downloader.MgrInstance().GetCertificateVisitor(wxConf.mchID)
	handler := notify.NewNotifyHandler(wxConf.apiV3Key, verifiers.NewSHA256WithRSAVerifier(certVisitor))

	// 3) 解析并解密 resource 到 PayScoreResource
	res := new(payments.Transaction)
	_, err = handler.ParseNotifyRequest(ctx, httpReq, res)
	if err != nil {
		return false, "", gerror.WrapCode(gcode.CodeOperationFailed, err, "ParseNotifyRequest 验签/解密失败")
	}

	if err != nil {
		return false, "", gerror.WrapCode(gcode.CodeOperationFailed, err, "outTradeNo to int error")
	}
	// 4) 幂等性校验，避免重复修改同一个订单而引发的数据不一致问题
	alreadyPaid, err := checkIdempotent(ctx, *res.OutTradeNo)
	if err != nil {
		return false, "", gerror.WrapCode(gcode.CodeOperationFailed, err, "checkIdempotent 幂等性校验失败")
	}

	// 5) 订单状态已修改
	if alreadyPaid {
		return true, *res.OutTradeNo, nil
	}

	// 6) 订单状态未修改
	return false, *res.OutTradeNo, nil
}

// 生成随机 nonce 字符串（hex）
func genNonceStr(n int) (string, error) {
	b := make([]byte, n/2+1)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	s := hex.EncodeToString(b)
	if len(s) > n {
		s = s[:n]
	}
	return s, nil
}
