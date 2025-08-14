package utility

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// 生成随机盐值
func GenerateSalt(length int) string {
	return grand.S(length, false)
}

// 密码加密 (双重MD5)
func EncryptPassword(password, salt string) string {
	// 加密(加密密码 + 加密盐)
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password) + gmd5.MustEncryptString(salt))
}

// 生成Token
func GenerateToken(ctx context.Context, userId int) (string, time.Time, error) {
	// 获取JWT配置
	jwtSecret := g.Cfg().MustGet(ctx, "jwt.secret").String()
	expireDays := g.Cfg().MustGet(ctx, "jwt.expire", 7).Int()

	// 计算过期时间
	expireTime := time.Now().Add(time.Duration(expireDays) * 24 * time.Hour)

	// 创建Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     expireTime.Unix(),
	})

	// 签名Token
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		g.Log().Errorf(ctx, "生成Token失败: %v", err)
		return "", time.Time{}, errors.New("系统错误")
	}

	return signedToken, expireTime, nil
}
