package router

import (
	"github.com/CountryMarket/CountryMarket-backend/controller"
	"github.com/CountryMarket/CountryMarket-backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitProductRouter(g *gin.RouterGroup) {
	g.GET("/tabList", controller.ProductGetTabList)
	g.GET("/tabProducts", controller.ProductGetTabProducts)
	g.POST("/tabModify", middleware.JWTMiddleware(), controller.ProductModifyTab)
	g.POST("/tabAdd", middleware.JWTMiddleware(), controller.ProductAddTab)
	g.POST("/tabDelete", middleware.JWTMiddleware(), controller.ProductDeleteTab)
	g.GET("/homeTab", controller.ProductGetHomeTab)
}
