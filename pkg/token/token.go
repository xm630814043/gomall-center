package token

import (
	"gomall-center/pkg/settings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Domain   string `json:"domain"`
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var secret []byte

func Setup() {
	secret = []byte(settings.AppConfig.Auth.JwtSecret)
}

func Generate(domain string, userId int64, username string) (string, error) {
	now := time.Now()
	expireTime := now.Add(settings.AppConfig.Auth.ExpireTime)
	claims := Claims{
		domain,
		userId,
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    domain,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(secret)
	return token, err
}

func Parse(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
