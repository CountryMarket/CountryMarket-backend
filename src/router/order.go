package router

import (
	"github.com/CountryMarket/CountryMarket-backend/controller"
	"github.com/CountryMarket/CountryMarket-backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitOrderRouter(g *gin.RouterGroup) {
	g.POST("/generateOrder", middleware.JWTMiddleware(), controller.OrderGenerateOrder)
	g.GET("/userOrder", middleware.JWTMiddleware(), controller.OrderGetUserOrder)
	g.GET("/orderInfo", middleware.JWTMiddleware(), controller.OrderGetOneOrder)
	g.GET("/shopOrder", middleware.JWTMiddleware(), controller.OrderGetShopOrder)
}
