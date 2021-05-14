package util

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 前端用户token过期时间
var StudentUserExpireDuration = time.Hour

// 学生用户的加密key
var StudentUserSecretKey = []byte("student_user_token")

type UserToken struct {
	jwt.StandardClaims
	// 自定义的用户信息
	Uid   int64 `json:"user_name"`
	Limit int
}

// 生成token
func GenToken(Uid int64, limit int, expireDuration time.Duration, secret_key []byte) (string, error) {

	user := UserToken{
		jwt.StandardClaims{
			// 现在 + 加上传的过期时间
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
			Issuer:    "micro_gin_vue",
		},
		Uid,
		limit,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	return token.SignedString(secret_key)

}

// 认证token
func AuthToken(tokenString string, secretKey []byte) (*UserToken, error) {
	if tokenString == "" {
		return nil, errors.New("token valid err")
	}

	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserToken{}, func(token *jwt.Token) (key interface{}, err error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	clasims, is_ok := token.Claims.(*UserToken)

	// 验证token
	if is_ok && token.Valid { // 正常的
		return clasims, nil
	}

	return nil, errors.New("token valid err")
}
