/*
* @desc:验证码处理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/28 9:01
 */

package captcha

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mojocn/base64Captcha"
	"github.com/tiger1103/gfast/v3/internal/app/common/service"
	captchaV2 "github.com/wenlng/go-captcha/captcha"
	"net/url"
)

func init() {
	service.RegisterCaptcha(New())
}

func New() service.ICaptcha {
	return &sCaptcha{
		driver: &base64Captcha.DriverString{
			Height:          80,
			Width:           240,
			NoiseCount:      50,
			ShowLineOptions: 20,
			Length:          4,
			Source:          "abcdefghjkmnpqrstuvwxyz23456789",
			Fonts:           []string{"chromohv.ttf"},
		},
		store: base64Captcha.DefaultMemStore,
	}
}

type sCaptcha struct {
	driver *base64Captcha.DriverString
	store  base64Captcha.Store
}

// GetCaptchaV2 创建点击验证码数据
func (s *sCaptcha) GetCaptchaV2(ctx context.Context) (dots map[int]captchaV2.CharDot, img, thumb, key string, err error) {
	capt := captchaV2.GetCaptcha()
	dots, img, thumb, key, err = capt.Generate()
	return
}

// CheckCaptchaV2 验证captchaV2数据
func (s *sCaptcha) CheckCaptchaV2(ctx context.Context, key string, dots string, removeKey ...bool) (err error) {
	dotsStr, err := gbase64.DecodeToString(dots)
	if err != nil {
		return err
	}
	// 进行Url转码，防止特殊字符问题
	dotsStr, err = url.QueryUnescape(dotsStr)
	dotsMap := gconv.Maps(dotsStr)
	if dotsMap == nil {
		return errors.New("提交的数据无效")
	}

	cacheDots := service.Cache().Get(ctx, "captchaV2_"+key)
	if cacheDots == nil {
		return errors.New("未找到验证数据")
	}
	var dotsMap2 map[int]captchaV2.CharDot
	err = cacheDots.Scan(&dotsMap2)
	if len(dotsMap) != len(dotsMap2) {
		return errors.New("人机验证失败")
	}
	g.Log().Info(ctx, dotsMap, dotsMap2)
	for i, dot := range dotsMap {
		checkStatus := captchaV2.CheckPointDistWithPadding(gconv.Int64(dot["x"]), gconv.Int64(dot["y"]), int64(dotsMap2[i].Dx), int64(dotsMap2[i].Dy), int64(dotsMap2[i].Width), int64(dotsMap2[i].Height), 10)
		if checkStatus == false {
			return errors.New("人机验证失败")
		}
	}
	if len(removeKey) > 0 && removeKey[0] {
		service.Cache().Remove(ctx, "captchaV2_"+key)
	}
	return
}

var (
	captcha = sCaptcha{
		driver: &base64Captcha.DriverString{
			Height:          80,
			Width:           240,
			NoiseCount:      50,
			ShowLineOptions: 20,
			Length:          4,
			Source:          "abcdefghjkmnpqrstuvwxyz23456789",
			Fonts:           []string{"chromohv.ttf"},
		},
		store: base64Captcha.DefaultMemStore,
	}
)

// GetVerifyImgString 获取字母数字混合验证码
func (s *sCaptcha) GetVerifyImgString(ctx context.Context) (idKeyC string, base64stringC string, err error) {
	driver := s.driver.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, s.store)
	idKeyC, base64stringC, _, err = c.Generate()
	return
}

// VerifyString 验证输入的验证码是否正确
func (s *sCaptcha) VerifyString(id, answer string) bool {
	c := base64Captcha.NewCaptcha(s.driver, s.store)
	answer = gstr.ToLower(answer)
	return c.Verify(id, answer, true)
}
