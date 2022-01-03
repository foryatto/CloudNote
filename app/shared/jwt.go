package shared

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

/*

jwt-go doc: https://pkg.go.dev/github.com/dgrijalva/jwt-go#section-readme

*/

var JwtUtil = jwtUtil{}

type jwtUtil struct {
	JwtExpireTime int64  // 过期时间
	JwtIssuer     string // 签发人
	JwtSecretKey  []byte // 密匙
}

func (j *jwtUtil) Set(expire int64, issuer, key string) {
	j.JwtExpireTime = expire
	j.JwtIssuer = issuer
	j.JwtSecretKey = []byte(key)
}

// CustomClaims 自定义Claims，添加字段uid
type CustomClaims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}

// Create 创建token，参数uid作为token携带的信息
func (j *jwtUtil) Create(uid string) (string, error) {
	claims := CustomClaims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: j.JwtExpireTime,
			Issuer:    j.JwtIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名，防篡改
	signedToken, err := token.SignedString(j.JwtSecretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// Parse 解析token
func (j *jwtUtil) Parse(signedToken string) (string, error) {
	token, err := jwt.ParseWithClaims(signedToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtSecretKey, nil
	})
	if err != nil {
		return "", err
	}
	// token.Valid 验证有效期
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.Uid, nil
	}
	return "", errors.New("token expired")
}
