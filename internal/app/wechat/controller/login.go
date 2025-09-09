/*
* @desc:用户登录
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/2/24 11:39
 */

package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/auth/response"
	resp "github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/phoneNumber/response"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/wechat"
	systemModel "github.com/tiger1103/gfast/v3/internal/app/system/model"
	systemService "github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/internal/app/wechat/service"
	"github.com/tiger1103/gfast/v3/library/libUtils"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

var Login = new(loginController)

type loginController struct {
	BaseController
}

func (c *loginController) WxUrl(ctx context.Context, req *wechat.WxUrlReq) (res *wechat.WxUrlRes, err error) {
	res = new(wechat.WxUrlRes)
	res.Url, err = service.Wechat().GetWxUrl(ctx)
	return
}

func (c *loginController) WxLogin(ctx context.Context, req *wechat.WxLoginReq) (res *wechat.WxLoginRes, err error) {
	var (
		data      *response.ResponseCode2Session
		phoneData *resp.ResponseGetUserPhoneNumber
		user      *systemModel.LoginUserRes
		token     string
	)
	data, err = service.Wechat().Session(ctx, req.Code)
	if err != nil {
		return
	}
	if data.ResponseMiniProgram.ErrCode != 0 {
		err = gerror.New(fmt.Sprintf("errCode: %d , errMsg: %s", data.ResponseMiniProgram.ErrCode, data.ResponseMiniProgram.ErrMsg))
		return
	}
	phoneData, err = service.Wechat().GetUserPhoneNumber(ctx, req.PhoneCode)
	if err != nil {
		return
	}
	if phoneData.ErrCode != 0 {
		err = gerror.New(fmt.Sprintf("errCode: %d , errMsg: %s", phoneData.ErrCode, phoneData.ErrMsg))
		return
	}
	user, token, err = c.doLogin(ctx, phoneData.PhoneInfo.PhoneNumber, data.OpenID)
	if err != nil {
		return
	}
	res = &wechat.WxLoginRes{
		UserInfo: user,
		Token:    token,
	}
	return
}

func (c *loginController) doLogin(ctx context.Context, phone, openId string) (user *systemModel.LoginUserRes, token string, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		ip := libUtils.GetClientIp(ctx)
		userAgent := libUtils.GetUserAgent(ctx)
		err = g.Try(ctx, func(ctx context.Context) {
			//警员服务器未获取到，从本地获取
			user, err = systemService.SysUser().GetUserByPhone(ctx, phone)
			liberr.ErrIsNil(ctx, err)
			if user == nil {
				// 保存登录失败的日志信息
				systemService.SysLoginLog().Invoke(gctx.New(), &systemModel.LoginLogParams{
					Status:    0,
					Username:  phone,
					Ip:        ip,
					UserAgent: userAgent,
					Msg:       "用户不存在",
					Module:    "微信小程序",
				})
				return
			}
			err = systemService.SysUser().UpdateLoginInfo(ctx, user.Id, ip, openId)
			liberr.ErrIsNil(ctx, err)
		})
		if user == nil {
			liberr.ErrIsNil(ctx, errors.New("用户不存在"))
		}
		// 报存登录成功的日志信息
		systemService.SysLoginLog().Invoke(gctx.New(), &systemModel.LoginLogParams{
			Status:    1,
			Username:  user.Mobile,
			Ip:        ip,
			UserAgent: userAgent,
			Msg:       "登录成功",
			Module:    "微信小程序",
		})
		key := gconv.String(user.Id) + "-" + gmd5.MustEncryptString(gconv.String(user.Id)) + gmd5.MustEncryptString(user.Mobile)
		if g.Cfg("wechat").MustGet(ctx, "gfToken.multiLogin").Bool() {
			key = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(gconv.String(user.Id)) + gmd5.MustEncryptString(user.Mobile+ip+userAgent)
		}
		user.UserPassword = ""
		user.UserSalt = ""
		token, err = service.GfToken().GenerateToken(ctx, key, user)
		liberr.ErrIsNil(ctx, err, "登录失败，后端服务出现错误")
		//用户在线状态保存
		systemService.SysUserOnline().Invoke(gctx.New(), &systemModel.SysUserOnlineParams{
			UserAgent: userAgent,
			Uuid:      gmd5.MustEncrypt(token),
			Token:     token,
			Username:  user.UserNickname,
			Ip:        ip,
		})
	})
	return
}
