package user

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/user/v1"
	user_info "shop-goframe-micro-service-refacotor/app/user/api/user_info/v1"
)

func (c *ControllerV1) WxMiniLogin(ctx context.Context, req *v1.WxMiniLoginReq) (res *v1.WxMiniLoginRes, err error) {
	// 发起授权
	miniprogram := wechat.NewWechat().GetMiniProgram(&miniConfig.Config{
		AppID:     g.Cfg().MustGet(nil, "wxMiniConf.appId").String(),
		AppSecret: g.Cfg().MustGet(nil, "wxMiniConf.secret").String(),
		Cache:     cache.NewMemory(),
	})
	authResult, err := miniprogram.GetAuth().Code2Session(req.Code)
	if err != nil || authResult.ErrCode != 0 || authResult.OpenID == "" {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "发起授权请求失败")
	}

	// 解析微信返回数据
	userData, err := miniprogram.GetEncryptor().Decrypt(authResult.SessionKey, req.EncryptedData, req.IV)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "解析数据失败")
	}

	// todo 搞清楚哪些内容放在 api，哪些内容放在 grpc
	fmt.Println("userDate,开发到这里了！！！")
	// 绑定用户或登录
	// 使用 gconv 自动转换结构体
	grpcReq := &user_info.WxMiniLoginReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC登录服务
	grpcRes, err := c.UserInfoClient.WxMiniLogin(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 使用gconv转换响应
	res = &v1.WxMiniLoginRes{}
	if err := gconv.Struct(grpcRes, res); err != nil {
		return nil, err
	}

	return res, nil
}
