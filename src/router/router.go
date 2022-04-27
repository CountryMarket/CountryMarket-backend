package router

import (
	"github.com/CountryMarket/CountryMarket-backend/util/response"
	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.RouterGroup) {
	g.GET("/hh", func(ctx *gin.Context) {
		response.Success(ctx, gin.H{
			"hello": "world",
		})
	})
	rUser := g.Group("/user")
	InitUserRouter(rUser)
}
