package utility

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// 自定义JWT声明
type CustomClaims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

const (
	JWTSecretKey = "yemcry@163.com"
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

// 生成JWT Token
func GenerateToken(userId int) (string, time.Time, error) {
	expireTime := time.Now().Add(24 * time.Hour)
	claims := CustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		return "", time.Time{}, err
	}

	return signedToken, expireTime, nil
}

// 解析JWT Token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecretKey), nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
