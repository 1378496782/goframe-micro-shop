/*
* @desc:警务采集登陆了
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/2/24 11:38
 */

package wechat

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

type WxUrlReq struct {
	g.Meta `path:"/wxUrl" tags:"微信接口/小程序登陆" method:"get" summary:"获取微信小程序url"`
}

type WxUrlRes struct {
	g.Meta `mime:"application/json"`
	Url    string `json:"url"`
}

type WxLoginReq struct {
	g.Meta    `path:"/wxLogin" tags:"微信接口/小程序登陆" method:"post" summary:"微信登录"`
	Code      string `p:"code" v:"required#code必须"`
	PhoneCode string `p:"phoneCode"`
}

type WxLoginRes struct {
	g.Meta   `mime:"application/json"`
	UserInfo *model.LoginUserRes `json:"userInfo"`
	Token    string              `json:"token"`
}

type UserLoginOutReq struct {
	g.Meta `path:"/logout" tags:"微信接口/小程序登陆" method:"get" summary:"退出登录"`
	commonApi.Author
}

type UserLoginOutRes struct {
}
