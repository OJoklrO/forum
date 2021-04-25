package app

import (
	"time"

	"forum/global"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	Name string `json:"name"`
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateJWTToken(id, name string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			Subject:   id,
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

// todo: learn jwt in depth

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	} else {
		claims, _ := tokenClaims.Claims.(*Claims)
		return claims, nil
	}
}
