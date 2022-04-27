package util

import (
	"errors"
	"github.com/CountryMarket/CountryMarket-backend/constant"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

type JwtClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func GenerateJWTToken(id string) (string, int64, error) {
	claims := &JwtClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(constant.JwtExpiresDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return t, 0, err
	}
	return t, claims.ExpiresAt, nil
}
func ParseJWTToken(tokenStr string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
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
func GetIdFromJWT(ctx *gin.Context) string {
	t, _ := ctx.Get("claims")
	claims := t.(*JwtClaims)
	return claims.Id
}
