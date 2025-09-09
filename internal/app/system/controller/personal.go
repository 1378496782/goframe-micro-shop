/*
* @desc:xxxx功能描述
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/11/3 10:32
 */

package controller

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libUtils"
)

var Personal = new(personalController)

type personalController struct {
}

func (c *personalController) GetPersonal(ctx context.Context, req *system.PersonalInfoReq) (res *system.PersonalInfoRes, err error) {
	res, err = service.Personal().GetPersonalInfo(ctx, req)
	return
}

func (c *personalController) EditPersonal(ctx context.Context, req *system.PersonalEditReq) (res *system.PersonalEditRes, err error) {
	ip := libUtils.GetClientIp(ctx)
	userAgent := libUtils.GetUserAgent(ctx)
	res = new(system.PersonalEditRes)
	res.UserInfo, err = service.Personal().EditPersonal(ctx, req)
	if err != nil {
		return
	}
	res.Token, err = c.genToken(ctx, res.UserInfo, ip, userAgent)
	return
}

func (c *personalController) ResetPwdPersonal(ctx context.Context, req *system.PersonalResetPwdReq) (res *system.PersonalResetPwdRes, err error) {
	res, err = service.Personal().ResetPwdPersonal(ctx, req)
	return
}

func (c *personalController) RefreshToken(ctx context.Context, req *system.RefreshTokenReq) (res *system.RefreshTokenRes, err error) {
	var (
		ip        = libUtils.GetClientIp(ctx)
		userAgent = libUtils.GetUserAgent(ctx)
	)
	res = new(system.RefreshTokenRes)
	res.UserInfo, err = service.SysUser().GetUserById(ctx, service.Context().GetUserId(ctx))
	if err != nil {
		return
	}
	if res.UserInfo == nil {
		err = errors.New("用户信息不存在")
		return
	}
	res.Token, err = c.genToken(ctx, res.UserInfo, ip, userAgent)
	return
}

func (c *personalController) genToken(ctx context.Context, userInfo *model.LoginUserRes, ip, userAgent string) (token string, err error) {
	key := gconv.String(userInfo.Id) + "-" + gmd5.MustEncryptString(userInfo.UserName) + gmd5.MustEncryptString(userInfo.UserPassword)
	if g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool() {
		key = gconv.String(userInfo.Id) + "-" + gmd5.MustEncryptString(userInfo.UserName) + gmd5.MustEncryptString(userInfo.UserPassword+ip+userAgent+gtime.Now().String())
	}

	token, err = service.GfToken().GenerateToken(ctx, key, userInfo)
	return
}
