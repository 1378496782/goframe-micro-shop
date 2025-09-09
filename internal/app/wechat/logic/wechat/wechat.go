/*
* @desc:微信小程序接口
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/3/22 9:08
 */

package wechat

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/auth/response"
	phoneNumberResp "github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/phoneNumber/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/urlLink/request"
	wxlinkResp "github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/urlLink/response"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tiger1103/gfast/v3/internal/app/wechat/service"
)

func init() {
	service.RegisterWechat(New())
}

func New() service.IWechat {
	var (
		ctx = gctx.New()
		cfg = g.Cfg("wechat")
	)
	miniProgramApp, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:     cfg.MustGet(ctx, "wechat.AppID").String(),  // 小程序appid
		Secret:    cfg.MustGet(ctx, "wechat.Secret").String(), // 小程序app secret
		HttpDebug: cfg.MustGet(ctx, "wechat.HttpDebug").Bool(),
		AESKey:    cfg.MustGet(ctx, "wechat.AesKey").String(),
		Log: miniProgram.Log{
			Level: cfg.MustGet(ctx, "wechat.Log.Leave").String(),
			File:  cfg.MustGet(ctx, "wechat.Log.File").String(),
			Error: cfg.MustGet(ctx, "wechat.Log.Error").String(),
		},
		// 可选，不传默认走程序内存
		Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{cfg.MustGet(ctx, "wechat.Cache.Addr").String()},
			Password: cfg.MustGet(ctx, "wechat.Cache.Password").String(),
			DB:       cfg.MustGet(ctx, "wechat.Cache.DB").Int(),
		}),
	})
	if err != nil {
		panic(err)
	}
	return &sWechat{
		miniProgramApp: miniProgramApp,
	}
}

type sWechat struct {
	miniProgramApp *miniProgram.MiniProgram
}

func (s *sWechat) Session(ctx context.Context, code string) (data *response.ResponseCode2Session, err error) {
	data, err = s.miniProgramApp.Auth.Session(ctx, code)
	return
}

func (s *sWechat) GetUserPhoneNumber(ctx context.Context, code string) (data *phoneNumberResp.ResponseGetUserPhoneNumber, err error) {
	data, err = s.miniProgramApp.PhoneNumber.GetUserPhoneNumber(ctx, code)
	return
}

// IsDebug 是否开发调试模式
func (s *sWechat) IsDebug(ctx context.Context) bool {
	debug := g.Cfg("wechat").MustGet(ctx, "wechat.isDebug").Bool()
	debug2 := g.Cfg("wechat").MustGet(ctx, "wechat.isDebug2").Bool()
	return debug || debug2
}

func (s *sWechat) GetWxUrl(ctx context.Context) (url string, err error) {
	var res *wxlinkResp.ResponseURLLinkGenerate
	res, err = s.miniProgramApp.URLLink.Generate(ctx, &request.URLLinkGenerate{
		EnvVersion:     "release",
		ExpireInterval: 1606737600,
		Path:           "pages/index/index",
		Query:          "",
	})
	if err != nil {
		return
	}
	if res.ErrMsg != "ok" {
		err = gerror.New("获取小程序url失败，err:" + res.ErrMsg + ",errCode:" + res.Code)
		return
	}
	url = res.URLLink
	return
}
