package util

import (
	"errors"
	"github.com/CountryMarket/CountryMarket-backend/constant"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"os"
	"time"
)

type JwtClaims struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	jwt.StandardClaims
}

func GenerateJWTToken(openid, sessionKey string) (string, int64, error) {
	claims := &JwtClaims{
		openid,
		sessionKey,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(constant.JwtExpiresDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("APPSECRET")))
	if err != nil {
		return t, 0, err
	}
	return t, claims.ExpiresAt, nil
}
func ParseJWTToken(tokenStr string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("APPSECRET")), nil
	})
	if err != nil {
		log.Println("token parse err: ", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
func GetClaimsFromJWT(ctx *gin.Context) (string, string) {
	t, _ := ctx.Get("claims")
	claims := t.(*JwtClaims)
	return claims.OpenId, claims.SessionKey
}
