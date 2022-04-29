package middleware

import (
	"github.com/CountryMarket/CountryMarket-backend/util"
	"github.com/CountryMarket/CountryMarket-backend/util/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Error(ctx, http.StatusBadRequest, "no permission，no token", nil)
			ctx.Abort()
			return
		}
		log.Print("token: ", authHeader)

		// 空格键分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Error(ctx, http.StatusBadRequest, "auth error", nil)
			ctx.Abort()
			return
		}

		// 解析 token
		claims, err := util.ParseJWTToken(parts[1])
		if err != nil {
			response.Error(ctx, http.StatusBadRequest, "invalid token", nil)
			ctx.Abort()
			return
		}

		// 将当前请求的 claims 信息保存到请求的上下文 ctx 上
		ctx.Set("claims", claims)
		ctx.Next() // 后续的处理函数可以用过 ctx.Get("claims") 来获取当前请求的用户信息

	}
}
